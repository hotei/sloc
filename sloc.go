// sloc.go (c) 2011, 2012 Scott Lawrence <bytbox@gmail.com>
//
// Portions (c) 2015 David Rook <hotei@gmail.com> use git blame for details
//
//	Changes by David Rook <hotei@gmail.com>
//	Version bumped to 0.1.3
//	Add a few more markdown .exts
//	Re-order as types/const/vars (personal preference only)
//	Added some documentation here and there.
//	Added MANY other languages (use -version -verbose to see the list)
//	Added -md flag (enclose output in ``` ... ``` )
//	change some names to something I like better
//		i --> fileinfo
//		n --> input directory path in main() and add() --> pathName?
//	Use mdr.GetArgs() so we can feed it with "find ."
//	Added -stdin flag to turn off local "." default
//	Added Spinner while working since large trees can take a while
//	Changed return method in add to allow printing errors if we get one.
//	Ignore backup files (*~) earlier
//
// ToDo
//
//	Wouldn't hurt to add a few tests.
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime/pprof"
	"sort"
	"strings"
	"text/tabwriter"
	//
	"github.com/hotei/mdr"
)

const VERSION = "sloc version 0.1.3 (c) Scott Lawrence, portions (c) David Rook"

/*
type Language struct {
	Namer
	Matcher
	Commenter
}
*/

type Commenter struct {
	LineComment  string
	StartComment string
	EndComment   string
	Nesting      bool
}

type Namer string

type Matcher func(string) bool

type Stats struct {
	FileCount    int
	TotalLines   int
	CodeLines    int
	BlankLines   int
	CommentLines int
}

type LData []LResult

type LResult struct {
	Name         string
	FileCount    int
	CodeLines    int
	CommentLines int
	BlankLines   int
	TotalLines   int
}

var info = map[string]*Stats{}

var files []string

var (
	flagCpuProfile string
	flagMd         bool
	flagNoOp       bool
	flagStdin      bool
	flagUseJson    bool
	flagVerbose    bool
	flagVersion    bool
)

// changed flags so we can have easier way to alias flags like -V and -version
func init() {
	flag.StringVar(&flagCpuProfile, "cpuprofile", "", "write cpu profile to file")
	flag.BoolVar(&flagUseJson, "json", false, "JSON-format output")
	flag.BoolVar(&flagVersion, "V", false, "display version info and exit")
	flag.BoolVar(&flagVersion, "version", false, "display version info and exit")
	flag.BoolVar(&flagVerbose, "v", false, "user sees more messages")
	flag.BoolVar(&flagVerbose, "verbose", false, "user sees more messages")
	flag.BoolVar(&flagNoOp, "n", false, "don't do count, just list files we look at")
	flag.BoolVar(&flagMd, "md", false, "output is markdown friendly (quoted)")
	flag.BoolVar(&flagStdin, "stdin", false, "read from Stdin, no default (.) arg")
}

func flagSetup() {
	// check flags for interactions/sanity
	if flagVerbose {
		Verbose = VerboseType(true)
	}
	Verbose.Printf("Verbose is set true\n")
	if flagVersion {
		fmt.Printf("Version %s\n", VERSION)
		if flagVerbose {
			for _, lang := range languages {
				fmt.Printf("%s \n", lang.Namer)
			}
		}
		usage()
		os.Exit(0)
	}
	// fmt.Printf("\n")  // bug or feature
}

func usage() {
	s := `usage:
	sloc 
	sloc [directory...]
		-cpuprofile="path"			write cpu profile to file
		-V                           print version info and exit
		-version                     print version info and exit
		-v                           use extra output detail
		-verbose                     use extra output detail
	`
	fmt.Printf("%s\n", s)
	os.Exit(0)
}

// TODO work properly with unicode
func (l Language) Update(c []byte, s *Stats) {
	s.FileCount++

	inComment := 0 // this is an int for nesting
	inLComment := false
	blank := true
	lc := []byte(l.LineComment)
	sc := []byte(l.StartComment)
	ec := []byte(l.EndComment)
	lp, sp, ep := 0, 0, 0

	for _, b := range c {
		if inComment == 0 && b == lc[lp] {
			lp++
			if lp == len(lc) {
				inLComment = true
				lp = 0
			}
		} else {
			lp = 0
		}
		if !inLComment && b == sc[sp] {
			sp++
			if sp == len(sc) {
				inComment++
				if inComment > 1 && !l.Nesting {
					inComment = 1
				}
				sp = 0
			}
		} else {
			sp = 0
		}
		if !inLComment && inComment > 0 && b == ec[ep] {
			ep++
			if ep == len(ec) {
				if inComment > 0 {
					inComment--
				}
				ep = 0
			}
		} else {
			ep = 0
		}

		if b != byte(' ') && b != byte('\t') && b != byte('\n') {
			blank = false
		}

		// BUG(srl): lines with comment don't count towards code
		// Note that lines with both code and comment count towards
		// each, but are not counted twice in the total.
		if b == byte('\n') {
			s.TotalLines++
			if inComment > 0 || inLComment {
				inLComment = false
				s.CommentLines++
			} else if blank {
				s.BlankLines++
			} else {
				s.CodeLines++
			}
			blank = true
			continue
		}
	}
}

func (l Namer) Name() string { return string(l) }

func (m Matcher) Match(fname string) bool { return m(fname) }

func mExt(exts ...string) Matcher {
	return func(fname string) bool {
		for _, ext := range exts {
			if ext == strings.ToLower(path.Ext(fname)) {
				return true
			}
		}
		return false
	}
}

func mName(names ...string) Matcher {
	return func(fname string) bool {
		for _, name := range names {
			if name == path.Base(fname) {
				return true
			}
		}
		return false
	}
}

func handleFile(fname string) {
	mdr.Spinner()
	var l Language
	found := false
	for _, lang := range languages {
		if lang.Match(fname) {
			found = true
			l = lang
			break
		}
	}
	if !found {
		return // ignore this file
	}

	fileinfo, ok := info[l.Name()]
	if !ok {
		// language has no existing stat struct so make a new one
		fileinfo = &Stats{}
		info[l.Name()] = fileinfo
	}
	c, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "  !Err---> file %s is not readable err:%v\n",
			fname, err)
		return
	}
	l.Update(c, fileinfo)
}

// add a path to the ones we want to inspect later
//	side effect is to append to files
//	note this is a recursive function
//	? better served by path/filepath/walk?
func add(pathName string) {
	mdr.Spinner()
	// Verbose.Printf("pathName %s tail %c\n", pathName, pathName[len(pathName)-1])
	// short-circuit if it's a backup file
	if pathName[len(pathName)-1] == '~' {
		return
	}
	// short-circuit known tarpits like ".git" and probably ".gvfs"
	// need a []string to keep list of these like in mdserver
	fi, err := os.Stat(pathName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "  !Err---> cant get info for path %s : %v\n",
			pathName, err)
		return
	}
	if fi.IsDir() {
		fs, err := ioutil.ReadDir(pathName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  !Err---> cant get read %s : %v\n",
				pathName, err)
			return
		}
		for _, f := range fs {
			if f.Name()[0] != '.' { // ignore self, parent and hidden files
				add(path.Join(pathName, f.Name())) // recursively do this
			}
		}
		return
	}
	// if its a regular file then add it to the files to process list
	if fi.Mode()&os.ModeType == 0 {
		files = append(files, pathName)
		return
	}
	// if it's not a regular file ignore it but print the mode?  How does this  help?
	println(fi.Mode())
	fmt.Fprintf(os.Stderr, "  !Err---> path %s is not a regular file\n", pathName)
}

func (d LData) Len() int { return len(d) }

func (d LData) Less(i, j int) bool {
	if d[i].CodeLines == d[j].CodeLines {
		return d[i].Name > d[j].Name
	}
	return d[i].CodeLines > d[j].CodeLines
}

func (d LData) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (r *LResult) Add(a LResult) {
	r.FileCount += a.FileCount
	r.CodeLines += a.CodeLines
	r.CommentLines += a.CommentLines
	r.BlankLines += a.BlankLines
	r.TotalLines += a.TotalLines
}

func printJSON() {
	bs, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}

func printInfo() {
	w := tabwriter.NewWriter(os.Stdout, 2, 8, 2, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "Language\tFiles\tCode\tComment\tBlank\tTotal\t")
	d := LData([]LResult{})
	total := &LResult{}
	total.Name = "Total"
	for n, i := range info {
		r := LResult{
			n,
			i.FileCount,
			i.CodeLines,
			i.CommentLines,
			i.BlankLines,
			i.TotalLines,
		}
		d = append(d, r)
		total.Add(r)
	}
	d = append(d, *total)
	sort.Sort(d)
	//d[0].Name = "Total"
	for _, i := range d {
		fmt.Fprintf(
			w,
			"%s\t%d\t%d\t%d\t%d\t%d\t\n",
			i.Name,
			i.FileCount,
			i.CodeLines,
			i.CommentLines,
			i.BlankLines,
			i.TotalLines)
	}
	w.Flush()
}

//
func main() {
	var args []string
	flag.Parse()
	flagSetup()
	// note - this must stay in main() so profile data is flushed at close
	if flagCpuProfile != "" {
		f, err := os.Create(flagCpuProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			return
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	if flagStdin {
		args = mdr.GetAllArgs()
	} else {
		args = flag.Args()
		if len(args) == 0 {
			args = append(args, `.`)
		}
	}
	if flagMd {
		fmt.Printf("```\n")
		fmt.Printf("SLOC output\n\n")
	}
	for _, filePathName := range args {
		add(filePathName)
	}

	if Verbose {
		for i := 0; i < len(files); i++ {
			fmt.Printf("File [%3d] %s\n", i+1, files[i])
		}
	}
	if flagNoOp {
		fmt.Printf("files[] has %d entries\n", len(files))
		fmt.Println("-n flag given - exiting")
		os.Exit(0)
	}

	for _, f := range files {
		handleFile(f)
	}

	if flagUseJson {
		printJSON()
	} else {
		printInfo()
	}
	if flagMd {
		fmt.Printf("```\n")
	}
}

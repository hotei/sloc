<center>
# sloc
</center>

sloc.go (c) 2011, 2012 Scott Lawrence <bytbox@gmail.com>

Portions (c) 2015 David Rook <hotei@gmail.com> use git blame for details

	Changes by David Rook <hotei1352@gmail.com>
	Version bumped to 0.1.3
	Add a few more markdown .exts
	Re-order as types/const/vars (personal preference only)
	Added some documentation here and there.
	Added MANY other languages (use -version -verbose to see the list)
	Added -md flag (enclose output in ``` ... ``` )
	change some names to something I like better
		i --> fileinfo
		n --> input directory path in main() and add() --> pathName?
	Use mdr.GetArgs() so we can feed it with "find ."
	Added -stdin flag to turn off local "." default
	Added Spinner while working since large trees can take a while
	Changed return method in add to allow printing errors if we get one.
	Ignore backup files (*~) earlier

ToDo

	UPDATE DOCS


github.com/hotei/sloc imports 

```
	C
	bufio
	bytes
	crypto
	crypto/md5
	crypto/sha256
	encoding
	encoding/base64
	encoding/json
	errors
	flag
	fmt
	github.com/hotei/datatable
	github.com/hotei/mdr
	hash
	hash/crc64
	internal/singleflight
	io
	io/ioutil
	log
	math
	math/rand
	net
	os
	os/exec
	path
	path/filepath
	reflect
	regexp
	regexp/syntax
	runtime
	runtime/pprof
	sort
	strconv
	strings
	sync
	sync/atomic
	syscall
	text/tabwriter
	time
	unicode
	unicode/utf16
	unicode/utf8
	unsafe


SLOC output

  Language  Files  Code  Comment  Blank  Total
     Total     11   997      119    261   1377
  Markdown      4   554        0    203    757
        Go      3   394      111     46    551
      Make      1    33        8      9     50
      Text      1    16        0      3     19
    Pascal      2     0        0      0      0
```


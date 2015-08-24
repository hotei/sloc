<center>
# sloc - Source Lines Of Code
</center>

<h3>   
<a href="http://godoc.org/github.com/hotei/sloc">
<img src="https://godoc.org/github.com/hotei/sloc?status.png" alt="GoDoc"/>
</a>
</h3>
<h2> Table of Contents</h2>

* <a href="#Sample Output">Sample Output</a>
* <a href="#Why should I use it">Why should I use it</a>
* <a href="#Installation">Installation</a>
* <a href="#Usage">Usage</a>
* <a href="#Limitations">Limitations</a>
* <a href="#Change Log">Change Log</a>
* <a href="#References">References</a>
* <a href="#Disclaimer">Disclaimer</a>
* <a href="#Licenses">Licenses</a>
* <a href="#autogen">local GODOC, dependency graph, deadcode, sloc</a>

------

<a href="http://travis-ci.org/hotei/sloc">
<img src="https://secure.travis-ci.org/hotei/sloc.png" alt="Build Status" /></a>Travis Status<br>

This document "README-sloc.md" is (c) 2015 David Rook.   

<a name="Sample Output">
### Sample Output
```
SLOC output for $HOME/mdr/GO

   Languages  Files     Code  Comment   Blank    Total  CodeLns
       Total  19826  5079185   811107  480229  6370521   100.0%
          Go  14134  3520492   534403  363020  4417915    69.3%
           C    947   326381   146560   42495   515436     8.1%
        Text    544   285006     1530   14595   301131     4.7%
         PDF    123   241839    28012    2124   271975     4.3%
        HTML    719   183067     2175   11670   196912     3.1%
         SVG    132   134237        0      75   134312     2.1%
         Rtf      2    61456        0       5    61461     1.0%
    Assembly    402    57484     3591    6914    67989     1.1%
         C++    575    54929    76486    5616   137031     2.2%
    Markdown    678    54875        0   20165    75040     1.2%
        JSON     89    38855        0      10    38865     0.6%
         XML     97    24945      359     707    26011     0.4%
        Java    194    20611    11006      32    31649     0.5%
         CSS    114    18250      884    1929    21063     0.3%
  JavaScript     91    16996     2618    3328    22942     0.4%
        Conf     55     9797        0     586    10383     0.2%
      Python     51     6360     1217    1269     8846     0.1%
        Make    272     4052      735    1574     6361     0.1%
        Rust     32     3206        0    1042     4248     0.1%
        YACC      6     3066      193     315     3574     0.1%
      Prolog    132     2120        0     566     2686     0.0%
         PHP      2     2098        8     586     2692     0.0%
       Shell    110     1812      634     422     2868     0.0%
        Bash     31     1178      382     283     1843     0.0%
         Lua     48     1172        0     186     1358     0.0%
        TOML    159     1092        0     178     1270     0.0%
        Perl     11      939      168     152     1259     0.0%
      Pascal      4      874        0       0      874     0.0%
        Tmpl     18      708        0     118      826     0.0%
       CMake     14      402       76     112      590     0.0%
        Ruby      6      284       54      70      408     0.0%
        YAML     16      262        0      48      310     0.0%
        Lisp     14      174       16      36      226     0.0%
           D      3      153        0       0      153     0.0%
         AWK      1       13        0       1       14     0.0%
```
<a name="Why should I use it">
### Why should I use it?

sloc can provide a picture of how a project was developed.  This can be critical
if you need to provide support for every language that's in the project's list.

<a name="Installation">
### Installation

If you have a working go installation on a Unix-like OS:

> ```go get github.com/hotei/sloc```

If go is not yet installed:
```
cd DestinationDirectory
git clone https://github.com/hotei/sloc.git
```

If you don't plan to instal go you can still download the git repository as
a zipfile and extract it to examine the contents.  However, if you just want a quick
look the best way is probably to inspect the GoDoc.Org link above.

</a>

<a name="Usage">
### Usage (incl option flags)

```
sloc                                | look in current directory for files to count
sloc $HOME/src                      | user's source directory
sloc -version -verbose              | list extra info, incl. languages recognized
find / -type f | sloc -stdin        | feed it files on stdin

options:
    - stdin                         | args will be taken from stdin, no local "." default
    - n                             | don't do count, just list files that would be counted
    - V or -version                 | print version and quit
    - v or -verbose                 | use more extensive user messages
    - md                            | make output markdown friendly with ``` quotes
    - json                          | output results in json format
    - cpuprofile=fname              | save cpu profile to fname
```

<a name="Limitations">
### Limitations

* It's possible that extensions can be mis-appropriated.  One popular IDE uses
.pro for config files.  That's also an extension indicating the Prolog language.
There probably are other name collisions I'm not aware of.
* The matching method is a little odd, as is the type Language struct. I 
may tinker with that once I'm a little more comfortable with the code.  It runs
fast enough now as it can process about a million lines per second. I'd like
to do a few timing runs and maybe a profile since Scott put the hooks in there
already.

<a name="Change Log">
### Change Log

* 2015-08-24 doc updates
* 2015-08-23 forked project and rebuilt with go 1.5

<a name="References">
### References

* [go language reference] [1] 
* [go standard library package docs] [2]
* [Source for sloc on github] [3]
* [Go projects list(s)] [7]
* [Excellent godoc howto by Nate Finch] [8]

[1]: http://golang.org/ref/spec/ "go reference spec"
[2]: http://golang.org/pkg/ "go package docs"
[3]: http://github.com/hotei/sloc "github.com/hotei/sloc"
[4]: http://golang.org/doc/go1compat.html "Go 1.x API contract"
[5]: http://blog.golang.org/2011/06/profiling-go-programs.html "Profiling go code"
[6]: http://golang.org/doc/articles/godoc_documenting_go_code.html "GoDoc HowTo"
[7]: https://github.com/golang/go/wiki/Projects "go project list"
[8]: https://github.com/natefinch/godocgo "Nate Finch's Tutorial for GoDoc"

Comments can be sent to David Rook  <hotei1352@gmail.com>  Any issues/bugs
can be mentioned in email or filed on github.

<a name="Disclaimer">
### Disclaimer
Any trademarks mentioned herein are the property of their respective owners.

<a name="Licenses">
### Licenses
The original code contained the following license:

> Copyright (c) 2011, 2012 Scott Lawrence <bytbox@gmail.com>
> 
> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included in
> all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
> THE SOFTWARE.

-----

> Those portions of the 'sloc' go program/package and demo programs written by
> David Rook are distributed under the Simplified BSD License.  You can identify
> those portion using git blame if so inclined.
> 
>
> Portions copyright (c) 2015 David Rook. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

----
<a name="autogen">
<center>
# Automatically Generated Documentation Follows
</center>



DO NOT EDIT BELOW - Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)

# .

sloc.go (c) 2011, 2012 Scott Lawrence <bytbox@gmail.com>

Portions (c) 2015 David Rook <hotei@gmail.com> use git blame for details


	Changes by David Rook <hotei@gmail.com>
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
	Print Percent of code lines each language comprises
	Removed requirement that files be all lowercase to match

ToDo


	Wouldn't hurt to add a few tests.








- - -
DO NOT EDIT ABOVE - Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)

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
```
```
deadcode results:

```
```
SLOC output

  Languages  Files  Code  Comment  Blank  Total  CodeLns
      Total     11  1004      124    174   1302   100.0%
   Markdown      4   555        0    115    670    51.5%
         Go      3   397      116     46    559    42.9%
       Make      1    36        8     10     54     4.1%
       Text      1    16        0      3     19     1.5%
     Pascal      2     0        0      0      0     0.0%
```
```
built with go version = go version go1.5 linux/amd64
```

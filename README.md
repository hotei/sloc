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
* <a href="#Change Log">Change Log</a>
* <a href="#References">References</a>
* <a href="#Disclaimer">Disclaimer</a>
* <a href="#Licenses">Licenses</a>
* <a href="#autogen">local GODOC, dependency graph, deadcode, sloc</a>

------

<a href="http://travis-ci.org/hotei/sloc">
<img src="https://secure.travis-ci.org/hotei/sloc.png" alt="Build Status" /></a>Travis Status<br>

This document "README-sloc.md" is (c) 2015 David Rook.   

`sloc` is a simple, do-one-thing-well program to calculate code statistics: the
number of lines in a project, and how much of that is code versus comment.

`sloc` cannot understand gitignore or hgignore, nor can it distinguish between
"real" source and auto-generated files, so for best results, run it on a fresh
repository with no compilation done.

You can generate JSON output with the `-json` flag, if that's easy to parse in
the programming/scripting language of your choice.

<a name="Sample Output">
### Sample Output

```
 $ sloc ~/misc/opt/go
        Language  Files    Code  Comment  Blank   Total
           Total   2808  512357    87177  67791  667325
              Go   2048  295054    62020  37973  395047
               C    474  166702    22330  21849  210881
            HTML     58   25627      183   4241   30051
        Assembly    131    9974      161   1491   11626
            YACC      6    5245      363    449    6057
          Python      6    2940      789    495    4224
      JavaScript      6    2526      496    585    3607
             XML      9     974       15     90    1079
           Shell      9     905      380    155    1440
             CSS      6     899       24    119    1042
            Perl      9     854      159    135    1148
            Bash     13     483      151    122     756
            Make     33     174      106     87     367
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
look the best way is probably to inspect the GoDoc.Org version at this link:
<a href="http://godoc.org/github.com/hotei/sloc">
<img src="https://godoc.org/github.com/hotei/sloc?status.png" alt="GoDoc"/>
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
	
<a name="Change Log">
### Change Log

* 2015-08-20 rebuilt with go 1.5

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

  Language  Files  Code  Comment  Blank  Total
     Total     11   946      121    173   1240
  Markdown      4   500        0    114    614
        Go      3   394      113     46    553
      Make      1    36        8     10     54
      Text      1    16        0      3     19
    Pascal      2     0        0      0      0
```
```
built with go version = go version go1.5 linux/amd64
```

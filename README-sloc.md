<center>
# ProgRAM
</center>

NOTE: Use Search&Replace to change ProgRAM to the new programs name.

The godoc button/link that follows is for ansiterm, you must replace the ansiterm
referece with ProGRAM once godoc.org has picked it up from github.com. This may
require godoc.org to be "prompted" to do same.

<h3>   
<a href="http://godoc.org/github.com/hotei/ansiterm">
<img src="https://godoc.org/github.com/hotei/ansiterm?status.png" alt="GoDoc"/>
</a>
</h3>

Travis tags are good but require registration at http://travis-ci.org and then
a bit of setup to put in a tarvis.yaml file.  See their homepage 
http://godoc.org and http://docs.travis-ci.com/user/languages/go/ for specifics.

<!---
<a href="http://travis-ci.org/hotei/ansiterm">
<img src="https://secure.travis-ci.org/hotei/ansiterm.png" alt="Build Status" /></a>
--->

This document "README-ProgRAM.md" is (c) 2015 David Rook. 

Note that it may have (possibly many?) more 
headings than necessary, as ideas were borrowed from multiple sources. Normally 
this is the first portion of README.md, where additional parts are added from
one or more of:
```
godoc2md (by Dave Cheny)
godepgraph (by Aram Hăvărneanu)
deadcode (by Remy Oudompheng)
```

Finally a tag is added for which compiler version was in vogue when the docs were created.
That MIGHT not be the same one used to actually generate the code as they can
get out of sync for various reasons.  This can be fixed, but a full retrofit will take
a while.


note - to embed an image inline it must be in same folder as the md file, or href'd to an
active server where the file is available. In our case the server is a go program
__mdserver__ running on 8281 on our local hardware, not available over the net.


http://127.0.0.1:8281/images --> ~/Desktop/GO/images

![Hotei](http://127.0.0.1:8281/images/chip.jpg "Hotei")	

for a link to an image just type the href, as : http://127.0.0.1:8281/images/Rex.png

ProgRAM is (c) 2015 David Rook - all rights reserved. The program
and related files in this repository are released under BSD 2-Clause License.
License details are at the end of this document. Bugs/issues can be reported on github.
Comments can be sent to <hotei1352@gmail.com>. "Thank you"s, constructive 
criticism and job offers are always welcome.  If you are so inclined you may
donate money to help continue the project by sending PayPal contributions to
my wife at diane@celticpapers.com. 

### Overview
# SLOC - Source Lines Of Code

`sloc` is a simple, do-one-thing-well program to calculate code statistics: the
number of lines in a project, and how much of that is code versus comment.

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

`sloc` cannot understand gitignore or hgignore, nor can it distinguish between
"real" source and auto-generated files, so for best results, run it on a fresh
repository with no compilation done.

You can generate JSON output with the `-json` flag, if that's easy to parse in
the programming/scripting language of your choice.

### Description

### What Is It?

### Why should I use it?

ProgRAM is a go program to solve problems

### How does it work?

### <font color=red> >>> Please Read This First <<< </font>

### Background Info on the problem

### Installation

If you have a working go installation on a Unix-like OS:

> ```go get github.com/hotei/ProgRAM```

If go is not yet installed:
```
cd DestinationDirectory
git clone https://github.com/hotei/ProgRAM.git
```

If you don't plan to instal go you can still download the git repository as
a zipfile and extract it to examine the contents.  However, if you just want a quick
look the best way is probably to inspect the GoDoc.Org version at this link:
<a href="http://godoc.org/github.com/hotei/ansiterm">
<img src="https://godoc.org/github.com/hotei/ansiterm?status.png" alt="GoDoc"/>
</a>

### Prerequisites

### Configuration

### Features

### Usage (incl option flags)

### Sample Output

### ScreenDumps

### Tests

### Example Programs

### Benchmarks

### Notes

### To do

NOTE:  "higher" relative priority is at top of list

1.  Added as needed (none active at the moment)

### Limitations

### Package contents

* Programs included
* Data files required
* Documentation provided

### Bugs

1.  Added as needed (none active at the moment)
`

### Development Environment
	Mint 17.1 Linux on i7/2500 mhz 8 'core' (4+4HT) HP Envy Laptop
	X11/R6
	gnu g++ compiler gcc version 4.8.2 (Ubuntu 4.8.2-19ubuntu1)
	go 1.5rc1


### Tests

```
Testing generally looks like this:


go test  // runs all the tests

go test -run=0001    // runs Test_0001 or you can use re2 ".*" for all

if all tests pass then go will allow you to run benchmarks

go test -test.bench=".*" to run all benchmarks
```
-----

### Profiling

Depending on how long the program is expected to run a profile may be in order.  

```
go tool pprof ./ProgRAM logit.prof

(then type help for options)

`web` is usually interesting but you may need to install Graphviz to use it.
$sudo apt-get install Graphviz

`svg` is useful with eog or chrome as the viewer.

```

	
### Change Log

* 2015-xx-xx rebuilt with go 1.5rc1
* 2015-xx-01 rebuilt with go 1.4.2
* 2014-xx-xx first commit to github.com
* 2009-11-07 translated ProgRAM to go
* 1969-09-30 wrote ProgRAM in Fortran II

### References

* [go language reference] [1] 
* [go standard library package docs] [2]
* [Source for ProgRAM on github] [3]
* [Go projects list(s)] [7]
* [Excellent godoc howto by Nate Finch] [8]

[1]: http://golang.org/ref/spec/ "go reference spec"
[2]: http://golang.org/pkg/ "go package docs"
[3]: http://github.com/hotei/ProgRAM "github.com/hotei/ProgRAM"
[4]: http://golang.org/doc/go1compat.html "Go 1.x API contract"
[5]: http://blog.golang.org/2011/06/profiling-go-programs.html "Profiling go code"
[6]: http://golang.org/doc/articles/godoc_documenting_go_code.html "GoDoc HowTo"
[7]: https://github.com/golang/go/wiki/Projects "go project list"
[8]: https://github.com/natefinch/godocgo "Nate Finch's Tutorial for GoDoc"

Comments can be sent to David Rook  <hotei1352@gmail.com>  Any issues/bugs
can be mentioned in email or filed on github.

### Disclaimer
Any trademarks mentioned herein are the property of their respective owners.

### License

The 'ProgRAM' go program/package and demo programs are distributed under the Simplified BSD License:

> Copyright (c) 2015 David Rook. All rights reserved.
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

<center>
# Automatically Generated Documentation Follows
</center>



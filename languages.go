// languages.go

package main

// interesting that it's an untyped struct...
type Language struct {
	Namer
	Matcher
	Commenter
}

var languages = []Language{

	Language{"Ada", mExt(".ada"), noComments},     // Ada 8888 needs work
	Language{"ALGOL", mExt(".algol"), noComments}, // ALGOL 8888 needs work
	Language{"APL", mExt(".apl"), noComments},     // Apl 8888 needs work
	Language{"Assembly", mExt(".asm", ".s"), semiComments},
	Language{"AWK", mExt(".awk"), noComments}, // AWK 8888 needs work
	Language{"Bash", mExt(".bash"), shComments},
	Language{"Basic", mExt(".bas"), noComments}, // Basic 8888 needs work
	Language{"BCPL", mExt(".???"), noComments},  // BCPL 8888 needs work
	Language{"C", mExt(".c", ".h"), cComments},
	Language{"C#", mExt(".c#", ".h#"), cComments}, // C# 8888 needs work
	Language{"C++", mExt(".cc", ".cpp", ".cxx", ".hh", ".hpp", ".hxx"), cComments},
	Language{"CMake", mName("CMakeLists.txt"), shComments}, // ??? 8888
	Language{"Cobol", mExt(".cobol"), noComments},          // Cobol 8888 needs work
	Language{"Conf", mExt(".conf"), noComments},            // Conf 8888 needs work
	Language{"CSS", mExt(".css"), cssComments},
	Language{"D", mExt(".d"), noComments},                               // D 8888 needs work
	Language{"Delphi", mExt(".?"), noComments},                          // Delphi 8888 needs work
	Language{"Eiffel", mExt(".efl"), noComments},                        // Eiffel 8888 needs work
	Language{"Erlang", mExt(".erl"), noComments},                        // Erlang 8888 needs work
	Language{"F", mExt(".f"), noComments},                               // F 8888 needs work
	Language{"Forth", mExt(".4th"), noComments},                         // Forth 8888 needs work
	Language{"Fortran", mExt(".f", ".for", ".f90", ".f95"), noComments}, // Fortran 8888 needs work
	Language{"Go", mExt(".go"), cComments},
	Language{"HAML", mExt(".haml"), noComments},
	Language{"Haskell", mExt(".hs", ".lhs"), hsComments},
	Language{"HTML", mExt(".htm", ".html", ".xhtml"), xmlComments},
	Language{"Jam", mName("Jamfile", "Jamrules"), shComments},
	Language{"Java", mExt(".java", ".jar", ".class"), cComments},
	Language{"JavaScript", mExt(".js"), cComments},
	Language{"JSON", mExt(".json"), noComments}, // JSON 8888 needs work
	Language{"Julia", mExt(".jul"), noComments}, // Julia 8888 needs work
	Language{"Lex", mExt(".l"), cComments},
	Language{"Lisp", mExt(".lsp", ".lisp"), semiComments},
	Language{"LilyPond", mExt(".ly"), noComments}, // LilyPond 8888 needs work
	Language{"Logo", mExt(".logo"), noComments},   // Logo 8888 needs work
	Language{"Lua", mExt(".lua"), noComments},     // Lua 8888 needs work
	Language{"M4", mName(".m4"), noComments},
	Language{"Make", mName("makefile", "Makefile", "MAKEFILE"), shComments},
	Language{"Markdown", mExt(".md", ".mdown", ".markdown"), noComments},
	Language{"ML", mExt(".ml"), noComments},                   // ML 8888 needs work
	Language{"Matlab", mExt(".mat"), noComments},              // Matlab 8888 needs work
	Language{"Mathematica", mExt(".mat"), noComments},         // Mathematica 8888 needs work
	Language{"Modula", mExt(".mod"), noComments},              // Modula 8888 needs work
	Language{"Modula-2", mExt(".mod2"), noComments},           // Modula-2 8888 needs work
	Language{"Modula-3", mExt(".mod2"), noComments},           // Modula-3 8888 needs work
	Language{"Mono", mExt(".mat"), noComments},                // Mono 8888 needs work
	Language{"Mumps", mExt(".mumps"), noComments},             // Mumps 8888 needs work
	Language{"Objective-C", mExt(".c/.h/.m/.mm"), noComments}, // Objective-C 8888 needs work
	Language{"Oberon", mExt(".???"), noComments},              // Oberon 8888 needs work
	Language{"Occam", mExt(".???"), noComments},               // Occam 8888 needs work
	Language{"Octave", mExt(".???"), noComments},              // Octave 8888 needs work
	Language{"OpenCL", mExt(".???"), noComments},              // OpenCL 8888 needs work
	Language{"Pascal", mExt(".pas"), cComments},
	Language{"PDF", mExt(".ps", ".pdf"), noComments}, // PDF 8888 needs work
	Language{"Perl", mExt(".pl", ".pm"), shComments},
	Language{"PHP", mExt(".php"), cComments},
	Language{"PL/1", mExt(".pl1", "plc"), noComments},       // Julia 8888 needs work
	Language{"Postscript", mExt(".ps", ".eps"), noComments}, // Postscript 8888 needs work
	Language{"Prolog", mExt(".pro", ".prolog"), noComments}, // Prolog 8888 needs work
	Language{"Python", mExt(".py"), pyComments},
	Language{"R", mExt(".r"), noComments},                                         // R 8888 needs work
	Language{"RatFor", mExt(".ratfor", ".rat4", ".rat5", ".ratfive"), noComments}, // Rational Fortran 8888 needs work
	Language{"Ruby", mExt(".rb"), shComments},
	Language{"Rust", mExt(".rst"), noComments}, // Rust 8888 needs work
	Language{"Rtf", mExt(".rtf"), noComments},  // Rtf 8888 needs work
	Language{"S", mExt(".s"), noComments},      // S 8888 needs work
	Language{"SASS", mExt(".sass"), cssComments},
	Language{"Scala", mExt(".scala"), cComments},
	Language{"Scheme", mExt(".scm", ".scheme"), semiComments}, // aka Guile?
	Language{"SCSS", mExt(".scss"), cssComments},
	Language{"Shell", mExt(".sh", ".zsh"), shComments},
	Language{"SQL", mExt(".sql"), sqlComments},
	Language{"Smalltalk", mExt(".???"), noComments}, // Smalltalk 8888 needs work
	Language{"SVG", mExt(".svg"), noComments},       // SVG 8888 needs work
	Language{"Text", mExt(".txt"), noComments},      // Text 8888 needs work
	Language{"Tex", mExt(".tex"), noComments},       // Tex 8888 needs work
	Language{"TCL", mExt(".tex"), noComments},       // TCL 8888 needs work
	Language{"Thrift", mExt(".thrift"), cComments},
	Language{"Tmpl", mExt(".tmpl"), noComments},         // Tmpl 8888 needs work
	Language{"TOML", mExt(".toml"), noComments},         // TOML 8888 needs work
	Language{"VBScript", mExt(".vbscript"), noComments}, // VBScript 8888 needs work
	Language{"XML", mExt(".xml"), xmlComments},
	Language{"YACC", mExt(".y"), cComments},
	Language{"YAML", mExt(".yaml"), noComments}, // YAML 8888 needs work
}

var (
	noComments   = Commenter{"\000", "\000", "\000", false}
	xmlComments  = Commenter{"\000", `<!--`, `-->`, false}
	cComments    = Commenter{`//`, `/*`, `*/`, false}
	cssComments  = Commenter{"\000", `/*`, `*/`, false}
	shComments   = Commenter{`#`, "\000", "\000", false}
	semiComments = Commenter{`;`, "\000", "\000", false}
	hsComments   = Commenter{`--`, `{-`, `-}`, true}
	sqlComments  = Commenter{`--`, "\000", "\000", false}
	pyComments   = Commenter{`#`, `"""`, `"""`, false}
)

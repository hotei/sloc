# Makefile for sloc program

PROG = sloc
VERSION = 0.1.3
TARDIR = 	$(HOME)/Desktop/TarPit/
DATE = 	`date "+%Y-%m-%d.%H_%M_%S"`
DOCOUT = README-$(PROG)-godoc.md

all:
	go build -v

install:
	go build -v
	go tool vet .
	go tool vet -shadow .
	gofmt -w *.go
# pick install for packages, cp for programs (or dont copy if you wish)
	go install
#	cp $(PROG) $(HOME)/bin

# note that godepgraph can be used to derive .travis.yml install: section
docs:
	godoc2md . > $(DOCOUT)
	godepgraph -md -p . >> $(DOCOUT)
	deadcode -md >> $(DOCOUT)
	sloc -md >> $(DOCOUT)
	cp README-$(PROG).md README.md
	cat $(DOCOUT) >> README.md
neat:
	go fmt ./...

dead:
	deadcode > problems.dead

index:
	cindex .

clean:
	go clean ./...
	rm -f *~ problems.dead count.out
	rm -f $(DOCOUT)

tar:
	echo $(TARDIR)$(PROG)_$(VERSION)_$(DATE).tar
	tar -ncvf $(TARDIR)$(PROG)_$(VERSION)_$(DATE).tar .

# Coverage test maker
#cover:
#	go test -covermode=count -coverprofile=count.out
#	cover -html=count.out

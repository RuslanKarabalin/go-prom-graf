PROG ?= template
TESTNUM ?= 01

run:
	go build -o $(PROG).x -compiler gc $(PROG)/$(PROG).go
	./$(PROG).x < $(PROG)/in$(TESTNUM).txt
	rm $(PROG).x

gen:
	mkdir -p $(PROG)
	cp template/template.go $(PROG)/$(PROG).go
	cp template/in01.txt $(PROG)/in01.txt

fmt:
	gofmt -s -w .

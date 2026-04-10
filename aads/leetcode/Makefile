PROG ?= template
TESTNUM ?= 01

run:
	go build -o $(PROG).x -compiler gc $(PROG)/$(PROG).go
	./$(PROG).x
	rm $(PROG).x

gen:
	mkdir -p $(PROG)
	cp template/template.go $(PROG)/$(PROG).go

fmt:
	gofmt -s -w .

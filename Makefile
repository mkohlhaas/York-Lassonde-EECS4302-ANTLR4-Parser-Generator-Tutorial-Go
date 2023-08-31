.PHONY: all clean

all: build

genParser: Expr.g4
	rm -rf parser
	antlr4 -Dlanguage=Go -o parser -no-listener -no-visitor $^

build: genParser
	go build

clean:
	rm -rf parser

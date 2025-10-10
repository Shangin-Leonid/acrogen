all: build run

build:
	go build -o acrgen.out acrgen.go

run:
	./acrgen.out src.txt russian_words.txt acrs.txt

clean:
	rm ./acrgen.out

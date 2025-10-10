all: build run

build:
	go build -o acrgen.out acrgen.go

run:
	./acrgen.out data/src.txt data/russian_words.txt acrs.txt

clean:
	rm ./acrgen.out

SRC_FILENAME ?= data/src.txt
DICT_FILENAME ?= data/russian_words.txt
OUTP_FILENAME ?= acrs.txt

INPUT_PARAMS = $(SRC_FILENAME) $(DICT_FILENAME) $(OUTP_FILENAME)

all: build run

build:
	go build -o acrgen.out *.go

run:
	./acrgen.out $(INPUT_PARAMS)

debug:
	dlv debug -- $(INPUT_PARAMS)

clean:
	rm ./acrgen.out

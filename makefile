all: build run

build:
	go build -o acrgen.out acrgen.go

run:
	./acrgen.out

clean:
	rm ./acrgen.out

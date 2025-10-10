package main

import (
	"acrgen/tfp"
	"fmt"
	"os"
)

func main() {
	argsWithoutProgName := os.Args[1:]
	if len(argsWithoutProgName) != 3 {
		fmt.Println("Неверное количество входных аргументов!")
		fmt.Println("Запустите программу заново, указав названия трёх  \".txt\" файлов: входного, с существующими словами-кандидатами и выходного")
		return
	}
	srcFilename := argsWithoutProgName[0]

	_, err := importSrcFromFile(srcFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// #
// Describes one source file entry (line), that represents a variant for acronym letter, its estimation and decoding (description).
// #
type SrcEntry struct {
	letter     rune
	estimation int
	decoding   []rune
}
type Src []SrcEntry

// #
// Parse source data file and import its content.
// #
func importSrcFromFile(srcFilename string) (src Src, err error) {

	var parseSrcFileLine tfp.LineParserFunc = func(line string) error {
		return nil
	}

	_, err = tfp.ParseFileLineByLine(srcFilename, parseSrcFileLine)

	if err != nil {
		return nil, err
	}

	return src, nil
}

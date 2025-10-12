package main

import (
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
	srcFilename, dictFilename := argsWithoutProgName[0], argsWithoutProgName[1]

	src, err := importSrcFromFile(srcFilename)
	if err != nil {
		fmt.Println(err)
		return
	}

	const ExpectedWordsAmount = 1532570 // 1'532'568 = amount of russian words in my collection
	dict, err := importDictionaryFromFile(dictFilename, ExpectedWordsAmount)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = generateAcronyms(src, dict)
	if err != nil {
		fmt.Println(err)
		return
	}
}

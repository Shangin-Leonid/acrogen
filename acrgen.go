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
	// srcFilename, dictFilename, outputFilename := argsWithoutProgName[0], argsWithoutProgName[1], argsWithoutProgName[2]
}

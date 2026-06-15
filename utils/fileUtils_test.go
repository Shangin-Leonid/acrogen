package utils_test

import (
	"fmt"

	"github.com/Shangin-Leonid/acrogen/utils"
)

func ExampleIsTextFileNameValid() {

	fmt.Println(utils.IsTextFileNameValid("filename"))
	fmt.Println(utils.IsTextFileNameValid("filename,txt"))
	fmt.Println(utils.IsTextFileNameValid("filename.ttt"))
	fmt.Println(utils.IsTextFileNameValid(".txt"))
	fmt.Println(utils.IsTextFileNameValid("f.json"))
	fmt.Println(utils.IsTextFileNameValid("file\xffname.txt"))

	fmt.Println(utils.IsTextFileNameValid("filename.txt"))
	fmt.Println(utils.IsTextFileNameValid("file.ru.txt"))

	// Output:
	// false
	// false
	// false
	// false
	// false
	// false
	// true
	// true
}

func ExampleWithoutExt() {

	fmt.Println(utils.WithoutExt("filename"))
	fmt.Println(utils.WithoutExt("filename.txt"))
	fmt.Println(utils.WithoutExt("f.json"))
	fmt.Println(utils.WithoutExt("file.ru.md"))

	// Output:
	// filename
	// filename
	// f
	// file.ru
}

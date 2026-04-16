package utils

import (
	"fmt"
	"strings"
)

func CharacterErr(obj string, idx int){
	fmt.Println(obj)
	fmt.Printf("%s^~~\n",strings.Repeat(" ", idx))
	fmt.Println(string(obj[idx]),"is not a Hexadecimal character")
}

func BinaryErr(idx int, msg string){
	//TODO: remove this when the table is being printed
	fmt.Println(strings.Repeat(" ", idx), "^~~")
	fmt.Println(msg)
}

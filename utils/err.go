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

func BinaryErr(binary string, idx int, msg string){
	//TODO: remove this when the table is being printed
	fmt.Println(binary)
	fmt.Printf("%s^~~\n",strings.Repeat(" ", idx))
	fmt.Println(msg)
}

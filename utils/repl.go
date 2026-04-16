package utils

import (
	"fmt"
	"strings"
)


func Repl(){
  for;;{

		var input string
		fmt.Printf(">")

		_, err := fmt.Scanln(&input)

		if err!= nil{
			fmt.Println("Error in input")
			continue
		}

		objectCodes := strings.Split(input, " ")

		for obj := range objectCodes{
			parseObj(strings.ToLower(objectCodes[obj]))
		}
	}
}

func parseObj(obj string){
  
	if obj  == "sic" {
		SicEnabled = true
		fmt.Println("Switched to SIC")
		return
	}
	if obj =="sic/xe" || obj == "sicxe" {
		fmt.Println("Switched to SIC/XE")
		SicEnabled = false
		return
	}


	// handle errors

	if( len(obj) != 6 && SicEnabled){
		fmt.Println("Err: Sic object format requires 6 Hexadecimal characters")
		return
	}
	if( len(obj) != 6 && len(obj) != 8){
		fmt.Println("Err: Sic/xe object format requires 6 Hexadecimal characters for format 3 or 8 Hexadecimal characters for format 4")
		return
	}

	// check each characters, for their range
	for i:=0; i< len(obj); i++{
    //check if its a number
    if(obj[i] >= '0' && obj[i]<='9'){ continue }
		if(obj[i] >= 'a' && obj[i]<='f'){ continue }
		CharacterErr(obj, i)
		return 
	}

	printBin(obj)

}

func printBin(obj string){
	// convert to binary string and check for err in format
	binary := make([]byte, len(obj)*4)
	runner:=0
  for idx:= range obj{
		bin := HexTable[obj[idx]]

		for s:=0; s<len(bin); s++{
			binary[runner + s] = bin[s]
		}
		runner+=len(bin)

	}

	if SicEnabled{
		labels := []string {"opcode", "x", "address"}
		sizes := []int {8,1,15}
    Table(labels, sizes, string(binary))
    

	} else{

		labels := []string{"opcode", "n", "i", "x", "b", "p", "e", "address"}

		address := 20

		fmt.Println("len: ", len(obj))
		if len(obj) == 6 {
			address = 12
		}

		sizes := []int {6,1,1,1,1,1,1,address}

		Table(labels, sizes, string(binary))

		b := 9
		p := 10
		e := 11

		//NOTE: these values are hardcoded :P
		//TODO: fix this, no hardcoding pls
		offset_e := 44
		offset_p := 32

		//TODO: pb check should be first, print the errors only after ^~ is done printing, a flag to check wether printing ^~ is done, to be used
		if address==12 && binary[e]=='1' {
			BinaryErr(offset_e, "Format 4 is specified but only 12 bits are used by the address field")

		} else if address==20 && binary[e]=='0' {
			BinaryErr(offset_e, "Format 3 is specified but 20 bits are used by the address field")

		}else if binary[p]=='1' && binary[b]=='1' {
			BinaryErr(offset_p, "Cant be both pc relative and base relative")
		}

	}

}

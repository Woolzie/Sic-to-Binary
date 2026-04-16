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
			parseObj(objectCodes[obj])
		}
	}
}

func parseObj(obj string){
  
	tempObj := strings.ToLower(obj)
	if tempObj  == "sic" {
		SicEnabled = true
		fmt.Println("Switched to SIC")
		return
	}
	if tempObj =="sic/xe" || tempObj == "sicxe" {
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
		if(obj[i] >= 'A' && obj[i]<='F'){ continue }
		CharacterErr(obj, i)
		return 
	}

	printBin(obj)

}

func printBin(obj string){
	// convert to binary string and check for err in format
	binary := make([]rune, len(obj)*4)
	runner:=0
  for idx:= range obj{
		bin := HexTable[rune(obj[idx])]

		for s:=0; s<len(bin); s++{
			binary[runner + s] = rune(bin[s])
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


		if len(obj) == 6 {
			address = 12
		}

		sizes := []int {6,1,1,1,1,1,1,address}

		Table(labels, sizes, string(binary))

		e:= 10
		if len(binary)==24 && binary[e]==1 {
			//TODO: this should be printed below the table
			BinaryErr(string(binary), e, "Format 4 is specified but only 12 bits are used by the address field")
		}
		if len(binary)==32 && binary[e]==0 {
			//TODO: this should be printed below the table
			BinaryErr(string(binary), e, "Format 3 is specified but 20 bits are used by the address field")
		}

	}

}

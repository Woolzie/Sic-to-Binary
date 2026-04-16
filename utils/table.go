package utils

import (
	"fmt"
	"strings"
)

//returns an array that contains the position of each bit?

func Table(labels []string, cellSize []int , binary string){
  if(len(labels)!=len(cellSize)){
		fmt.Printf("DEBUG: %s\n","labels and sizes dont match ")
		return
	}

	labelPadding, sizePadding := makePadding(labels, cellSize)

	for i:= range labels{
		pad := strings.Repeat(" ", labelPadding[i])
    fmt.Printf(" %s%s%s", pad, labels[i], pad)
	}
	fmt.Println()

	left := 0
	right:= 0

	//TODO: add space between each half byte (for readability)

	for i:= range cellSize{
		pad := strings.Repeat(" ", sizePadding[i])

		right += cellSize[i]
		fmt.Printf("|%s%s%s", pad, binary[left: right ], pad)
		left = right
	}
	fmt.Println("|")
	fmt.Println()

}


func makePadding(labels []string, cellSize []int) ( []int, []int ){
	labelPadding := make([]int, len(cellSize))
	cellPadding := make([]int, len(cellSize))

	//TODO: remove the spaces variable to make it simpler
	for i:= range labels{
		labelSize := len(labels[i])

		//TODO: dont change this logic here, trust
		// spaces := ( cellSize[i] / 4 )*2 //two spaces for seperation
		spaces :=0

		if labelSize < cellSize[i] {
			labelPadding[i] = ( cellSize[i] + spaces - labelSize )/2 + 2 
			cellPadding[i] = 2
		}else{
			cellPadding[i] = (labelSize - cellSize[i] - spaces )/2 + 2
			labelPadding[i] = 2
		}
	}

	return labelPadding, cellPadding
}

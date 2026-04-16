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

	labelPadding, cellPadding := makePadding(labels, cellSize)

	//NOTE: this is a hacky fix to make the labels algin with the cells
	fmt.Printf(" ")
	if len(labels) > 3{
		fmt.Printf(" ")
	}
	for i:= range labels{
		pad := strings.Repeat(" ", labelPadding[i])
    fmt.Printf("%s%s%s ", pad, labels[i], pad)
	}
	fmt.Println()

	left := 0
	right:= 0
	cellBegin:= 0

	//TODO: add space between each half byte (for readability)

	//INFO: printing within the cell
	for i:= range cellSize{

		pad := strings.Repeat(" ", cellPadding[i])
		fmt.Printf("|%s", pad)

		// prints the half byte value and space repeatedly
    for j:= cellSize[i]/4; j>=0; j--{

			right += 4
      if right - cellBegin > cellSize[i]{
        right = cellSize[i] + cellBegin
			}

			fmt.Printf("%s", binary[left:right])
			if j > 0 {
				fmt.Printf(" ")
			}

			left = right
		}

		fmt.Printf("%s", pad)
		cellBegin+= cellSize[i]

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
		spaces := ( cellSize[i] / 4 ) //two spaces for seperation

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

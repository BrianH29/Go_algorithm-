package solveif

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//Quadrant question
func Quadrant() {
	var x, y int

	fmt.Println("입력:")
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanln(reader, &x)
	checkErr(err)
	fmt.Println("입력2:")
	_, err = fmt.Fscanln(reader, &y)
	checkErr(err)

	if x > 0 && y > 0 {
		fmt.Println(1)
	} else if x < 0 && y > 0 {
		fmt.Println(2)
	} else if x < 0 && y < 0 {
		fmt.Println(3)
	} else if x > 0 && y < 0 {
		fmt.Println(4)
	}

	/*
		if x > 0 {
			if y > 0 {
				fmt.Println(1)
			} else if y < 0 {
				fmt.Println(4)
			}
		} else if x < 0 {
			if y > 0 {
				fmt.Println(2)
			} else if y < 0 {
				fmt.Println(3)
			}
		}
	*/
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

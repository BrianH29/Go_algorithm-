package solvefor

import "fmt"

func diamondStar() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < (3-i)*2-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 	    *
	// 	   ***
	//    *****
	//   *******
	//    *****
	// 	   ***
	// 	    *
}

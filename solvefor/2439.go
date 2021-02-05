package solvefor

import "fmt"

func printStar() {
	var n int

	fmt.Println("input No: ")
	_, err := fmt.Scanf("%d\n", &n)
	checkErr(err)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//     *
	//    **
	//   ***
	//  ****
	// *****
}

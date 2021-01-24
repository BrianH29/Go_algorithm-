package solvefor

import (
	"fmt"
	"log"
)

func increaseStar() {
	var n int

	fmt.Println("input no : ")
	_, err := fmt.Scanf("%d\n", &n)
	checkErr(err)

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

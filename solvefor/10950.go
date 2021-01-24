package solvefor

import (
	"fmt"
	"log"
)

func add() {
	var T, a, b int

	fmt.Println("input limit : ")
	_, err := fmt.Scanf("%d\n", &T)
	checkErr(err)

	for i := 0; i < T; i++ {
		fmt.Println("input first num: ")
		_, err := fmt.Scanf("%d\n", &a)
		checkErr(err)

		fmt.Println("input second num: ")
		_, err = fmt.Scanf("%d\n", &b)
		checkErr(err)

		fmt.Println(a + b)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

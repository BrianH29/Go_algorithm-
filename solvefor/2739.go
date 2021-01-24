package solvefor

import (
	"fmt"
	"log"
)

func mulitplyTable() {
	var no int

	fmt.Println("input number : ")
	_, err := fmt.Scanf("%d", &no)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 9; i++ {
		fmt.Printf("%d * %d = %d\n", no, i, (no * i))
	}
}

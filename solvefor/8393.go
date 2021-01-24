package solvefor

import (
	"fmt"
	"log"
)

func increaseInputNo() {
	var n int

	_, err := fmt.Scanf("%d\n", &n)
	checkErr(err)

	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//CheckErr checking err
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package solveif

import "fmt"

func yoonneun() {
	var year int

	_, err := fmt.Scanf("%d", &year)
	if err != nil {
		fmt.Println("잘못 입력하셨습니다.")
	}

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

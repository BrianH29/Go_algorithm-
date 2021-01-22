package solveif

import "fmt"

func compareAB() {
	var a, b int

	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		fmt.Println("잘못 입력하셨습니다.")
	}

	if a < b {
		fmt.Println("<")
	} else if a > b {
		fmt.Println(">")
	} else if a == b {
		fmt.Println("==")
	}
}

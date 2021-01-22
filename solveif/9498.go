package solveif

import "fmt"

func score() {
	var score int
	var grade string

	_, err := fmt.Scanf("%d", &score)
	if err != nil {
		fmt.Println("잘못 입력하셨습니다.")
	}

	if 90 <= score {
		grade = "A"
	} else if 80 <= score {
		grade = "B"
	} else if 70 <= score {
		grade = "C"
	} else if 60 <= score {
		grade = "D"
	} else {
		grade = "F"
	}
	fmt.Println(grade)
}

package solveif

import (
	"fmt"
	"log"
)

func alarm() {
	var h, m int

	fmt.Println("input hour : ")
	//reader := bufio.NewReader(os.Stdin)
	//_, err := fmt.Fscanln(reader, &h)
	_, err := fmt.Scanf("%d\n", &h)
	checkErr(err)

	fmt.Println("input minute : ")
	_, err = fmt.Scanf("%d", &m)
	checkErr(err)

	//h = 60min
	if m < 45 {
		h--
		m = (60 - (45 - m))

		if h < 0 {
			h = 23
		}
		fmt.Println(h, " ", m)
	} else {
		fmt.Println(h, " ", (60 - m))
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

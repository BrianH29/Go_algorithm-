package solveArray

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func minMax() {
	var n int

	fmt.Println("inputNo : ")
	reader := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanln(reader, &n)
	checkErr(err)

	var inputs = make([]int, n)
	fmt.Printf("input %d random digit :", &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &inputs[i])
	}

	sort.Slice(inputs, func(i, j int) bool {
		return inputs[i] < inputs[j]
	})

	fmt.Printf("%d %d\n", inputs[0], inputs[len(inputs)-1])
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

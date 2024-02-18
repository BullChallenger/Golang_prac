package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	r := rand.Intn(100)
	cnt := 1

	for {
		fmt.Println("숫자를 입력하세요 : ")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요")
		} else {
			if n > r {
				fmt.Println("입력하신 숫자가 더 큽니다")
			} else if n < r {
				fmt.Println("입력하신 숫자가 더 작습니다")
			} else {
				fmt.Println("숫자가 일치합니다. 시도한 횟수 : ", cnt)
				break
			}
			cnt++
		}
	}
}

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}

	return n, err
}

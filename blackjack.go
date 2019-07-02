package main

import (
	"fmt"
	"math/rand"
	"time"
)

func judge(sumD int, sumP int) {
	if sumD > sumP {
		fmt.Printf("You lose.")
	}
	if sumD < sumP {
		fmt.Printf("You win!")
	}
}

func drawCard(hand *[]int, sum *int) {
	rand.Seed(time.Now().UnixNano())

	newCard := rand.Intn(10) + 1
	*hand = append(*hand, newCard)
	*sum += newCard
}

func main() {
	var dealer []int
	var player []int
	sumD := 0
	sumP := 0
	dealer = make([]int, 0, 10)
	player = make([]int, 0, 10)
	for i := 0; i < 2; i++ {
		drawCard(&dealer, &sumD)
		drawCard(&player, &sumP)
	}
	fmt.Println("dealer: ", dealer[0])
	fmt.Println("player: ", player, "合計:", sumP)

	status := ""
	for status != "Stand" {
		fmt.Println("Hit or Stand?")
		fmt.Scanf("%s", &status)
		if status == "Hit" {
			drawCard(&player, &sumP)
			fmt.Println("player: ", player, "合計:", sumP)
			if sumP > 21 {
				fmt.Println("You Lose")
				return
			}

		}
	}
	//ディーラーは17以上になるまで引き続ける
	for sumD <= 16 {
		drawCard(&dealer, &sumD)
		if sumD > 21 {
			fmt.Println("You win")
			return
		}

	}
	fmt.Println("dealer: ", dealer, "合計:", sumD)
	fmt.Println("player: ", player, "合計:", sumP)

	judge(sumD, sumP)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)
func judge(sumD,sumP,dealer,player){
    if sumD >= 22 && sumP >=22 {
        fmt.Printf("This game is draw.\n")
    }
    if sumD >= 22 && sumP <= 21 {
        fmt.Printf("Dealer is burst!\nYou win!\n")
    }
    if sumD <= 21 && sumP >= 22 {
        fmt.Printf("Player is burst!\nYou lose.\n")
    }
    if sumD <= 21 && sumP <= 21 {
        if sumD >sumP {
            fmt.Printf("You lose.")
        }
        if sumD < sumP {
            fmt.Printf("You win!")
        }
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
		drawCard(&player, &sumP)
		fmt.Println("player: ", player, "合計:", sumP)
		if sumP > 21 {
			fmt.Println("You Lose")
			return
		}
	}
	//ディーラーは17以上になるまで引き続ける
	for sumD <= 16 {
		drawCard(&dealer, &sumD)
	}
	fmt.Println("dealer: ", dealer, "合計:", sumD)
	fmt.Println("player: ", player, "合計:", sumP)

}

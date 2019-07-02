package main

import (
	"fmt"
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

func drawCard(hand [10]int)

func main() {
	var dealer []int
	var player []int
	dealer = make([]int, 0, 10)
	player = make([]int, 0, 10)
	sumD := 0
	sumP := 0
	fmt.Println(dealer[1])
	fmt.Println(player)
	fmt.Printf("%T", sumD)
	fmt.Printf("%T", sumP)

}

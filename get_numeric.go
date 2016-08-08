package go_game_1_input

import (
	"fmt"
)

func Get_numeric(userText string) int {
	var i int
	fmt.Print(userText, " ")
	_, err := fmt.Scanf("%d", &i)
	if err == nil {
		return i
	} else {
		return -1
	}
}

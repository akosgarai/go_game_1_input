package go_game_1_input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Get_string(textToPrint string) string {
	Clear_screen()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(textToPrint, "\n")
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}

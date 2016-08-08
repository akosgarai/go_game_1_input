package go_game_1_input

import (
	"os"
	"os/exec"
)

func Clear_screen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

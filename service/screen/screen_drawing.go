package screen_service

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func InitScreen(isActive bool, durationMilis int) {
	//header
	fmt.Printf("%.42s\n", "╒════════════════════════════════════════╕")
	fmt.Printf("%-15s%12s%15s\n", "│---", "Auto clicker", "---│")
	fmt.Printf("%42s\n", "╞════════════════════════════════════════╡")

	//body
	fmt.Printf("│%40s│\n", "")
	fmt.Printf("│Time duration%12sSec:%d Milis:%d00│\n", "", durationMilis/10, durationMilis%10)
	fmt.Printf("│Clicker active%26v│\n", isActive)

	//footer
	fmt.Printf("%42s", "╘════════════════════════════════════════╛")
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

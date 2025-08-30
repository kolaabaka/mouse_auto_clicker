package main

import (
	"context"
	"fmt"
	click_service "mouse_auto_clicker/service"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	isActive := false
	durationMilis := 10

	ClearScreen()
	InitScreen(isActive, durationMilis)

	runeChanel := make(chan rune)
	var ctx context.Context
	var cancelFunc func()

	clicker := click_service.ClickerConfig{Delimetr: durationMilis}
	go click_service.TakePressedValue(runeChanel)

	for {
		switch <-runeChanel {
		case 's':
			ClearScreen()
			ctx = context.Background()
			ctx, cancelFunc = context.WithCancel(ctx)
			isActive = true
			InitScreen(isActive, durationMilis)
			go clicker.ClickingStart(ctx)
		case 'i':
			durationMilis += 1
			ClearScreen()
			InitScreen(isActive, durationMilis)
			clicker.IncreaseTiming()
		case 'r':
			if durationMilis-1 > 0 {
				durationMilis -= 1
			}
			ClearScreen()
			InitScreen(isActive, durationMilis)
			clicker.ReduceTiming()
		case 'p':
			ClearScreen()
			isActive = false
			InitScreen(isActive, durationMilis)
			if cancelFunc != nil {
				cancelFunc()
			}
		case 'x':
			return
		default:
		}
	}
}

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

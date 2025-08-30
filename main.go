package main

import (
	"context"
	click_service "mouse_auto_clicker/service"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	runeChanel := make(chan rune)
	var ctx context.Context
	var cancelFunc func()

	clicker := click_service.ClickerConfig{Delimetr: 1}
	go click_service.TakePressedValue(runeChanel)

	for {
		switch <-runeChanel {
		case 's':
			ctx = context.Background()
			ctx, cancelFunc = context.WithCancel(ctx)
			go clicker.ClickingStart(ctx)
		case 'e':
			clicker.IncreaseTiming()
		case 'p':
			if cancelFunc != nil {
				cancelFunc()
			}
		case 't':
			clicker.TestReuse()
		case 'x':
			return
		default:
		}
	}
}

func ClearScreen() {
	for {
		time.Sleep(time.Millisecond * 500)
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
}

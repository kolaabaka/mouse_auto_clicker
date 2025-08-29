package main

import (
	"context"
	click_service "mouse_auto_clicker/service"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	p := click_service.ClickerConfig{Delimetr: 10}

	contextBackground := context.Background()
	contextBackground, _ = context.WithTimeout(contextBackground, time.Second*3)
	go p.ClickingStart(contextBackground, &wg)
	wg.Wait()
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

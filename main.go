package main

import (
	"context"
	"mouse_auto_clicker/configuration"

	click_service "mouse_auto_clicker/service/click"
	screen_service "mouse_auto_clicker/service/screen"
)

func main() {

	config := configuration.MakeConfig()
	isActive := false
	durationMilis := int(config.Duration)

	screen_service.ClearScreen()
	screen_service.InitScreen(isActive, durationMilis)

	runeChanel := make(chan rune)
	var ctx context.Context
	var cancelFunc func()

	clicker := click_service.ClickerConfig{Delimetr: durationMilis}
	go click_service.TakePressedValue(runeChanel)

	for {
		switch <-runeChanel {
		case 's':
			screen_service.ClearScreen()
			ctx = context.Background()
			ctx, cancelFunc = context.WithCancel(ctx)
			isActive = true
			screen_service.InitScreen(isActive, durationMilis)
			go clicker.ClickingStart(int(config.MousePosY), int(config.MousePosY), int(config.Duration), ctx)
		case 'i':
			durationMilis += 1
			screen_service.ClearScreen()
			screen_service.InitScreen(isActive, durationMilis)
			clicker.IncreaseTiming()
		case 'r':
			if durationMilis-1 > 0 {
				durationMilis -= 1
			}
			screen_service.ClearScreen()
			screen_service.InitScreen(isActive, durationMilis)
			clicker.ReduceTiming()
		case 'p':
			screen_service.ClearScreen()
			isActive = false
			screen_service.InitScreen(isActive, durationMilis)
			if cancelFunc != nil {
				cancelFunc()
			}
		case 'x':
			return
		default:
		}
	}
}

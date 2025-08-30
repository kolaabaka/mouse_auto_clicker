package click_service

import "github.com/eiannone/keyboard"

func TakePressedValue(runeChanel chan rune) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		runeChanel <- char
	}
}

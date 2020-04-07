package input_svc

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

func PrintConfirmationMenu(prompt string) {
	fmt.Println(prompt)
	fmt.Println("Y for yes, or N for no")
}

func PrintPrompt() {
	fmt.Print("> ")
}

func UserConfirms(prompt string) (bool, error) {
	PrintConfirmationMenu(prompt)
	for {
		err := keyboard.Open()
		if err != nil {
			return false, err
		}
		char, key, err := keyboard.GetKey()
		keyboard.Close()
		if err != nil {
			return false, err
		}

		fmt.Printf("%c\n", char)

		if key == keyboard.KeyEsc || char == 'n' || char == 'N' {
			return false, nil
		} else if char == 'y' || char == 'Y' {
			return true, nil
		}
	}
}

package main

import (
	"fmt"
	red "red/src/prog"
)

func main() {
	red.ClearScreen()
	fmt.Print("\033[?1049h")

	fmt.Println("=== BIENVENUE DANS REDQUEST ===")
	fmt.Println()
	fmt.Println("Création de votre personnage...")

	c1 := red.CharacterCreation()

	fmt.Println("\nPersonnage créé avec succès !")
	red.DisplayInfo(c1)

	red.Menu(&c1)
}

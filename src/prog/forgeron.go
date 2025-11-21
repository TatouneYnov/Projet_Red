package red

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Forgeron(c *Character) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║         ⚒️ FORGERON REINHARDT ⚒️           ║")
		fmt.Println("╚══════════════════════════════════════════╝")

		fmt.Println("\n\033[34m🛡️ \"HONNEUR ET GLOIRE, MON AMI! QUE PUIS-JE FORGER POUR VOUS AUJOURD'HUI?\"\033[0m")
		fmt.Println("\033[34m🔨 \"JE SUIS LE BOUCLIER DE MES COMPAGNONS, ET MES CRÉATIONS SONT AUSSI SOLIDES QUE MOI!\"\033[0m")

		fmt.Println("\n╔════════════════════════════════════════════════════════╗")
		fmt.Println("║ [1] 🧢 Casquette de Mario+       🪙  5 pièces           ║")
		fmt.Println("║     Nécessite : 1 Plume de coussin, 1 Cuir de Pumba    ║")
		fmt.Println("║     Bonus : +20❤️                                       ║")
		fmt.Println("║ [2] 👚 Tunique de Link+          🪙  5 pièces           ║")
		fmt.Println("║     Nécessite : 2 Moumouth de daronne, 1 Peau de Troll ║")
		fmt.Println("║     Bonus : +45❤️                                       ║")
		fmt.Println("║ [3] 👞 Bottes de Sonic+          🪙  5 pièces           ║")
		fmt.Println("║     Nécessite : 1 Moumouth de daronne, 1 Cuir de Pumba ║")
		fmt.Println("║     Bonus : +12⚡                                      ║")
		fmt.Println("║ [4] 🚪 Quitter l'atelier                               ║")
		fmt.Println("╚════════════════════════════════════════════════════════╝")
		fmt.Print("Votre choix : ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		var ok bool
		switch choice {
		case "1":
			ok = true
			if c.Argent < 5 {
				fmt.Println("❌ PAS ASSEZ DE PIÈCES, CAMARADE!")
				ok = false
			}
			if c.Inventaire["Plume de coussin"] < 1 {
				fmt.Println("❌ IL VOUS MANQUE UNE PLUME DE COUSSIN!")
				ok = false
			}
			if c.Inventaire["Cuir de Pumba"] < 1 {
				fmt.Println("❌ IL VOUS MANQUE UN CUIR DE PUMBA!")
				ok = false
			}
			if !ok {
				fmt.Print("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
				continue
			}
			if !CanAddItem(c, 1) {
				fmt.Println("❌ INVENTAIRE PLEIN! JE NE PEUX PAS PORTER PLUS!")
				fmt.Print("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
				continue
			}
			c.Argent -= 5
			c.Inventaire["Plume de coussin"]--
			c.Inventaire["Cuir de Pumba"]--
			AddInventory(c, "Casquette de Mario+", 1)
			fmt.Println("🎉 MAMMA MIA! CASQUETTE DE MARIO AMÉLIORÉE FORGÉE AVEC EXCELLENCE!")
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
		case "2":
			ok = true
			if c.Argent < 5 {
				fmt.Println("❌ PAS ASSEZ DE PIÈCES D'OR, CAMARADE!")
				ok = false
			}
			if c.Inventaire["Moumouth de daronne"] < 2 {
				fmt.Println("❌ IL VOUS FAUT 2 MOUMOUTH DE DARONNE!")
				ok = false
			}
			if c.Inventaire["Peau de Troll"] < 1 {
				fmt.Println("❌ IL VOUS MANQUE UNE PEAU DE TROLL!")
				ok = false
			}
			if !ok {
				fmt.Print("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
				continue
			}
			if !CanAddItem(c, 1) {
				fmt.Println("❌ INVENTAIRE PLEIN! PLUS DE PLACE POUR CE CHEF D'ŒUVRE!")
				fmt.Print("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
				continue
			}
			c.Argent -= 5
			c.Inventaire["Moumouth de daronne"] -= 2
			c.Inventaire["Peau de Troll"]--
			AddInventory(c, "Tunique de Link+", 1)
			fmt.Println("🎉 HYAAA! TUNIQUE DE LINK AMÉLIORÉE FORGÉE AVEC COURAGE!")
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
		case "3":
			ok = true
			if c.Argent < 5 {
				fmt.Println("❌ PAS ASSEZ DE PIÈCES D'OR, CAMARADE!")
				ok = false
			}
			if c.Inventaire["Moumouth de daronne"] < 1 {
				fmt.Println("❌ IL VOUS MANQUE UN MOUMOUTH DE DARONNE!")
				ok = false
			}
			if c.Inventaire["Cuir de Pumba"] < 1 {
				fmt.Println("❌ IL VOUS MANQUE UN CUIR DE PUMBA!")
				ok = false
			}
			if !ok {
				fmt.Print("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
				continue
			}
			if !CanAddItem(c, 1) {
				fmt.Println("❌ INVENTAIRE PLEIN! IL FAUT FAIRE DE LA PLACE!")
				fmt.Print("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
				continue
			}
			c.Argent -= 5
			c.Inventaire["Moumouth de daronne"]--
			c.Inventaire["Cuir de Pumba"]--
			AddInventory(c, "Bottes de Sonic+", 1)
			fmt.Println("🎉 GOTTA GO FASTER! BOTTES DE SONIC AMÉLIORÉES FORGÉES AVEC VITESSE!")
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
		case "4":
			fmt.Println("⚒️ \"LA JUSTICE SERA RENDUE! REVENEZ QUAND VOUS VOUDREZ, CAMARADE!\"")
			fmt.Println("🚪 Vous quittez l'atelier de Reinhardt.")
			return
		default:
			fmt.Println("❌ Choix invalide. Veuillez choisir 1, 2, 3 ou 4.")
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
		}
		fmt.Println()
	}
}

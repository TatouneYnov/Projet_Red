package red

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Combat(c *Character, m Monster) {
	reader := bufio.NewReader(os.Stdin)

	c.Mana = c.ManaMax


	tourCombat := 1

	ClearScreen()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║              ⚔️  COMBAT ⚔️                ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Printf("\n🔴 Un %s apparaît ! Préparez-vous au combat !\n", m.Nom)
	fmt.Print("\nAppuyez sur Entrée pour commencer le combat...")

	FlushInputBuffer(reader)

	_, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erreur de lecture : ", err)
	}

	FlushInputBuffer(reader)

	joueurTour := c.Initiative >= m.Initiative
	fmt.Printf("\n📝 Initiative du joueur : %d | Initiative du %s : %d\n", c.Initiative, m.Nom, m.Initiative)
	if joueurTour {
		fmt.Println("🟢 Vous avez l'initiative et commencez le combat !")
		fmt.Println("(Vous commencez car votre initiative est supérieure ou égale à celle du monstre)")
	} else {
		fmt.Printf("🔴 Le %s a l'initiative et commence le combat !\n", m.Nom)
		fmt.Printf("(Le %s commence car son initiative est supérieure à la vôtre)\n", m.Nom)
	}

	for {
		if joueurTour {
			CharacterTurn(c, &m, reader, tourCombat, true)
			joueurTour = false
		} else {
			if m.Nom == "Gobelin" {
				GoblinPattern(tourCombat, m, c)
			} else {
				dommageEnnemi := rand.Intn(3) + m.Attaque
				c.PvActuelle -= dommageEnnemi
				if c.PvActuelle < 0 {
					c.PvActuelle = 0
				}
				fmt.Println()
				fmt.Println("┌──────────────────────────────────────────┐")
				fmt.Printf("│ 👹 %s attaque %s                  │\n", m.Nom, c.Nom)
				fmt.Printf("│ 🗡️  Inflige %d points de dégâts          │\n", dommageEnnemi)
				fmt.Println("└──────────────────────────────────────────┘")
			}
			joueurTour = true
		}
		tourCombat++

		if c.PvActuelle <= 0 {
			fmt.Println("\n💀 Vous avez été vaincu...")
			fmt.Printf("Le %s vous a terrassé.\n", m.Nom)
			fmt.Print("\nAppuyez sur Entrée pour voir votre fiche personnage...")
			reader.ReadString('\n')
			DisplayInfo(*c)
			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
			return
		}

		if m.PvActuelle <= 0 {
			fmt.Println("\n🏆 Victoire !")
			fmt.Printf("Vous avez vaincu le %s !\n", m.Nom)

			expGagnee := m.Exp
			fmt.Printf("✨ Vous gagnez %d points d'expérience !\n", expGagnee)
			c.ExpActuelle += expGagnee
			for c.ExpActuelle >= c.ExpMax {
				c.ExpActuelle -= c.ExpMax
				c.Niveau++
				c.PvMax += 10
				c.PvActuelle = c.PvMax
				c.ExpMax += 10 * c.Niveau
				fmt.Printf("\n⬆️  Niveau augmenté ! Vous êtes maintenant niveau %d.\n", c.Niveau)
				fmt.Printf("❤️  PV max augmenté à %d.\n", c.PvMax)
				fmt.Printf("🔸 Il vous faut %d XP pour le prochain niveau.\n", c.ExpMax)
			}

			argentGagne := m.Argent
			c.Argent += argentGagne
			fmt.Printf("💰 Vous avez gagné %d pièces d'or !\n", argentGagne)

			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
			return
		}
	}
}

func TrainingFight(c *Character) {
	reader := bufio.NewReader(os.Stdin)
	gobelin := InitGoblin()

	tourCombat := 1

	ClearScreen()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║         ⚔️  COMBAT D'ENTRAÎNEMENT ⚔️       ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Printf("\n🔴 Un %s d'entraînement apparaît ! Préparez-vous au combat !\n", gobelin.Nom)
	fmt.Print("\nAppuyez sur Entrée pour commencer le combat...")

	FlushInputBuffer(reader)

	_, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erreur de lecture : ", err)
	}

	FlushInputBuffer(reader)

	fmt.Printf("\n📝 Initiative du joueur : %d | Initiative du %s : %d\n", c.Initiative, gobelin.Nom, gobelin.Initiative)
	var joueurTour bool
	if gobelin.Initiative > c.Initiative {
		fmt.Printf("🔴 Le %s a l'initiative et commence le combat !\n", gobelin.Nom)
		fmt.Printf("(Le %s commence car son initiative est supérieure à la vôtre)\n", gobelin.Nom)
		joueurTour = false
	} else {
		fmt.Println("🟢 Vous avez l'initiative et commencez le combat !")
		fmt.Println("(Vous commencez car votre initiative est supérieure ou égale à celle du monstre)")
		joueurTour = true
	}
	for {
		if joueurTour {
			playerAction := false
			for !playerAction {
				actionResult := CharacterTurn(c, &gobelin, reader, tourCombat, true)
				if actionResult == 1 {
					playerAction = true
				} else if actionResult == 0 {
					playerAction = false
				} else if actionResult == 2 {
					return
				} else if actionResult == -1 {
					fmt.Println("\n❗ Choix invalide. Veuillez réessayer.")
					time.Sleep(500 * time.Millisecond)
					playerAction = false
				}
			}
			joueurTour = false
			tourCombat++
		} else {
			GoblinPattern(tourCombat, gobelin, c)
			joueurTour = true
		}
		for {
			if c.PvActuelle <= 0 {
				fmt.Println("\n💀 Vous avez été vaincu...")
				fmt.Printf("Le %s vous a terrassé.\n", gobelin.Nom)
				fmt.Print("\nAppuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
				return
			}

			if gobelin.PvActuelle <= 0 {
				fmt.Println("\n🏆 Victoire !")
				fmt.Printf("Vous avez vaincu le %s d'entraînement!\n", gobelin.Nom)

				expGagnee := gobelin.Exp
				fmt.Printf("✨ Vous gagnez %d points d'expérience !\n", expGagnee)
				c.ExpActuelle += expGagnee
				for c.ExpActuelle >= c.ExpMax {
					c.ExpActuelle -= c.ExpMax
					c.Niveau++
					c.PvMax += 10
					c.PvActuelle = c.PvMax
					c.ExpMax += 10 * c.Niveau
					fmt.Printf("\n⬆️  Niveau augmenté ! Vous êtes maintenant niveau %d.\n", c.Niveau)
					fmt.Printf("❤️  PV max augmenté à %d.\n", c.PvMax)
					fmt.Printf("🔸 Il vous faut %d XP pour le prochain niveau.\n", c.ExpMax)
				}

				argentGagne := gobelin.Argent
				c.Argent += argentGagne
				fmt.Printf("💰 Vous avez gagné %d pièces d'or !\n", argentGagne)

				fmt.Print("\nAppuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
				return
			}

			ClearScreen()
			fmt.Println("╔══════════════════════════════════════════╗")
			fmt.Printf("║     ⚔️  COMBAT D'ENTRAÎNEMENT - TOUR %d ⚔️ ║\n", tourCombat)
			fmt.Println("╚══════════════════════════════════════════╝")

			playerHealthPercent := float64(c.PvActuelle) / float64(c.PvMax) * 100
			monsterHealthPercent := float64(gobelin.PvActuelle) / float64(gobelin.PvMax) * 100
			playerManaPercent := float64(c.Mana) / float64(c.ManaMax) * 100

			playerHealthBar := ""
			monsterHealthBar := ""
			playerManaBar := ""

			for i := 0; i < 20; i++ {
				if float64(i)*5 < playerHealthPercent {
					playerHealthBar += "█"
				} else {
					playerHealthBar += "░"
				}

				if float64(i)*5 < monsterHealthPercent {
					monsterHealthBar += "█"
				} else {
					monsterHealthBar += "░"
				}

				if float64(i)*5 < playerManaPercent {
					playerManaBar += "█"
				} else {
					playerManaBar += "░"
				}
			}

			fmt.Println("\n┌─────────────────────────────┐  ┌─────────────────────────────┐")
			fmt.Printf("│   🧙 \033[1m%-10s\033[0m [Niv.%2d]     │   │ 👹 \033[1m%-10s\033[0m                  │\n", c.Nom, c.Niveau, gobelin.Nom)
			fmt.Printf("│     \033[32m%-20s\033[0m     │  │ \033[31m%-20s\033[0m     │\n", playerHealthBar, monsterHealthBar)
			fmt.Printf("│     PV: %-3d/%-3d (%.0f%%)        │  │ PV: %-3d/%-3d (%.0f%%)        │\n",
				c.PvActuelle, c.PvMax, playerHealthPercent,
				gobelin.PvActuelle, gobelin.PvMax, monsterHealthPercent)
			fmt.Printf("│     \033[36m%-20s\033[0m     │  │                             │\n", playerManaBar)
			fmt.Printf("│     MP: %-3d/%-3d (%.0f%%)        │  │                             │\n",
				c.Mana, c.ManaMax, playerManaPercent)
			fmt.Println("└─────────────────────────────┘  └─────────────────────────────┘")

			playerAction := false
			for !playerAction {
				ClearScreen()
				fmt.Println("╔══════════════════════════════════════════╗")
				fmt.Printf("║      ⚔️  COMBAT D'ENTRAÎNEMENT - TOUR %d ⚔️ ║\n", tourCombat)
				fmt.Println("╚══════════════════════════════════════════╝")

				fmt.Println("\n┌─────────────────────────────┐  ┌─────────────────────────────┐")
				fmt.Printf("│   🧙 \033[1m%-10s\033[0m [Niv.%2d]    │  │  👹 \033[1m%-10s\033[0m              │\n", c.Nom, c.Niveau, gobelin.Nom)
				fmt.Printf("│    \033[32m%-20s\033[0m     │  │    \033[31m%-20s\033[0m     │\n", playerHealthBar, monsterHealthBar)
				fmt.Printf("│    PV: %-3d/%-3d (%.0f%%)        │  │    PV: %-3d/%-3d (%.0f%%)        │\n",
					c.PvActuelle, c.PvMax, playerHealthPercent,
					gobelin.PvActuelle, gobelin.PvMax, monsterHealthPercent)
				fmt.Printf("│    \033[36m%-20s\033[0m     │  │                             │\n", playerManaBar)
				fmt.Printf("│    MP: %-3d/%-3d (%.0f%%)        │  │                             │\n",
					c.Mana, c.ManaMax, playerManaPercent)
				fmt.Println("└─────────────────────────────┘  └─────────────────────────────┘")

				actionResult := CharacterTurn(c, &gobelin, reader, tourCombat, true)

				if actionResult == 1 {
					playerAction = true
				} else if actionResult == 0 {
					playerAction = false
				} else if actionResult == 2 {
					return
				} else if actionResult == -1 {
					fmt.Println("\n❗ Choix invalide. Veuillez réessayer.")
					time.Sleep(500 * time.Millisecond)
					playerAction = false
				}
			}

			if gobelin.PvActuelle <= 0 {
				continue
			}

			ClearScreen()

			playerHealthPercent = float64(c.PvActuelle) / float64(c.PvMax) * 100
			monsterHealthPercent = float64(gobelin.PvActuelle) / float64(gobelin.PvMax) * 100
			playerManaPercent = float64(c.Mana) / float64(c.ManaMax) * 100

			playerHealthBar = ""
			monsterHealthBar = ""
			playerManaBar = ""

			for i := 0; i < 20; i++ {
				if float64(i)*5 < playerHealthPercent {
					playerHealthBar += "█"
				} else {
					playerHealthBar += "░"
				}

				if float64(i)*5 < monsterHealthPercent {
					monsterHealthBar += "█"
				} else {
					monsterHealthBar += "░"
				}

				if float64(i)*5 < playerManaPercent {
					playerManaBar += "█"
				} else {
					playerManaBar += "░"
				}
			}

			fmt.Println("╔══════════════════════════════════════════╗")
			fmt.Printf("║     ⚔️  COMBAT D'ENTRAÎNEMENT - TOUR %d ⚔️ ║\n", tourCombat)
			fmt.Println("╚══════════════════════════════════════════╝")

			fmt.Println("\n┌─────────────────────────────┐  ┌─────────────────────────────┐")
			fmt.Printf("│   🧙 \033[1m%-10s\033[0m [Niv.%2d]    │  │  👹 \033[1m%-10s\033[0m              │\n", c.Nom, c.Niveau, gobelin.Nom)
			fmt.Printf("│    \033[32m%-20s\033[0m     │  │    \033[31m%-20s\033[0m     │\n", playerHealthBar, monsterHealthBar)
			fmt.Printf("│    PV: %-3d/%-3d (%.0f%%)        │  │    PV: %-3d/%-3d (%.0f%%)        │\n",
				c.PvActuelle, c.PvMax, playerHealthPercent,
				gobelin.PvActuelle, gobelin.PvMax, monsterHealthPercent)
			fmt.Printf("│    \033[36m%-20s\033[0m     │  │                             │\n", playerManaBar)
			fmt.Printf("│    MP: %-3d/%-3d (%.0f%%)        │  │                             │\n",
				c.Mana, c.ManaMax, playerManaPercent)
			fmt.Println("└─────────────────────────────┘  └─────────────────────────────┘")

			fmt.Println("\n╔══════════════════════════════════════════╗")
			fmt.Println("║           👹 TOUR DU MONSTRE 👹          ║")
			fmt.Println("╚══════════════════════════════════════════╝")

			GoblinPattern(tourCombat, gobelin, c)

			tourCombat++

			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			FlushInputBuffer(reader)
			reader.ReadString('\n')
		}
	}
}

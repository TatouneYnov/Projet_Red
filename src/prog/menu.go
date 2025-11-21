package red

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Menu(c *Character) {
	reader := bufio.NewReader(os.Stdin)

	for {
		ClearScreen()

		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║            🎮 MENU PRINCIPAL 🎮          ║")
		fmt.Println("╠══════════════════════════════════════════╣")

		nomText := fmt.Sprintf("  Bienvenue, %s  ", c.Nom)
		var nomSpaces string
		if len(nomText) < 40 {
			nomSpaces = strings.Repeat(" ", 40-len(nomText))
		} else {
			nomSpaces = "   "
		}
		fmt.Printf("║%s%s║\n", nomText, nomSpaces)

		classeText := fmt.Sprintf("  Classe:    %s      ", c.Classe)
		var classeSpaces string
		if len(classeText) < 40 {
			classeSpaces = strings.Repeat(" ", 40-len(classeText))
		} else {
			classeSpaces = "   "
		}
		fmt.Printf("║%s%s║\n", classeText, classeSpaces)

		pvText := fmt.Sprintf("  PV: %d/%d", c.PvActuelle, c.PvMax)
		var pvSpaces string
		if len(pvText) < 40 {
			pvSpaces = strings.Repeat(" ", 40-len(pvText))
		} else {
			pvSpaces = "   "
		}
		fmt.Printf("║%s%s║\n", pvText, pvSpaces)

		argentText := fmt.Sprintf("  💰 Argent: %d pièces               ", c.Argent)
		var argentSpaces string
		if len(argentText) < 40 {
			argentSpaces = strings.Repeat(" ", 40-len(argentText))
		} else {
			argentSpaces = "    "
		}
		fmt.Printf("║%s%s║\n", argentText, argentSpaces)

		var questText string
		if c.MainQuest.Completed && c.MainQuest.FinalBossDefeated {
			questText = "  👑 Quête principale: ACCOMPLIE + BOSS FINAL VAINCU"
		} else if c.MainQuest.Completed {
			questText = "  🏆 Quête principale: TERMINÉE (Boss Final disponible)"
		} else {
			questText = fmt.Sprintf("  🏅 Quête principale: %d/10           ", c.MainQuest.Progression)
		}
		var questSpaces string
		if len(questText) < 40 {
			questSpaces = strings.Repeat(" ", 40-len(questText))
		} else {
			questSpaces = "    "
		}
		fmt.Printf("║%s%s║\n", questText, questSpaces)

		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║                                          ║")
		fmt.Println("║  \033[36m[1]\033[0m 📊 Afficher les informations        ║")
		fmt.Println("║  \033[36m[2]\033[0m 📈 Statistiques et compétences      ║")
		fmt.Println("║  \033[36m[3]\033[0m 🎒 Accéder à l'inventaire           ║")
		fmt.Println("║  \033[33m[4]\033[0m 🏪 Visiter le marchand              ║")
		fmt.Println("║  \033[33m[5]\033[0m ⚒️  Visiter le forgeron              ║")
		fmt.Println("║  \033[35m[6]\033[0m 🎯 Combat d'entraînement            ║")
		fmt.Println("║  \033[31m[7]\033[0m 🏅 Quête principale                 ║")
		fmt.Println("║  \033[31m[8]\033[0m 🚪 Quitter RedQuest               ║")
		fmt.Println("║                                          ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		fmt.Print("\n🔹 Votre choix: ")

		FlushInputBuffer(reader)
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			ClearScreen()
			DisplayInfo(*c)
			fmt.Print("\n📌 Appuyez sur Entrée pour retourner au menu...")
			FlushInputBuffer(reader)
			reader.ReadString('\n')
		case "2":
			clearScreen()
			StatsMenu(c)
		case "3":
			clearScreen()
			AccessInventory(c)
			waitForEnter()
		case "4":
			clearScreen()
			Marchand(c)
			waitForEnter()
		case "5":
			clearScreen()
			Forgeron(c)
			waitForEnter()
		case "6":
			clearScreen()
			fmt.Println("\n🎯 Vous vous dirigez vers l'arène d'entraînement...")
			TrainingFight(c)
			if c.PvActuelle <= 0 {
				c.PvActuelle = 1
				fmt.Println("\n🧙 Un guérisseur vous a sauvé in extremis !")
				fmt.Println("   Vous êtes encore faible, mais vous survivrez.")
			}
			waitForEnter()
		case "7":
			clearScreen()
			questSubMenu(c)
		case "8":
			clearScreen()
			fmt.Println("\n🌟 ════════════════════════════════════ 🌟")
			fmt.Println("   À bientôt, brave aventurier ! 🗡️✨")
			fmt.Println("🌟 ════════════════════════════════════ 🌟")
			fmt.Print("\033[?1049l")
			return
		default:
			fmt.Println("\n❌ Choix invalide. Veuillez choisir entre 1 et 8.")
			fmt.Println("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
		}
	}
}

func questSubMenu(c *Character) {
	reader := bufio.NewReader(os.Stdin)

	for {
		ClearScreen()

		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║           🏅 QUÊTE PRINCIPALE 🏅         ║")
		fmt.Println("╠══════════════════════════════════════════╣")

		var questStatus string
		if c.MainQuest.Completed && c.MainQuest.FinalBossDefeated {
			questStatus = "  👑 Statut: \033[33mTOUT COMPLÉTÉ\033[0m (Boss Final Vaincu)"
		} else if c.MainQuest.Completed {
			questStatus = "  🏆 Statut: \033[32mTERMINÉE\033[0m (\033[31mBoss Final disponible\033[0m)"
		} else {
			questStatus = fmt.Sprintf("  📊 Progression: %d/10 monstres vaincus", c.MainQuest.Progression)
		}

		statusLen := len(questStatus) - 25
		var questSpaces string
		if statusLen < 40 {
			questSpaces = strings.Repeat(" ", 40-statusLen)
		} else {
			questSpaces = " "
		}

		fmt.Printf("║%s%s║\n", questStatus, questSpaces)

		if !c.MainQuest.Completed {
			nextMonster := InitQuestMonster(c.MainQuest.Progression + 1)
			monsterText := fmt.Sprintf("  🔸 Prochain: %s", nextMonster.Nom)

			monsterLen := len(monsterText)
			var monsterSpaces string
			if monsterLen < 40 {
				monsterSpaces = strings.Repeat(" ", 40-monsterLen)
			} else {
				monsterSpaces = " "
			}

			fmt.Printf("║%s%s║\n", monsterText, monsterSpaces)
		} else if c.MainQuest.Completed && !c.MainQuest.FinalBossDefeated {
			bossText := "  🔥 Boss Final: Print Alphabet"
			bossSpaces := strings.Repeat(" ", 40-len(bossText))
			fmt.Printf("║%s%s║\n", bossText, bossSpaces)
		}

		fmt.Println("╠══════════════════════════════════════════╣")
		fmt.Println("║                                          ║")
		fmt.Println("║  \033[36m[1]\033[0m 📜 Voir détails de la quête         ║")

		if c.MainQuest.Completed && c.MainQuest.FinalBossDefeated {
			fmt.Println("║  \033[33m[2]\033[0m 🏆 Revivre les combats de légende    ║")
		} else if c.MainQuest.Completed {
			fmt.Println("║  \033[31m[2]\033[0m 🔥 AFFRONTER LE BOSS FINAL        ║")
		} else {
			fmt.Println("║  \033[35m[2]\033[0m ⚔️  Combattre le prochain monstre    ║")
		}

		fmt.Println("║  \033[36m[3]\033[0m 🔙 Retourner au menu principal      ║")
		fmt.Println("║                                          ║")
		fmt.Println("╚══════════════════════════════════════════╝")
		fmt.Print("\n🔹 Votre choix: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			clearScreen()
			DisplayQuestStatus(*c)
			waitForEnter()
		case "2":
			clearScreen()
			if c.MainQuest.Completed {
				fmt.Println("\n🔥 Vous vous préparez à affronter Print Alphabet, le boss final...")
			} else {
				fmt.Println("\n⚔️ Vous partez affronter le prochain monstre de la quête...")
			}
			MainQuestFight(c)
			waitForEnter()
		case "3":
			return
		default:
			fmt.Println("\n❌ Choix invalide. Veuillez choisir entre 1 et 3.")
			waitForEnter()
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func clearScreen() {
	ClearScreen()
}

func waitForEnter() {
	fmt.Print("\n📌 Appuyez sur Entrée pour retourner au menu...")
	reader := bufio.NewReader(os.Stdin)
	FlushInputBuffer(reader)
	reader.ReadString('\n')
}

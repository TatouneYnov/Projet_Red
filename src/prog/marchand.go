package red

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Marchand(c *Character) {
	reader := bufio.NewReader(os.Stdin)

	for {
		ClearScreen()
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║          🦝 TOM NOOK'S BOUTIQUE 🦝       ║")
		fmt.Println("╚══════════════════════════════════════════╝")

		fmt.Println("\n🦝 \"Oui, oui! Bienvenue dans mon humble boutique, hm?\"")
		fmt.Println("    \"Les prix sont fixes, mais la qualité est garantie, oui, oui!\"")

		fmt.Println("\n┌──────────────────────────────────────────┐")
		fmt.Println("│               📦 BOUTIQUE 📦             │")
		fmt.Println("├══════════════════════════════════════════┤")
		fmt.Println("│            ⚔️  ARMES ⚔️                    │")
		fmt.Println("│  \033[91m[1]\033[0m  ⚔️  Épée Rouillée     (+2💥)   8 🪙  │")
		fmt.Println("│  \033[91m[2]\033[0m  ⚔️  Épée en Fer      (+4💥)   15 🪙  │")
		fmt.Println("│  \033[93m[3]\033[0m  🪓 Hache de Bataille (+6💥)  25 🪙  │")
		fmt.Println("│  \033[93m[4]\033[0m  ⚔️  Épée en Acier    (+8💥)   35 🪙  │")
		fmt.Println("│  \033[95m[5]\033[0m  🏹 Arc Elfique       (+10💥) 50 🪙  │")
		fmt.Println("│  \033[95m[6]\033[0m  🔮 Bâton Magique     (+12💥) 60 🪙  │")
		fmt.Println("│  \033[96m[7]\033[0m  ⚔️  Épée Légendaire  (+15💥)  80 🪙  │")
		fmt.Println("├──────────────────────────────────────────┤")
		fmt.Println("│           👕 ÉQUIPEMENTS 👕              │")
		fmt.Println("│  \033[34m[8]\033[0m  🎖️ Casque de Chevalier (+15❤️) 15 🪙  │")
		fmt.Println("│  \033[34m[9]\033[0m  �️ Armure Renforcée   (+35❤️)  30 🪙  │")
		fmt.Println("│  \033[34m[10]\033[0m � Bottes de Forgeron  (+8⚡) 20 🪙  │")
		fmt.Println("├──────────────────────────────────────────┤")
		fmt.Println("│          🧪 UTILITAIRES 🧪               │")
		fmt.Println("│  \033[32m[11]\033[0m 🧪 Potion de vie              3 🪙  │")
		fmt.Println("│  \033[36m[12]\033[0m 🧪 Potion Bleue de Mana       5 🪙  │")
		fmt.Println("│  \033[35m[13]\033[0m ☠️  Toxine de Teemo            6 🪙  │")
		fmt.Println("│  \033[33m[14]\033[0m 📘 Livre: Hadouken           25 🪙  │")
		fmt.Println("│  \033[33m[15]\033[0m � Livre: Blizzard            28 🪙  │")
		fmt.Println("│  \033[33m[16]\033[0m � Livre: Éclair de Zeus      30 🪙  │")
		fmt.Println("│  \033[33m[17]\033[0m 📙 Livre: FUS RO DAH         35 🪙  │")
		fmt.Println("│  \033[33m[18]\033[0m 📗 Livre: Tornade            32 🪙  │")
		fmt.Println("│  \033[33m[19]\033[0m 🎒 Sac d'inventaire          30 🪙  │")
		fmt.Println("├──────────────────────────────────────────┤")
		fmt.Println("│           🛠️  CRAFT 🛠️                     │")
		fmt.Println("│  \033[36m[20]\033[0m 🐺 Moumouth de daronne       25 🪙  │")
		fmt.Println("│  \033[36m[21]\033[0m 👹 Peau de Troll             35 🪙  │")
		fmt.Println("│  \033[36m[22]\033[0m 🐗 Cuir de Pumba             20 🪙  │")
		fmt.Println("│  \033[36m[23]\033[0m 🪶 Plume de coussin           15 🪙  │")
		fmt.Println("├──────────────────────────────────────────┤")
		fmt.Println("│  \033[31m[24]\033[0m 🚪 Quitter la boutique             │")
		fmt.Println("└──────────────────────────────────────────┘")

		fmt.Printf("\n💰 Votre bourse: %d pièces\n", c.Argent)
		fmt.Print("💰 Que souhaitez-vous acheter ? ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			if c.Argent >= 5 {
				fmt.Println("\n🦝 \"Ah! Une épée rouillée, parfait pour débuter!\"")
				fmt.Println("    \"Elle vous coûtera 5 pièces d'or, oui oui.\"")

				c.Argent -= 5
				AddInventory(c, "Épée Rouillée", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cette épée coûte 5 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "2":
			if c.Argent >= 12 {
				fmt.Println("\n🦝 \"Excellente épée en fer! Bien plus solide!\"")
				fmt.Println("    \"Elle vous coûtera 12 pièces d'or.\"")

				c.Argent -= 12
				AddInventory(c, "Épée en Fer", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cette épée coûte 12 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "3":
			if c.Argent >= 18 {
				fmt.Println("\n🦝 \"Une hache redoutable! Pour les vrais guerriers!\"")
				fmt.Println("    \"Elle vous coûtera 18 pièces d'or.\"")

				c.Argent -= 18
				AddInventory(c, "Hache de Bataille", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cette hache coûte 18 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "4":
			if c.Argent >= 25 {
				fmt.Println("\n🦝 \"Magnifique épée en acier! Qualité supérieure!\"")
				fmt.Println("    \"Elle vous coûtera 25 pièces d'or.\"")

				c.Argent -= 25
				AddInventory(c, "Épée en Acier", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cette épée coûte 25 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "5":
			if c.Argent >= 35 {
				fmt.Println("\n🦝 \"Un arc elfique légendaire! Très rare!\"")
				fmt.Println("    \"Il vous coûtera 35 pièces d'or.\"")

				c.Argent -= 35
				AddInventory(c, "Arc Elfique", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cet arc coûte 35 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "6":
			if c.Argent >= 40 {
				fmt.Println("\n🦝 \"Un bâton magique puissant! Pour les mages!\"")
				fmt.Println("    \"Il vous coûtera 40 pièces d'or.\"")

				c.Argent -= 40
				AddInventory(c, "Bâton Magique", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Ce bâton coûte 40 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "7":
			if c.Argent >= 50 {
				fmt.Println("\n🦝 \"L'épée légendaire! Le summum de la perfection!\"")
				fmt.Println("    \"Elle vous coûtera 50 pièces d'or.\"")

				c.Argent -= 50
				AddInventory(c, "Épée Légendaire", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cette épée coûte 50 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "8":
			if c.Argent >= 15 {
				fmt.Println("\n🦝 \"Un casque de chevalier forgé avec honneur!\"")
				fmt.Println("    \"Il vous coûtera 15 pièces d'or.\"")

				c.Argent -= 15
				AddInventory(c, "Casque de Chevalier", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Ce casque coûte 15 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "9":
			if c.Argent >= 30 {
				fmt.Println("\n🦝 \"Une armure renforcée de qualité supérieure!\"")
				fmt.Println("    \"Elle vous coûtera 30 pièces d'or.\"")

				c.Argent -= 30
				AddInventory(c, "Armure Renforcée", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cette armure coûte 30 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "10":
			if c.Argent >= 20 {
				fmt.Println("\n🦝 \"Des bottes forgées avec expertise artisanale!\"")
				fmt.Println("    \"Elles vous coûteront 20 pièces d'or.\"")

				c.Argent -= 20
				AddInventory(c, "Bottes de Forgeron", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Ces bottes coûtent 20 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "11":
			if c.Argent >= 3 {
				fmt.Println("\n🦝 \"Ah! Une potion de vie, un classique, hm?\"")
				fmt.Println("    \"Elle vous coûtera 3 pièces d'or, oui oui.\"")

				c.Argent -= 3
				AddInventory(c, "Potion de vie", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cette potion coûte 3 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "12":
			if c.Argent >= 5 {
				fmt.Println("\n🦝 \"Ah! Une potion bleue de mana, excellente pour les mages, hm?\"")
				fmt.Println("    \"Elle restaure votre énergie magique!\"")
				fmt.Println("    \"Elle vous coûtera 5 pièces d'or.\"")

				c.Argent -= 5
				AddInventory(c, "Potion Bleue de Mana", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cette potion coûte 5 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "13":
			if c.Argent >= 6 {
				fmt.Println("\n🦝 \"Oh! La Toxine de Teemo, très populaire, hm?\"")
				fmt.Println("    \"Faites attention avec ça, c'est très dangereux, oui oui!\"")
				fmt.Println("    \"Elle vous coûtera 6 pièces d'or.\"")

				c.Argent -= 6
				AddInventory(c, "Toxine de Teemo", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cette toxine coûte 6 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "14":
			if c.Argent >= 25 {
				fmt.Println("\n🦝 \"Le livre des arts martiaux d'Hadouken! Très puissant!\"")
				fmt.Println("    \"Il vous coûtera 25 pièces d'or.\"")

				c.Argent -= 25
				AddInventory(c, "Livre de Sort : Hadouken", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Ce livre coûte 25 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "16":
			if c.Argent >= 30 {
				fmt.Println("\n🦝 \"Le livre divin de Zeus! Pouvoir électrisant!\"")
				fmt.Println("    \"Il vous coûtera 30 pièces d'or.\"")

				c.Argent -= 30
				AddInventory(c, "Livre de Sort : Éclair de Zeus", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Ce livre coûte 30 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "15":
			if c.Argent >= 28 {
				fmt.Println("\n🦝 \"Le livre de glace Blizzard! Brrr, très froid!\"")
				fmt.Println("    \"Il vous coûtera 28 pièces d'or.\"")

				c.Argent -= 28
				AddInventory(c, "Livre de Sort : Blizzard", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Ce livre coûte 28 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "17":
			if c.Argent >= 35 {
				fmt.Println("\n🦝 \"Le légendaire cri draconique! FUS RO DAH!\"")
				fmt.Println("    \"Il vous coûtera 35 pièces d'or.\"")

				c.Argent -= 35
				AddInventory(c, "Livre de Sort : FUS RO DAH", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Ce livre coûte 35 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "18":
			if c.Argent >= 32 {
				fmt.Println("\n🦝 \"Le pouvoir des vents de Hyrule! Magnifique!\"")
				fmt.Println("    \"Il vous coûtera 32 pièces d'or.\"")

				c.Argent -= 32
				AddInventory(c, "Livre de Sort : Tornade", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Ce livre coûte 32 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "19":
			if c.Argent >= 30 {
				fmt.Println("\n🦝 \"Ah! Un sac pour agrandir votre inventaire, très pratique!\"")
				fmt.Println("    \"Il vous coûtera 30 pièces d'or.\"")

				if c.UpgradesUsed >= 3 {
					fmt.Println("\n🦝 \"Oh, mais je vois que vous avez déjà atteint la limite d'amélioration, hm?\"")
					fmt.Println("    \"Votre sac ne peut plus être agrandi davantage !\"")
				} else {
					c.Argent -= 30
					UpgradeInventorySlot(c)
					fmt.Println("\n🎉 Transaction réussie !")
					fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
					fmt.Printf("🎒 Votre inventaire peut maintenant contenir %d objets.\n", c.InventorySize)
					fmt.Printf("📊 Améliorations utilisées : %d/3\n", c.UpgradesUsed)
				}
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cette amélioration coûte 30 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "20":
			if c.Argent >= 25 {
				fmt.Println("\n🦝 \"Moumouth de daronne, extrêmement rare!\"")
				fmt.Println("    \"Ce matériau précieux vous coûtera 25 pièces d'or.\"")

				c.Argent -= 25
				AddInventory(c, "Moumouth de daronne", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cet objet coûte 25 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "21":
			if c.Argent >= 35 {
				fmt.Println("\n🦝 \"Peau de Troll, matériau légendaire!\"")
				fmt.Println("    \"Cette peau exceptionnelle vous coûtera 35 pièces d'or.\"")

				c.Argent -= 35
				AddInventory(c, "Peau de Troll", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cette peau coûte 35 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "22":
			if c.Argent >= 20 {
				fmt.Println("\n🦝 \"Cuir de Pumba, matériau de luxe!\"")
				fmt.Println("    \"Ce cuir premium vous coûtera 20 pièces d'or.\"")

				c.Argent -= 20
				AddInventory(c, "Cuir de Pumba", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Ce cuir coûte 20 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "23":
			if c.Argent >= 15 {
				fmt.Println("\n🦝 \"Plume de coussin, matériau enchantée!\"")
				fmt.Println("    \"Cette plume magique vous coûtera 15 pièces d'or.\"")

				c.Argent -= 15
				AddInventory(c, "Plume de coussin", 1)

				fmt.Println("\n🎉 Transaction réussie !")
				fmt.Printf("💰 Il vous reste %d pièces.\n", c.Argent)
			} else {
				fmt.Println("\n🦝 \"Pas assez de pièces!\"")
				fmt.Println("    \"Cette plume coûte 15 pièces d'or.\"")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "24":
			fmt.Println("\n🦝 \"Au revoir et revenez bientôt, hm?\"")
			fmt.Println("    \"N'hésitez pas à revenir pour de nouveaux objets!\"")
			return

		default:
			fmt.Println("\n❌ Choix invalide. Veuillez choisir entre 1 et 21.")
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
		}

		fmt.Println()
	}
}

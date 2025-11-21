package red

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AccessInventory(c *Character) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("╔══════════════════════════════════════════╗")
		fmt.Println("║              🎒 INVENTAIRE 🎒            ║")
		fmt.Println("╚══════════════════════════════════════════╝")

		if len(c.Inventaire) == 0 {
			fmt.Println("\n📦 Votre inventaire est vide.")
			fmt.Println("   Explorez le monde pour trouver des objets !")
			fmt.Print("\nAppuyez sur Entrée pour retourner au menu...")
			reader.ReadString('\n')
			return
		}

		fmt.Println("\n🎁 Contenu de votre inventaire :")
		fmt.Println("┌──────────────────────────────────────────┐")
		for item, quantity := range c.Inventaire {
			fmt.Printf("│ ✨ %-25s x%d          │\n", item, quantity)
		}
		fmt.Println("└──────────────────────────────────────────┘")

		fmt.Println("\n╔══════════════════════════════════════════╗")
		fmt.Println("║               🎮 ACTIONS 🎮              ║")
		fmt.Println("╠══════════════════════════════════════════╣")

		potionQuantity := c.Inventaire["Potion de vie"]
		manaQuantity := c.Inventaire["Potion Bleue de Mana"]
		poisonQuantity := c.Inventaire["Toxine de Teemo"]

		bookQty := c.Inventaire["Livre de Sort : Hadouken"] +
			c.Inventaire["Livre de Sort : Éclair de Zeus"] +
			c.Inventaire["Livre de Sort : Blizzard"] +
			c.Inventaire["Livre de Sort : FUS RO DAH"] +
			c.Inventaire["Livre de Sort : Tornade"]

		chapeauQty := c.Inventaire["Casquette de Mario"]
		tuniqueQty := c.Inventaire["Tunique de Link"]
		bottesQty := c.Inventaire["Bottes de Sonic"]

		epeRouilleeQty := c.Inventaire["Épée Rouillée"]
		epeFerQty := c.Inventaire["Épée en Fer"]
		hacheBatailleQty := c.Inventaire["Hache de Bataille"]
		epeAcierQty := c.Inventaire["Épée en Acier"]
		arcElfiqueQty := c.Inventaire["Arc Elfique"]
		batonMagiqueQty := c.Inventaire["Bâton Magique"]
		epeLegendaireQty := c.Inventaire["Épée Légendaire"]

		hasEquipment := chapeauQty > 0 || tuniqueQty > 0 || bottesQty > 0 ||
			epeRouilleeQty > 0 || epeFerQty > 0 || hacheBatailleQty > 0 ||
			epeAcierQty > 0 || arcElfiqueQty > 0 || batonMagiqueQty > 0 || epeLegendaireQty > 0
		if potionQuantity > 0 {
			fmt.Printf("║  \033[32m[1]\033[0m 🧪 Utiliser une Potion de vie (%d)   ║\n", potionQuantity)
		} else {
			fmt.Println("║  \033[90m[1]\033[0m 🧪 Pas de potion de vie             ║")
		}
		if manaQuantity > 0 {
			fmt.Printf("║  \033[36m[2]\033[0m 🧪 Utiliser une Potion de mana (%d)║\n", manaQuantity)
		} else {
			fmt.Println("║  \033[90m[2]\033[0m 🧪 Pas de potion de mana            ║")
		}
		if poisonQuantity > 0 {
			fmt.Printf("║  \033[35m[3]\033[0m ☠️  Utiliser une Potion de poison (%d)║\n", poisonQuantity)
		} else {
			fmt.Println("║  \033[90m[3]\033[0m ☠️  Pas de potion de poison          ║")
		}

		if bookQty > 0 {
			fmt.Printf("║  \033[33m[4]\033[0m 📘 Lire le Livre de Sort (%d)         ║\n", bookQty)
		} else {
			fmt.Println("║  \033[90m[4]\033[0m 📘 Aucun livre disponible           ║")
		}

		if hasEquipment {
			fmt.Println("║  \033[34m[5]\033[0m ⚔️  Gérer l'équipement               ║")
		} else {
			fmt.Println("║  \033[90m[5]\033[0m ⚔️  Pas d'équipement disponible      ║")
		}

		fmt.Println("║  \033[31m[6]\033[0m 🚪 Retourner au menu principal      ║")
		fmt.Println("╚══════════════════════════════════════════╝")

		fmt.Print("\n🔹 Que voulez-vous faire ? ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			if potionQuantity > 0 {
				fmt.Printf("\n💊 Vous vous apprêtez à utiliser une Potion de vie...")
				fmt.Printf("\n🩸 État actuel : %d/%d PV\n", c.PvActuelle, c.PvMax)

				if c.PvActuelle >= c.PvMax {
					fmt.Println("\n❌ Vous êtes déjà en pleine santé !")
					fmt.Print("Appuyez sur Entrée pour continuer...")
					reader.ReadString('\n')
				} else {
					TakePot(c)
					fmt.Print("\nAppuyez sur Entrée pour continuer...")
					reader.ReadString('\n')
				}
			} else {
				fmt.Println("\n❌ Vous n'avez pas de Potion de vie à utiliser.")
				fmt.Print("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
			}

		case "2":
			if manaQuantity > 0 {
				fmt.Printf("\n🔮 Vous vous apprêtez à utiliser une Potion de mana...")
				fmt.Printf("\n🔵 État actuel : %d/%d MP\n", c.Mana, c.ManaMax)

				if c.Mana >= c.ManaMax {
					fmt.Println("\n❌ Votre mana est déjà au maximum !")
					fmt.Print("Appuyez sur Entrée pour continuer...")
					reader.ReadString('\n')
				} else {
					manaRestored := 30
					if c.Mana+manaRestored > c.ManaMax {
						manaRestored = c.ManaMax - c.Mana
					}
					c.Mana += manaRestored

					c.Inventaire["Potion Bleue de Mana"]--
					if c.Inventaire["Potion Bleue de Mana"] == 0 {
						delete(c.Inventaire, "Potion Bleue de Mana")
					}

					fmt.Printf("\n✅ Vous avez restauré %d points de mana !\n", manaRestored)
					fmt.Printf("🔵 Mana actuelle : %d/%d MP\n", c.Mana, c.ManaMax)
					fmt.Print("Appuyez sur Entrée pour continuer...")
					reader.ReadString('\n')
				}
			} else {
				fmt.Println("\n❌ Vous n'avez pas de Potion de mana à utiliser.")
				fmt.Print("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
			}

		case "3":
			if poisonQuantity > 0 {
				fmt.Printf("\n☠️  Vous vous apprêtez à utiliser une Potion de poison...")
				fmt.Printf("\n🩸 État actuel : %d/%d PV\n", c.PvActuelle, c.PvMax)
				fmt.Println("\n⚠️  ATTENTION : Cette potion va vous infliger 30 dégâts au total !")
				fmt.Print("🤔 Êtes-vous sûr de vouloir continuer ? (o/n) : ")

				confirmInput, _ := reader.ReadString('\n')
				confirm := strings.TrimSpace(strings.ToLower(confirmInput))

				if confirm == "o" || confirm == "oui" {
					PoisonPot(c)
				} else {
					fmt.Println("\n🛡️ Sage décision ! Vous rangez la potion.")
				}
				fmt.Print("\nAppuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
			} else {
				fmt.Println("\n❌ Vous n'avez pas de Potion de poison à utiliser.")
				fmt.Print("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
			}

		case "4":
			if bookQty > 0 {
				fmt.Println("\n╔══════════════════════════════════════════╗")
				fmt.Println("║         📚 CHOISIR UN LIVRE 📚           ║")
				fmt.Println("╚══════════════════════════════════════════╝")
				fmt.Println("Quels livres souhaitez-vous utiliser ?")

				optionIndex := 1
				bookOptions := make(map[int]string)

				if c.Inventaire["Livre de Sort : Hadouken"] > 0 {
					fmt.Printf("[%d] 🔥 Livre : Hadouken (%d disponible)\n", optionIndex, c.Inventaire["Livre de Sort : Hadouken"])
					bookOptions[optionIndex] = "Livre de Sort : Hadouken"
					optionIndex++
				}
				if c.Inventaire["Livre de Sort : Éclair de Zeus"] > 0 {
					fmt.Printf("[%d] ⚡ Livre : Éclair de Zeus (%d disponible)\n", optionIndex, c.Inventaire["Livre de Sort : Éclair de Zeus"])
					bookOptions[optionIndex] = "Livre de Sort : Éclair de Zeus"
					optionIndex++
				}
				if c.Inventaire["Livre de Sort : Blizzard"] > 0 {
					fmt.Printf("[%d] 🧊 Livre : Blizzard (%d disponible)\n", optionIndex, c.Inventaire["Livre de Sort : Blizzard"])
					bookOptions[optionIndex] = "Livre de Sort : Blizzard"
					optionIndex++
				}
				if c.Inventaire["Livre de Sort : FUS RO DAH"] > 0 {
					fmt.Printf("[%d] 💫 Livre : FUS RO DAH (%d disponible)\n", optionIndex, c.Inventaire["Livre de Sort : FUS RO DAH"])
					bookOptions[optionIndex] = "Livre de Sort : FUS RO DAH"
					optionIndex++
				}
				if c.Inventaire["Livre de Sort : Tornade"] > 0 {
					fmt.Printf("[%d] 🌪️ Livre : Tornade (%d disponible)\n", optionIndex, c.Inventaire["Livre de Sort : Tornade"])
					bookOptions[optionIndex] = "Livre de Sort : Tornade"
					optionIndex++
				}

				fmt.Printf("[%d] 🚪 Annuler\n", optionIndex)

				fmt.Print("\nVotre choix : ")
				bookChoice, _ := reader.ReadString('\n')
				bookChoice = strings.TrimSpace(bookChoice)

				var bookIndex int
				fmt.Sscanf(bookChoice, "%d", &bookIndex)

				if bookIndex >= 1 && bookIndex < optionIndex {
					selectedBook := bookOptions[bookIndex]
					UseSpecificSpellBook(c, selectedBook)

					c.Inventaire[selectedBook]--
					if c.Inventaire[selectedBook] == 0 {
						delete(c.Inventaire, selectedBook)
					}
				} else {
					fmt.Println("📚 Vous rangez vos livres.")
				}

				fmt.Print("\nAppuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
			} else {
				fmt.Println("\n❌ Vous n'avez pas de Livre de Sort à utiliser.")
				fmt.Print("Appuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
			}

		case "5":
			if hasEquipment {
				ClearScreen()
				fmt.Println("╔══════════════════════════════════════════╗")
				fmt.Println("║           👕 ÉQUIPEMENT 👕              ║")
				fmt.Println("╚══════════════════════════════════════════╝")

				fmt.Println("\n📊 Équipements actuels :")
				fmt.Printf("🧢 Tête : %s\n", c.Equipement.Tete)
				fmt.Printf("👚 Torse : %s\n", c.Equipement.Torse)
				fmt.Printf("👞 Pieds : %s\n", c.Equipement.Pieds)
				fmt.Printf("⚔️  Arme : %s\n", c.Equipement.Arme)

				fmt.Println("\nChoisissez un équipement à porter :")
				optionIndex := 1
				equipOptions := make(map[int]string)

				if chapeauQty > 0 {
					fmt.Printf("%d. Casquette de Mario (Tête, +10 PV max)\n", optionIndex)
					equipOptions[optionIndex] = "Casquette de Mario"
					optionIndex++
				}
				if tuniqueQty > 0 {
					fmt.Printf("%d. Tunique de Link (Torse, +25 PV max)\n", optionIndex)
					equipOptions[optionIndex] = "Tunique de Link"
					optionIndex++
				}
				if bottesQty > 0 {
					fmt.Printf("%d. Bottes de Sonic (Pieds, +5 Initiative)\n", optionIndex)
					equipOptions[optionIndex] = "Bottes de Sonic"
					optionIndex++
				}

				if epeRouilleeQty > 0 {
					fmt.Printf("%d. Épée Rouillée (Arme, +2 dégâts)\n", optionIndex)
					equipOptions[optionIndex] = "Épée Rouillée"
					optionIndex++
				}
				if epeFerQty > 0 {
					fmt.Printf("%d. Épée en Fer (Arme, +4 dégâts)\n", optionIndex)
					equipOptions[optionIndex] = "Épée en Fer"
					optionIndex++
				}
				if hacheBatailleQty > 0 {
					fmt.Printf("%d. Hache de Bataille (Arme, +6 dégâts)\n", optionIndex)
					equipOptions[optionIndex] = "Hache de Bataille"
					optionIndex++
				}
				if epeAcierQty > 0 {
					fmt.Printf("%d. Épée en Acier (Arme, +8 dégâts)\n", optionIndex)
					equipOptions[optionIndex] = "Épée en Acier"
					optionIndex++
				}
				if arcElfiqueQty > 0 {
					fmt.Printf("%d. Arc Elfique (Arme, +10 dégâts)\n", optionIndex)
					equipOptions[optionIndex] = "Arc Elfique"
					optionIndex++
				}
				if batonMagiqueQty > 0 {
					fmt.Printf("%d. Bâton Magique (Arme, +12 dégâts)\n", optionIndex)
					equipOptions[optionIndex] = "Bâton Magique"
					optionIndex++
				}
				if epeLegendaireQty > 0 {
					fmt.Printf("%d. Épée Légendaire (Arme, +15 dégâts)\n", optionIndex)
					equipOptions[optionIndex] = "Épée Légendaire"
					optionIndex++
				}

				fmt.Printf("%d. Retour à l'inventaire\n", optionIndex)

				fmt.Print("\nVotre choix : ")
				equipChoice, _ := reader.ReadString('\n')
				equipChoice = strings.TrimSpace(equipChoice)

				var equipIndex int
				fmt.Sscanf(equipChoice, "%d", &equipIndex)

				if equipIndex >= 1 && equipIndex < optionIndex {
					selectedItem := equipOptions[equipIndex]
					pvBonus := EquipItem(c, selectedItem)

					UpdateMaxHP(c)
					UpdateInitiative(c)

					if pvBonus > 0 {
						fmt.Printf("💪 Vos PV maximum sont maintenant de %d.\n", c.PvMax)
					}
					fmt.Printf("⚡ Votre initiative est maintenant de %d.\n", c.Initiative)
				} else {
					fmt.Println("Retour à l'inventaire...")
				}
			} else {
				fmt.Println("\n❌ Vous n'avez aucun équipement à porter.")
			}
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')

		case "6":
			fmt.Println("\n🚪 Retour au menu principal...")
			return

		default:
			fmt.Println("\n❌ Choix invalide. Veuillez choisir entre 1 et 6.")
			fmt.Print("Appuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
		}

		ClearScreen()
	}
}

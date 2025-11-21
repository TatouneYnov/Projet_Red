package red

import "fmt"

func EquipItem(c *Character, itemName string) int {
	if c.Inventaire[itemName] <= 0 {
		fmt.Println("❌ Vous ne possédez pas cet objet dans votre inventaire.")
		return 0
	}

	var slot string
	var oldItem string
	var pvBonus int

	switch itemName {
	case "Casquette de Mario+":
		slot = "Tete"
		oldItem = c.Equipement.Tete
		pvBonus = 20

	case "Tunique de Link+":
		slot = "Torse"
		oldItem = c.Equipement.Torse
		pvBonus = 45

	case "Bottes de Sonic+":
		slot = "Pieds"
		oldItem = c.Equipement.Pieds
		pvBonus = 0

	case "Épée Rouillée":
		slot = "Arme"
		oldItem = c.Equipement.Arme
		pvBonus = 0
	case "Épée en Fer":
		slot = "Arme"
		oldItem = c.Equipement.Arme
		pvBonus = 0
	case "Épée en Acier":
		slot = "Arme"
		oldItem = c.Equipement.Arme
		pvBonus = 0
	case "Épée Légendaire":
		slot = "Arme"
		oldItem = c.Equipement.Arme
		pvBonus = 0
	case "Hache de Bataille":
		slot = "Arme"
		oldItem = c.Equipement.Arme
		pvBonus = 0
	case "Arc Elfique":
		slot = "Arme"
		oldItem = c.Equipement.Arme
		pvBonus = 0
	case "Bâton Magique":
		slot = "Arme"
		oldItem = c.Equipement.Arme
		pvBonus = 0

	case "Casque de Chevalier":
		slot = "Tete"
		oldItem = c.Equipement.Tete
		pvBonus = 15
	case "Armure Renforcée":
		slot = "Torse"
		oldItem = c.Equipement.Torse
		pvBonus = 35
	case "Bottes de Forgeron":
		slot = "Pieds"
		oldItem = c.Equipement.Pieds
		pvBonus = 0

	default:
		fmt.Println("❌ Cet objet n'est pas équipable.")
		return 0
	}

	if oldItem != "Aucun" && oldItem != "Vêtements simples" && oldItem != "Sandales usées" && oldItem != "Mains nues" {
		if c.Inventaire == nil {
			c.Inventaire = make(map[string]int)
		}
		c.Inventaire[oldItem]++
		fmt.Printf("📦 Vous avez retiré %s et l'avez remis dans votre inventaire.\n", oldItem)
	}

	switch slot {
	case "Tete":
		c.Equipement.Tete = itemName
	case "Torse":
		c.Equipement.Torse = itemName
	case "Pieds":
		c.Equipement.Pieds = itemName
	case "Arme":
		c.Equipement.Arme = itemName
	}

	c.Inventaire[itemName]--
	if c.Inventaire[itemName] <= 0 {
		delete(c.Inventaire, itemName)
	}

	switch slot {
	case "Arme":
		damageBonus := GetWeaponDamageBonus(itemName)
		fmt.Printf("✅ Vous avez équipé %s ! (+%d dégâts)\n", itemName, damageBonus)
	case "Pieds":
		initiativeBonus := GetBootsInitiativeBonus(itemName)
		fmt.Printf("✅ Vous avez équipé %s ! (+%d initiative)\n", itemName, initiativeBonus)
	default:
		if pvBonus > 0 {
			fmt.Printf("✅ Vous avez équipé %s ! (+%d PV max)\n", itemName, pvBonus)
		} else {
			fmt.Printf("✅ Vous avez équipé %s !\n", itemName)
		}
	}
	return pvBonus
}

func GetWeaponDamageBonus(weaponName string) int {
	switch weaponName {
	case "Épée Rouillée":
		return 2
	case "Épée en Fer":
		return 4
	case "Hache de Bataille":
		return 6
	case "Épée en Acier":
		return 8
	case "Arc Elfique":
		return 10
	case "Bâton Magique":
		return 12
	case "Épée Légendaire":
		return 15
	default:
		return 0
	}
}

func GetBootsInitiativeBonus(bootsName string) int {
	switch bootsName {
	case "Bottes de Sonic+":
		return 12
	case "Bottes de Forgeron":
		return 8
	default:
		return 0
	}
}

func UpdateMaxHP(c *Character) {
	baseHP := 100

	c.PvMax = baseHP

	if c.Equipement.Tete == "Casquette de Mario" {
		c.PvMax += 10
	}
	if c.Equipement.Torse == "Tunique de Link" {
		c.PvMax += 25
	}

	if c.PvActuelle > c.PvMax {
		c.PvActuelle = c.PvMax
	}
}

func UpdateInitiative(c *Character) {
	baseInitiative := c.Agilite

	c.Initiative = baseInitiative

	c.Initiative += GetBootsInitiativeBonus(c.Equipement.Pieds)
}

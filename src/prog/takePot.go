package red

import "fmt"

func TakePot(c *Character) {
	potionKey := "Potion de vie"
	quantity, exists := c.Inventaire[potionKey]

	if !exists || quantity <= 0 {
		fmt.Println("❌ Vous n'avez pas de potion à utiliser.")
		return
	}

	c.Inventaire[potionKey]--
	if c.Inventaire[potionKey] <= 0 {
		delete(c.Inventaire, potionKey)
	}

	fmt.Printf("🩸 PV avant soin : %d / %d\n", c.PvActuelle, c.PvMax)

	soin := 50
	if c.PvActuelle+soin > c.PvMax {
		c.PvActuelle = c.PvMax
	} else {
		c.PvActuelle += soin
	}
	fmt.Println("🧪 Vous avez utilisé une potion de vie !")
	fmt.Printf("❤️ PV après soin : %d / %d\n", c.PvActuelle, c.PvMax)
}

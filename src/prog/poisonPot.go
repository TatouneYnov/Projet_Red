package red

import (
	"fmt"
	"time"
)

func PoisonPot(c *Character) {
	potionKey := "Toxine de Teemo"
	quantity, exists := c.Inventaire[potionKey]

	if !exists || quantity <= 0 {
		fmt.Println("❌ Vous n'avez pas de toxine à utiliser.")
		return
	}

	c.Inventaire[potionKey]--
	if c.Inventaire[potionKey] <= 0 {
		delete(c.Inventaire, potionKey)
	}

	fmt.Println("☠️  Vous appliquez la toxine de Teemo...")
	fmt.Println("🤢 Le poison de League of Legends commence à faire effet !")

	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)

		damage := 10
		c.PvActuelle -= damage

		if c.PvActuelle < 0 {
			c.PvActuelle = 0
		}

		fmt.Printf("💀 Toxine (seconde %d) : -%d PV\n", i, damage)
		fmt.Printf("❤️  PV actuels : %d/%d\n", c.PvActuelle, c.PvMax)

		if c.PvActuelle <= 0 {
			fmt.Println("☠️  La toxine vous a terrassé ! Le capitaine Teemo en service !")
			IsDead(c)
			break
		}
	}

	if c.PvActuelle > 0 {
		fmt.Println("🩹 Les effets de la toxine se dissipent...")
	}
}

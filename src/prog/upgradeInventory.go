package red

import "fmt"

func UpgradeInventorySlot(c *Character) bool {
	if c.UpgradesUsed >= 3 {
		fmt.Println("❌ Vous avez déjà utilisé toutes vos améliorations d'inventaire (3/3).")
		return false
	}

	c.InventorySize += 10
	c.UpgradesUsed++

	fmt.Printf("✅ Votre inventaire a été amélioré ! Nouvelle capacité : %d objets.\n", c.InventorySize)
	fmt.Printf("📊 Améliorations utilisées : %d/3\n", c.UpgradesUsed)

	return true
}

package red

type Equipment struct {
	Tete       string
	Torse      string
	Pieds      string
	Arme       string
	Initiative int
}

type MainQuest struct {
	Progression       int
	Completed         bool
	FinalBossDefeated bool
}

type Character struct {
	Nom           string
	Classe        string
	Niveau        int
	PvMax         int
	PvActuelle    int
	ExpActuelle   int
	ExpMax        int
	ExpTotal      int
	Force         int
	Agilite       int
	Intelligence  int
	PointsStats   int
	Inventaire    map[string]int
	Skill         []string
	Argent        int
	Equipement    Equipment
	InventorySize int
	UpgradesUsed  int
	Initiative    int
	MainQuest     MainQuest
	Mana          int
	ManaMax       int
}

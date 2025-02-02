package main

type Spell struct {
	name      string
	mana      int
	damage    int
	hitPoints int
	hasEffect bool
	effect    Effect
}

func getSpells() []Spell {
	return []Spell{
		{
			name:      "Magic Missile",
			mana:      53,
			damage:    4,
			hasEffect: false,
			effect:    Effect{},
		},
		{
			name:      "Drain",
			mana:      73,
			damage:    2,
			hitPoints: 2,
			hasEffect: false,
			effect:    Effect{},
		},
		{
			name:      "Shield",
			mana:      113,
			hasEffect: true,
			effect: Effect{
				name:   "Shield",
				timer:  6,
				armour: 7,
			},
		},
		{
			name:      "Poison",
			mana:      173,
			hasEffect: true,
			effect: Effect{
				name:   "Poison",
				timer:  6,
				damage: 3,
			},
		},
		{
			name:      "Recharge",
			mana:      229,
			hasEffect: true,
			effect: Effect{
				name:  "Recharge",
				timer: 5,
				mana:  101,
			},
		},
	}
}

package main

import "fmt"

// Method chain: Define Creature with Name, Attack, Defense; Modifier interface with Add(Modifier), Handle(); CreatureModifier with *Creature & next Modifier; DoubleAttackModifier embeds CreatureModifier
// Broker Chain: CoR + Mediator + Observer + CQS

type Creature struct {
	Name    string
	Attack  int
	Defense int
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack int, defense int) *Creature {
	return &Creature{name, attack, defense}
}

type Modifier interface {
	Add(Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{c, nil}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "'attack...")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

func main() {
	// Create Creature
	c := &Creature{Name: "Goblin", Attack: 1, Defense: 1}

	fmt.Println(c)

	root := NewDoubleAttackModifier(c)
	root.Handle()

	fmt.Println(c)
}

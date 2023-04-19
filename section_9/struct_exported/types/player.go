// Package athletes
package athletes

import (
	"strings"
)

// Info...
type Info struct {
	Country    string
	HairColour string
}

// Player...
type Player struct {
	Name  string
	Sport string
	Age   int
	Info
}

// ToLowerCase...
func (p *Player) ToLowerCase() *Player {
	p.Name = strings.ToLower(p.Name)
	p.Sport = strings.ToLower(p.Sport)
	p.Country = strings.ToLower(p.Country)
	p.HairColour = strings.ToLower(p.HairColour)

	return p
}

package athletes

import (
	"strings"
)

// Info ...your comments
type Info struct {
	Country    string
	HairColour string
}

// Player ...your comments
type Player struct {
	Name  string
	Sport string
	Age   int
	Info
}

// ToLowerCase ...your comments
func (p *Player) ToLowerCase() *Player {

	p.Name = strings.ToLower(p.Name)
	p.Sport = strings.ToLower(p.Sport)
	p.Country = strings.ToLower(p.Country)
	p.HairColour = strings.ToLower(p.HairColour)

	return p

}

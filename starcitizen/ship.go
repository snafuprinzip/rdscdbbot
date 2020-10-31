package starcitizen

import (
	"fmt"

	"github.com/go-yaml/yaml"
)

var (
	// Ships contains all available ships
	Ships []Ship
)

// Ship represents a single ship
type Ship struct {
	fullName    string
	Vendor      string
	Chassis     string
	Model       string `yaml:",omitempty"`
	Size        string
	Shiptype    string `yaml:"type"`
	Focus       string
	Length      float64
	Mass        int
	MinCrew     int
	MaxCrew     int
	MaxSpeed    int
	MaxAFBSpeed int
	Pitch       int
	Yaw         int
	Roll        int
	Cargo       int
	Fuel        int
	QuantumFuel float64
	Price       int
	Shop        string
	URL         string
	Shields     []Component
	Weapons     []Component
}

// CreateNewShips creates a new ship db
func CreateNewShips() {
	fmt.Println("creating new ship database...")
	Ships = make([]Ship, 0)

	s := Ship{
		Vendor:      "Aegis",
		Chassis:     "Avenger",
		Model:       "Stalker",
		Size:        "Small",
		Shiptype:    "Combat",
		Focus:       "Interdiction",
		Length:      22.5,
		Mass:        50040,
		MinCrew:     1,
		MaxCrew:     1,
		MaxSpeed:    184,
		MaxAFBSpeed: 1307,
		Pitch:       65,
		Yaw:         65,
		Roll:        120,
		Cargo:       0,
		Fuel:        90000,
		QuantumFuel: 583.33,
		Price:       882200,
		Shop:        "New Deal, Lorville",
		URL:         "https://robertsspaceindustries.com/pledge/ships/aegis-avenger/Avenger-Stalker",
	}
	s.fullName = charReplacer.Replace(s.Vendor + "-" + s.Chassis +
		"-" + s.Model)

	Ships = append(Ships, s)

	return
}

func (s Ship) String() string {
	return fmt.Sprintf(
		"%s %s %s\n"+
			"%s / %s\n"+
			"%s (%.2f m)\n"+
			"Mass:\t%d kg\n"+
			"Crew:\t%d / %d\n"+
			"max. Speed (AFB): %d (%d)\n"+
			"P/Y/R:\t%d / %d / %d degrees\n"+
			"Cargo:\t%d SCU\n"+
			"Fuel (Quantum): %d (%.2f)\n"+
			"%s (aUEC %d)\n"+
			"%s\n",
		s.Vendor, s.Chassis, s.Model,
		s.Shiptype, s.Focus,
		s.Size, s.Length, s.Mass,
		s.MinCrew, s.MaxCrew, s.MaxSpeed, s.MaxAFBSpeed,
		s.Pitch, s.Yaw, s.Roll,
		s.Cargo, s.Fuel, s.QuantumFuel,
		s.Shop, s.Price, s.URL,
	)
}

// Yaml returns a yaml encoded byte array of the Ship
func (s Ship) Yaml() []byte {
	y, err := yaml.Marshal(&s)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return y
}

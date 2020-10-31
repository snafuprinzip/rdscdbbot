package starcitizen

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-yaml/yaml"
)

var (
	// Components contains all Components
	Components []Component

	// charReplacer is used for the generation of unique ids and filenames
	charReplacer = strings.NewReplacer(" ", "_", ".", "")
)

// Component represents a single generic ship component
type Component struct {
	fullName             string
	Manufacturer         string
	Function             string
	Name                 string
	Class                string
	Grade                int
	Size                 int
	PowerDrawRequestTime float64
	PowerToEMP           float64
	TemperatureToIR      float64
}

func (c Component) String() string {
	return fmt.Sprintf("%s %s\n"+
		"%s\t(Size:\t%d)\n"+
		"%s\t(Grade:\t%d)\n"+
		"Power draw request time: %.2f s\n"+
		"Power to EM:\t\t%.2f\n"+
		"Temperature to IR:\t%.2f\n", c.Manufacturer, c.Name,
		c.Function, c.Size,
		c.Class, c.Grade,
		c.PowerDrawRequestTime, c.PowerToEMP,
		c.TemperatureToIR)
}

// Yaml returns a yaml encoded byte array of the Component
func (c Component) Yaml() []byte {
	y, err := yaml.Marshal(&c)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(y))
	return y
}

// CreateNewComponents creates a new components db
func CreateNewComponents() {
	fmt.Println("creating new components database...")
	Components = make([]Component, 0)

	c := Component{
		Manufacturer:         "Lightning Power Ltd.",
		Function:             "Power Plant",
		Name:                 "Zapjet",
		Class:                "Civilian",
		Grade:                4,
		Size:                 1,
		PowerDrawRequestTime: 11.25,
		PowerToEMP:           2.31,
		TemperatureToIR:      2.31,
	}

	c.fullName = charReplacer.Replace(c.Manufacturer + "-" + c.Name)

	Components = append(Components, c)
	return
}

// GetComponent returns a component
func GetComponent(name string) *Component {
	for _, c := range Components {
		if name == c.Name {
			return &c
		}
	}
	return nil
}

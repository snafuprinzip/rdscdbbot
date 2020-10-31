package starcitizen

import (
	"fmt"

	"github.com/go-yaml/yaml"
)

// HardPoints is a list of the ships available hardpoints
var HardPoints []HardPoint

// HardPoint represents a ship slot to attach components to
type HardPoint struct {
	Kind      string
	Size      int
	component string
}

func (hp HardPoint) String() string {
	return fmt.Sprintf("%s (size: %d)\n", hp.Kind, hp.Size)
}

// Yaml returns a yaml encoded HardPoint struct
func (hp HardPoint) Yaml() []byte {
	y, err := yaml.Marshal(hp)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return y
}

// CreateNewHardPoints is only a temporary solution
func CreateNewHardPoints() {
	fmt.Println("creating hardpoints...")
	HardPoints = make([]HardPoint, 0)

	hp := HardPoint{
		Kind:      "Power Plant",
		Size:      1,
		component: "Lightning_Power_Ltd-Zapjet",
	}
	HardPoints = append(HardPoints, hp)
	return
}

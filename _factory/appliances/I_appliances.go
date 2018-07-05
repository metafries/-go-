package appliances

import (
	"errors"
)

// Log errors when they occur

// Appliance : The main interface used to describe appliances
type Appliance interface {
	Start()
	GetPurpose() string
}

const ( // Appliance Types
	STOVE = iota
	FRIDGE
	MICROWAVE
)

func CreateAppliance(aType int) (Appliance, error) {
	switch aType {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	case MICROWAVE:
		return new(Microwave), nil
	default:
		return nil, errors.New("Invalid Appliance Type")
	}
}

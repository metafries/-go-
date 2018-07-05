package appliances

// Fridge : Contain a string representing the type name
type Fridge struct {
	typeName string
}

// Start : The Fridge struct implements the start() function
func (fr *Fridge) Start() {
	fr.typeName = " Fridge "
}

// GetPurpose : The Fridge struct implements the GetPurpose() function
func (fr *Fridge) GetPurpose() string {
	return "I am a" + fr.typeName + "I cool stuff down!"
}

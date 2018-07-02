package appliances

// Stove : Contain a string representing the type name
type Stove struct {
	typeName string
}

// Start : The Stove struct implements the start() function
func (sv *Stove) Start() {
	sv.typeName = " Stove "
}

// GetPurpose : The Stove struct implements the GetPurpose() function
func (sv *Stove) GetPurpose() string {
	return "I am a" + sv.typeName + "I cook food!"
}

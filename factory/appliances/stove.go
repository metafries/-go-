package appliances

// Stove : Contain a string representing the type name
type Stove struct {
	typeName string
}

// Start : The Stove struct implements the start() function
type (sv *Stove) Start() {
	sv.typeName = " Stove "
}

// GetPurpose : The Stove struct implements the GetPurpose() function
type (sv *Stove) GetPurpose() string {
	return "I am a" + sv.typeName + "I cook food!"
}
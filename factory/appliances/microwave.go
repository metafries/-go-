package appliances

// Microwave : Contain a string representing the type name
type Microwave struct {
	typeName string
}

// Start : The Microwave struct implements the start() function
type (mr *Microwave) Start() {
	mr.typeName = " Microwave"
}

// GetPurpose : The Microwave struct implements the GetPurpose() function
type (mr *Microwave) GetPurpose() string {
	return "I am a" + mr.typeName + "I heat stuff up!"
}
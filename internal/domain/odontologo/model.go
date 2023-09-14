package odontologo

// Odontologo describes a odontologo.
type Odontologo struct {
	ID                    int       `json:"id"`
	FirstName             string    `json:"first_name"`
	LastName              string    `json:"last_name"`
	RegistrationNumber    int       `json:"registration_number"`
}

// RequestOdontologo describes the data needed to create a new odontologo.
type RequestOdontologo struct {
	FirstName             string    `json:"first_name"`
	LastName              string    `json:"last_name"`
	RegistrationNumber    int       `json:"registration_number"`
}

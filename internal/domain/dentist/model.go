package dentist

// Dentist describes a dentist.
type Dentist struct {
	ID                 int    `json:"id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	RegistrationNumber int    `json:"registration_number"`
}

// NewDentist describes the data needed to create a new Dentist.
type NewDentist struct {
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	RegistrationNumber int    `json:"registration_number"`
}

package user

// User is a user
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastNameName"`
}

// ByID fetches a user by ID
func ByID(id string) (User, error) {

	// mock
	return User{
		ID:        "abcde12345",
		FirstName: "Mike",
		LastName:  "Donnici",
	}, nil

}

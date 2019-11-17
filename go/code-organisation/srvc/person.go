package srvc

import "fmt"

// Person represents a Person entity and is managed 
// via a PersonStore. This is the type that will 
// implement the CRUD interface(?)
type Person struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
	Age  string `json:"age"`

	// The attached store manages the Person data
	store PersonStore
}

const (
	PersonErrorNoID = "Person.ID is missing"
)



// Validate imposes rules on what is a valid Person value.
// Use the tags validation package.
func (p *Person) Validate() error {
	// ... do the validation
	return nil
}

// Create saves a validated Person value to the store
func (p *Person) Create() error {
	var err error
	err = p.Validate()
	if err != nil {
		return fmt.Errorf("Person.Validate() err = %s", err)
	}
	return p.store.Save(p)
}

// Read is effectively a search and might look for 
// real person data based on any of the field values
// specified in the person value. Obviously, ID is the
// most important and, if present, would be the sole 
// search criteria. 
func (p *Person) Read() error {
	// Read from store by ID
	if p.ID == "" {
		return p.store.Fetch(p)
	}
	return fmt.Errorf("Person.Read() err = %s", PersonErrorNoID)
}

// Update saves a value of the type, assuming it already exists.
func (p *Person) Update() error {
	_, err := p.store.ByID(p.ID)
	if err != nil {
		return fmt.Errorf("Person.Update() err = %s", err)
	}
	return p.store.Save(p)
}

// Delete deletes a value, imposing any required rules
func (p *Person) Delete() error {
	return p.store.Remove(p)
}

// save stores a Person value in whatever way is needed and saves
// the results to person.
func (ps PersonStore) save(person *Person) error {
	// Save to db and return the fresh value. The process of
	// saving may add new things as well, like an id for example.
	// ... do the save....
	// assign values to the person arg
	return nil
}

// Fetch retrieves a Person record based on whatever
// relevant data is available and / or required.
func (ps PersonStore) Fetch(person *Person) error {
	// should assign all values to person but we
	// will lose the store?
	person, err := ps.byID(person.ID)
	return err
}
// Remove performs the delete operation for the type,
// in whatever way is appropriate, across one or more
// data stores.
func (ps PersonStore) Remove(person *Person) error {
	// ... remove from wherever
	return nil
}

// byID fetches a record by its primary identifier
func (ps PersonStore) byID(id string) (*Person, error) {
	// DB lookup
	return &Person{}, nil
}

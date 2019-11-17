package srvc

// PersonStore deals with the nitty gritty of managing the
// data for a Person type. This is the interface between
// requests for Person types and the data source of those types
type PersonStore struct {
	store Store
}

// NewPersonStore sets up a PersonStore
func NewPersonStore(store Store) *PersonStore {
	return &PersonStore {
		store: store,
	}
}

// Save stores a Person value in whatever way is needed and saves
// the results to person.
func (ps PersonStore) Save(person *Person) error {
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




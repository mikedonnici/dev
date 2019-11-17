package srvc

// Store holds database connections
type Store struct {
	Connections []Connection
}

// Connection is a database connection
type Connection struct {
	Name string
}

// DataManager describes the  'VCRUD' interface
type DataManager interface {
	Validate() error
	Create() error
	Read() error
	Update() error
	Delete() error
}

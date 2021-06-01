package db

// Generalized PIR database interface to be implemented
// TODO: Determine concrete function signatures for this interface
type PIRDB interface {
	// Initialized the PIR database with some data
	Init() error
	// Updates the PIR database with new data
	Update() error
	// Removes data from the PIR database
	Remove() error
	// Gets data from the PIR database
	RetrieveData() error
	// Retrieves Manifest file for the DB
	GetManifestFile() (*ManifestFile, error)
}

// Manifest file interface meant to be implemented as part of a PIRDB implementation
// Manifest files are used to give PIR clients a view of a PIR server's database in order
// properly query for their needed data.
// TODO: Finalize function signatures
type ManifestFile interface {
	// Serializes a Manifest file into some desirable format
	Serialize() (*byte, error)
	Deserialize() (*byte, error)
	OutputToFile() (*byte, error)
}

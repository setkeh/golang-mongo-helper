package mongo

// Config Mongo DB
type Config struct {
	Host       []string
	Username   string
	Password   string
	Database   string
	Collection string
}

// Query Model to Make Querys Generic
type Query struct {
	QueryKey   string
	QueryValue string
}

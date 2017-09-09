package mongo

// Config Mongo DB
type Config struct {
	Host     []string
	Username string
	Password string
	Database string
}

// Query Model to Make Querys Generic
type Query struct {
	Collection string
	QueryKey   string
	QueryValue string
}

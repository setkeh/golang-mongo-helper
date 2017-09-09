# golang-mongo-helper
Mongo DB helper For various projects

This is a Mongo DB Helper im building for working on a couple of other projects.


Basic Usage
------------

There are Currently two Configuration Structs one for configuring the connection and one for configuring the query and they look like,

##### Configure the Connection
```
type Config struct {
	Host     []string
	Username string
	Password string
	Database string
}
```

##### Configure the Query
```
type Query struct {
	Collection string
	QueryKey   string
	QueryValue string
}
```

##### Mongo Connection Code
```
dbconf := mongo.Config{
		Host:     []string{"localhost:27017"},
		Username: "UserName",
		Password: "Password",
		Database: "atlas",
	}

	session, err := mongo.Connect(dbconf)
	if err != nil {
		log.Fatal(err)
	}
```

##### Running a Query
````
query := mongo.Query{
		Collection: "users",
		QueryKey:   "Username",
		QueryValue: "Testuser1",
	}

	result, err := mongo.DoQuery(session, dbconf, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
```


TODO:
------------

Add SSL Support.
Probably loads more.
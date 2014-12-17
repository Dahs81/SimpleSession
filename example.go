package main

import "github.com/Dahs81/simple-mgo-db/db"

func main() {
	d := db.New()
	// d.SetEnv(db.Env{Host: "MONGO_HOST_TEST", Port: "MONGO_PORT_TEST"})\
	d.SetHost("127.0.0.1")
	d.SetPort(27017)
	d.Start("test1")
}

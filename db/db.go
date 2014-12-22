package db

import (
	"fmt"
	"strconv"

	"github.com/swhite24/envreader"

	"gopkg.in/mgo.v2"
)

var (
	envSet     bool
	connection string
)

type (
	// SimpleSession - Wrapper around Session
	SimpleSession struct {
		Session *mgo.Session
		Host    string
		Port    string
		User    string
		Pass    string
		Env     Env
		DbName  string
	}

	// Env - Stores the Environment variables common to mongodb
	Env struct {
		Host string
		Port string
		User string
		Pass string
	}
)

// New - Creates a new instance of SimpleSession
func New() *SimpleSession {

	envSet = false

	return &SimpleSession{
		Session: nil,
		User:    "",          // Default has no username
		Pass:    "",          // Default has password
		Host:    "localhost", // Default Host
		Port:    "27017",     // Default Port
		Env:     Env{},
	}
}

// Start - Starts the mongo connection
func (s *SimpleSession) Start(name string) {
	s.DbName = name

	fmt.Println(s.Env)

	fmt.Println(envSet)

	// Read the env vars that you want to use
	dbEnv := envreader.Read(s.Env.Host, s.Env.Port, s.Env.User, s.Env.Pass)

	fmt.Println("dbEnv: ", dbEnv)
	connection = "mongodb://"

	// if no Env.User/Pass, skip setting that part of the connection
	if envSet {
		s.Host = dbEnv[s.Env.Host]
		s.Port = dbEnv[s.Env.Port]

		if dbEnv[s.Env.User] != "" && dbEnv[s.Env.Pass] != "" {
			connection += dbEnv[s.Env.User] + ":" + dbEnv[s.Env.Pass] + "@"
		}
	}

	if s.User != "" && s.Pass != "" {
		connection += s.User + ":" + s.Pass + "@"
	}

	connection += s.Host + ":"
	connection += s.Port + "/" + name

	fmt.Println("CONN: ", connection)

	session, err := mgo.Dial(connection)
	if err != nil {
		panic(err)
	}

	// defer session.Close()

	s.Session = session
}

// SetDbName -
func (s *SimpleSession) SetDbName(name string) {
	s.DbName = name
}

// SetEnv - Add environment variables to a SimpleSession
//
// Usage:
// myEnv := map[string]string{"TEST_DB_USER": "username", "TEST_DB_PASS": "password"}
// session.SetEnv(myEnv)
//
func (s *SimpleSession) SetEnv(e Env) {
	s.Env = e
	envSet = true
}

// SetHost -
func (s *SimpleSession) SetHost(host string) {
	s.Host = host
}

// SetPort -
func (s *SimpleSession) SetPort(port int) {
	s.Port = strconv.Itoa(port)
}

// SetUser -
func (s *SimpleSession) SetUser(user string) {
	s.User = user
}

// SetPass -
func (s *SimpleSession) SetPass(pass string) {
	s.Pass = pass
}

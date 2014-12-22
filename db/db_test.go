package db

import "testing"

func TestNew(t *testing.T) {
	d := New()

	if d.Session != nil {
		t.Errorf("The default Session value should be %s ", d.Session)
	}

	if d.Host != "localhost" {
		t.Errorf("The default Host value should be %s ", d.Host)
	}

	if d.Port != "27017" {
		t.Errorf("The default Port value should be %s ", d.Port)
	}

	if d.User != "" {
		t.Errorf("The default User value should be %s ", d.User)
	}

	if d.Pass != "" {
		t.Errorf("The default Pass value should be %s ", d.Pass)
	}

	if envSet != false {
		t.Errorf("Error")
	}
}

func TestStart(t *testing.T) {
	d := New()
	d.Start("test1")

	if d.DbName != "test1" {
		t.Errorf("Error")
	}

	if connection != "mongodb://localhost:27017/test1" {
		t.Errorf("Error")
	}
}

func TestSetEnv(t *testing.T) {

	d := New()
	env := Env{Host: "MONGO_TEST_HOST", Port: "MONGO_TEST_PORT", User: "", Pass: ""}
	d.SetEnv(env)

	if envSet != true {
		t.Errorf("Error")
	}
}

func TestSetHost(t *testing.T) {
	d := New()
	d.SetHost("127.0.0.1")

	if d.Host != "127.0.0.1" {
		t.Errorf("Host was not correctly set to %s ", d.Host)
	}
}

func TestSetPort(t *testing.T) {
	d := New()
	d.SetPort(55555)

	if d.Port != "55555" {
		t.Errorf("Port was not correctly set to %s ", d.Port)
	}
}

func TestSetUser(t *testing.T) {
	d := New()
	d.SetUser("username")

	if d.User != "username" {
		t.Errorf("User was not correctly set to %s ", d.User)
	}
}

func TestSetPass(t *testing.T) {
	d := New()
	d.SetPass("myPassword")

	if d.Pass != "myPassword" {
		t.Errorf("Pass was not correctly set to %s ", d.Pass)
	}
}

func TestSetDbName(t *testing.T) {
	d := New()
	d.SetDbName("test3")

	if d.DbName != "test3" {
		t.Errorf("DbName was not correctly set to %s ", d.DbName)
	}
}

## SimpleSession

A wrapper around mgo.Session that allows the user to specify ENV variables (or not).

There are two ways to setup your mongo session.  

You can use ENV variable by calling the SetEnv function.  If you do this, you will need to export the names of the ENV variables in you ~/.bashrc or ~/.profile.  Be sure to source this file.  If you use the SetEnv function, you should not use the other setter functions.

If you do not want to use ENV variables, you can set up your connection path by using the SetHost, SetPath, SetUser (optional), and SetPass(optional) functions.  If you use these setters, you should not use the SetEnv function.

### Examples

##### Using environment variables

```
package main

import "github.com/Dahs81//db"

func main() {
  d := db.New()
  d.SetEnv(db.Env{Host: "MONGO_HOST_TEST", Port: "MONGO_PORT_TEST"})
  d.Start("test1")
}
```

#### Using Setters

```
package main

import "github.com/Dahs81/SimpleSession/db"

func main() {
  d := db.New()
  d.SetHost("127.0.0.1") // localhost
  d.SetHost(27017)
  d.Start("test1")
}
```

##### Using Setters with Username and Password

```
package main

import "github.com/Dahs81/SimpleSession/db"

func main() {
  d := db.New()
  d.SetHost("127.0.0.1") // localhost
  d.SetHost(27017)
  d.SetUser("username")
  d.SetPass("password")
  d.Start("test1")
}
```

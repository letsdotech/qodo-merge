package main

var users = map[string]string{}

func saveUser(id, name string) {
	go func() {
		users[id] = name // unsafe concurrent write
	}()
}

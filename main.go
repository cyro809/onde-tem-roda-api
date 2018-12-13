package main

import (
	"fmt"

	"github.com/onde-tem-roda-api/manager"
)

//olhar pacote flag

// criar um pacote cmd com um main para instrumentar os command lines
func main() {
	// user := manager.SelectUserByID(1)
	users := manager.SelectAllUsers()
	fmt.Println(users)
}

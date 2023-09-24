package main

import (
	"fmt"

	"github.com/marcoagpegoraro/marco_blog/initializers"
)

func init() {
	initializers.LoadEnvPackages()
	initializers.ConnectToDatabase()
}

func main() {
	fmt.Println("Hello World2")
}

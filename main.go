package main

import (
	"fmt"

	handlers "github.com/osoriojuan/go-integration-grcp/handlers"
)

func main() {

	fmt.Println("Starting app...")
	handlers.StartDefaultRouter()
	handlers.ConfigureRoutesAvailables()
	handlers.Serve()
	fmt.Println("Successfully started.")
}

package http

import (
	"fmt"

	"github.com/matias179/go-fiber-clone/internal/config"
)

// New start the app.
func New() {
	config.SetEnvironmentVariables()
	fmt.Println("calling new here!")
}

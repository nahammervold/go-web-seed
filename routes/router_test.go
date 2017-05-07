package routes

import (
	"fmt"
	"testing"
)

// TestGetRouter Returns a ServeMux pointer with the routes attached to it
func TestGetRouter(t *testing.T) {

	mux := GetRouter()
	if mux == nil {
		t.Fatal("Mux wasn't instantiated")
	}

	fmt.Printf("%s", "Mux was instantiated properly\n")
}

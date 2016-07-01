package azuregraph

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// TestNewDispatcher tests NewDispatcher
func TestNewDispatcher(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Errorf("Failed to load .env file")
		return
	}
	dispatcher, err := NewDispatcher(os.Getenv("TENANT_ID"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		t.Errorf("NewDispatcher() returns error: %s", err)
		return
	}
	if dispatcher.tokenType != "Bearer" {
		t.Errorf("Invalid Dispatcher type")
		return
	}
	u, err := dispatcher.getEndpoint("user")
	if err != nil {
		t.Errorf("Invalid resource type")
		return
	}
	if u.Scheme != "https" {
		t.Errorf("Invalid endpoint")
		return
	}
	if _, err := dispatcher.UserGet(os.Getenv("TEST_USER")); err != nil {
		t.Errorf("UserGet() failed: %s", err)
		return
	}
	users, err := dispatcher.UserList()
	if err != nil {
		t.Errorf("UserList() failed: %s", err)
		return
	}
	fmt.Println(users)
}

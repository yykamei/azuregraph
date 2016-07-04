package azuregraph

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// TestGetUsers tests NewDispatcher, UserGet and UserList
func TestGetUsers(t *testing.T) {
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
	if _, _, err = dispatcher.UserList(&OdataQuery{Top: 1}); err != nil {
		t.Errorf("UserList() failed: %s", err)
		return
	}
}

// TestGetGroups tests NewDispatcher, GroupGet and GroupList
func TestGetGroups(t *testing.T) {
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
	u, err := dispatcher.getEndpoint("group")
	if err != nil {
		t.Errorf("Invalid resource type")
		return
	}
	if u.Scheme != "https" {
		t.Errorf("Invalid endpoint")
		return
	}
	if _, err := dispatcher.GroupGet(os.Getenv("TEST_GROUP_ID")); err != nil {
		t.Errorf("GroupGet() failed: %s", err)
		return
	}
	if _, _, err = dispatcher.GroupList(&OdataQuery{Top: 1}); err != nil {
		t.Errorf("GroupList() failed: %s", err)
		return
	}
}

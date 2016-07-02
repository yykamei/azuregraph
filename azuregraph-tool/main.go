package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/yykamei/azuregraph/azuregraph"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app          = kingpin.New("azuregraph-tool", "A command-line tool for Azure AD Graph API.")
	tenantID     = app.Flag("tenantid", "Tenant ID (Required)").Required().Short('T').String()
	clientID     = app.Flag("clientid", "OAuth2 Client ID (Required)").Required().Short('C').String()
	clientSecret = app.Flag("clientsecret", "OAuth2 Client secret file path(Required)").Required().Short('S').File()
	debug        = app.Flag("debug", "Enable debug mode.").Bool()

	user     = app.Command("user", "Manage users.")
	userGet  = user.Command("get", "Get a user.")
	userList = user.Command("list", "List users.")
	// userCreate           = user.Command("create", "Create a user.")
	// userUpdate           = user.Command("update", "Update a user.")
	// userDelete           = user.Command("delete", "Delete a user.")
	// userGetManager       = user.Command("getmanager", "Get a user's manager.")
	// userAssignManager    = user.Command("assignmanager", "Assign a user's manager.")
	// userGetDirectReports = user.Command("getdirectreports", "Get user's direct reports.")
	// userGetGroups        = user.Command("getgroups", "Get user's groups and directory role memberships.")

	// group             = app.Command("group", "Manage groups.")
	// groupGet          = group.Command("get", "Get a group.")
	// groupList         = group.Command("list", "List groups.")
	// groupCreate       = group.Command("create", "Create a group.")
	// groupUpdate       = group.Command("update", "Update a group.")
	// groupDelete       = group.Command("delete", "Delete a group.")
	// groupGetMembers   = group.Command("getmembers", "Get members.")
	// groupAddMembers   = group.Command("addmembers", "Add members.")
	// groupDeleteMember = group.Command("deletemember", "Delete a member.")
)

func createDispatcher() *azuregraph.Dispatcher {
	buf, err := ioutil.ReadAll(*clientSecret)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Failed to open client secret file: `%s'\n",
			err,
		)
		os.Exit(1)
		return nil
	}
	secret := strings.TrimSpace(string(buf))
	dispatcher, err := azuregraph.NewDispatcher(*tenantID, *clientID, secret)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Failed to start API: `%s'\n",
			err,
		)
		os.Exit(1)
		return nil
	}
	return dispatcher
}

func main() {
	registerUser(app)
	// registerGroup(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

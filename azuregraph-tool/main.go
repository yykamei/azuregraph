package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/yykamei/azuregraph/azuregraph"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app          = kingpin.New("azure-ad-graphapi-tool", "A command-line tool for Azure AD Graph API.")
	tenantID     = app.Flag("tenantid", "Tenant ID (Required)").Required().Short('T').String()
	clientID     = app.Flag("clientid", "OAuth2 Client ID (Required)").Required().Short('C').String()
	clientSecret = app.Flag("clientsecret", "OAuth2 Client secret file path(Required)").Required().Short('S').File()
	debug        = app.Flag("debug", "Enable debug mode.").Bool()

	// user <command> [options] [<args>...]

	user = app.Command("user", "Manage users.")

	userGet     = user.Command("get", "Get a user.")
	userGetName = userGet.Arg("userid", "user principal name or object ID").Required().String()

	userList          = user.Command("list", "List users.")
	userListFilter    = userList.Flag("filter", "$filter query for listing users.").Short('f').String()
	userListTop       = userList.Flag("top", "$top query for listing users.").Short('t').Int()
	userListOrderBy   = userList.Flag("orderby", "$orderby query for listing users.").Short('o').String()
	userListSkipToken = userList.Flag("skiptoken", "$skiptoken query for listing users.").Short('s').String()

	userCreate = user.Command("create", "Create a user.")

	userUpdate = user.Command("update", "Update a user.")

	userDelete = user.Command("delete", "Delete a user.")

	userGetManager = user.Command("getmanager", "Get a user's manager.")

	userAssignManager = user.Command("assignmanager", "Assign a user's manager.")

	userGetDirectReports = user.Command("getdirectreports", "Get user's direct reports.")

	userGetGroups = user.Command("getgroups", "Get user's groups and directory role memberships.")

	// group <command> [options] [<args>...]

	group = app.Command("group", "Manage groups.")

	groupGet     = group.Command("get", "Get a group.")
	groupGetName = groupGet.Arg("objectid", "Object ID").Required().String()

	groupList          = group.Command("list", "List groups.")
	groupListFilter    = groupList.Flag("filter", "$filter query for listing groups.").Short('f').String()
	groupListTop       = groupList.Flag("top", "$top query for listing groups.").Short('t').Int()
	groupListOrderBy   = groupList.Flag("orderby", "$orderby query for listing groups.").Short('o').String()
	groupListSkipToken = groupList.Flag("skiptoken", "$skiptoken query for listing groups.").Short('s').String()

	groupCreate = group.Command("create", "Create a group.")

	groupUpdate = group.Command("update", "Update a group.")

	groupDelete = group.Command("delete", "Delete a group.")

	groupGetMembers = group.Command("getmembers", "Get members.")

	groupAddMembers = group.Command("addmembers", "Add members.")

	groupDeleteMember = group.Command("deletemember", "Delete a member.")

	// contact = app.Command("contact", "Manage contacts.")

	// directoryRoles = app.Command("directoryroles", "Manage directory roles.")

	// domain = app.Command("domain", "Manage domains.")
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

func createOdataQuery(filter *string, top *int, orderby *string, skiptoken *string) *azuregraph.OdataQuery {
	query := azuregraph.OdataQuery{}
	if filter != nil {
		query.Filter = *filter
	}
	if top != nil {
		query.Top = *top
	}
	if orderby != nil {
		query.OrderBy = *orderby
	}
	if skiptoken != nil {
		query.SkipToken = *skiptoken
	}
	return &query
}

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case userGet.FullCommand():
		dispatcher := createDispatcher()
		u, err := dispatcher.UserGet(*userGetName)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to User GET API: `%s'\n",
				err,
			)
			os.Exit(1)
		}
		buf, err := json.MarshalIndent(u, "", "  ")
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Encode User Type: `%s'\n",
				err,
			)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(buf))
	case userList.FullCommand():
		dispatcher := createDispatcher()
		query := createOdataQuery(userListFilter, userListTop, userListOrderBy, userListSkipToken)
		us, skiptoken, err := dispatcher.UserList(query)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to User GET API: `%s'\n",
				err,
			)
			os.Exit(1)
		}
		buf, err := json.MarshalIndent(us, "", "  ")
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Encode User Type: `%s'\n",
				err,
			)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(buf))
		if skiptoken != nil {
			fmt.Println()
			fmt.Println("If you want to get a next page, you can specify following the skiptoken flag.")
			fmt.Println("You must quote it with double quotes characters.")
			fmt.Println()
			fmt.Printf("%#v\n", *skiptoken)
		}
	case userCreate.FullCommand():
		println("OK")
	case userUpdate.FullCommand():
		println("OK")
	case userDelete.FullCommand():
		println("OK")
	case userGetManager.FullCommand():
		println("OK")
	case userAssignManager.FullCommand():
		println("OK")
	case userGetDirectReports.FullCommand():
		println("OK")
	case userGetGroups.FullCommand():
		println("OK")
	case groupGet.FullCommand():
		dispatcher := createDispatcher()
		g, err := dispatcher.GroupGet(*groupGetName)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Group GET API: `%s'\n",
				err,
			)
			os.Exit(1)
		}
		buf, err := json.MarshalIndent(g, "", "  ")
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Encode Group Type: `%s'\n",
				err,
			)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(buf))
	case groupList.FullCommand():
		dispatcher := createDispatcher()
		query := createOdataQuery(groupListFilter, groupListTop, groupListOrderBy, groupListSkipToken)
		gs, skiptoken, err := dispatcher.GroupList(query)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Group GET API: `%s'\n",
				err,
			)
			os.Exit(1)
		}
		buf, err := json.MarshalIndent(gs, "", "  ")
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Encode Group Type: `%s'\n",
				err,
			)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(buf))
		if skiptoken != nil {
			fmt.Println()
			fmt.Println("If you want to get a next page, you can specify following the skiptoken flag.")
			fmt.Println("You must quote it with double quotes characters.")
			fmt.Println()
			fmt.Printf("%#v\n", *skiptoken)
		}
	}
}

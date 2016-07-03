package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yykamei/azuregraph/azuregraph"

	"gopkg.in/alecthomas/kingpin.v2"
)

func registerUser(app *kingpin.Application) {
	var (
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
	)
	registerUserGet(userGet)
	registerUserList(userList)
}

func registerUserGet(cmd *kingpin.CmdClause) {
	userid := cmd.Arg("userid", "user principal name or object ID").Required().String()
	cmd.Action(func(ctx *kingpin.ParseContext) error {
		dispatcher := createDispatcher()
		user, err := dispatcher.UserGet(*userid)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to User GET API: `%s'\n",
				err,
			)
			return err
		}
		buf, err := json.MarshalIndent(user, "", "  ")
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Encode User Type: `%s'\n",
				err,
			)
			return err
		}
		fmt.Printf("%s\n", string(buf))
		return nil
	})
}

func registerUserList(cmd *kingpin.CmdClause) {
	params := azuregraph.OdataQuery{}
	cmd.Flag("filter", "$filter query for listing users.").Short('f').StringVar(&params.Filter)
	cmd.Flag("top", "$top query for listing users.").Short('t').IntVar(&params.Top)
	cmd.Flag("orderby", "$orderby query for listing users.").Short('o').StringVar(&params.OrderBy)
	cmd.Flag("skiptoken", "$skiptoken query for listing users.").Short('s').StringVar(&params.SkipToken)

	cmd.Action(func(ctx *kingpin.ParseContext) error {
		dispatcher := createDispatcher()
		users, skiptoken, err := dispatcher.UserList(&params)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to User GET API: `%s'\n",
				err,
			)
			return err
		}
		buf, err := json.MarshalIndent(users, "", "  ")
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Encode User Type: `%s'\n",
				err,
			)
			return err
		}
		fmt.Printf("%s\n", string(buf))
		if skiptoken != nil {
			fmt.Println()
			fmt.Println("If you want to get a next page, you can specify following the skiptoken flag.")
			fmt.Println("You must quote it with double quotes characters.")
			fmt.Println()
			fmt.Printf("%#v\n", *skiptoken)
		}
		return nil
	})
}

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yykamei/azuregraph/azuregraph"

	"gopkg.in/alecthomas/kingpin.v2"
)

func registerGroup(app *kingpin.Application) {
	var (
		group     = app.Command("group", "Manage groups.")
		groupGet  = group.Command("get", "Get a group.")
		groupList = group.Command("list", "List groups.")
		// groupCreate       = group.Command("create", "Create a group.")
		// groupUpdate       = group.Command("update", "Update a group.")
		// groupDelete       = group.Command("delete", "Delete a group.")
		// groupGetMembers   = group.Command("getmembers", "Get members.")
		// groupAddMembers   = group.Command("addmembers", "Add members.")
		// groupDeleteMember = group.Command("deletemember", "Delete a member.")
	)
	registerGroupGet(groupGet)
	registerGroupList(groupList)
}

func registerGroupGet(cmd *kingpin.CmdClause) {
	objectid := cmd.Arg("objectid", "Object ID").Required().String()
	cmd.Action(func(ctx *kingpin.ParseContext) error {
		dispatcher := createDispatcher()
		group, err := dispatcher.GroupGet(*objectid)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Group GET API: `%s'\n",
				err,
			)
			return err
		}
		buf, err := json.MarshalIndent(group, "", "  ")
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Encode Group Type: `%s'\n",
				err,
			)
			return err
		}
		fmt.Printf("%s\n", string(buf))
		return nil
	})
}

func registerGroupList(cmd *kingpin.CmdClause) {
	params := azuregraph.OdataQuery{}
	cmd.Flag("filter", "$filter query for listing groups.").Short('f').StringVar(&params.Filter)
	cmd.Flag("top", "$top query for listing groups.").Short('t').IntVar(&params.Top)
	cmd.Flag("orderby", "$orderby query for listing groups.").Short('o').StringVar(&params.OrderBy)
	cmd.Flag("skiptoken", "$skiptoken query for listing groups.").Short('s').StringVar(&params.SkipToken)

	cmd.Action(func(ctx *kingpin.ParseContext) error {
		dispatcher := createDispatcher()
		groups, skiptoken, err := dispatcher.GroupList(&params)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Group GET API: `%s'\n",
				err,
			)
			return err
		}
		buf, err := json.MarshalIndent(groups, "", "  ")
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Encode Group Type: `%s'\n",
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

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yykamei/azuregraph/azuregraph"

	"gopkg.in/alecthomas/kingpin.v2"
)

func registerSubscribedSku(app *kingpin.Application) {
	var (
		subscribedSku     = app.Command("subscribedsku", "Manage subscribedSkus.")
		subscribedSkuGet  = subscribedSku.Command("get", "Get a subscribedSku.")
		subscribedSkuList = subscribedSku.Command("list", "List subscribedSkus.")
	)
	registerSubscribedSkuGet(subscribedSkuGet)
	registerSubscribedSkuList(subscribedSkuList)
}

func registerSubscribedSkuGet(cmd *kingpin.CmdClause) {
	objectid := cmd.Arg("objectid", "Object ID").Required().String()
	cmd.Action(func(ctx *kingpin.ParseContext) error {
		dispatcher := createDispatcher()
		subscribedSku, err := dispatcher.SubscribedSkuGet(*objectid)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to SubscribedSku GET API: `%s'\n",
				err,
			)
			return err
		}
		buf, err := json.MarshalIndent(subscribedSku, "", "  ")
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Encode SubscribedSku Type: `%s'\n",
				err,
			)
			return err
		}
		fmt.Printf("%s\n", string(buf))
		return nil
	})
}

func registerSubscribedSkuList(cmd *kingpin.CmdClause) {
	params := azuregraph.OdataQuery{}
	cmd.Flag("filter", "$filter query for listing subscribedSkus.").Short('f').StringVar(&params.Filter)
	cmd.Flag("top", "$top query for listing subscribedSkus.").Short('t').IntVar(&params.Top)
	cmd.Flag("orderby", "$orderby query for listing subscribedSkus.").Short('o').StringVar(&params.OrderBy)
	cmd.Flag("skiptoken", "$skiptoken query for listing subscribedSkus.").Short('s').StringVar(&params.SkipToken)

	cmd.Action(func(ctx *kingpin.ParseContext) error {
		dispatcher := createDispatcher()
		subscribedSkus, skiptoken, err := dispatcher.SubscribedSkuList(&params)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to SubscribedSku GET API: `%s'\n",
				err,
			)
			return err
		}
		buf, err := json.MarshalIndent(subscribedSkus, "", "  ")
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Failed to Encode SubscribedSku Type: `%s'\n",
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

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
	registerGroup(app)
	registerSubscribedSku(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

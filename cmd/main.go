package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

const (
	homeFolderName = ".plex-cli"
)

var (
	title          string
	baseURL        string
	token          string
	isVerbose      bool
	plexConn       *plex.Plex
	storeDirectory string
)

type server struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (s server) Serialize() ([]byte, error) {
	return json.Marshal(s)
}

func unserializeServer(serializedServer []byte) (server, error) {
	var s server

	err := json.Unmarshal(serializedServer, &s)

	return s, err
}

func main() {
	app := cli.NewApp()

	app.Name = "plex-cli"
	app.Usage = "Interact with your plex server and plex.tv from the command line"
	app.Version = "0.0.1"

	if dir, err := homedir.Dir(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		storeDirectory = dir
	}

	app.Commands = []cli.Command{
		{
			Name:   "test",
			Usage:  "Test your connection to your Plex Media Server",
			Action: test,
		},
		{
			Name:   "end",
			Usage:  "End a transcode session",
			Action: endTranscode,
		},
		{
			Name:   "server-info",
			Usage:  "Print info about your servers - ip, machine id, access tokens, etc",
			Action: getServersInfo,
		},
		{
			Name:   "sections",
			Usage:  "Print info about your server's sections",
			Action: getSections,
		},
		{
			Name:   "link",
			Usage:  "authorize an app (e.g. amazon fire app) with a 4 character `code`. REQUIRES a plex token",
			Action: linkApp,
		},
		{
			Name:   "library",
			Usage:  "display your libraries",
			Action: getLibraries,
		},
		{
			Name:   "request-pin",
			Usage:  "request a pin (4 character code) from plex.tv to link account to an app. Use this to recieve an id to check for an auth token",
			Action: requestPIN,
		},
		{
			Name:  "check-pin",
			Usage: "check status of pin (4 character code) from plex.tv to link account to an app. Use this to recieve an auth token. REQUIRES an id",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "poll",
					Usage: "check every second if pin is authorized",
				},
			},
			Action: checkPIN,
		},
		{
			Name:   "signin",
			Usage:  "use your username and password to receive a plex auth token",
			Action: signIn,
		},
		{
			Name:   "sessions",
			Usage:  "display info on users currently consuming media",
			Action: getSessions,
		},
		{
			Name:   "pick-server",
			Usage:  "choose a server to interact with",
			Action: pickServer,
		},
		{
			Name:   "webhooks",
			Usage:  "display webhooks associated with your account",
			Action: webhooks,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "add",
					Usage: "create a new webhook",
				},
				cli.BoolFlag{
					Name:  "delete",
					Usage: "delete a webhook",
				},
			},
		},
		{
			Name:   "search",
			Usage:  "search for media information on your server",
			Action: search,
		},
		{
			Name:   "episode",
			Usage:  "get metadata of an episode of a show",
			Action: getEpisode,
		},
		{
			Name:   "on-deck",
			Usage:  "display titles of media that is on deck",
			Action: getOnDeck,
		},
		{
			Name:   "unlock",
			Usage:  "remove lock on pid file",
			Action: unlock,
		},
		{
			Name:   "stop",
			Usage:  "stop playback on device",
			Action: stopPlayback,
		},
		{
			Name:   "account",
			Usage:  "get account info from plex.tv",
			Action: getAccountInfo,
		},
		{
			Name:   "metadata",
			Usage:  "get metadata of media on plex server",
			Action: getMetadata,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

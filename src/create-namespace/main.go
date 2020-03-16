package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	// "strings"

	"onyxia-cli/src/create-namespace/api"
	. "onyxia-cli/src/create-namespace/configuration"
	"onyxia-cli/src/create-namespace/oidc"

	"github.com/urfave/cli/v2"
)

func main() {
	ReadConfig()
	// Disable SSL validation
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	app := &cli.App{
		Version:              "v0.1",
		Name:                 "onyxia-cli",
		Usage:                "CLI for Onyxia services",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "accessToken",
				Aliases:  []string{"token"},
				EnvVars:  []string{"KC_ACCESS_TOKEN"},
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "whoami",
				Usage: "Show user's informations",
				Action: func(c *cli.Context) error {
					tokenString := c.String("accessToken")
					oidc.DisplayID(tokenString)
					return nil
				},
			},
			{
				Name:  "kub",
				Usage: "Setup kubernetes",
				Subcommands: []*cli.Command{
					{
						Name:  "create",
						Usage: "Create resources",
						Subcommands: []*cli.Command{
							{
								Name:  "namespace",
								Usage: "Create namespace",
								Flags: []cli.Flag{
									&cli.StringFlag{
										Name:     "namespaceName",
										Aliases:  []string{"name", "n"},
										Usage:    "Name of the namespace to create",
										Required: true,
									},
									&cli.StringFlag{
										Name:        "namespaceOwnerType",
										Aliases:     []string{"ownerType", "t"},
										Required:    false,
										Usage:       "The type of ownership. Can be either USER or GROUP.",
										DefaultText: "USER",
										Value:       "USER",
									},
									&cli.StringFlag{
										Name:        "namespaceOwner",
										Aliases:     []string{"owner", "o"},
										Required:    false,
										Usage:       "The owner of the namespace.",
										DefaultText: "YOU",
									},
								},
								Action: func(c *cli.Context) error {
									namespaceName := c.String("namespaceName")
									tokenString := c.String("accessToken")
									ownerType := c.String("namespaceOwnerType")
									ownerName := c.String("namespaceOwner")
									if ownerType == "USER" {
										api.CreateNamespace(tokenString, namespaceName, ownerType, "oidc"+oidc.GetID(tokenString).Preferred_username)
									} else {
										if ownerName == "" {
											fmt.Println("Veuillez choisir un groupe parmi ceux-ci :")
											oidc.DisplayGroups(oidc.GetID(tokenString), false)
										} else {
											api.CreateNamespace(tokenString, namespaceName, ownerType, "oidc"+ownerName)
										}
									}
									return nil
								},
								BashComplete: func(c *cli.Context) {
									if c.NArg() > 0 {
										return
									}
									if os.Args[len(os.Args)-2] == "-o" || os.Args[len(os.Args)-2] == "--owner" {
										oidc.DisplayGroups(oidc.GetID(c.String("accessToken")), true)
									} else if os.Args[len(os.Args)-2] == "-t" || os.Args[len(os.Args)-2] == "--ownerType" {
										for _, elem := range []string{"GROUP", "USER"} {
											fmt.Println(elem)
										}
									}
								},
							},
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

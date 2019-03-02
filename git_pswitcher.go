package main

import (
	"fmt"
	"os"

	"github.com/acrap/git_pswitcher/git"
	"github.com/acrap/git_pswitcher/profile"

	"github.com/urfave/cli"
)

var (
	name  string
	email string
	force bool
)

// our main function
func main() {
	app := cli.NewApp()
	app.Version = "0.1"
	app.Description = "git_pswitcher is an utility to keep git profiles and easily switch between them"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name",
			Value:       "",
			Usage:       "name of a git user",
			Destination: &name,
		},
		cli.StringFlag{
			Name:        "email",
			Value:       "",
			Usage:       "email of a git user",
			Destination: &email,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "get all stored profiles",
			Action: func(c *cli.Context) error {
				var profiles []profile.Profile
				var err error
				db := profile.CreateDefaultJsonFileDb()
				if profiles, err = db.GetProfiles(); err != nil {
					fmt.Printf("Can't read profiles from file: %s\n", err)
					return err
				}
				for _, elem := range profiles {
					fmt.Printf("Name <%s>, Email <%s>\n", elem.Name, elem.Email)
				}
				if len(profiles) == 0 {
					fmt.Printf("Profile database is empty\n")
				}
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "remove a profile from db by name. Specify the name with --name flag",
			Action: func(c *cli.Context) error {
				db := profile.CreateDefaultJsonFileDb()
				if err := db.RemoveProfile(name); err != nil {
					fmt.Printf("%v\n", err)
					return err
				}
				fmt.Printf("Profile <%s> has been removed\n", name)
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a profile to the list. Use --name and --email to set values",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "force", Destination: &force},
			},
			Action: func(c *cli.Context) error {
				var err error
				prof := &profile.Profile{}
				if err = prof.SetEmail(email); err != nil {
					fmt.Printf("%v\n", err)
					return err
				}
				prof.SetName(name)
				db := profile.CreateDefaultJsonFileDb()
				if err = db.AddProfile(*prof, force); err != nil {
					fmt.Printf("Can't add new profile: %s\n", err)
					return err
				}
				println("Added new profile")
				return nil
			},
		},
		{
			Name:    "switch",
			Aliases: []string{"s"},
			Usage:   "switch to a profile. Just set the name with --name",
			Action: func(c *cli.Context) error {
				var err error
				var profiles []profile.Profile
				var chosenOne *profile.Profile
				db := profile.CreateDefaultJsonFileDb()

				if profiles, err = db.GetProfiles(); err != nil {
					fmt.Printf("can't get profiles %v\n", err)
					return err
				}
				for _, elem := range profiles {
					if elem.Name == name {
						chosenOne = &elem
						break
					}
				}
				if chosenOne == nil {
					fmt.Printf("can't switch to profile %v\n", fmt.Errorf("can't find a profile"))
				}
				if err = git.SwitchToProfile(*chosenOne); err != nil {
					fmt.Printf("can't switch to profile %v\n", err)
					return err
				}
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Use --help to get information about usage")
		return nil
	}

	app.Run(os.Args)
}

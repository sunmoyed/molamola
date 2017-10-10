package main

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"

	"github.com/sunmoyed/molamola/server/pkg/log"
	"github.com/sunmoyed/molamola/server/pkg/user"
)

var logger = log.DefaultLogger

const (
	dataFlag string = "data"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Mola Mola CLI"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  dataFlag,
			Value: "data",
			Usage: "Data directory",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "User",
			Action:  userAct,
			Subcommands: []cli.Command{
				{
					Name:   "add",
					Usage:  "Add a user",
					Action: userAddAct,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "password",
							Value: "",
							Usage: "Password",
						},
					},
				},
				{
					Name:   "remove",
					Usage:  "Remove a user",
					Action: userRemoveAct,
				},
			},
		},
	}

	app.RunAndExitOnError()
}

func userAct(c *cli.Context) error {
	datadir := c.GlobalString(dataFlag)
	us, usErr := user.NewUserState(datadir)
	if usErr != nil {
		return usErr
	}
	users, usersErr := us.GetUsers()
	if usersErr != nil {
		return usersErr
	}
	for uname, _ := range users {
		fmt.Println(uname)
	}
	return nil
}

func userAddAct(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("invalid number of arguments")
	}
	username := c.Args().Get(0)
	logger.Printf("adding user %s", username)

	datadir := c.GlobalString(dataFlag)
	logger.Printf("datadir %s", datadir)

	us, usErr := user.NewUserState(datadir)
	if usErr != nil {
		return usErr
	}
	return us.AddUser(username, c.String("password"))
}

func userRemoveAct(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("invalid number of arguments")
	}
	username := c.Args().Get(0)
	logger.Printf("deleting user %s", username)

	datadir := c.GlobalString(dataFlag)
	logger.Printf("datadir %s", datadir)

	us, usErr := user.NewUserState(datadir)
	if usErr != nil {
		return usErr
	}
	return us.DelUser(username)
}

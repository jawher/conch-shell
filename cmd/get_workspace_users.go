package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mkideal/cli"
	"github.com/olekukonko/tablewriter"
	"os"
)

type getWorkspaceUsersArgs struct {
	cli.Helper
	Id string `cli:"*id,uuid" usage:"ID of the workspace (required)"`
}

var GetWorkspaceUsersCmd = &cli.Command{
	Name: "get_workspace_users",
	Desc: "Get a list of users for the given workspace ID",
	Argv: func() interface{} { return new(getWorkspaceUsersArgs) },
	Fn: func(ctx *cli.Context) error {
		args, _, api, err := GetStarted(&getWorkspaceUsersArgs{}, ctx)

		if err != nil {
			return err
		}

		argv := args.Local.(*getWorkspaceUsersArgs)

		users, err := api.GetWorkspaceUsers(argv.Id)
		if err != nil {
			return err
		}

		if args.Global.JSON == true {

			j, err := json.Marshal(users)

			if err != nil {
				return err
			}

			fmt.Println(string(j))

		} else {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Name", "Email", "Role"})

			table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
			table.SetCenterSeparator("|")

			for _, u := range users {
				table.Append([]string{u.Name, u.Email, u.Role})
			}

			table.Render()
		}
		return nil
	},
}
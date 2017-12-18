package cli

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/spf13/cobra"
)

func newServerResetpasswordCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "reset-password [flags] <id>",
		Short:            "Reset password of a server",
		Args:             cobra.ExactArgs(1),
		TraverseChildren: true,
		RunE:             cli.wrap(runServerResetpassword),
	}
	return cmd
}

func runServerResetpassword(cli *CLI, cmd *cobra.Command, args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("invalid server id")
	}

	ctx := context.Background()
	server := &hcloud.Server{ID: id}
	result, _, err := cli.Client().Server.ResetPassword(ctx, server)
	if err != nil {
		return err
	}
	if err := <-waitAction(ctx, cli.Client(), result.Action); err != nil {
		return err
	}
	fmt.Printf("Password of server %d reset to: %s\n", id, result.RootPassword)
	return nil
}

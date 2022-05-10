package dataimports

import (
	"fmt"
	"github.com/planetscale/cli/internal/cmdutil"
	ps "github.com/planetscale/planetscale-go/planetscale"
	"github.com/spf13/cobra"
)

func DetachExternalDatabaseCmd(ch *cmdutil.Helper) *cobra.Command {
	var flags struct {
		name string
	}

	detachExternalDatabaseReq := &ps.DetachExternalDatabaseRequest{}

	cmd := &cobra.Command{
		Use:     "make-primary [options]",
		Short:   "mark PlanetScale's database as the primary, and the external database as replica",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			detachExternalDatabaseReq.Organization = ch.Config.Organization
			detachExternalDatabaseReq.DatabaseName = flags.name

			client, err := ch.Client()
			if err != nil {
				return err
			}

			_, err = client.DataImports.DetachExternalDatabase(ctx, detachExternalDatabaseReq)
			if err != nil {
				switch cmdutil.ErrCode(err) {
				case ps.ErrNotFound:
					return fmt.Errorf("unable to detach external database for PlanetScale database %s", flags.name)
				default:
					return cmdutil.HandleError(err)
				}
			}

			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&flags.name, "name", "", "")

	return cmd
}

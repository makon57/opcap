package cmd

import (
	"fmt"
	// "strings"
	// "text/tabwriter"

	"github.com/opdev/opcap/internal/bundles"
	"github.com/spf13/cobra"
)

func bundlesCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "bundles",
		Short: "List the bundles of each oeprator",
		RunE: func(cmd *cobra.Command, args []string) error {
			bundles, err := bundles.ListBundles()
			if err != nil {
				return err
			}

			for _, bundle := range bundles {
				fmt.Println(bundle)
			}
			return nil
		},
	}
	return &cmd
}

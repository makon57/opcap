package cmd

import (
	"context"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/opdev/opcap/internal/bundle"

	"github.com/spf13/cobra"
)

var listBundlesFlags struct {
	bundlesDir  string
	bundlesRepo string
}

func listBundlesCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "bundles",
		Short: "List all bundles and versions",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := listBundles(cmd.Context(), cmd.OutOrStdout())
			if err != nil {
				return err
			}

			return nil
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&listBundlesFlags.bundlesDir, "bundles-path", "",
		"specifies the source directory with bundles")
	flags.StringVar(&listBundlesFlags.bundlesRepo, "bundles-repo", "",
		"specifies the repo in which to clone bundles from")

	return &cmd
}

func listBundles(ctx context.Context, out io.Writer) error {
	bundles, err := bundle.List(listBundlesFlags.bundlesRepo, listBundlesFlags.bundlesDir)
	if err != nil {
		return err
	}

	headings := "Package Name\tStarting CSV\tVersion\tDefault Channel\tOcpVersions"
	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, headings)
	for _, bundle := range bundles {
		packageInfo := []string{bundle.StartingCSV, bundle.PackageName, bundle.Channel, bundle.Version, bundle.OcpVersions}
		fmt.Fprintln(w, strings.Join(packageInfo, "\t"))
	}
	w.Flush()

	return err
}

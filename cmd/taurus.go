package cmd

import (
	"os"

	"github.com/croomes/fiord/cmd/formatter"
	"github.com/croomes/fiord/fio"
	"github.com/spf13/cobra"
)

type taurusOptions struct {
	reportURL string
}

func init() {
	RootCmd.AddCommand(newTaurusCmd())
}

func newTaurusCmd() *cobra.Command {

	opts := taurusOptions{}

	cmd := &cobra.Command{
		Use:   "taurus",
		Short: "Taurus final-status output",
		Long: `Taurus output can be used to produce reports that are compatible
with the Jenkins Performance Plugin.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTaurus(opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.reportURL, "report-url", "", "URL to full report")

	return cmd
}

func runTaurus(opts taurusOptions) error {
	report, err := fio.Decode(os.Stdin)
	if err != nil {
		return err
	}

	ctx := formatter.Context{
		Output: os.Stdout,
		Format: formatter.NewTaurusFormat(),
	}

	return formatter.TaurusWrite(ctx, report, opts.reportURL)
}

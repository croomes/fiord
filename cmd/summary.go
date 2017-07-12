package cmd

import (
	"os"

	"github.com/croomes/fiord/cmd/formatter"
	"github.com/croomes/fiord/fio"
	"github.com/spf13/cobra"
)

type summaryOptions struct {
	input  string
	quiet  bool
	format string
}

func init() {
	RootCmd.AddCommand(newSummaryCmd())
}

func newSummaryCmd() *cobra.Command {

	opts := summaryOptions{}

	cmd := &cobra.Command{
		Use:   "summary",
		Short: "Basic job summary",
		Long: `Produces a Human-readable summary.  The format may be set to
table, raw, or a custom Go template.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSummary(opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.input, "input", "i", "", "Input file if not stdin")
	flags.BoolVarP(&opts.quiet, "quiet", "q", false, "Only display IOPS")
	flags.StringVar(&opts.format, "format", "table", "Pretty-print report using a Go template")

	return cmd
}

func runSummary(opts summaryOptions) error {

	input := os.Stdin
	if opts.input != "" {
		var err error
		input, err = os.Open(opts.input)
		if err != nil {
			return err
		}
	}

	report, err := fio.Decode(input)
	if err != nil {
		return err
	}

	ctx := formatter.Context{
		Output: os.Stdout,
		Format: formatter.NewSummaryFormat(opts.format, opts.quiet),
	}

	return formatter.SummaryWrite(ctx, report)
}

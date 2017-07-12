package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/croomes/fiord/backends/influxdb"
	"github.com/croomes/fiord/cmd/formatter"
	"github.com/croomes/fiord/fio"
)

const (
	DefaultInfluxdbURI = "http://172.28.128.3:8086"
	DefaultInfluxdbDB  = "fio"
)

type influxdbOptions struct {
	uri    string
	db     string
	tags   string
	quiet  bool
	format string
}

func init() {
	RootCmd.AddCommand(newInfluxdbCmd())
}

func newInfluxdbCmd() *cobra.Command {

	opts := influxdbOptions{}

	cmd := &cobra.Command{
		Use:   "influxdb",
		Short: "Publish report to InfluxDB",
		Long:  `Sends report to InfluxDB`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInfluxdb(opts)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.quiet, "quiet", "q", false, "Only display IOPS")
	flags.StringVar(&opts.format, "format", "influxdb", "Pretty-print report using a Go template")
	flags.StringVar(&opts.uri, "uri", DefaultInfluxdbURI, "<scheme>://<ip>:<port> of InfluxDB API")
	flags.StringVar(&opts.db, "db", DefaultInfluxdbDB, "InfluxDB database")
	flags.StringVar(&opts.tags, "tags", "", "Additional tags in key=value format, comma-separated")

	return cmd
}

func runInfluxdb(opts influxdbOptions) error {
	report, err := fio.Decode(os.Stdin)
	if err != nil {
		return err
	}

	idb := influxdb.Init(opts.uri, opts.db)

	ctx := formatter.Context{
		Output: idb,
		Format: formatter.NewInfluxdbFormat(opts.format, opts.quiet),
	}

	if err := formatter.InfluxdbWrite(ctx, report, opts.tags); err != nil {
		return err
	}

	return idb.Flush()
}

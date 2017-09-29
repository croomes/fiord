package formatter

import (
	"fmt"

	"github.com/croomes/fiord/fio"
)

const (
	InfluxDBFormatKey          = "influxdb"
	defaultInfluxdbQuietFormat = "{{.IOPS}}"
	defaultInfluxdbFormat      = "{{.IOPS}}{{.Bandwidth}}"
)

// NewInfluxdbFormat returns a format for use with a influxdb Context
func NewInfluxdbFormat(source string, quiet bool) Format {
	if source == InfluxDBFormatKey {
		if quiet {
			return defaultInfluxdbQuietFormat
		}
		return defaultInfluxdbFormat
	}
	return Format(source)
}

// InfluxdbWrite writes formatted summaries using the Context
func InfluxdbWrite(ctx Context, report fio.Report, tags string) error {
	if tags != "" {
		tags = "," + tags
	}
	render := func(format func(subContext subContext) error) error {
		for _, job := range report.Jobs {
			// InfluxDB wants nanosecond timestamps, convert here
			if err := format(&influxdbContext{ts: report.Timestamp * 1000000000, tags: tags, v: job}); err != nil {
				return err
			}
		}
		return nil
	}
	return ctx.Write(&influxdbContext{}, render)
}

type influxdbContext struct {
	HeaderContext
	tags string
	ts   int64
	v    fio.Job
}

func (c *influxdbContext) IOPS() string {
	return fmt.Sprintf("iops,op=read,job=%s%s value=%.2f %d\n", c.v.JobName, c.tags, c.v.Read.IOPS, c.ts) +
		fmt.Sprintf("iops,op=write,job=%s%s value=%.2f %d\n", c.v.JobName, c.tags, c.v.Write.IOPS, c.ts) +
		fmt.Sprintf("iops,op=trim,job=%s%s value=%.2f %d\n", c.v.JobName, c.tags, c.v.Trim.IOPS, c.ts)
}

func (c *influxdbContext) Bandwidth() string {
	return fmt.Sprintf("throughput,op=read,job=%s%s value=%d %d\n", c.v.JobName, c.tags, c.v.Read.Bandwidth, c.ts) +
		fmt.Sprintf("throughput,op=write,job=%s%s value=%d %d\n", c.v.JobName, c.tags, c.v.Write.Bandwidth, c.ts) +
		fmt.Sprintf("throughput,op=trim,job=%s%s value=%d %d\n", c.v.JobName, c.tags, c.v.Trim.Bandwidth, c.ts)
}

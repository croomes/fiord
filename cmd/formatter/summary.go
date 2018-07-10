package formatter

import (
	"fmt"

	"github.com/croomes/fiord/fio"
)

const (
	defaultSummaryQuietFormat = "{{.IOPS}}"
	defaultSummaryTableFormat = "table {{.Name}}\t{{.IOPS}}\t{{.Throughput}}"
	defaultSummaryCSVFormat   = "csv {{.Name}},{{.ReadIOPS}},{{.ReadThroughput}},{{.WriteIOPS}},{{.WriteThroughput}}"

	summaryNameHeader            = "JOB NAME"
	summaryIOPSHeader            = "IO/S R/W/T"
	summaryReadIOPSHeader        = "READ IOPS"
	summaryWriteIOPSHeader       = "WRITE IOPS"
	summaryThroughputHeader      = "MB/S R/W/T"
	summaryReadThroughputHeader  = "READ MB/S"
	summaryWriteThroughputHeader = "WRITE MB/S"
)

// NewSummaryFormat returns a format for use with a summary Context
func NewSummaryFormat(source string, quiet bool) Format {
	switch source {
	case TableFormatKey:
		if quiet {
			return defaultSummaryQuietFormat
		}
		return defaultSummaryTableFormat
	case CSVFormatKey:
		return defaultSummaryCSVFormat
	case RawFormatKey:
		if quiet {
			return `name: {{.Name}}`
		}
		return `name: {{.Name}}\n`
	}
	return Format(source)
}

// SummaryWrite writes formatted summaries using the Context
func SummaryWrite(ctx Context, summary fio.Report) error {
	render := func(format func(subContext subContext) error) error {
		for _, job := range summary.Jobs {
			if err := format(&summaryContext{v: job}); err != nil {
				return err
			}
		}
		return nil
	}
	return ctx.Write(&summaryContext{}, render)
}

type summaryContext struct {
	HeaderContext
	v fio.Job
}

func (c *summaryContext) Name() string {
	c.AddHeader(summaryNameHeader)
	return c.v.JobName
}

func (c *summaryContext) IOPS() string {
	c.AddHeader(summaryIOPSHeader)
	return fmt.Sprintf("%.2f/%.2f/%.2f", c.v.Read.IOPS, c.v.Write.IOPS, c.v.Trim.IOPS)
}

func (c *summaryContext) ReadIOPS() string {
	c.AddHeader(summaryReadIOPSHeader)
	return fmt.Sprintf("%.2f", c.v.Read.IOPS)
}

func (c *summaryContext) WriteIOPS() string {
	c.AddHeader(summaryWriteIOPSHeader)
	return fmt.Sprintf("%.2f", c.v.Write.IOPS)
}

func (c *summaryContext) Throughput() string {
	c.AddHeader(summaryThroughputHeader)
	return fmt.Sprintf("%.2f/%.2f/%.2f", KtoMB(c.v.Read.Bandwidth), KtoMB(c.v.Write.Bandwidth), KtoMB(c.v.Trim.Bandwidth))
}

func (c *summaryContext) ReadThroughput() string {
	c.AddHeader(summaryReadThroughputHeader)
	return fmt.Sprintf("%.2f", KtoMB(c.v.Read.Bandwidth))
}

func (c *summaryContext) WriteThroughput() string {
	c.AddHeader(summaryWriteThroughputHeader)
	return fmt.Sprintf("%.2f", KtoMB(c.v.Write.Bandwidth))
}

package formatter

import (
	"fmt"

	"github.com/croomes/fiord/fio"
)

const (
	defaultSummaryQuietFormat = "{{.IOPS}}"
	defaultSummaryTableFormat = "table {{.Name}}\t{{.IOPS}}\t{{.Bandwidth}}"

	summaryNameHeader      = "JOB NAME"
	summaryIOPSHeader      = "IO/S R/W/T"
	summaryBandwidthHeader = "MB/S R/W/T"
)

// NewSummaryFormat returns a format for use with a summary Context
func NewSummaryFormat(source string, quiet bool) Format {
	switch source {
	case TableFormatKey:
		if quiet {
			return defaultSummaryQuietFormat
		}
		return defaultSummaryTableFormat
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

func (c *summaryContext) Bandwidth() string {
	c.AddHeader(summaryBandwidthHeader)
	return fmt.Sprintf("%.2f/%.2f/%.2f", KtoMB(c.v.Read.Bandwidth), KtoMB(c.v.Write.Bandwidth), KtoMB(c.v.Trim.Bandwidth))
}

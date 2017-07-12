package formatter

import (
	"encoding/xml"
	"fmt"

	"github.com/croomes/fiord/backends/taurus"
	"github.com/croomes/fiord/fio"
)

const (
	TaurusFormatKey          = "taurus"
	defaultTaurusQuietFormat = "{{.Group}}"
	defaultTaurusFormat      = "{{.Group}}"
)

// NewTaurusFormat returns a format for use with a summary Context
func NewTaurusFormat() Format {
	return TaurusFormatKey
}

// TaurusWrite writes formatted summaries using the Context
func TaurusWrite(ctx Context, report fio.Report, reportURL string) error {
	t, err := TaurusReport(report, reportURL)
	if err != nil {
		return err
	}

	enc := xml.NewEncoder(ctx.Output)
	enc.Indent("  ", "    ")
	if err := enc.Encode(t); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	return nil
}

func TaurusReport(r fio.Report, reportURL string) (*taurus.Report, error) {
	t := &taurus.Report{
		FinalStatus: &taurus.FinalStatus{
			ReportURL: &taurus.ReportURL{Text: reportURL},
		},
	}
	for _, job := range r.Jobs {
		group, err := TaurusGroup(job)
		if err != nil {
			fmt.Printf("ERROR: parsing job: %v", err)
		}
		t.FinalStatus.Group = append(t.FinalStatus.Group, group)
	}
	return t, nil
}

func TaurusGroup(j fio.Job) (*taurus.Group, error) {
	g := &taurus.Group{
		AttrLabel: j.JobName + "_read",
		Avg_ct: &taurus.Avg_ct{
			AttrValue: fmt.Sprintf("%.2f", j.Read.SubmitLatency.Mean),
			Name:      &taurus.Name{Text: "avg_ct"},
			Value:     &taurus.Value{Text: fmt.Sprintf("%.2f", j.Read.SubmitLatency.Mean)},
		},
		Avg_rt: &taurus.Avg_rt{
			AttrValue: fmt.Sprintf("%.2f", j.Read.CompletionLatency.Mean),
			Name:      &taurus.Name{Text: "avg_rt"},
			Value:     &taurus.Value{Text: fmt.Sprintf("%.2f", j.Read.CompletionLatency.Mean)},
		},
		Stdev_rt: &taurus.Stdev_rt{
			AttrValue: fmt.Sprintf("%.2f", j.Read.CompletionLatency.StdDev),
			Name:      &taurus.Name{Text: "stdev_rt"},
			Value:     &taurus.Value{Text: fmt.Sprintf("%.2f", j.Read.CompletionLatency.StdDev)},
		},
		Avg_lt: &taurus.Avg_lt{
			AttrValue: fmt.Sprintf("%.2f", j.Read.Latency.Mean),
			Name:      &taurus.Name{Text: "avg_lt"},
			Value:     &taurus.Value{Text: fmt.Sprintf("%.2f", j.Read.Latency.Mean)},
		},
		Bytes: &taurus.Bytes{
			AttrValue: fmt.Sprintf("%d", j.Read.IOBytes),
			Name:      &taurus.Name{Text: "bytes"},
			Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.IOBytes)},
		},
		Throughput: &taurus.Throughput{
			AttrValue: fmt.Sprintf("%.2f", j.Read.BandwidthMean),
			Name:      &taurus.Name{Text: "throughput"},
			Value:     &taurus.Value{Text: fmt.Sprintf("%.2f", j.Read.BandwidthMean)},
		},
		Concurrency: &taurus.Concurrency{
			AttrValue: fmt.Sprintf("%.2f", j.Read.IOPS),
			Name:      &taurus.Name{Text: "concurrency"},
			Value:     &taurus.Value{Text: fmt.Sprintf("%.2f", j.Read.IOPS)},
		},
		Perc: []*taurus.Perc{
			&taurus.Perc{
				AttrParam: "1.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P1),
				Name:      &taurus.Name{Text: "perc/1.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P1)},
			},
			&taurus.Perc{
				AttrParam: "5.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P5),
				Name:      &taurus.Name{Text: "perc/5.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P5)},
			},
			&taurus.Perc{
				AttrParam: "10.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P10),
				Name:      &taurus.Name{Text: "perc/10.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P10)},
			},
			&taurus.Perc{
				AttrParam: "20.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P20),
				Name:      &taurus.Name{Text: "perc/20.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P20)},
			},
			&taurus.Perc{
				AttrParam: "30.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P30),
				Name:      &taurus.Name{Text: "perc/30.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P30)},
			},
			&taurus.Perc{
				AttrParam: "40.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P40),
				Name:      &taurus.Name{Text: "perc/40.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P40)},
			},
			&taurus.Perc{
				AttrParam: "50.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P50),
				Name:      &taurus.Name{Text: "perc/50.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P50)},
			},
			&taurus.Perc{
				AttrParam: "60.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P60),
				Name:      &taurus.Name{Text: "perc/60.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P60)},
			},
			&taurus.Perc{
				AttrParam: "70.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P70),
				Name:      &taurus.Name{Text: "perc/70.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P70)},
			},
			&taurus.Perc{
				AttrParam: "80.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P80),
				Name:      &taurus.Name{Text: "perc/80.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P80)},
			},
			&taurus.Perc{
				AttrParam: "90.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P90),
				Name:      &taurus.Name{Text: "perc/90.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P90)},
			},
			&taurus.Perc{
				AttrParam: "95.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P95),
				Name:      &taurus.Name{Text: "perc/95.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P95)},
			},
			&taurus.Perc{
				AttrParam: "99.0",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P99),
				Name:      &taurus.Name{Text: "perc/99.0"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P99)},
			},
			&taurus.Perc{
				AttrParam: "99.5",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P9950),
				Name:      &taurus.Name{Text: "perc/99.5"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P9950)},
			},
			&taurus.Perc{
				AttrParam: "99.95",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P9995),
				Name:      &taurus.Name{Text: "perc/99.95"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P9995)},
			},
			&taurus.Perc{
				AttrParam: "99.99",
				AttrValue: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P9999),
				Name:      &taurus.Name{Text: "perc/99.99"},
				Value:     &taurus.Value{Text: fmt.Sprintf("%d", j.Read.CompletionLatency.Percentile.P9999)},
			},
		},
	}
	return g, nil
}

// group, err := c.ParseGroup(job)
// if err != nil {
// 	fmt.Printf("ERROR: parsing job: %v", err)
// }
// report.Group = append(report.Group, group)
// // 		fmt.Sprintf("iops,op=write,job=%s value=%.2f %d\n", c.v.JobName, c.v.Write.IOPS, c.ts) +
// // 		fmt.Sprintf("iops,op=trim,job=%s value=%.2f %d\n", c.v.JobName, c.v.Trim.IOPS, c.ts)

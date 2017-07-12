package fio

// Report is the output of fio when run with `--output-format=json`
type Report struct {
	FIOVersion string `json:"fio version"`
	Timestamp  int    `json:"timestamp"`
	Time       string `json:"time"`
	Jobs       []Job  `json:"jobs"`
}

// Job describes a test run
type Job struct {
	JobName string    `json:"jobname"`
	GroupID int       `json:"groupid"`
	Error   int       `json:"error"`
	ETA     int       `json:"eta"`
	Elapsed int       `json:"elapsed"`
	Read    *IOReport `json:"read"`
	Write   *IOReport `json:"write"`
	Trim    *IOReport `json:"trim"`
	UsrCPU  float32   `json:"usr_cpu"`
	SysCPU  float32   `json:"sys_cpu"`
	Context int       `json:"ctx"`
}

// IOReport describes stats for differnet IO profiles
type IOReport struct {
	IOBytes            int               `json:"io_bytes"`
	Bandwidth          int               `json:"bw"`
	IOPS               float32           `json:"iops"`
	Runtime            int               `json:"runtime"`
	TotalIOs           int               `json:"total_ios"`
	ShortIOs           int               `json:"short_ios"`
	DropIOs            int               `json:"drop_ios"`
	SubmitLatency      SubmitLatency     `json:"slat"`
	CompletionLatency  CompletionLatency `json:"clat"`
	Latency            Summary           `json:"lat"`
	BandwidthMin       int               `json:"bw_min"`
	BandwidthMax       int               `json:"bw_max"`
	BandwidthAggregate float32           `json:"bw_agg"`
	BandwidthMean      float32           `json:"bw_mean"`
	BandwidthDeviation float32           `json:"bw_dev"`
}

// SubmitLatency is the amount of time it took to submit IO to the kernel for
// processing.
type SubmitLatency struct {
	Summary
}

// CompletionLatency is the amount of time that passes between submission to the
// kernel and when the IO is complete, not including submission latency.
type CompletionLatency struct {
	Summary
	Percentile Percentile `json:"percentile"`
}

// Summary is a high-level measurement..
type Summary struct {
	Min    int     `json:"min"`
	Max    int     `json:"max"`
	Mean   float32 `json:"mean"`
	StdDev float32 `json:"stddev"`
}

// Percentile metrics
type Percentile struct {
	P1    int `json:"1.000000"`
	P5    int `json:"5.000000"`
	P10   int `json:"10.000000"`
	P20   int `json:"20.000000"`
	P30   int `json:"30.000000"`
	P40   int `json:"40.000000"`
	P50   int `json:"50.000000"`
	P60   int `json:"60.000000"`
	P70   int `json:"70.000000"`
	P80   int `json:"80.000000"`
	P90   int `json:"90.000000"`
	P95   int `json:"95.000000"`
	P99   int `json:"99.000000"`
	P9950 int `json:"99.500000"`
	P9990 int `json:"99.900000"`
	P9995 int `json:"99.950000"`
	P9999 int `json:"99.990000"`
}

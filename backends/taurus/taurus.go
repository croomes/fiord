package taurus

type Report struct {
	FinalStatus *FinalStatus `xml:" FinalStatus,omitempty" json:"FinalStatus,omitempty"`
}

type FinalStatus struct {
	Group     []*Group   `xml:" Group,omitempty" json:"Group,omitempty"`
	ReportURL *ReportURL `xml:" ReportURL,omitempty" json:"ReportURL,omitempty"`
}

type Group struct {
	AttrLabel   string       `xml:" label,attr"  json:",omitempty"`
	Avg_ct      *Avg_ct      `xml:" avg_ct,omitempty" json:"avg_ct,omitempty"`
	Avg_lt      *Avg_lt      `xml:" avg_lt,omitempty" json:"avg_lt,omitempty"`
	Avg_rt      *Avg_rt      `xml:" avg_rt,omitempty" json:"avg_rt,omitempty"`
	Bytes       *Bytes       `xml:" bytes,omitempty" json:"bytes,omitempty"`
	Concurrency *Concurrency `xml:" concurrency,omitempty" json:"concurrency,omitempty"`
	Fail        *Fail        `xml:" fail,omitempty" json:"fail,omitempty"`
	Perc        []*Perc      `xml:" perc,omitempty" json:"perc,omitempty"`
	Rc          []*Rc        `xml:" rc,omitempty" json:"rc,omitempty"`
	Stdev_rt    *Stdev_rt    `xml:" stdev_rt,omitempty" json:"stdev_rt,omitempty"`
	Succ        *Succ        `xml:" succ,omitempty" json:"succ,omitempty"`
	Throughput  *Throughput  `xml:" throughput,omitempty" json:"throughput,omitempty"`
}

type ReportURL struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Avg_ct struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	Name      *Name  `xml:" name,omitempty" json:"name,omitempty"`
	Value     *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Avg_lt struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	Name      *Name  `xml:" name,omitempty" json:"name,omitempty"`
	Value     *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Avg_rt struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	Name      *Name  `xml:" name,omitempty" json:"name,omitempty"`
	Value     *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Bytes struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	Name      *Name  `xml:" name,omitempty" json:"name,omitempty"`
	Value     *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Concurrency struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	Name      *Name  `xml:" name,omitempty" json:"name,omitempty"`
	Value     *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Fail struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	Name      *Name  `xml:" name,omitempty" json:"name,omitempty"`
	Value     *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Perc struct {
	AttrParam string `xml:" param,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	Name      *Name  `xml:" name,omitempty" json:"name,omitempty"`
	Value     *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Rc struct {
	AttrParam string `xml:" param,attr"  json:",omitempty"`
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	Name      *Name  `xml:" name,omitempty" json:"name,omitempty"`
	Value     *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Stdev_rt struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	Name      *Name  `xml:" name,omitempty" json:"name,omitempty"`
	Value     *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Succ struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	Name      *Name  `xml:" name,omitempty" json:"name,omitempty"`
	Value     *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Throughput struct {
	AttrValue string `xml:" value,attr"  json:",omitempty"`
	Name      *Name  `xml:" name,omitempty" json:"name,omitempty"`
	Value     *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}

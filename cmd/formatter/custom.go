package formatter

import (
	"strings"
)

type subContext interface {
	TabHeader() string
	CSVHeader() string
	AddHeader(header string)
}

// HeaderContext provides the subContext interface for managing headers
type HeaderContext struct {
	header []string
}

// TabHeader returns the header as a tab-delimited string
func (c *HeaderContext) TabHeader() string {
	if c.header == nil {
		return ""
	}
	return strings.Join(c.header, "\t")
}

// CSVHeader returns the header as a comma-delimited string
func (c *HeaderContext) CSVHeader() string {
	if c.header == nil {
		return ""
	}
	return strings.Join(c.header, ",")
}

// AddHeader adds another column to the header
func (c *HeaderContext) AddHeader(header string) {
	if c.header == nil {
		c.header = []string{}
	}
	c.header = append(c.header, strings.ToUpper(header))
}

func stripNamePrefix(ss []string) []string {
	sss := make([]string, len(ss))
	for i, s := range ss {
		sss[i] = s[1:]
	}

	return sss
}

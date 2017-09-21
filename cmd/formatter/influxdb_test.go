package formatter

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/croomes/fiord/cmd/formatter"
	"github.com/croomes/fiord/fio"
)

func TestInfluxdbWrite(t *testing.T) {
	input, err := os.Open("../../data/input.json")
	if err != nil {
		t.Fatalf("Can't read test input file: %v", err)
	}

	report, err := fio.Decode(input)
	if err != nil {
		t.Fatalf("Can't decode test input file: %v", err)
	}

	out := bytes.NewBufferString("")

	ctx := formatter.Context{
		Output: out,
		Format: formatter.NewInfluxdbFormat(InfluxDBFormatKey, false),
	}

	if err := formatter.InfluxdbWrite(ctx, report, ""); err != nil {
		t.Errorf("Error formatting as influxdb: %v", err)
	}

	// Debug
	// fmt.Printf("%s\n", out.String())

	ref, err := ioutil.ReadFile("../../data/influxdb.out")
	if err != nil {
		t.Fatalf("Can't read reference output data: %v", err)
	}

	if strings.TrimSpace(string(ref)) != strings.TrimSpace(out.String()) {
		t.Errorf("Output did not match expected response")
	}
}

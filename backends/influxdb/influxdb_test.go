package influxdb

import (
	"fmt"
	"testing"

	"github.com/croomes/fiord/testutil"
)

func TestInit(t *testing.T) {
	i := Init("http://1.2.3.4", "db")

	if i.uri != "http://1.2.3.4" {
		t.Errorf("expected %s, got %s", "http://1.2.3.4", i.uri)
	}
	if i.db != "db" {
		t.Errorf("expected %s, got %s", "db", i.db)
	}
}

func TestDefaultInit(t *testing.T) {
	i := Init("", "")

	if i.uri != "" {
		t.Errorf("expected %s, got %s", "", i.uri)
	}
	if i.db != "" {
		t.Errorf("expected %s, got %s", "", i.db)
	}
}

func TestWrite(t *testing.T) {

	i := Init("", "")

	i.Write([]byte("buffer data"))

	if i.buffer.String() != "buffer data" {
		t.Errorf("expected buffer data to match %q", "buffer data")
	}
}

func TestFlush(t *testing.T) {

	body := "test data"
	port := testutil.GetPort()

	i := Init(fmt.Sprintf("http://127.0.0.1:%d", port), "")
	i.buffer.Write([]byte(body))

	_, teardown := testutil.NewTestingServer(port, 200, body)
	defer teardown()

	err := i.Flush()
	if err != nil {
		t.Fatalf("error during flush: %v", err)
	}
}

package influxdb

import (
	"bytes"
	"fmt"
	"net/http"
)

type InfluxDB struct {
	uri    string
	db     string
	buffer *bytes.Buffer
}

func Init(uri, db string) *InfluxDB {
	return &InfluxDB{
		uri:    uri,
		db:     db,
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func (i *InfluxDB) Write(data []byte) (int, error) {
	return i.buffer.Write(data)
}

func (i *InfluxDB) Flush() error {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/write?db=%s", i.uri, i.db), i.buffer)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil
	// }
	// fmt.Printf("Response:\n%s\n-------\n", body)

	return nil
}

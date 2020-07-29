package main

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func main() {
	// create new client with default option for server url authenticate by token
	client := influxdb2.NewClient("http://localhost:8086", "my-token")
	// // user blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking("my-org", "mydb")

	// create point using fluent style
	p := influxdb2.NewPointWithMeasurement("test").
		AddTag("test_id", "test123").
		AddField("value", 200.0).
		SetTime(time.Now())
	writeAPI.WritePoint(context.Background(), p)

	// Or write directly line protocol
	// line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%v", 23.5, 45)
	// writeAPI.WriteRecord(context.Background(), line)

	// get query client
	queryAPI := client.QueryAPI("my-org")
	// get parser flux query result
	result, err := queryAPI.Query(context.Background(), `from(bucket:"mydb")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "test")`)
	if err == nil {
		// Use Next() to iterate over query result lines
		for result.Next() {
			// Observe when there is new grouping key producing new table
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// read result
			fmt.Printf("row: %s\n", result.Record().String())
		}
		if result.Err() != nil {
			fmt.Printf("Query error: %s\n", result.Err().Error())
		}
	} else {
		fmt.Println("error in api query: ", err)
	}
	// Ensures background processes finishes
	client.Close()
}

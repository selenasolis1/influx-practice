# influx-practice
This repo is intended to be used as a template for an influxdb server linked to chronograf with a go client. The go client utilizes the [influxdb-client-go](github.com/influxdata/influxdb-client-go) package.

## influxdb
information regarding influxdb terminology and concepts can be found here...
[influxdb key concepts](https://docs.influxdata.com/influxdb/v1.8/concepts/key_concepts/) 

## Chronograf


## How to run
```
$ docker-compose up
$ cd basic
$ go run main.go
```
The main program is set up to demonstrate a write and a query to influxdb.

## Additional Information
There are many ways to write and query data to influxdb. 

#### curl
Here are a few useful example POST and GET requests to be to made to influxdb server.
```
curl -i -X POST "http://localhost:8086/write?db=mydb" --data-binary "test,test_id=test123 value=100"
```
```
curl -G "http://localhost:8086/query?pretty=true" --data-urlencode "db=mydb" --data-urlencode "q=SELECT * FROM test"
```
Click for the full [API reference](https://docs.influxdata.com/influxdb/v1.8/tools/api/)

#### influxQL
```
$ docker exec -it <containerID> bash
# influx
> <InfluxQLCommands>
```

[InfluxQL reference](https://docs.influxdata.com/influxdb/v1.8/query_language/spec/)

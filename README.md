# elog
Publish log to elasticsearch

```go
/*
* Initial log
* param1: elasticseach address
* param2: index
* param3: doc type
* param4: application name
*/
elog.Init("http://127.0.0.1:9200", "log", "doc", "web.app")

/*
* Panic
* application will be exit 1
*/
elog.Panic("Could not connect to database")

elog.Fatal("Could not connect to api")

elog.Error("Could not handle request")

elog.Warning("this request require customer id")

elog.Info("Application start at port 8080")

elog.Debug("Writing log")
```

Example log data
```json
{
   "Severity": "info",
   "Message": "Starting Balance Calculator",
   "Timestamp": "2018-12-31T18:13:57.925186+07:00",
   "Application": "balance.calculator"
}
``` 

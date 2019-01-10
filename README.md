# elog
Publish log to elasticsearch

```go
/*
* Initial log
* param1: elasticseach address
* param2: index
* param3: doc type
* param4: application name
* param5: software version
*/
elog.Init("http://127.0.0.1:9200", "log", "doc", "web.app", "1.0")

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
    "Message": "GET /",
    "Timestamp": "2019-01-05T18:25:34.566485073Z",
    "Application": "line-messaging",
    "Version": "0.0.5"
}
``` 

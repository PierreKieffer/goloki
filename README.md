# goloki

Go middleware for Grafana Loki's HTTP API 



* [Install](#install)
* [Log structure](#log-structure)
* [Usage](#usage)
	* [Log](#log)
	* [Push](#push)
	* [Labels](#labels)



## Install 

```bash 
go get github.com/PierreKieffer/goloki
```

## Log structure 
```json 
{
  "streams": [
    {
      "stream": {
        "label": "value"
      },
      "values": [
          [ "<unix epoch in nanoseconds>", "<log line>" ],
          [ "<unix epoch in nanoseconds>", "<log line>" ]
      ]
    }
  ]
}
```

## Usage
```go
import (
	"fmt"
	"github.com/PierreKieffer/goloki"
)
```

### Log 
- Single log 
```go 
log := goloki.Log("log line")
```

- Group of logs 
```go
l1 := "log line 1"
l2 := "log line 2"

log := goloki.Log([]string{l1, l2}, labels)
```
### Push 
```go
err := log.Push("http://loki:3100")
if err != nil {
	// ... 
}
```

### Labels 
```go
var labels = make(map[string]interface{})
labels["level"] = "INFO"
labels["foo"] = "bar"

log := goloki.Log("log line", labels)

log.Push("http://loki:3100")
```





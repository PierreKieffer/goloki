# goloki

Go middleware for Loki HTTP API 


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
```go 
log := goloki.Log("log line")
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





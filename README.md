# log15-redis-handler
	
## How to send log messages to Logstash via Redis

	package main

	import (
		"os"

		"github.com/bolcom/log15-redis-handler"
		"github.com/inconshreveable/log15"
	)

	func main() {
		logger := log15.New()
		f := redis15.LogstashFormat{}
		f.Application = "daxy"
		f.Role = "webservice"
		f.Instance = "ps34"
		f.SourceHost = "MacErnest"
		f.Pretty = true // e.g. hook to verbose flag

		redisHandler , _ := NewRedisHandler([]string{"localhost:6379}, "logs", f) 
		// must handle err

		logger.SetHandler(redisHandler)
		logger.Debug("hello", "pi", 3.14)
	}

Example output

	{
			"@source_host": "MacErnest",
			"@timestamp": "2017-02-17T09:57:37.028587577+01:00",
			"@fields": {
					"application": "daxy",
					"file": "redis15.go",
					"instance": "ps34",
					"level": "DEBUG",
					"line": "22",
					"pi": "3.14",
					"role": "webservice"
			},
			"@message": "hello"
	}

Note that most values are converted to a String. 
Use NoStringConvert on the LogstashFormat value to specify the exceptions.
For application logging, the `logtype`, `level` and `message` are reserved.

© 2017 bol.com. Apache v2 Licensed. Contributions are welcome.
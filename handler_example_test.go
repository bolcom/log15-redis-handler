/*
Copyright 2017 bol.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package redis15

import (
	"os"
	"testing"

	"github.com/inconshreveable/log15"
)

/**
{
    "@source_host": "MacErnest",
    "@timestamp": "2014-03-21T10:52:05.495118455+01:00",
    "@fields": {
        "level": "INFO",
        "threadid": 02628,
        "file": "glog_logstash_test.go",
        "line": 60,
        "instance": "ps34",
        "role": "webservice",
        "application": "daxy"
    },
    "@message": "hello"
}
**/
func TestJSONOutput(t *testing.T) {
	logger := log15.New()
	f := LogstashFormat{}
	f.Application = "daxy"
	f.Role = "webservice"
	f.Instance = "ps34"
	f.SourceHost = "MacErnest"
	f.Pretty = true // hook to verbose flag

	// use Caller handler on verbose only
	caller := log15.CallerFileHandler(log15.StreamHandler(os.Stdout, f))
	logger.SetHandler(caller)
	logger.Debug("hello", "pi", 3.14)
}

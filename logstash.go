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
	"encoding/json"
	"time"

	"github.com/inconshreveable/log15"
)

type LogstashFormat struct {
	SourceHost  string
	Instance    string
	Role        string
	Application string
	Pretty      bool
	OmitEmpty   bool
	// NoStringConvertSet holds keys for which no string conversion must be applied
	// Not go-routine safe
	NoStringConvertSet map[string]bool
}

// NoStringConvert registers keys for which the context value should not be converted to a string.
func (f *LogstashFormat) NoStringConvert(keys ...string) {
	if f.NoStringConvertSet == nil {
		f.NoStringConvertSet = map[string]bool{}
	}
	for _, each := range keys {
		f.NoStringConvertSet[each] = true
	}
}

type logstashRecord struct {
	SourceHost string                 `json:"@source_host"`
	Timestamp  time.Time              `json:"@timestamp"`
	Fields     map[string]interface{} `json:"@fields"`
	Message    string                 `json:"@message"`
}

// Format implements log15.Format
func (f LogstashFormat) Format(r *log15.Record) (data []byte) {
	lr := logstashRecord{}
	lr.SourceHost = f.SourceHost
	lr.Timestamp = r.Time
	lr.Fields = map[string]interface{}{
		"level":       level(r.Lvl),
		"instance":    f.Instance,
		"role":        f.Role,
		"application": f.Application,
	}
	lr.Message = r.Msg
	mapContextToFields(r.Ctx, f.OmitEmpty, f.NoStringConvertSet, lr.Fields)
	if f.Pretty {
		data, _ = json.MarshalIndent(lr, "", "\t")
	} else {
		data, _ = json.Marshal(lr)
	}
	return
}

// level returns a log4j label for a given lvl.
func level(l log15.Lvl) string {
	switch l {
	case log15.LvlInfo:
		return "INFO"
	case log15.LvlDebug:
		return "DEBUG"
	case log15.LvlWarn:
		return "WARN"
	case log15.LvlCrit:
		return "FATAL"
	case log15.LvlError:
		return "ERROR"
	}
	return l.String()
}

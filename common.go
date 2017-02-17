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
	"fmt"
	"io"
	"strings"

	"github.com/inconshreveable/log15"
)

func mapContextToFields(context []interface{}, omitEmpty bool, noConvertSet map[string]bool, fields map[string]interface{}) {
	var v interface{}
	for i := 0; i < len(context); i += 2 {
		// handle odd key count
		if i == len(context)-1 {
			break
		}
		// skip nil values
		if v = context[i+1]; v == nil {
			continue
		}
		// expect string keys
		k, ok := context[i].(string)
		if !ok {
			// convert if not
			k = fmt.Sprintf("%v", context[i])
		}
		if isReservedKey(k) {
			k = "_" + k
		}
		// caller is a special context key which is set when
		// the stack is inspected by a Caller* Handler
		if "caller" == k {
			ln := strings.Split(v.(string), ":")
			fields["file"] = ln[0]
			fields["line"] = ln[1]
			continue
		}
		// check for reserved keys for which types are important
		if noConvertSet != nil {
			// check presence
			if _, ok := noConvertSet[k]; ok {
				fields[k] = v
				continue
			}
		}
		// always use string value
		// (mixed types for the same field causes problems)
		var valueString string
		if s, ok := v.(string); ok {
			valueString = s
		} else {
			valueString = fmt.Sprintf("%v", v)
		}
		// optionally skip empty values
		if omitEmpty && len(valueString) == 0 {
			continue
		}
		fields[k] = valueString
	}
}

var newline = []byte("\n")

// StreamWithNewlineHandler writes log records to an io.Writer
// with the given format and appends a newline (\n).
//
// StreamWithNewlineHandler wraps itself with LazyHandler and SyncHandler
// to evaluate Lazy objects and perform safe concurrent writes.
func StreamWithNewlineHandler(wr io.Writer, fmtr log15.Format) log15.Handler {
	h := log15.FuncHandler(func(r *log15.Record) error {
		_, err := wr.Write(fmtr.Format(r))
		if err == nil {
			_, err = wr.Write(newline)
		}
		return err
	})
	return log15.LazyHandler(log15.SyncHandler(h))
}

func isReservedKey(key string) bool {
	return "message" == key ||
		"level" == key ||
		"logtype" == key
}

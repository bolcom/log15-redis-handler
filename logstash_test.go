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
	"reflect"
	"testing"

	"github.com/inconshreveable/log15"
)

func TestLogstashFormatNoStringConvert(t *testing.T) {
	f := LogstashFormat{}
	f.NoStringConvert("statuscode")
	r := log15.Record{}
	r.Ctx = []interface{}{"statuscode", 200}
	var hash map[string]interface{}
	json.Unmarshal(f.Format(&r), &hash)
	fields, ok := hash["@fields"].(map[string]interface{})
	if !ok {
		t.Error("fields is not a map")
	}
	if got, want := fields["statuscode"], float64(200); got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestLogstashFormatApplication(t *testing.T) {
	f := LogstashFormat{Application: "app"}
	r := log15.Record{}
	var hash map[string]interface{}
	json.Unmarshal(f.Format(&r), &hash)
	fields, ok := hash["@fields"].(map[string]interface{})
	if !ok {
		t.Error("fields is not a map")
	}
	if got, want := fields["application"], "app"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestLogstashFormat_NoStringConvert(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		f    *LogstashFormat
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.f.NoStringConvert(tt.args.keys...)
	}
}

func TestLogstashFormat_Format(t *testing.T) {
	type args struct {
		r *log15.Record
	}
	tests := []struct {
		name     string
		f        LogstashFormat
		args     args
		wantData []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if gotData := tt.f.Format(tt.args.r); !reflect.DeepEqual(gotData, tt.wantData) {
			t.Errorf("%q. LogstashFormat.Format() = %v, want %v", tt.name, gotData, tt.wantData)
		}
	}
}

func Test_level(t *testing.T) {
	type args struct {
		l log15.Lvl
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := level(tt.args.l); got != tt.want {
			t.Errorf("%q. level() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

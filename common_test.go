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
	"bytes"
	"reflect"
	"testing"

	"github.com/inconshreveable/log15"
)

func Test_mapContextToFields(t *testing.T) {
	type args struct {
		context      []interface{}
		omitEmpty    bool
		noConvertSet map[string]bool
		fields       map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mapContextToFields(tt.args.context, tt.args.omitEmpty, tt.args.noConvertSet, tt.args.fields)
	}
}

func TestStreamWithNewlineHandler(t *testing.T) {
	type args struct {
		fmtr log15.Format
	}
	tests := []struct {
		name   string
		args   args
		want   log15.Handler
		wantWr string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		wr := &bytes.Buffer{}
		if got := StreamWithNewlineHandler(wr, tt.args.fmtr); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. StreamWithNewlineHandler() = %v, want %v", tt.name, got, tt.want)
		}
		if gotWr := wr.String(); gotWr != tt.wantWr {
			t.Errorf("%q. StreamWithNewlineHandler() = %v, want %v", tt.name, gotWr, tt.wantWr)
		}
	}
}

func Test_isReservedKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := isReservedKey(tt.args.key); got != tt.want {
			t.Errorf("%q. isReservedKey() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

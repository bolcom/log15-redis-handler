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
	"reflect"
	"testing"
	"time"

	"github.com/inconshreveable/log15"
)

func TestNewBufferingHandler(t *testing.T) {
	type args struct {
		delegate log15.Handler
		capacity int
		interval time.Duration
	}
	tests := []struct {
		name string
		args args
		want *BufferingHandler
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewBufferingHandler(tt.args.delegate, tt.args.capacity, tt.args.interval); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewBufferingHandler() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBufferingHandler_Log(t *testing.T) {
	type args struct {
		r *log15.Record
	}
	tests := []struct {
		name    string
		b       *BufferingHandler
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.b.Log(tt.args.r); (err != nil) != tt.wantErr {
			t.Errorf("%q. BufferingHandler.Log() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestBufferingHandler_start(t *testing.T) {
	tests := []struct {
		name string
		b    *BufferingHandler
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.b.start()
	}
}

func TestBufferingHandler_flush(t *testing.T) {
	type args struct {
		records []*log15.Record
	}
	tests := []struct {
		name string
		b    *BufferingHandler
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.b.flush(tt.args.records)
	}
}

func TestBufferingHandler_Flush(t *testing.T) {
	tests := []struct {
		name string
		b    *BufferingHandler
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.b.Flush()
	}
}

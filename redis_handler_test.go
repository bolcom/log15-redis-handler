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

	"github.com/inconshreveable/log15"
)

func TestNewRedisHandler(t *testing.T) {
	type args struct {
		endpoints []string
		key       string
		format    log15.Format
	}
	tests := []struct {
		name    string
		args    args
		want    *RedisHandler
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := NewRedisHandler(tt.args.endpoints, tt.args.key, tt.args.format)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. NewRedisHandler() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewRedisHandler() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_pick(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := pick(tt.args.list); got != tt.want {
			t.Errorf("%q. pick() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestRedisHandler_Close(t *testing.T) {
	tests := []struct {
		name string
		h    *RedisHandler
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.h.Close()
	}
}

func TestRedisHandler_Log(t *testing.T) {
	type args struct {
		r *log15.Record
	}
	tests := []struct {
		name    string
		h       *RedisHandler
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.h.Log(tt.args.r); (err != nil) != tt.wantErr {
			t.Errorf("%q. RedisHandler.Log() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestRedisHandler_rpush(t *testing.T) {
	type args struct {
		r *log15.Record
	}
	tests := []struct {
		name    string
		h       *RedisHandler
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.h.rpush(tt.args.r); (err != nil) != tt.wantErr {
			t.Errorf("%q. RedisHandler.rpush() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestRedisHandler_tryRecoverFromError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		h       *RedisHandler
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.h.tryRecoverFromError(tt.args.err); (err != nil) != tt.wantErr {
			t.Errorf("%q. RedisHandler.tryRecoverFromError() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

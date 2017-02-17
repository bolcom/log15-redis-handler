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
	"fmt"
	"sync"
	"testing"

	"github.com/inconshreveable/log15"
)

func TestThatCommandKeyAndContextIsSendOverConnection(t *testing.T) {
	c := new(capturingConnection)
	h := &RedisHandler{new(sync.Mutex), "test", []string{"0:0"}, c, log15.JsonFormat()}

	l := log15.New("app", "test")
	l.SetHandler(h)
	l.Info("msg", "k", "v")
	if c.command != "RPUSH" {
		t.Errorf("wrong command: [%s]", c.command)
	}
	if len(c.args) == 0 {
		t.Fatal("missing args")
	}
	if c.args[0] != "test" {
		t.Error("wrong key")
	}
	doc := make(map[string]interface{})
	err := json.Unmarshal(c.args[1].([]byte), &doc)
	if err != nil {
		t.Error("no json?")
	}
	if doc["app"] != "test" {
		t.Error("missing app->test")
	}
}

func TestConnectionFailureWhenLogging(t *testing.T) {
	c := new(capturingConnection)
	c.err = fmt.Errorf("something is wrong")
	h := &RedisHandler{new(sync.Mutex), "test", []string{"0:0"}, c, log15.JsonFormat()}
	l := log15.New("app", "test")
	l.SetHandler(h)
	l.Info("msg", "k", "v")
	t.Logf("%#v", c)
	if got, want := c.returnedErr, true; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestForReal(t *testing.T) {
	t.Skip()
	h, err := NewRedisHandler([]string{"localhost:6379"}, "logs", LogstashFormat{})
	if err != nil {
		t.Fatal(err)
	}
	defer h.Close()
	l := log15.New("app", "test")
	l.SetHandler(h)
	l.Info("msg", "k", "v")
}

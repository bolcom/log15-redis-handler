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
	"math/rand"
	"sync"

	"github.com/garyburd/redigo/redis"
	"github.com/inconshreveable/log15"
)

// RedisHandler is a log15.Handler implementation that sends JSON documents to a Redis queue (key).
type RedisHandler struct {
	mu         *sync.Mutex
	key        string
	endpoints  []string
	connection RedisConnection
	format     log15.Format
}

// RedisConnection is the interface of functions we use from redis.
type RedisConnection interface {
	Send(string, ...interface{}) error
	Close() error
	Err() error
	Flush() error
}

// NewRedisHandler return a new connected RedisHandler.
// endpoints is a list of strings using the format host:port
func NewRedisHandler(endpoints []string, key string, format log15.Format) (*RedisHandler, error) {
	if len(endpoints) == 0 {
		return nil, fmt.Errorf("no endpoints given, handler not available")
	}
	connection, err := redis.Dial("tcp", pick(endpoints))

	return &RedisHandler{new(sync.Mutex), key, endpoints, connection, format}, err
}

// pick returns a randomly element
func pick(list []string) string {
	return list[rand.Intn(len(list))]
}

// Close closes the open connection to redis.
func (h *RedisHandler) Close() {
	if h.connection.Err() != nil {
		h.connection.Close()
	}
}

// Log implements log15.Handler
func (h *RedisHandler) Log(r *log15.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	// needs to flush otherwise the buffer in the conn.Redis will overflow after 4096 bytes
	defer h.connection.Flush()

	var err error
	err = h.rpush(r)
	if err == nil {
		return nil
	}

	for i := 0; i < 10; i++ {
		err = h.tryRecoverFromError(err)
		if err == nil {
			err = h.rpush(r)
			if err != nil {
				continue
			}
		}
	}
	if err != nil {
		return fmt.Errorf("No connection estabalished for redis handler: %s", err)
	}
	return nil
}

func (h *RedisHandler) rpush(r *log15.Record) error {
	return h.connection.Send("RPUSH", h.key, h.format.Format(r))
}

// tryRecoverFromError inspects the error to see if a reconnect is needed. Returns true if succeeded.
func (h *RedisHandler) tryRecoverFromError(err error) error {
	c, err := redis.Dial("tcp", pick(h.endpoints))
	if err != nil {
		return err
	}
	h.connection = c
	return nil
}

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
	"sync"
	"time"

	"github.com/inconshreveable/log15"
)

type BufferingHandler struct {
	mutex         *sync.RWMutex
	records       []*log15.Record
	maxRecords    int
	handler       log15.Handler
	flushInterval time.Duration
}

func NewBufferingHandler(delegate log15.Handler, capacity int, interval time.Duration) *BufferingHandler {
	h := &BufferingHandler{
		mutex:         new(sync.RWMutex),
		records:       []*log15.Record{},
		maxRecords:    capacity,
		handler:       delegate,
		flushInterval: interval,
	}
	h.start()
	return h
}

// Log implements log15.Handler
func (b *BufferingHandler) Log(r *log15.Record) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.records = append(b.records, r)
	if len(b.records) == b.maxRecords {
		go b.flush(b.records)
		b.records = []*log15.Record{}
	}
	return nil
}

// start spawns a goroutine the periodically checks the buffered records and flushes them
func (b *BufferingHandler) start() {
	go func() {
		for {
			time.Sleep(b.flushInterval)
			b.mutex.RLock()
			if len(b.records) > 0 {
				go func(list []*log15.Record) {
					b.flush(list)
				}(b.records)
				b.mutex.RUnlock()

				b.mutex.Lock()
				b.records = []*log15.Record{}
				b.mutex.Unlock()
			} else {
				b.mutex.RUnlock()
			}
		}
	}()
}

// flush passes each record to the delegate handler
// this runs inside the mutex critical region
func (b *BufferingHandler) flush(records []*log15.Record) {
	for _, each := range records {
		b.handler.Log(each)
	}
}

// Flush can be called to explicitly send all pending records.
func (b *BufferingHandler) Flush() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.flush(b.records)
}

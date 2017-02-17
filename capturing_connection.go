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

import "log"

type capturingConnection struct {
	command, key string
	args         []interface{}
	err          error
	returnedErr  bool
}

func (c *capturingConnection) Close() error { return nil }
func (c *capturingConnection) Flush() error { return nil }
func (c *capturingConnection) Err() error   { return nil }

func (c *capturingConnection) Send(cmd string, args ...interface{}) error {
	if c.err != nil {
		c.returnedErr = true
		return c.err
	}
	c.command = cmd
	c.key = args[0].(string)
	c.args = args
	log.Print(cmd)
	for i, a := range args {
		log.Printf("%d=%s", i, a)
	}
	return nil
}

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

import "testing"

func Test_capturingConnection_Close(t *testing.T) {
	tests := []struct {
		name    string
		c       *capturingConnection
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.c.Close(); (err != nil) != tt.wantErr {
			t.Errorf("%q. capturingConnection.Close() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_capturingConnection_Flush(t *testing.T) {
	tests := []struct {
		name    string
		c       *capturingConnection
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.c.Flush(); (err != nil) != tt.wantErr {
			t.Errorf("%q. capturingConnection.Flush() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_capturingConnection_Err(t *testing.T) {
	tests := []struct {
		name    string
		c       *capturingConnection
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.c.Err(); (err != nil) != tt.wantErr {
			t.Errorf("%q. capturingConnection.Err() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_capturingConnection_Send(t *testing.T) {
	type args struct {
		cmd  string
		args []interface{}
	}
	tests := []struct {
		name    string
		c       *capturingConnection
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.c.Send(tt.args.cmd, tt.args.args...); (err != nil) != tt.wantErr {
			t.Errorf("%q. capturingConnection.Send() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

/*
Copyright 2019 The Kubernetes Authors.

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

// This file is generated. DO NOT EDIT.

package legacyflag

import(
	"github.com/spf13/pflag"
	
)

// Int64Value is a reference to a registered int64 flag value.
type Int64Value struct {
	name string
	value int64
	fs *pflag.FlagSet
}

// Int64Var registers a flag for int64 against the FlagSet, and returns
// a Int64Value reference to the registered flag value.
func (fs *FlagSet) Int64Var(name string, def int64, usage string) *Int64Value {
	v := &Int64Value{
		name: name,
		fs: fs.fs,
	}
	fs.fs.Int64Var(&v.value, name, def, usage)
	return v
}

// Set copies the flag value to the target if the flag was set.
func (v *Int64Value) Set(target *int64) {
	if v.fs.Changed(v.name) {
		*target = v.value
	}
}

// Apply calls the apply func with the flag value if the flag was set.
func (v *Int64Value) Apply(apply func(value int64)) {
	if v.fs.Changed(v.name) {
		apply(v.value)
	}
}

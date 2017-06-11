package main

import "testing"

func TestList_Insert(t *testing.T) {
	type args struct {
		key interface{}
	}

	tests := []struct {
		name string
		l    *List
		args args
	}{{name: "Hola", l: &List{}, args: args{key: 1}}}

	for _, tt := range tests {
		tt.l.Insert(tt.args.key)
	}
}

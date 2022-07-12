package goroutine_map

import (
	"reflect"
	"testing"
)

func TestNewGMap(t *testing.T) {
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name string
		args args
		want *GMap
	}{
		{
			name: "NewGMap",
			args: args{
				m: map[string]string{
					"test": "test",
				},
			},
			want: &GMap{
				m: map[string]string{
					"test": "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGMap(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGMap_Add(t *testing.T) {
	type fields struct {
		m map[string]string
	}
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "GMap_Add",
			fields: fields{
				m: map[string]string{},
			},
			args: args{
				key: "new_key",
				val: "new_val",
			},
			want: true,
		},
		{
			name: "GMap_Add",
			fields: fields{
				m: map[string]string{
					"new_key": "new_val",
				},
			},
			args: args{
				key: "new_key",
				val: "new_val",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GMap{
				m: tt.fields.m,
			}
			if got := g.Add(tt.args.key, tt.args.val); got != tt.want {
				t.Errorf("GMap.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGMap_Get(t *testing.T) {
	type fields struct {
		m map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "GMap_Get",
			fields: fields{
				m: map[string]string{
					"test": "test",
				},
			},
			args: args{
				key: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GMap{
				m: tt.fields.m,
			}
			if got := g.Get(tt.args.key); got != tt.want {
				t.Errorf("GMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGMap_Del(t *testing.T) {
	type fields struct {
		m map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "GMap_Del",
			fields: fields{
				m: map[string]string{
					"del_key": "val",
				},
			},
			args: args{
				key: "del_key",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GMap{
				m: tt.fields.m,
			}
			g.Del(tt.args.key)
		})
	}
}

func TestGMap_Len(t *testing.T) {
	type fields struct {
		m map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "GMap_Len",
			fields: fields{
				m: map[string]string{
					"key": "val",
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GMap{
				m: tt.fields.m,
			}
			if got := g.Len(); got != tt.want {
				t.Errorf("GMap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

package strsort

import (
	"sort"
	"testing"
)

// func TestNewStrSort(t *testing.T) {
// 	type args struct {
// 		m []string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *StrSort
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewStrSort(tt.args.m); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewStrSort() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
func TestNewMSort(t *testing.T) {
	ss := []string{"liu", "fei", "laa", "lia", "zhang", "liuf", "li", "zhan"}

	s := NewStrSort(ss)
	sort.Sort(s)
	if s.sMap[0] != "fei" {
		t.Fatalf("msort is bad")
	}
	if s.sMap[1] != "laa" {
		t.Fatalf("msort is bad")
	}
	if s.sMap[2] != "li" {
		t.Fatalf("msort is bad")
	}
	if s.sMap[3] != "lia" {
		t.Fatalf("msort is bad")
	}
	if s.sMap[4] != "liu" {
		t.Fatalf("msort is bad")
	}
	if s.sMap[5] != "liuf" {
		t.Fatalf("msort is bad")
	}
	if s.sMap[6] != "zhan" {
		t.Fatalf("msort is bad")
	}
	if s.sMap[7] != "zhang" {
		t.Fatalf("msort is bad")
	}
}

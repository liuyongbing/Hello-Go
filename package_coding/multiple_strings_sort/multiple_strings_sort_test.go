package multiple_strings_sort

import (
	"sort"
	"testing"
)

// func TestNewStrMapSort(t *testing.T) {
// 	type args struct {
// 		m []string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *StrMapSort
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewStrMapSort(tt.args.m); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewStrMapSort() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
func TestNewMSort(t *testing.T) {
	ss := []string{"liu", "fei", "laa", "lia", "zhang", "liuf", "li", "zhan"}

	m := NewMSort(ss)
	sort.Sort(m)
	if m.s[0] != "fei" {
		t.Fatalf("msort is bad")
	}
	if m.s[1] != "laa" {
		t.Fatalf("msort is bad")
	}
	if m.s[2] != "li" {
		t.Fatalf("msort is bad")
	}
	if m.s[3] != "lia" {
		t.Fatalf("msort is bad")
	}
	if m.s[4] != "liu" {
		t.Fatalf("msort is bad")
	}
	if m.s[5] != "liuf" {
		t.Fatalf("msort is bad")
	}
	if m.s[6] != "zhan" {
		t.Fatalf("msort is bad")
	}
	if m.s[7] != "zhang" {
		t.Fatalf("msort is bad")
	}
}

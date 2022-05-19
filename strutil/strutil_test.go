package strutil

import (
	"reflect"
	"testing"
)

func TestContainsAnd(t *testing.T) {
	type args struct {
		str    string
		substr []string
	}
	tests := []struct {
		name           string
		args           args
		wantIsMatching bool
	}{
		{
			name: "Number",
			args: args{
				str:    "12345667124123",
				substr: []string{"1", "2", "3", "4"},
			},
			wantIsMatching: true,
		}, {
			name: "Alphabet",
			args: args{
				str:    "asdqweArcfdADIQEWJRIJnvciSAjAjs",
				substr: []string{"A", "B", "C", "D"},
			},
			wantIsMatching: false,
		}, {
			name: "Number+Alphabet",
			args: args{
				str:    "12AFSsda3456671asdSADASD24123",
				substr: []string{"1", "F", "4", "D", "3456671asdSADASD"},
			},
			wantIsMatching: true,
		}, {
			name: "Chinese",
			args: args{
				str:    "啊扫地机年会上的你啊时代的你saddle你撒旦",
				substr: []string{"你", "啊", "撒", "旦", "啊", "机年会上的你啊时代的"},
			},
			wantIsMatching: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsMatching := ContainsAnd(tt.args.str, tt.args.substr...); gotIsMatching != tt.wantIsMatching {
				t.Errorf("ContainsAnd() = %v, want %v", gotIsMatching, tt.wantIsMatching)
			}
		})
	}
}
func TestContainsOr(t *testing.T) {
	type args struct {
		str    string
		substr []string
	}
	tests := []struct {
		name           string
		args           args
		wantIsMatching bool
	}{
		{
			name: "Number",
			args: args{
				str:    "12345667124123",
				substr: []string{"1", "2", "3", "4"},
			},
			wantIsMatching: true,
		}, {
			name: "Alphabet",
			args: args{
				str:    "asdqweArcfdADIQEWJRIJnvciSAjAjs",
				substr: []string{"A", "B", "C", "D"},
			},
			wantIsMatching: true,
		}, {
			name: "Number+Alphabet",
			args: args{
				str:    "12AFSsda3456671asdSADASD24123",
				substr: []string{"1", "F", "4", "D", "3456671asdSADASD"},
			},
			wantIsMatching: true,
		}, {
			name: "Chinese",
			args: args{
				str:    "啊扫地机年会上的你啊时代的你saddle你撒旦",
				substr: []string{"你", "啊", "撒", "旦", "啊", "机年会上的你啊时代的"},
			},
			wantIsMatching: true,
		}, {
			name: "Chinese",
			args: args{
				str:    "啊扫地机年会上的你啊时代的你saddle你撒旦",
				substr: []string{"1", "2", "3", "bv撒", "a1", "机年会上5啊时代的"},
			},
			wantIsMatching: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsMatching := ContainsOr(tt.args.str, tt.args.substr...); gotIsMatching != tt.wantIsMatching {
				t.Errorf("ContainsOr() = %v, want %v", gotIsMatching, tt.wantIsMatching)
			}
		})
	}
}
func TestContains(t *testing.T) {
	type args struct {
		str    string
		substr string
	}
	tests := []struct {
		name           string
		args           args
		wantIsMatching bool
	}{
		{name: "", args: args{str: "你好123asdsdaASDoe123!#@*%)!!\"", substr: "\""}, wantIsMatching: true},
		{name: "", args: args{"你好123asdsdaASDoe123!#@*%)!!\"", ")!!"}, wantIsMatching: true},
		{name: "", args: args{"你好123asdsdaASDoe123!#@*%)!!\"", "你好123asdsdaASDoe123!#@*%)!!\""}, wantIsMatching: true},
		{name: "", args: args{"你好123asdsdaASDoe123!#@*%)!!\"", "你好"}, wantIsMatching: true},
		{name: "", args: args{"你好123asdsdaASDoe123!#@*%)!!\"", "你"}, wantIsMatching: true},
		{name: "", args: args{"你好123asdsdaASDoe123!#@*%)!!\"", "你11"}, wantIsMatching: false},
		{name: "", args: args{"你好123asdsdaASDoe123!#@*%)!!\"", ""}, wantIsMatching: false},
		{name: "", args: args{"你好123asdsdaASDoe123!#@*%)!!\"", `\`}, wantIsMatching: false},
		{name: "", args: args{"你好123asdsdaASDoe123!#@*%)!!\"", `"`}, wantIsMatching: true},
		{name: "", args: args{"你好123asdsdaASDoe123!#@*%)!!\"", `ASDoe`}, wantIsMatching: true},
		{name: "", args: args{"你好123asdsdaASDoe123!#@*%)!!\"", "!#@*%"}, wantIsMatching: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsMatching := Contains(tt.args.str, tt.args.substr); gotIsMatching != tt.wantIsMatching {
				t.Errorf("Contains() = %v, want %v", gotIsMatching, tt.wantIsMatching)
			}
		})
	}
}

func TestSplitLoop(t *testing.T) {
	type args struct {
		str       string
		splitChar rune
		interval  int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name:       "",
			args:       args{"123,456,789,012,345,678,890", rune(','), 4},
			wantResult: []string{"123,456,789,012", "345,678,890"},
		},
		{
			name:       "",
			args:       args{",,,,,,", rune(','), 4},
			wantResult: []string{",,,", ",,"},
		},
		{
			name:       "",
			args:       args{"你好吗，你还在吗，你没事吧？", rune('你'), 2},
			wantResult: []string{"你好吗，", "还在吗，你没事吧？"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := SplitLoop(tt.args.str, tt.args.splitChar, tt.args.interval); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("SplitLoop() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

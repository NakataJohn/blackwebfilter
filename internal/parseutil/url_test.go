package parseutil

import (
	"reflect"
	"testing"
)

// func Test_getSecondLevelDomain(t *testing.T) {
// 	type args struct {
// 		urlStr string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name:    "good",
// 			args:    args{urlStr: "https://www.hebgh.org/#/login"},
// 			want:    ".hebgh.org",
// 			wantErr: false,
// 		},
// 		{
// 			name:    "good1",
// 			args:    args{urlStr: "https://wzjk.axtx.com.cn/#/login"},
// 			want:    ".axtx.com.cn",
// 			wantErr: false,
// 		},
// 		{
// 			name:    "good2",
// 			args:    args{urlStr: "https://wzjkax.heboylove.blogspot.co.at"},
// 			want:    ".heboylove.blogspot.co.at",
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := GetDomain(tt.args.urlStr)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("getDomain() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("getDomain() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestGetSubDomains(t *testing.T) {
	type args struct {
		urlStr string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "good",
			args: args{urlStr: "https://www.hebgh.org/#/login"},
			want: []string{".hebgh.org"},
		},
		{
			name: "good1",
			args: args{urlStr: "https://wzjk.axtx.com.cn/#/login"},
			want: []string{".axtx.com.cn"},
		},
		{
			name: "good2",
			args: args{urlStr: "https://wzjkax.heboylove.blogspot.co.at"},
			want: []string{".heboylove.blogspot.co.at", ".blogspot.co.at", ".co.at", ".at"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSubDomains(tt.args.urlStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubDomains() = %v, want %v", got, tt.want)
			}
		})
	}
}

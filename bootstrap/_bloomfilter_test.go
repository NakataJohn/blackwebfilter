package bootstrap

import "testing"

func TestFilterTest(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "good",
			args: args{data: "aaaa"},
			want: true,
		},
		{
			name: "bad",
			args: args{data: ".10gb.ru"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterTest(tt.args.data); got != tt.want {
				t.Errorf("FilterTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

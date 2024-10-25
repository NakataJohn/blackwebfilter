package bootstrap

import "testing"

func TestTrieTest(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "good",
			args: args{
				word: "深圳市東方富海科技有限公司",
			},
			want: 0,
		},
		{
			name: "good1",
			args: args{
				word: "亚洲国产成人久久综合一区,欧美日韩视频在线播放,久久毛片免费看一区二区三区,国产福利在线观看第二区",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrieTest(tt.args.word); got != tt.want {
				t.Errorf("TrieTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

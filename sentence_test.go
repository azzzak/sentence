package sentence

import (
	"testing"
)

var pattern = "{{pluraln .Bottles `бутылка|бутылки|бутылок`}} пива {{plural .Bottles `стояла|стояли|стояло`}} на столе"

func TestSentence_MustRender(t *testing.T) {
	type args struct {
		pattern string
		data    interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok 1 bottle",
			args: args{
				pattern: pattern,
				data: struct {
					Bottles int
				}{
					Bottles: 1,
				},
			},
			want: "1 бутылка пива стояла на столе",
		},
		{
			name: "ok 2 bottles",
			args: args{
				pattern: pattern,
				data: struct {
					Bottles int
				}{
					Bottles: 2,
				},
			},
			want: "2 бутылки пива стояли на столе",
		},
		{
			name: "ok 5 bottles",
			args: args{
				pattern: pattern,
				data: struct {
					Bottles int
				}{
					Bottles: 5,
				},
			},
			want: "5 бутылок пива стояло на столе",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := New()
			got := s.MustRender(tt.args.pattern, tt.args.data)
			if got != tt.want {
				t.Errorf("Sentence.MustRender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSentence_Concurrency(t *testing.T) {
	t.Run("concurrent", func(t *testing.T) {
		s, _ := New()
		for i := 0; i < 25; i++ {
			go func() {
				_ = s.MustRender(pattern, struct {
					Bottles int
				}{
					Bottles: 5,
				})
			}()
		}
	})
}

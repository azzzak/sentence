package sentence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pluralEnglish(t *testing.T) {
	minutes := []string{"minute", "minutes"}

	type args struct {
		n     int
		forms []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{-5, minutes}, want: "minutes"},
		{args: args{-2, minutes}, want: "minutes"},
		{args: args{-1, minutes}, want: "minute"},
		{args: args{0, minutes}, want: "minutes"},
		{args: args{1, minutes}, want: "minute"},
		{args: args{2, minutes}, want: "minutes"},
		{args: args{10, minutes}, want: "minutes"},
		{args: args{121, minutes}, want: "minutes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pluralEnglish(tt.args.n, tt.args.forms)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_pluralRussian(t *testing.T) {
	minutes := []string{"минуту", "минуты", "минут"}

	type args struct {
		n     int
		forms []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{-5, minutes}, want: "минут"},
		{args: args{-2, minutes}, want: "минуты"},
		{args: args{-1, minutes}, want: "минуту"},
		{args: args{0, minutes}, want: "минут"},
		{args: args{1, minutes}, want: "минуту"},
		{args: args{2, minutes}, want: "минуты"},
		{args: args{3, minutes}, want: "минуты"},
		{args: args{4, minutes}, want: "минуты"},
		{args: args{5, minutes}, want: "минут"},
		{args: args{6, minutes}, want: "минут"},
		{args: args{7, minutes}, want: "минут"},
		{args: args{8, minutes}, want: "минут"},
		{args: args{9, minutes}, want: "минут"},
		{args: args{10, minutes}, want: "минут"},
		{args: args{11, minutes}, want: "минут"},
		{args: args{12, minutes}, want: "минут"},
		{args: args{13, minutes}, want: "минут"},
		{args: args{14, minutes}, want: "минут"},
		{args: args{15, minutes}, want: "минут"},
		{args: args{16, minutes}, want: "минут"},
		{args: args{17, minutes}, want: "минут"},
		{args: args{18, minutes}, want: "минут"},
		{args: args{19, minutes}, want: "минут"},
		{args: args{20, minutes}, want: "минут"},
		{args: args{21, minutes}, want: "минуту"},
		{args: args{22, minutes}, want: "минуты"},
		{args: args{23, minutes}, want: "минуты"},
		{args: args{24, minutes}, want: "минуты"},
		{args: args{25, minutes}, want: "минут"},
		{args: args{26, minutes}, want: "минут"},
		{args: args{50, minutes}, want: "минут"},
		{args: args{51, minutes}, want: "минуту"},
		{args: args{52, minutes}, want: "минуты"},
		{args: args{55, minutes}, want: "минут"},
		{args: args{59, minutes}, want: "минут"},
		{args: args{80, minutes}, want: "минут"},
		{args: args{81, minutes}, want: "минуту"},
		{args: args{100, minutes}, want: "минут"},
		{args: args{101, minutes}, want: "минуту"},
		{args: args{105, minutes}, want: "минут"},
		{args: args{111, minutes}, want: "минут"},
		{args: args{112, minutes}, want: "минут"},
		{args: args{113, minutes}, want: "минут"},
		{args: args{114, minutes}, want: "минут"},
		{args: args{115, minutes}, want: "минут"},
		{args: args{120, minutes}, want: "минут"},
		{args: args{121, minutes}, want: "минуту"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pluralRussian(tt.args.n, tt.args.forms)
			assert.Equal(t, tt.want, got)
		})
	}
}

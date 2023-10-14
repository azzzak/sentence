package sentence

import (
	"reflect"
	"testing"
)

func Test_initPlural(t *testing.T) {
	type args struct {
		num  int
		item reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok string 1",
			args: args{
				num:  1,
				item: reflect.ValueOf("минута|минуты|минут"),
			},
			want:    "минута",
			wantErr: false,
		},
		{
			name: "ok string 2",
			args: args{
				num:  2,
				item: reflect.ValueOf("минута|минуты|минут"),
			},
			want:    "минуты",
			wantErr: false,
		},
		{
			name: "ok string 5",
			args: args{
				num:  5,
				item: reflect.ValueOf("минута|минуты|минут"),
			},
			want:    "минут",
			wantErr: false,
		},
		{
			name: "ok slice 5",
			args: args{
				num:  5,
				item: reflect.ValueOf([]string{"минута", "минуты", "минут"}),
			},
			want:    "минут",
			wantErr: false,
		},
		{
			name: "err",
			args: args{
				num:  5,
				item: reflect.ValueOf("минута|минуты"),
			},
			wantErr: true,
		},
		{
			name: "err wrong type",
			args: args{
				num:  5,
				item: reflect.ValueOf([]int{1, 2}),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := initPlural("ru", "|")

			got, err := fn(tt.args.num, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("plural() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("plural() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initPluraln(t *testing.T) {
	type args struct {
		num  int
		item reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok string 1",
			args: args{
				num:  1,
				item: reflect.ValueOf("минута|минуты|минут"),
			},
			want:    "1 минута",
			wantErr: false,
		},
		{
			name: "ok string 2",
			args: args{
				num:  2,
				item: reflect.ValueOf("минута|минуты|минут"),
			},
			want:    "2 минуты",
			wantErr: false,
		},
		{
			name: "ok string 5",
			args: args{
				num:  5,
				item: reflect.ValueOf("минута|минуты|минут"),
			},
			want:    "5 минут",
			wantErr: false,
		},
		{
			name: "ok slice 5",
			args: args{
				num:  5,
				item: reflect.ValueOf([]string{"минута", "минуты", "минут"}),
			},
			want:    "5 минут",
			wantErr: false,
		},
		{
			name: "err",
			args: args{
				num:  5,
				item: reflect.ValueOf("минута|минуты"),
			},
			wantErr: true,
		},
		{
			name: "err wrong type",
			args: args{
				num:  5,
				item: reflect.ValueOf([]int{1, 2}),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := initPluraln(initPlural("ru", "|"))

			got, err := fn(tt.args.num, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("pluraln() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pluraln() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initAny(t *testing.T) {
	type args struct {
		item reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want1   string
		want2   string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				item: reflect.ValueOf("гофер|гошник"),
			},
			want1:   "гофер",
			want2:   "гошник",
			wantErr: false,
		},
		{
			name: "ok slice",
			args: args{
				item: reflect.ValueOf([]string{"гофер", "гошник"}),
			},
			want1:   "гофер",
			want2:   "гошник",
			wantErr: false,
		},
		{
			name: "err",
			args: args{
				item: reflect.ValueOf([]int{1, 2}),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := initAny("|")

			got, err := fn(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("any() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want1 && got != tt.want2 {
				t.Errorf("any() = %v, want %v or %v", got, tt.want1, tt.want2)
			}
		})
	}
}

func Test_initfAny(t *testing.T) {
	type args struct {
		items []reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want1   string
		want2   string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				items: []reflect.Value{
					reflect.ValueOf("гофер"),
					reflect.ValueOf("гошник"),
				},
			},
			want1:   "гофер",
			want2:   "гошник",
			wantErr: false,
		},
		{
			name: "err",
			args: args{
				items: []reflect.Value{
					reflect.ValueOf(1),
					reflect.ValueOf(2),
				},
			},
			want1:   "1",
			want2:   "2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := initAnyf()

			got, err := fn(tt.args.items...)
			if (err != nil) != tt.wantErr {
				t.Errorf("any() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want1 && got != tt.want2 {
				t.Errorf("any() = %v, want %v or %v", got, tt.want1, tt.want2)
			}
		})
	}
}

package pluralizer

import (
	"reflect"
	"testing"
)

func TestToPlural(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name  string
		args  func(t *testing.T) args
		want1 string
	}{
		{
			name: "Regular noun",
			args: func(_ *testing.T) args {
				return args{word: "book"}
			},
			want1: "books",
		},
		{
			name: "Noun ending in 'y'",
			args: func(_ *testing.T) args {
				return args{word: "city"}
			},
			want1: "cities",
		},
		{
			name: "Noun ending in 'f'",
			args: func(_ *testing.T) args {
				return args{word: "leaf"}
			},
			want1: "leaves",
		},
		{
			name: "Irregular noun",
			args: func(_ *testing.T) args {
				return args{word: "child"}
			},
			want1: "children",
		},
		{
			name: "Noun already in plural form",
			args: func(_ *testing.T) args {
				return args{word: "dogs"}
			},
			want1: "dogs",
		},
		{
			name: "Uncountable noun",
			args: func(_ *testing.T) args {
				return args{word: "information"}
			},
			want1: "information",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := ToPlural(tArgs.word)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToPlural got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

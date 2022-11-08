package collections

import (
	"reflect"
	"testing"
)

func TestArrayList_Get(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "case1",
			fields: fields{data: []int{1}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayList[int]{
				data: tt.fields.data,
			}
			if got := a.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

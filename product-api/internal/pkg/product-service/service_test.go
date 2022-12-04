package product_service

import (
	"reflect"
	"testing"

	"github.com/soa/product-api/internal/models"
)

func Test_getCategorys(t *testing.T) {
	type args struct {
		cats []models.Category
		m    map[int64]models.Category
		id   int64
	}
	catMap := map[int64]models.Category{
		1: {
			ID:          1,
			Name:        "1",
			Description: "",
			ParentID:    2,
			Level:       1,
		},
		2: {
			ID:          2,
			Name:        "2",
			Description: "",
			ParentID:    0,
			Level:       2,
		},
	}
	tests := []struct {
		name string
		args args
		want []models.Category
	}{
		{
			name: "From first",
			args: args{
				cats: nil,
				m:    catMap,
				id:   1,
			},
			want: []models.Category{
				{1,
					"1",
					"",
					2,
					1,
				},
				{
					2,
					"2",
					"",
					0,
					2,
				},
			},
		},
		{
			name: "From second",
			args: args{
				cats: nil,
				m:    catMap,
				id:   2,
			},
			want: []models.Category{
				{
					2,
					"2",
					"",
					0,
					2,
				},
			},
		},
		{
			name: "Empty",
			args: args{
				cats: nil,
				m:    catMap,
				id:   0,
			},
			want: []models.Category{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCategorys(tt.args.cats, tt.args.m, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCategorys() = %v, want %v", got, tt.want)
			}
		})
	}
}

package services

import (
	"reflect"
	"testing"
)

func TestPackSizeService_CalculatePackSizeByOrderAmount(t *testing.T) {
	type fields struct {
		PackSizes []int
	}
	type args struct {
		orderItems int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[int]int
		wantErr bool
	}{
		{
			name: "Test with valid pack sizes from email",
			fields: fields{
				PackSizes: []int{23, 31, 53},
			},
			args: args{
				orderItems: 500000,
			},
			want: map[int]int{
				53: 9433,
				31: 1,
				23: 1,
			},
			wantErr: false,
		},
		{
			name: "Test with valid pack sizes from pdf",
			fields: fields{
				PackSizes: []int{250, 500, 1000, 2000, 5000},
			},
			args: args{
				orderItems: 12001,
			},
			want: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPackSizeService(tt.fields.PackSizes)
			got, err := p.CalculatePackSizeByOrderAmount(tt.args.orderItems)
			if (err != nil) != tt.wantErr {
				t.Errorf("PackSizeService.CalculatePackSizeByOrderAmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackSizeService.CalculatePackSizeByOrderAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

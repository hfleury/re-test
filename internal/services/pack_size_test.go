package services

import (
	"reflect"
	"testing"
)

func TestPackSizeService_CalculatePackSizeByOrderAmount(t *testing.T) {
	type args struct {
		orderItems int
		packSizes  []int
	}
	tests := []struct {
		name    string
		args    args
		want    map[int]int
		wantErr bool
	}{
		{
			name: "Test with valid pack sizes from email",
			args: args{
				orderItems: 500000,
				packSizes:  []int{23, 31, 53},
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
			args: args{
				orderItems: 12001,
				packSizes:  []int{250, 500, 1000, 2000, 5000},
			},
			want: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
			wantErr: false,
		},
		{
			name: "Test with invalid input (zero)",
			args: args{
				orderItems: 0,
				packSizes:  []int{250, 500, 1000},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test with small order amount not divisible",
			args: args{
				orderItems: 7,
				packSizes:  []int{5},
			},
			want: map[int]int{
				5: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewPackSizeService()
			got, err := service.CalculatePackSizeByOrderAmount(tt.args.orderItems, tt.args.packSizes)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculatePackSizeByOrderAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

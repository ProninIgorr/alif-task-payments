package services

import (
	"testing"

	"github.com/ProninIgorr/alif-task-payments/internal/models"
)

func TestCheckCategory(t *testing.T) {
	type args struct {
		product *models.Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "HaveCategory",
			args: args{
				product: &models.Product{
					Category:    "Smartphone",
					Amount:      1000,
					Installment: 12,
				},
			},
			wantErr: false,
		},

		{
			name: "NotHave",
			args: args{
				product: &models.Product{
					Category:    "WashinMachine",
					Amount:      1000,
					Installment: 12,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckCategory(tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("CheckCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckInstallment(t *testing.T) {
	type args struct {
		product *models.Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "HaveInstallment",
			args: args{
				product: &models.Product{
					Category:    "Smartphone",
					Amount:      1000,
					Installment: 12,
				},
			},
			wantErr: false,
		},

		{
			name: "NotHave",
			args: args{
				product: &models.Product{
					Category:    "Smartphone",
					Amount:      1000,
					Installment: 2,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckInstallment(tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("CheckInstallment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTotalAmount(t *testing.T) {
	type args struct {
		product *models.Product
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "TotalAmount",
			args: args{
				product: &models.Product{
					Category:    "Smartphone",
					Amount:      1000,
					Installment: 12,
				},
			},
			want: 1060,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalAmount(tt.args.product); got != tt.want {
				t.Errorf("TotalAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentPerMonth(t *testing.T) {
	type args struct {
		product *models.Product
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "PaymentPerMonth",
			args: args{
				product: &models.Product{
					Category:    "Smartphone",
					Amount:      1000,
					Installment: 12,
				},
			},
			want: 88.33333333333333,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PaymentPerMonth(tt.args.product); got != tt.want {
				t.Errorf("PaymentPerMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

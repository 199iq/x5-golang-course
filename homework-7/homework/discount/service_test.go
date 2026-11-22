package discount

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidatePercent(t *testing.T) {
	tests := []struct {
		name    string
		percent int
		wantErr error
	}{
		{name: "minimum", percent: 0},
		{name: "valid", percent: 10},
		{name: "maximum", percent: 100},
		{name: "below minimum", percent: -1, wantErr: ErrInvalidPercent},
		{name: "above maximum", percent: 101, wantErr: ErrInvalidPercent},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := ValidatePercent(tt.percent)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestApply(t *testing.T) {
	tests := []struct {
		name    string
		price   int
		percent int
		want    int
		wantErr error
	}{
		{name: "regular discount", price: 1000, percent: 15, want: 850},
		{name: "no discount", price: 1000, percent: 0, want: 1000},
		{name: "full discount", price: 1000, percent: 100, want: 0},
		{name: "negative price", price: -1, percent: 10, wantErr: ErrNegativePrice},
		{name: "negative percent", price: 1000, percent: -1, wantErr: ErrInvalidPercent},
		{name: "percent above maximum", price: 1000, percent: 101, wantErr: ErrInvalidPercent},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Apply(tt.price, tt.percent)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				require.Zero(t, got)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFinalPrice(t *testing.T) {
	tests := []struct {
		name    string
		items   []Item
		want    int
		wantErr error
	}{
		{name: "empty list", items: []Item{}, want: 0},
		{
			name: "multiple items",
			items: []Item{
				{Price: 1000, DiscountPercent: 10},
				{Price: 500, DiscountPercent: 20},
			},
			want: 1300,
		},
		{
			name: "invalid price after valid item",
			items: []Item{
				{Price: 100, DiscountPercent: 10},
				{Price: -1, DiscountPercent: 10},
			},
			wantErr: ErrNegativePrice,
		},
		{
			name: "invalid discount after valid item",
			items: []Item{
				{Price: 100, DiscountPercent: 10},
				{Price: 500, DiscountPercent: 101},
			},
			wantErr: ErrInvalidPercent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := FinalPrice(tt.items)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				require.Zero(t, got)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

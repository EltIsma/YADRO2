package domain

import (
	"testing"
	"time"
)

func TestRevenueCalculation(t *testing.T) {
	t1, _ := time.Parse("15:04", "10:00")
	t2, _ := time.Parse("15:04", "11:00")
	t3, _ := time.Parse("15:04", "11:00")
	t4, _ := time.Parse("15:04", "00:00")
	t5, _ := time.Parse("15:04", "08:01")
	t6, _ := time.Parse("15:04", "08:01")

	tests := []struct {
		table       *Table
		eventTime   time.Time
		price       int
		wantRevenue int
		wantTime    time.Time
	}{
		{
			table: &Table{
				busy:    false,
				time:    t1,
				revenue: 0,
				start:   t1,
			},
			eventTime:   t2,
			price:       10,
			wantRevenue: 10,
			wantTime:    t3,
		},
		{
			table: &Table{
				busy:    false,
				time:    t4,
				revenue: 0,
				start:   t4,
			},
			eventTime:   t5,
			price:       10,
			wantRevenue: 90,
			wantTime:    t6,
		},
	}

	for _, test := range tests {
		test.table.revenueCalculation(test.eventTime, test.price)
		if test.table.revenue != test.wantRevenue || test.table.time != test.wantTime {
			t.Errorf("revenueCalculation(%v, %v, %v) = %v, %v, want %v, %v", test.table, test.eventTime, test.price, test.table.revenue, test.table.time, test.wantRevenue, test.wantTime)
		}
	}
}

package stats

import "github.com/bdaler/bank/pkg/types"

func Avg(payments []types.Payment) types.Money {
	count := len(payments)
	sum := 0

	for i := 0; i < count; i++ {
		sum += int(payments[i].Amount)
	}
	avg := sum / count

	return types.Money(avg)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}

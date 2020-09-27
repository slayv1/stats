package stats

import (
	"github.com/slayv1/bank/pkg/types"
)

//Avg func
func Avg(payments []types.Payment) types.Money{

	var sum types.Money 
	for _, v := range payments {
		sum += v.Amount
	}
	return sum/types.Money(len(payments))
	
}

//TotalInCategory func
func TotalInCategory(payments []types.Payment, category types.Category) (total types.Money) {

	for _, v := range payments {
	  if category == v.Category {
		total += v.Amount
	  }
	}
	return
  }


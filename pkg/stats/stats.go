package stats

import (
	"github.com/slayv1/bank/v2/pkg/types"
)

//Avg func return average amount from slice Payment
func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	i := 0
	for _, v := range payments {
	  if v.Status == types.StatusFail {
		continue
	  }
	  sum += v.Amount
	  i++
	}
	return sum / types.Money(i)
  }
  
  //TotalInCategory  returned total sum in one category
  func TotalInCategory(payments []types.Payment, category types.Category) (total types.Money) {
  
	for _, v := range payments {
	  if v.Status == types.StatusFail {
		continue
	  }
	  if category == v.Category {
		total += v.Amount
	  }
	}
	return
  }

	
  


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

  //CategoriesAvg func
func CategoriesAvg(payments []types.Payment) ( map[types.Category]types.Money) {

	mp := make(map[types.Category]types.Money)
	for _, v := range payments {
  
	  if _, er := mp[v.Category]; er {
		continue
	  }
	  filtered := FilterByCategory(payments, v.Category)
	  mp[v.Category]=Avg(filtered)
	}
  
  
	return mp
  }
  
  
  //FilterByCategory func
  func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment{
  
	var filtered []types.Payment
  
	for _, v := range payments {
	  if v.Category == category {
		filtered = append(filtered, v)
	  }
	}
	return filtered
  }

	
  
//PeriodsDynamic func
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	mp := map[types.Category]types.Money{}
 
	if len(first)>=len(second){
 
		 for k := range first {
			 mp[k]=second[k]-first[k]
		 }
		 return mp
 
	}
	
	 for k := range second {
		 mp[k]=second[k]-first[k]
	 }
 
	return mp
 }
 

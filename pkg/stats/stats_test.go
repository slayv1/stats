package stats

import (
	"reflect"
	"testing"
	"github.com/slayv1/bank/v2/pkg/types"
)

func TestCategoriesAvg_nil(t *testing.T) {
	var payments []types.Payment
	res := CategoriesAvg(payments)
  
	if len(res) != 0 {
	  t.Errorf(" got > %v want > nil", res)
	}
  }
  func TestCategoriesAvg_one(t *testing.T) {
	payments := []types.Payment{
	  {
		ID:       1,
		Category: "cafe",
		Amount:   100_00,
	  },
	  {
		ID:       2,
		Category: "cafe",
		Amount:   100_00,
	  },
	}
	  expected := map[types.Category]types.Money{
	  "cafe":100_00,
	}
  
  
	res := CategoriesAvg(payments)
  
  
	if !reflect.DeepEqual(expected, res) {
	  t.Errorf(" got > %v want > nil", res)
	} 
  }
  
  func TestCategoriesAvg_multiple(t *testing.T) {
	payments := []types.Payment{
	  {
		ID:       1,
		Category: "cafe",
		Amount:   100_00,
	  },
	  {
		ID:       2,
		Category: "cafe",
		Amount:   100_00,
	  },
	  {
		ID:       3,
		Category: "restaurant",
		Amount:   10_00,
	  },
	  {
		ID:       2,
		Category: "restaurant",
		Amount:   100_00,
	  },
	}
	  expected := map[types.Category]types.Money{
	  "cafe":100_00,
	  "restaurant":55_00,
	}
  
  
	res := CategoriesAvg(payments)
  
  
	if !reflect.DeepEqual(expected, res) {
	  t.Errorf(" got > %v want > nil", res)
	} 
  
  }
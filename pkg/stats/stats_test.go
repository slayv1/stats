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



  func TestPeriodsDynamic_negative(t *testing.T) {
	first := map[types.Category]types.Money{
	  "cafe": 20,
	  "auto": 14,
	}
	second := map[types.Category]types.Money{
	  "cafe": 15,
	  "auto": 7,
	}
	want := map[types.Category]types.Money{
	  "cafe": -5,
	  "auto": -7,
	}
  
	got := PeriodsDynamic(first, second)
  
	if !reflect.DeepEqual(want, got) {
	  t.Errorf(" got > %v \n want > %v", got, want)
	} 
  
  }
  
  func TestPeriodsDynamic_positive(t *testing.T) {
	first := map[types.Category]types.Money{
	  "cafe": 20,
	  "auto": 14,
	}
	second := map[types.Category]types.Money{
	  "cafe": 35,
	  "auto": 17,
	}
	want := map[types.Category]types.Money{
	  "cafe": 15,
	  "auto": 3,
	}
  
	got := PeriodsDynamic(first, second)
  
	if !reflect.DeepEqual(want, got) {
	  t.Errorf(" got > %v \n want > %v", got, want)
	} 
  
  }
  func TestPeriodsDynamic_notEqualMap(t *testing.T) {
	first := map[types.Category]types.Money{
	  "cafe": 20,
	  "auto": 14,
	}
	second := map[types.Category]types.Money{
	  "cafe": 35,
  
	}
	want := map[types.Category]types.Money{
	  "cafe": 15,
	  "auto": -14,
	}
  
	got := PeriodsDynamic(first, second)
  
	if !reflect.DeepEqual(want, got) {
	  t.Errorf(" got > %v \n want > %v", got, want)
	} 
  
  }
  
  
  func TestPeriodsDynamic_OneMoreElem(t *testing.T) {
	first := map[types.Category]types.Money{
	  "cafe": 20,
	  "auto": 14,
	}
	second := map[types.Category]types.Money{
	  "cafe": 35,
	  "auto": 17,
	  "mobile": 17,
	}
	want := map[types.Category]types.Money{
	  "cafe": 15,
	  "auto": 3,
	  "mobile": 17,
	}
  
	got := PeriodsDynamic(first, second)
  
	if !reflect.DeepEqual(want, got) {
	  t.Errorf(" got > %v \n want > %v", got, want)
	} 
  
  }

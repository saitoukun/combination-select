package compute

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//動的計画法を解く関数
func TestSolveDP(t *testing.T) {
	var expectValue = 3
	var targetWeight = 7
	items := Items{
		Item{1,1,"item1"},
		Item{2,1,"item2"},
		Item{3,1,"item3"},
		Item{4,1,"item4"},
		Item{5,1,"item5"},
	}
	//weight7で価値が最大の組み合わせの数は3つ(1,2,4)
	result, _ := solveDP(len(items), targetWeight, items)
	assert.Equal(t, expectValue, result)
}


func TestCombinationSelect(t *testing.T) {
	var targetWeight = 7
	items := Items{
		Item{5,1,"item1"},
		Item{4,1,"item2"},
		Item{3,1,"item3"},
		Item{2,1,"item4"},
		Item{1,1,"item5"},
	}

	result := CombinationSelect(targetWeight, items)
	var resultSum = 0
	for i := 0; i < len(result); i++ {
		resultSum = resultSum + result[i].Weight
	}
	assert.Equal(t, targetWeight, resultSum)
}

func TestCombinationSelect2(t *testing.T) {
	var targetWeight = 3
	items := Items{
		Item{5,1,"item1"},
		Item{4,1,"item2"},
		Item{3,1,"item3"},
		Item{2,1,"item4"},
		Item{1,1,"item5"},
	}
	result := CombinationSelect(targetWeight, items)

	var resultSum = 0
	for i := 0; i < len(result); i++ {
		resultSum = resultSum + result[i].Weight
	}

	assert.Equal(t, targetWeight, resultSum)
}
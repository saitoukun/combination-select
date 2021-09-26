package compute

import (
	"math/rand"
	"time"
)

type Item struct {
	Weight int `json:"weight"`
	Value  int `json:"value"`
	Name string `json:"name"`
}
type Items []Item

type Trace struct {
	x, y int
}

// CombinationSelect
// 曲の配列と時間からDPを作成し結果を返す
func CombinationSelect(targetWeight int, items Items) Items {
	_, traces := solveDP(len(items), targetWeight, items)
	solvedIndex := solveTrace(len(items), targetWeight, traces)

	resultItems := []Item{}
	for _, value := range solvedIndex {
		resultItems = append(resultItems, items[value-1])
	}
	return resultItems
}

// N: itemの数
// M: 合計の重み
func solveDP(N int,M int, items Items) (int, [][]Trace) {
	dp, traces := createDPTable(N, M)

	for i:= 1; i <= N; i++ {
		for j := 0; j <= M; j++ {
			if dp[i-1][j] >= 0 {
				dp[i][j] = dp[i-1][j]
				traces[i][j] = Trace{i - 1, j}
			}
			if j >= items[i-1].Weight && dp[i-1][j-items[i-1].Weight] >= 0 && dp[i][j] < dp[i-1][j-items[i-1].Weight]+items[i-1].Value {
				dp[i][j] = dp[i-1][j-items[i-1].Weight] + items[i-1].Value
				traces[i][j] = Trace{i-1, j-items[i-1].Weight}
			}
		}
	}
	return dp[N][M], traces
}

// N: itemの数
// M: 合計の重み
func solveTrace(N int, M int, traces [][]Trace) []int {
	pnt := Trace{N, M}
	routes := []Trace{}
	res := []int{}

	for pnt.x != 0 {
		routes = append(routes, pnt)
		pre := pnt
		pnt = traces[pnt.x][pnt.y]

		if pre.y != pnt.y {
			res = append(res, pre.x)
		}
	}

	//fmt.Println(routes)
	//fmt.Println(res)

	return res
}

// N: itemの数
// M: 合計の重み
func createDPTable(N int, M int) ([][]int, [][]Trace) {
	dp := make([][]int, N+1)
	traces := make([][]Trace, N+1)

	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, M+1)
		traces[i] = make([]Trace, M+1)
	}

	for i := 0; i <= N; i++ {
		for j := 1; j <= M; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	return dp, traces
}

//配列をシャッフルする
func shuffle(items Items) Items {
	var newItems = items
	for i := range newItems {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		newItems[i], newItems[j] = newItems[j], newItems[i]
	}
	return newItems
}

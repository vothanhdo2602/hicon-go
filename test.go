package hicon

import (
	"fmt"
	"github.com/vothanhdo2602/hicon/external/util/log"
)

func getServerIndex(n int32, arrival []int32, burstTime []int32) []int32 {
	// Write your code here
	var (
		result  []int32
		servers = make([][]int32, n)
	)

	for i := 0; i < len(arrival); i++ {
		for j := 0; j < len(servers); j++ {
			if len(servers[j]) == 0 {
				servers[j] = []int32{arrival[i], arrival[i] + burstTime[i], int32(i)}
				result = append(result, int32(j)+1)
				break
			}

			timeComplete := arrival[i] + burstTime[i]
			if timeComplete <= servers[j][1] && arrival[i] <= servers[j][1] {
				servers[j] = []int32{arrival[i], timeComplete, int32(i)}
				result = append(result, int32(j)+1)
				break
			}
		}
	}

	fmt.Println("@@@@@@@@@@@@@@ ", result)
	return result
}

func main() {
	log.Init()

	//var (
	//	ctx = context.Background()
	//)

	arrival := []int32{2, 4, 1, 8, 9}
	burst := []int32{7, 9, 2, 4, 5}
	fmt.Println(arrival, burst)
	getServerIndex(3, arrival, burst)
}

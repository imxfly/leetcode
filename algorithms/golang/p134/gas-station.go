package p134

// URL: https://leetcode-cn.com/problems/gas-station/

// CanCompleteCircuit gas station
func CanCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)

	for i := 0; i <= length; i++ {
		ni := i % length // 边界处理
		if gas[ni] < cost[ni] || gas[ni]+gas[(ni+1)%length] < cost[ni]+cost[(ni+1)%length] {
			// 选择出发的加油站，先剔除不符合要求的加油站
			continue
		}
		j := ni
		sum := gas[ni] - cost[ni]
		for ; j-ni < length; j++ {
			nj := (j + 1) % length
			if sum+gas[nj] < cost[nj] {
				// 只有当前剩余油量加上下一站的油量能够支撑去下一站所需油量时才行
				break
			}
			sum += gas[nj] - cost[nj]
		}
		if j >= ni+length-1 {
			return ni
		}
	}

	return -1
}

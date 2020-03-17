package p447

// URL: https://leetcode-cn.com/problems/number-of-boomerangs/

func dist(i []int, j []int) int {
	return (i[0]-j[0])*(i[0]-j[0]) + (i[1]-j[1])*(i[1]-j[1])
}

// NumberOfBoomerangs number of boomerangs
func NumberOfBoomerangs(points [][]int) int {
	num := 0 // number of boomerangs

	for i := range points {
		ht := make(map[int]int, len(points))
		for j := range points {
			if i != j {
				dist := dist(points[i], points[j])
				if cnt, ok := ht[dist]; ok {
					num += cnt * 2
				}
				ht[dist]++
			}
		}
	}

	return num
}

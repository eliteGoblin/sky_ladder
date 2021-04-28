package _55_jump_game

import "github.com/eliteGoblin/sky_ladder/common"

func canJump(nums []int) bool {
	for i := 0; i < len(nums); {
		most := i + nums[i]
		for j := i; j <= most; j++ {
			most = common.Max(most, j+nums[j])
			if most >= len(nums)-1 {
				return true
			}
		}
		if i == most {
			return false
		}
		i = most
	}
	return false
}

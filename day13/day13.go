package day13

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

func FirstPart(lines []string) (int, error) {
	var minWaitTime, busNumber = math.MaxInt64, 0
	timestamp, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, err
	}
	shuttles := strings.Split(lines[1], ",")
	for _, s := range shuttles {
		if s == "x" {
			continue
		}
		number, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		division := timestamp / number
		minTime := division * number
		if division < timestamp {
			minTime += number
		}
		minTime -= timestamp
		if minTime < minWaitTime {
			minWaitTime = minTime
			busNumber = number
		}
	}
	return busNumber * minWaitTime, nil
}

func SecondPart(lines []string) (int, error) {
	shuttles := strings.Split(lines[1], ",")
	var nums []int
	var remainders []int
	for i, s := range shuttles {
		if s == "x" {
			continue
		}
		number, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		nums = append(nums, number)
		remainders = append(remainders, number-i)
	}
	return chineseRemainderTheorem(nums, remainders), nil
}

func chineseRemainderTheorem(nums, remainders []int) int {
	prod := greatestCommonDivisor(nums)
	inv := make([]int, len(nums))
	pp := make([]int, len(nums))
	for i := range nums {
		pp[i] = prod / nums[i]
		inv[i] = modInverse(pp[i], nums[i])
	}
	var res int
	for i := range nums {
		res += remainders[i] * pp[i] * inv[i]
	}
	return res % prod
}

func modInverse(a, m int) int {
	return int(big.NewInt(0).ModInverse(big.NewInt(int64(a)), big.NewInt(int64(m))).Int64())
}

func greatestCommonDivisor(nums []int) int {
	var res = nums[0]
	for i := 1; i < len(nums); i++ {
		res *= nums[i]
	}
	return res
}

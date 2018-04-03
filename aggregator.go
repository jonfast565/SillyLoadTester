package main

import "time"

func AggregateInt64Slice(slice []int64, fn func(int64, int64) int64) int64 {
	var result int64
	for i := 0; i < len(slice); i++ {
		result = fn(slice[i], result)
	}
	return result
}

func DurationToInt64(duration time.Duration) int64 {
	return duration.Nanoseconds()
}

func DurationSliceToInt64Slice(slice []time.Duration) []int64 {
	result := make([]int64, len(slice))
	for i := 0; i < len(slice); i++ {
		result[i] = DurationToInt64(slice[i])
	}
	return result
}

func Int64ToDuration(integer int64) time.Duration {
	return time.Duration(integer)
}

func (loadTester *LoadTester) GetAverageDuration() time.Duration {
	durationNanos := DurationSliceToInt64Slice(loadTester.requestDurations)
	addition := func(d1 int64, d2 int64) int64 {
		return d1 + d2
	}
	aggregationResult := AggregateInt64Slice(durationNanos, addition)
	result := aggregationResult / int64(len(durationNanos))
	return Int64ToDuration(result)
}
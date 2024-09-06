package tokenBucket

import (
	"testing"
	"time"
)

func Test_RateLimiter(t *testing.T) {
	type fields struct {
		capacity    int
		refill      int
		refreshRate time.Duration
	}

	testCases := []struct {
		name   string
		fields fields
		want   []bool
	}{
		{
			"user with 10 req per sec",
			fields{
				capacity:    10,
				refill:      5,
				refreshRate: time.Second,
			},
			[]bool{true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false},
		},
		{
			"user with 10 req per sec",
			fields{
				capacity:    5,
				refill:      5,
				refreshRate: time.Second,
			},
			[]bool{true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		},
	}
	for _, tt := range testCases {
		t.Run("test_name", func(t *testing.T) {
			rl := NewTokenBucket(tt.fields.capacity, tt.fields.refill, tt.fields.refreshRate)
			got := runRequests(rl)

			for i := 0; i < 20; i++ {
				if got[i] != tt.want[i] {
					t.Errorf("%v .TokenBucket.HaveTokens() = %v, want %v", i, got[i], tt.want[i])
				}
			}

		})
	}

}

func runRequests(rl RateLimiter) []bool {
	result := []bool{}
	for i := 0; i < 20; i++ {
		result = append(result, rl.HaveTokens(1))
	}
	return result
}

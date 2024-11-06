package arithmetics

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/heliobmartins/concurrencyworkshop/internal/pattern/fanoutin"
)

func SequentialSum(inputSize int) int {
	sum := 0
	for i := 1; i <= inputSize; i++ {
		sum += process(i)
	}
	return sum
}

// ParallelSum implement this method.
func ParallelSum(inputSize int) int {
	sum := atomic.Int32{}
	var wg sync.WaitGroup
	for i := 1; i <= inputSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum.Add(int32(process(i)))
		}()
	}
	wg.Wait()
	return int(sum.Load())
}

// ParallelSum implement this method.
func ParallelSum2(inputSize int) int {
	ctx := context.Background()
	var jobs []fanoutin.Job[int]
	for i := 1; i < inputSize; i++ {
		jobs = append(jobs, fanoutin.Job[int]{
			ID:    i,
			Value: i,
		})
	}
	results := fanoutin.FanOut(ctx, jobs, processAsync)
	sum := 0
	for result := range results {
		if result.Err != nil {
			continue
		}
		sum += result.Value
	}
	return sum
}

func processAsync(ctx context.Context, num int) (int, error) {
	return process(num), nil
}

func process(num int) int {
	time.Sleep(time.Millisecond) // simulate processing time
	return num * num
}

//455773333 ns/op
//595574854

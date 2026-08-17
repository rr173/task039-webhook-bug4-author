package webhook

import (
	"sync"
	"testing"
)

func TestProbeConcurrentDownstreamAccess(t *testing.T) {
	s := New("probe-secret")
	var wg sync.WaitGroup
	for worker := 0; worker < 8; worker++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			for i := 0; i < 2000; i++ {
				if (worker+i)%2 == 0 {
					s.SetDownstream("http://127.0.0.1:8080/a")
				} else {
					_ = s.Downstream()
				}
			}
		}(worker)
	}
	wg.Wait()
}

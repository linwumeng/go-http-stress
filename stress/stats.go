package stress

import (
	"fmt"
	"time"

	sts "github.com/montanaflynn/stats"
)

type stats struct {
	ch          <-chan *data
	done        chan<- bool
	concurrency int
	n           int
	clock       int
}

func (s *stats) start(t int, url string) {
	ticker := time.NewTicker(time.Second)

	s.printHead(t, url)
	slice := []*data{}

	go func(timer <-chan time.Time, first bool) {
		for {
			select {
			case <-timer:
				if first {
					first = false
					timer = time.Tick(time.Duration(t) * time.Second)
					s.printRow(t, slice)
					continue
				}
				s.printRow(t, slice)
			case p, ok := <-s.ch:
				if !ok {
					s.printRow(t, slice)
					s.done <- true
					return
				}

				slice = append(slice, p)
			}
		}
	}(ticker.C, true)
}

func (s *stats) printHead(t int, url string) {
	fmt.Printf("URL:      %v\n", url)
	fmt.Printf("协程数:   %v\n", s.concurrency)
	fmt.Printf("总请求数: %v\n", s.concurrency*s.n)
	fmt.Printf("输出间隔: %vs\n", t)
	fmt.Println("─────┬────────┬────────┬────────┬────────┬────────┬────────┬────────┬────────")
	result := fmt.Sprintf(" 窗口│ 请求数 │ 成功数 │ 错误率 │中位耗时│最长耗时│最短耗时│95%% 耗时│ TPS")
	// result := fmt.Sprintf(" 耗时│ 并发数 │ 成功数 │ 失败率 │   qps  │最长耗时│最短耗时│平均耗时│ 错误码")
	fmt.Println(result)
	fmt.Println("─────┼────────┼────────┼────────┼────────┼────────┼────────┼────────┼────────")
}

func (s *stats) printRow(i int, data []*data) {
	if len(data) == 0 {
		return
	}

	s.clock++
	suc, fal, rt := s.count(data)
	er := float32(fal) / float32(len(data))
	m, _ := sts.Median(rt)
	max, _ := sts.Max(rt)
	min, _ := sts.Min(rt)
	p95, _ := sts.Percentile(rt, 95)
	sum, _ := sts.Sum(rt)
	tps := float32(1000) * float32(len(data)) * float32(s.concurrency) / float32(sum)

	fmt.Printf(" %4v|%8v|%8v|%8.2f|%6vms|%6vms|%6vms|%6vms|%8.2f\n", s.clock, len(data), suc, er, int32(m), int32(max), int32(min), int32(p95), tps)
}

func (s *stats) count(data []*data) (int, int, []float64) {
	suc, fal, rt := 0, 0, make([]float64, len(data))
	for i, p := range data {
		if p.code/100 == 2 {
			suc++
		} else {
			fal++
		}

		rt[i] = float64(p.rt)
	}

	return suc, fal, rt
}

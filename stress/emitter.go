package stress

import "sync"

type Emitter struct {
	concurrency int
	count       int
}

func NewEmitter(c, n int) *Emitter {
	return &Emitter{
		concurrency: c,
		count:       n,
	}
}

func (e *Emitter) Emit(url string, t int) {
	wg := sync.WaitGroup{}
	wg.Add(e.concurrency)
	ch := make(chan *data, 100)
	done := make(chan bool)
	for i := 0; i < e.concurrency; i++ {
		go func() {
			r := &requestor{
				url: url,
				n:   e.count,
				wg:  &wg,
				ch:  ch,
			}

			r.get()
		}()
	}

	stats := &stats{ch: ch, done: done, concurrency: e.concurrency, n: e.count}
	stats.start(t, url)

	wg.Wait()
	close(ch)
	<-done
}

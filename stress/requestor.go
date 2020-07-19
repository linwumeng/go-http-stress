package stress

import (
	"crypto/tls"
	"net/http"
	"sync"
	"time"
)

type data struct {
	code int
	rt   int64
}
type requestor struct {
	url string
	n   int
	wg  *sync.WaitGroup
	ch  chan<- *data
}

type timerRoundtripper struct {
	roundtripper http.RoundTripper
	ch           chan<- *data
}

func (t *timerRoundtripper) RoundTrip(req *http.Request) (*http.Response, error) {
	d := data{code: http.StatusInternalServerError}
	now := time.Now()
	resp, e := t.roundtripper.RoundTrip(req)

	if e == nil {
		d.code = resp.StatusCode
	}
	d.rt = time.Now().Sub(now).Milliseconds()
	t.ch <- &d

	return resp, e
}

func (r *requestor) get() {
	tr := &timerRoundtripper{
		roundtripper: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		ch: r.ch,
	}
	client := &http.Client{Transport: tr}

	for i := 0; i < r.n; i++ {
		resp, _ := client.Get(r.url)
		if resp != nil {
			resp.Body.Close()
		}
	}

	r.wg.Done()
}

package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	mutex     sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {

	if size <= 0 {
		return nil, errors.New("Size value too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {

	case r, ok := <-p.resources:

		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resource")

		return p.factory()
	}
}
func (p *Pool) Release(r io.Closer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release:", "In Pool")
	default:
		log.Println("Release:", "Pool is Full,Closing resource")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		return
	}
	p.closed = true

	// 必须close p.resources的原因是避免之后的for range 陷入死锁一直等待
	close(p.resources)

	for r := range p.resources {
		r.Close()
	}

}

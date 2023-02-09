 package writeasync

import (
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"io"
	"sync"
	"errors"
	"math"
	"fmt"
	"time"
)

type WriteAsync struct {
	wr      	io.Writer
	lock    	sync.RWMutex
	wg      	sync.WaitGroup
	pool  		*bytebufferpool.Pool
	channel 	chan *bytebufferpool.ByteBuffer 
	tmp_buff 	*bytebufferpool.ByteBuffer 
	tmp_length  int
	tmp_err     error
}

func (this *WriteAsync) Close() {
	this.lock.Lock()
	close(this.channel)
	this.wg.Wait()
	this.lock.Unlock()
}

var GetQueueError = errors.New("writing failure, is locked")
var WriteError = errors.New("writing failure, queue is full")

func (this *WriteAsync) Write(data []byte) (int, error) {
	if this.lock.TryRLock() {
		defer this.lock.RUnlock()

		this.tmp_buff = this.pool.Get()
		this.tmp_length, this.tmp_err = this.tmp_buff.Write(data)
		if this.tmp_err != nil {
			this.pool.Put(this.tmp_buff)
			return this.tmp_length, this.tmp_err
		}
		
		select {
		case this.channel <- this.tmp_buff:
			break
		case <- time.After(time.Millisecond * 20):
			this.pool.Put(this.tmp_buff)
			return this.tmp_length, WriteError
		}
		return len(data), nil
	} else {
		return -1, GetQueueError
	}
}

func (this *WriteAsync) pop() {
	defer this.wg.Done()

	var buff  *bytebufferpool.ByteBuffer 
	var err error

	for buff = range this.channel {
		_, err = this.wr.Write(buff.Bytes())
		if err != nil {
			fmt.Printf("error: %s, message: %s", err.Error(), buff.String())
		}
		this.pool.Put(buff)
	}
}

var MaxBufferSize = math.MaxUint16 * 16

func NewWriterAsync(w io.Writer) *WriteAsync {
	if w == nil {
		return nil
	}

	out := WriteAsync{
		pool:    new(bytebufferpool.Pool),
		wr:        w,
		channel: make(chan *bytebufferpool.ByteBuffer, MaxBufferSize),
	}

	out.wg.Add(1)
	go out.pop()

	return &out
}


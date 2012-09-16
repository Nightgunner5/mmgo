package terrain

import (
	"github.com/banthar/gl"
	"time"
)

func init() {
	go func() {
		for {
			time.Sleep(10 * time.Second)

			threshold := time.Now().Unix() - 20
			chunkLock.Lock()
			for coord, chunk := range chunks {
				if chunk.lastGet < threshold {
					go chunk.gc()
					delete(chunks, coord)
				}
			}
			chunkLock.Unlock()
		}
	}()
}

func (c *Chunk) markGet() {
	c.lastGet = time.Now().Unix()
}

func (c *Chunk) gc() {
	if c.DisplayList != 0 {
		gl.DeleteLists(c.DisplayList, 1)
	}
}

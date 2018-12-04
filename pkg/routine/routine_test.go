package routine

import (
	//"PushServer/pkg/setting"
	"log"
	"testing"
	"time"
)

//模板推送功能测试
func TestIterator(t *testing.T) {

	gatefs := NewGatefs(3)

	for i := 0; i < 4; i++ {
		go func(i int) {
			gatefs.Enter()
			log.Printf("%d", i)
		}(i)
	}

	time.Sleep(10 * time.Second)
}

//用于限制进程的数量
package routine

type gate chan bool

type Gatefs struct {
	gate gate
}

func NewGatefs(num int) *Gatefs {
	return &Gatefs{
		gate: make(gate, num),
	}
}

func (g Gatefs) Enter() {
	g.gate <- true
}

func (g Gatefs) Leave() {
	<-g.gate
}

package process

import "fmt"

type Batch struct {
	task      func()
	completed bool
	id        int
}

func (b *Batch) Add(task func()) {
	b.task = task
	b.completed = false
}

func (b *Batch) SetId(id int) {
	b.id = id
}

func (b *Batch) execute(c chan int) {
	fmt.Println("waiting...")
	b.task()
	fmt.Println(fmt.Sprintf("hi %d", b.id))
	b.completed = true
	c <- -1
}

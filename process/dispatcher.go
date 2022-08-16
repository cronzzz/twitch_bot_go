package process

type Dispatcher struct {
	batches []*Batch
	process chan int
	quit    chan int
}

func (d *Dispatcher) Add(batch *Batch) {
	d.batches = append(d.batches, batch)
}

func (d *Dispatcher) SetProcessChannel(p chan int) {
	d.process = p
}

func (d *Dispatcher) SetQuitChannel(q chan int) {
	d.quit = q
}

func (d *Dispatcher) getBatches() []*Batch {
	return d.batches
}

func (d *Dispatcher) Dispatch() {
	max_routines := 100
	routines := 0
	semaphoreChannel := make(chan int)
	for {
		select {
		case <-d.process:
			if len(d.getBatches()) > 0 {
				if routines < max_routines {
					limit := max_routines - routines
					if len(d.getBatches()) < limit {
						limit = len(d.getBatches())
					}
					for _, batch := range d.getBatches()[:limit] {
						if !batch.completed {
							go batch.execute(semaphoreChannel)
							routines++
						}
					}
				}
				for {
					routines = routines + <-semaphoreChannel
					if routines < 50 {
						break
					}
				}
				completedBacthes := 0
				for _, batch := range d.getBatches() {
					if batch.completed {
						completedBacthes++
					} else {
						break
					}
				}
				if len(d.batches) > 0 {
					d.batches = d.batches[completedBacthes:]
				}
			}
		case <-d.quit:
			for _, batch := range d.getBatches() {
				if !batch.completed {
					go batch.execute(semaphoreChannel)
				}
			}
			return
		}
	}
}

package react

// Define reactor, cell and canceler types here.
// These types will implement the Reactor, Cell and Canceler interfaces, respectively.
type canceler struct {
	id           int
	registeredTo *cell
}

type reactor struct {
	computes []*cell
}

type cell struct {
	data      int
	compute   func() int
	callbacks map[int]func(int)
	reactor   *reactor
}

func (c *canceler) Cancel() {
	delete(c.registeredTo.callbacks, c.id)
}

func (c *cell) Value() int {
	if c.compute != nil {
		return c.compute()
	}
	return c.data
}

func (c *cell) SetValue(value int) {
	if c.data != value {
		c.data = value
		c.reactor.update()
	}
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	c.callbacks[len(c.callbacks)] = callback
	return &canceler{id: len(c.callbacks) - 1, registeredTo: c}
}

func New() Reactor {
	return &reactor{computes: make([]*cell, 0)}
}

func (r *reactor) CreateInput(initial int) InputCell {
	return &cell{data: initial, reactor: r}
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	newCell := &cell{}

	newCell.compute = func() int { return compute(dep.Value()) }
	newCell.data = newCell.compute()
	newCell.callbacks = make(map[int]func(int))
	r.computes = append(r.computes, newCell)
	return newCell
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	newCell := &cell{}

	newCell.compute = func() int { return compute(dep1.Value(), dep2.Value()) }
	newCell.data = newCell.compute()
	newCell.callbacks = make(map[int]func(int))
	r.computes = append(r.computes, newCell)
	return newCell
}

func (r *reactor) update() {
	for _, c := range r.computes {
		if c.compute() == c.data {
			continue
		}
		for _, cb := range c.callbacks {
			cb(c.Value())
		}
	}
}

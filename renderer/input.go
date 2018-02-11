package renderer

type Input interface {
	HasNextObject() bool
	NextObject() Object
	Reset()
	Add(object Object)
}

type input struct {
	objects []Object
	current int
}

func NewInput() Input {
	return &input {
		objects: make([]Object, 0),
	}
}

func (c input)HasNextObject() bool {
	return c.current < len(c.objects)
}

func (c *input)NextObject()Object {
	if c.HasNextObject(){
		result := c.objects[c.current]
		c.current++
		return result
	}
	return nil
}

func (c *input)Reset() {
	c.current = 0
}

func (c *input)Add(object Object) {
	c.objects = append(c.objects, object)
}

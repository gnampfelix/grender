package renderer

type Input interface {
	HasNextObject() bool
	NextObject() Object
	Reset()
}

type cubeInput struct {
	cube Object
	current int
}

func NewCubeInput() Input {
	return &cubeInput {
		cube: NewCube(),
	}
}

func (c cubeInput)HasNextObject() bool {
	return c.current == 0
}

func (c *cubeInput)NextObject()Object {
	if c.current == 0 {
		c.current = 1
		return c.cube
	}
	return nil
}

func (c *cubeInput)Reset() {
	c.current = 0
}

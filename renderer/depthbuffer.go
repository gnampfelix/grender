package renderer

type DepthBuffer interface {
  SetDepthIfCloser(depth float64, x, y int) bool
}

func NewMapBuffer()DepthBuffer {
  return &mapDepth{buffer: make(map[int]map[int]float64)}
}

type mapDepth struct {
  buffer map[int]map[int]float64
}

func (m *mapDepth)SetDepthIfCloser(depth float64, x, y int) bool {
  _, ok := m.buffer[x]
  if !ok {
    m.buffer[x] = make(map[int]float64)
    m.buffer[x][y] = depth
    return true
  }
  currentCol, ok := m.buffer[x][y]
  if !ok {
    m.buffer[x][y] = depth
    return true
  }

  if depth < currentCol {
    m.buffer[x][y] = depth
    return true
  }
  return false
}

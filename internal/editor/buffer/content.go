package buffer

type Content struct {
	currentLine int
	lines       map[int]string
}

func NewContent() *Content {
	return &Content{
		currentLine: 0,
		lines:       make(map[int]string, 0),
	}
}

func (c *Content) Line(number int) string {
	return c.lines[number]
}

func (c *Content) Lines() map[int]string {
	return c.lines
}

func (c *Content) SetLine(number int, line string) {
	c.lines[number] = line
}

func (c *Content) String() string {
	var result string
	for _, line := range c.lines {
		result += line + "\n"
	}
	return result
}

func (c *Content) WriteLine(line string) {
	c.lines[c.currentLine] = line
	c.currentLine++
}

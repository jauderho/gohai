package memory

type Memory struct{}

const name = "memory"

func (self *Memory) Name() string {
	return name
}

func (self *Memory) Collect() (result any, err error) {
	result, err = getMemoryInfo()
	return
}

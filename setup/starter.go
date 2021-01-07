package setup

type IStarterInterface interface {
	SetUp()
	Bootstrap()
}

type starter struct {
	elements []IStarterInterface
}

func (s *starter) SetUp() {

	if len(s.elements) == 0 {
		return
	}

	for i := 0; i < len(s.elements); i++ {
		s.elements[i].SetUp()
	}

	for i := 0; i < len(s.elements); i++ {
		s.elements[i].Bootstrap()
	}
}

func (s *starter) Register(starter IStarterInterface) {
	s.elements = append(s.elements, starter)
}

func NewStarter() *starter {
	return &starter{}
}

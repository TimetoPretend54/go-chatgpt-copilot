package mock

type AnimalService struct {
	ExpectedOutput string
	ExpectedErr    error
}

func (m *AnimalService) DoSomething(test string) (string, error) {
	return m.ExpectedOutput, m.ExpectedErr
}

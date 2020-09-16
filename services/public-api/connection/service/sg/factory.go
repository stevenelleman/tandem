package sg

type SgConnectionFactory struct{}

func NewSgConnFactory() *SgConnectionFactory {
	return &SgConnectionFactory{}
}

func (f *SgConnectionFactory) Connection() *SgConnection {
	return &SgConnection{}
}

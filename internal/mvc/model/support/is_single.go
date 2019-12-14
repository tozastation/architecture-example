package support

type (
	IsSingle struct {value bool}
)

func NewIsSingle(isSingle bool) (*IsSingle, error) {
	n := &IsSingle{isSingle}
	return n, nil
}

func (n IsSingle) Value() bool {
	return n.value
}
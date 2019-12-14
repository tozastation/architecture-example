package support

type (
	Price struct {value uint}
)

func NewPrice() (*Price, error) {
	n := &Price{}
	return n, nil
}

func (n Price) Value() uint {
	return n.value
}
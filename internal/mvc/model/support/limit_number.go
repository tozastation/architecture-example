package support

type LimitNumber struct {
	isLimit bool
	value uint
}

func NewLimitNumber(limit uint) (*LimitNumber, error) {
	n := &LimitNumber{value:limit}
	return n, nil
}

func (n LimitNumber) Value() uint {
	return n.value
}
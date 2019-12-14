package support

import (
	"fmt"
)

type SearchOption struct {
	single *IsSingle
	limit  *LimitNumber
}

func NewSearchSingleOption() (*SearchOption, error) {
	isSingle, err := NewIsSingle(true)
	if err != nil {
		return nil, err
	}
	n := &SearchOption{single: isSingle, limit: nil}
	return n, nil
}

func NewSearchManyOption(limit LimitNumber) (*SearchOption, error) {
	isSingle, err := NewIsSingle(false)
	if err != nil {
		return nil, err
	}
	n := &SearchOption{single: isSingle, limit: &limit}
	return n, nil
}

func (n SearchOption) String() string {
	return fmt.Sprintf("%d", n.limit.Value())
}

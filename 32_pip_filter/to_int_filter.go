package _2_pip_filter

import (
	"errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data should be []string")

type ToIntFilter struct {}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tif *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	var ret []int
	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err !=nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}

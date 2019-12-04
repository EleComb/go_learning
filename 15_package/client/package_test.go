package client

import (
	"github.com/EleComb/go_learning/15_package/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
	//t.Log(series.square(5)) // 不可访问
}

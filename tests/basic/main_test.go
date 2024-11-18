package basic

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAddOne(t *testing.T) {

	//var (
	//	intput = 1
	//	output = 2
	//)
	//
	//actual := AddOne(1)
	//
	//if actual != output {
	//	t.Errorf("AddOne(%d), input %d, actual = %d", intput, output, actual)
	//}

	assert.Equal(t, AddOne(2), 3, ":3")

}

//
// testify golang

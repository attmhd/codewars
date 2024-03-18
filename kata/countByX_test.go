package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountByX(t *testing.T) {

	result := CountByX(2, 5)
	exp := []int{2, 4, 5, 8, 10}
	// if reflect.DeepEqual(result, exp) != true {
	// 	t.Error("tidak sesuai")
	// }

	assert.Equal(t, result, exp, "Harusnya Sama")

	// fmt.Println("succes ", result, exp)

}

package testdata

import "fmt"

func preferAnyVariables() {
	var a any
	var b interface{} // want `Use of 'interface{}' detected. Replace with 'any'.`

	var c []any
	var d []interface{} // want `Use of 'interface{}' detected. Replace with 'any'.`

	var e map[any]int
	var f map[interface{}]int // want `Use of 'interface{}' detected. Replace with 'any'.`

	var g map[int]any
	var h map[int]interface{} // want `Use of 'interface{}' detected. Replace with 'any'.`

	var i map[any]interface{} // want `Use of 'interface{}' detected. Replace with 'any'.`
	var j map[interface{}]any // want `Use of 'interface{}' detected. Replace with 'any'.`

	fmt.Println(a, b, c, d, e, f, g, h, i, j)
}

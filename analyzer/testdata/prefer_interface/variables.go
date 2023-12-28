package testdata

import "fmt"

func preferInterfaceVariables() {
	var a any // want `Use of 'any' detected. Replace with 'interface{}'.`
	var b interface{}

	var c []any // want `Use of 'any' detected. Replace with 'interface{}'.`
	var d []interface{}

	var e map[any]int // want `Use of 'any' detected. Replace with 'interface{}'.`
	var f map[interface{}]int

	var g map[int]any // want `Use of 'any' detected. Replace with 'interface{}'.`
	var h map[int]interface{}

	var i map[any]interface{} // want `Use of 'any' detected. Replace with 'interface{}'.`
	var j map[interface{}]any // want `Use of 'any' detected. Replace with 'interface{}'.`

	fmt.Println(a, b, c, d, e, f, g, h, i, j)
}

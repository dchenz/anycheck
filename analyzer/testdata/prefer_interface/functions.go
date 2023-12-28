package testdata

func preferInterfaceFunctions1() any { // want `Use of 'any' detected. Replace with 'interface{}'.`
	return nil
}

func preferInterfaceFunctions2() (
	any, // want `Use of 'any' detected. Replace with 'interface{}'.`
	interface{},
) {
	return nil, nil
}

func preferInterfaceFunctions3(a any) { // want `Use of 'any' detected. Replace with 'interface{}'.`
}

func preferInterfaceFunctions4(
	a any, // want `Use of 'any' detected. Replace with 'interface{}'.`
	b interface{},
) {
}

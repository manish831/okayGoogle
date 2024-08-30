package main

func main() {}

func add(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	return nil
}

func genericAdd[T int | float64 | string](a, b T) T {
	return a+b // shows eeror if T is of type 'any'
	// because + is not defined for sabhi datatype.

	// So, + will be accepatable as long as T would be 
	// int 'or' float64 'or' string..
}
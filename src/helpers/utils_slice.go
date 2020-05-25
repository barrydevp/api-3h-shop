package helpers

func AppendInterfaceSlice(dst []interface{}, src ...interface{}) []interface{} {
	for _, item := range src {
		dst = append(dst, item)
	}

	return dst
}

package codec

func Join[T any](enc *JsonEncoder, key string, data []T, quote bool) {
	n := len(data)
	if n == 0 {
		enc.Raw(key, empty)
		return
	}

	enc.Key(key)
	enc.Arr("")

	for i := 0; i < n; i++ {
		item := data[i]
		if quote {
			enc.Val(strutil.String(item))
			enc.WriteByte(',')
			continue
		}

		enc.WriteString(strutil.String(item))
		enc.WriteByte(',')
	}
	enc.End("]")
	enc.WriteByte(',')
}
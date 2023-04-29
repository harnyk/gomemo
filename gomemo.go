package gomemo

func Memoize[
	IN comparable,
	OUT any,
](
	getter func(input IN) (OUT, error),
) func(input IN) (OUT, error) {
	var cache = make(map[IN]OUT)
	return func(input IN) (OUT, error) {
		if val, ok := cache[input]; ok {
			return val, nil
		}

		val, err := getter(input)
		if err != nil {
			return val, err
		}

		cache[input] = val
		return val, nil
	}
}

func MemoizeWithHasher[
	IN comparable,
	OUT any,
	KEY comparable,
](
	getter func(input IN) (OUT, error),
	hasher func(input IN) KEY,
) func(input IN) (OUT, error) {
	var cache = make(map[KEY]OUT)
	return func(input IN) (OUT, error) {
		hash := hasher(input)
		if val, ok := cache[hash]; ok {
			return val, nil
		}

		val, err := getter(input)
		if err != nil {
			return val, err
		}

		cache[hash] = val
		return val, nil
	}
}

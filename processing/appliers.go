package processing

var IntAppliers map[string]func(int) (int, error)

func init() {
	IntAppliers = make(map[string]func(int) (int, error))
	IntAppliers["double"] = func(i int) (int, error) {
		return i * 2, nil
	}

	IntAppliers["pow2"] = func(i int) (int, error) {
		return i * i, nil
	}

	IntAppliers["inc100th"] = func(i int) (int, error) {
		var res int
		for i := 0; i < 100000; i++ {
			res += 1
		}
		return res, nil
	}
}

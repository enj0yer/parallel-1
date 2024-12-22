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
	IntAppliers["prime"] = func(n int) (int, error) {
		if n < 2 {
			return 0, nil
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return 0, nil
			}
		}
		return 1, nil
	}
	IntAppliers["fib"] = func(n int) (int, error) {
		if n <= 1 {
			return n, nil
		}

		prev, curr := 0, 1
		for i := 2; i <= n; i++ {
			prev, curr = curr, prev+curr
		}
		return curr, nil
	}
}

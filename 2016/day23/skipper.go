package main

type registers struct {
	a int
	b int
	c int
	d int
}

type skipper struct {
	cache map[int][]map[string]int
}

func newSkipper() skipper {
	return skipper{
		cache: map[int][]map[string]int{},
	}
}

func (s *skipper) add(step int, registerValues map[string]int) (patternFound bool, intervals map[string]int) {
	// Get or create step cache.
	var stepCache []map[string]int
	var exists = false
	if stepCache, exists = s.cache[step]; !exists {
		stepCache = []map[string]int{}
	}

	// Append latest.
	stepCache = append(
		stepCache,
		map[string]int{
			"a": registerValues["a"],
			"b": registerValues["b"],
			"c": registerValues["c"],
			"d": registerValues["d"]})

	// Validate cache.
	const iterationsRequired = 500
	if len(stepCache) > iterationsRequired {
		intervals = map[string]int{
			"a": stepCache[1]["a"] - stepCache[0]["a"],
			"b": stepCache[1]["b"] - stepCache[0]["b"],
			"c": stepCache[1]["c"] - stepCache[0]["c"],
			"d": stepCache[1]["d"] - stepCache[0]["d"],
		}

		for i := 1; i < len(stepCache)-1; i++ {
			fitsPattern := stepCache[i+1]["a"]-stepCache[i]["a"] == intervals["a"] &&
				stepCache[i+1]["b"]-stepCache[i]["b"] == intervals["b"] &&
				stepCache[i+1]["c"]-stepCache[i]["c"] == intervals["c"] &&
				stepCache[i+1]["d"]-stepCache[i]["d"] == intervals["d"]

			// The gap between elements is not consistent.
			// Clear the cache.
			if !fitsPattern {
				stepCache = []map[string]int{}
				break
			}
		}
	}

	// Cache latest version.
	s.cache[step] = stepCache

	// At this point if the cache contains more than 5 items there is an established pattern.
	if len(stepCache) > iterationsRequired {
		return true, intervals
	}

	return false, map[string]int{}
}

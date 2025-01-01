package main

import (
	"fmt"
	"math"
)

type routeCacheKey struct {
	from rune
	to   rune
}

var routeCache = map[routeCacheKey]string{
	{'^', '^'}: "A",
	{'^', '>'}: "v>A",
	{'^', 'v'}: "vA",
	{'^', '<'}: "v<A",
	{'^', 'A'}: ">A",
	{'>', '^'}: "<^A",
	{'>', '>'}: "A",
	{'>', 'v'}: "<A",
	{'>', '<'}: "<<A",
	{'>', 'A'}: "^A",
	{'v', '^'}: "^A",
	{'v', '>'}: ">A",
	{'v', 'v'}: "A",
	{'v', '<'}: "<A",
	{'v', 'A'}: "^>A",
	{'<', '^'}: ">^A",
	{'<', '>'}: ">>A",
	{'<', 'v'}: ">A",
	{'<', '<'}: "A",
	{'<', 'A'}: ">>^A",
	{'A', '^'}: "<A",
	{'A', '>'}: "vA",
	{'A', 'v'}: "<vA",
	{'A', '<'}: "v<<A",
	{'A', 'A'}: "A",
}

func buildRouteCache() {

	// DEBUG: Remove later.
	return

	keys := "^>v<A"

	for _, from := range keys {
		for _, to := range keys {
			routes := getRoutes(directionKeypad, directionPoints, from, to)

			if len(routes) == 1 {
				fmt.Println("  ", string(from), string(to), "cached", routes[0])
				routeCache[routeCacheKey{from, to}] = routes[0]
				continue
			}

			if getLen(routes[0]) > getLen(routes[1]) {
				routeCache[routeCacheKey{from, to}] = routes[1]
				fmt.Println("  ", string(from), string(to), "cached", routes[1])
			} else {
				routeCache[routeCacheKey{from, to}] = routes[0]
				fmt.Println("  ", string(from), string(to), "cached", routes[0])
			}

		}
	}
}

func getLen(route string) int {
	result := math.MaxInt

	for _, level0 := range getDirectionalRoutes(route) {
		for _, level1 := range getDirectionalRoutes(level0) {
			for _, level2 := range getDirectionalRoutes(level1) {
				for _, level3 := range getDirectionalRoutes(level2) {
					if len(level3) < result {
						result = len(level3)
					}
				}
			}
		}
	}

	return result
}

func getExtendedCachedRoute(route string) string {
	var result string

	from := 'A'
	for _, to := range route {
		result += getCachedRoute(from, to)
		from = to
	}

	return result
}

func getCachedRoute(from, to rune) string {
	return routeCache[routeCacheKey{from, to}]
}

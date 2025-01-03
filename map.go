package main

func main() {
	m := map[string]int{
		"ichiziku": 1,
		"zukkini":  2,
		"mikan":    3,
		"nashi":    4,
		"ringo":    5,
	}
	for k, v := range m {
		println(k, v)
	}
}

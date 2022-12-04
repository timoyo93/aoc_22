package main

func CreatePriorities() map[string]int {
	m := make(map[string]int)
	i := 1
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, c := range chars {
		m[string(c)] = i
		i++
	}
	return m
}

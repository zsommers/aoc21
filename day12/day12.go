package day12

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func small(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return !unicode.IsUpper(r)
}

func readInput(input []string) map[string][]string {
	g := map[string][]string{}
	for _, s := range input {
		ns := strings.Split(s, "-")
		if len(ns) != 2 {
			panic(fmt.Sprintf("can't parse '%s'", s))
		}
		a := ns[0]
		b := ns[1]
		if _, ok := g[a]; !ok {
			g[a] = []string{b}
		} else {
			g[a] = append(g[a], b)
		}
		if _, ok := g[b]; !ok {
			g[b] = []string{a}
		} else {
			g[b] = append(g[b], a)
		}
	}
	return g
}

func copy(m map[string]bool) map[string]bool {
	n := map[string]bool{}
	for k, v := range m {
		n[k] = v
	}
	return n
}

func path(g map[string][]string, node string, visited map[string]bool) []string {
	paths := []string{}
	if small(node) {
		visited[node] = true
	}

	for _, destination := range g[node] {
		if destination == "end" {
			paths = append(paths, node+",end")
			continue
		}
		if visited[destination] {
			continue
		}

		for _, p := range path(g, destination, copy(visited)) {
			paths = append(paths, node+","+p)
		}
	}

	return paths
}

func path2(g map[string][]string, node string, visited map[string]bool, canDouble bool) []string {
	paths := []string{}

	for _, destination := range g[node] {
		if destination == "end" {
			paths = append(paths, node+",end")
			continue
		}
		if visited[destination] {
			continue
		}

		if canDouble && small(node) && node != "start" {
			for _, p := range path2(g, destination, copy(visited), false) {
				paths = append(paths, node+","+p)
			}
		}
		cp := copy(visited)
		if small(node) {
			cp[node] = true
		}
		for _, p := range path2(g, destination, cp, canDouble) {
			paths = append(paths, node+","+p)
		}
	}

	return paths
}

func A(input []string) int {
	g := readInput(input)
	paths := path(g, "start", map[string]bool{})
	return len(paths)
}

func B(input []string) int {
	g := readInput(input)
	paths := path2(g, "start", map[string]bool{}, true)
	ps := map[string]bool{}
	for _, p := range paths {
		ps[p] = true
	}
	return len(ps)
}

package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	data := 0
	got := generateRandomElements(data)
	want := 0
	if len(got) != want {
		t.Errorf("generateRandomElements(%v) = %d, want %d", data, len(got), want)
	}
	data = 10
	got = generateRandomElements(data)
	want = 10
	if len(got) != want {
		t.Errorf("generateRandomElements(%v) = %d, want %d", data, len(got), want)
	}
	data = 20
	got = generateRandomElements(data)
	for i := 0; i < len(got); i++ {
		if got[i] <= 0 {
			t.Errorf("generateRandomElements: элемент %d = %d, want > 0", i, got[i])
		}
	}
}

func TestMaximum(t *testing.T) {
	data := []int{}
	got := maximum(data)
	want := 0
	if got != want {
		t.Errorf("maximum(%v) = %d, want %d", data, got, want)
	}
	data = []int{1}
	got = maximum(data)
	want = 1
	if got != want {
		t.Errorf("maximum(%v) = %d, want %d", data, got, want)
	}
	data = []int{3, 7, 10, 15}
	got = maximum(data)
	want = 15
	if got != want {
		t.Errorf("maximum(%v) = %d, want %d", data, got, want)
	}
}

func TestMaxChunks(t *testing.T) {
	data := []int{5, 2, 9, 1, 7, 3, 8, 4, 6, 10}
	got := maxChunks(data)
	if got != 10 {
		t.Errorf("maxChunks(%v) = %d, want 10", data, got)
	}
}

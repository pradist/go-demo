package calculate

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAddTable(t *testing.T) {
	tests := []struct {
		x, y int
		want int
	}{
		{2, 3, 5},
		{1, 2, 3},
		{0, 0, 0},
	}
	for _, tt := range tests {
		got := Add(tt.x, tt.y)
		if got != tt.want {
			t.Errorf("got %d, want %d", got, tt.want)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10, 20)
	}
}

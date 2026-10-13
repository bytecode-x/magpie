package settings

import "testing"

func TestDefaultWorkingWindow(t *testing.T) {
	s := Settings{}
	if got := s.Compact(); got != 512000 {
		t.Fatalf("default compaction window = %d; want 512000", got)
	}
	for _, n := range []int{0, 128000, 272000, 400000, 511999, 512000, 512001, 1000000} {
		if got, want := s.Working(n), min(n, 512000); got != want {
			t.Errorf("Working(%d) = %d; want %d", n, got, want)
		}
	}
	for _, n := range []int{272000, 400000, 512000, 600000} {
		s := Settings{CompactAt: n}
		if s.Compact() != n || s.Working(1000000) != n {
			t.Errorf("explicit compaction window %d: %d/%d", n, s.Compact(), s.Working(1000000))
		}
	}
	s = Settings{CompactAt: 272000, FullContext: true}
	if s.Compact() != 0 || s.Working(1000000) != 1000000 {
		t.Errorf("full window: %d/%d", s.Compact(), s.Working(1000000))
	}
}

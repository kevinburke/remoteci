package remoteci

import (
	"testing"
	"time"
)

func TestEffectiveCost(t *testing.T) {
	cost := GetEffectiveCost(1 * time.Hour)
	if cost != 7897 {
		t.Errorf("expected 1 hour cost to be %d, was %d", 7897, cost)
	}
	cost = GetEffectiveCost(30 * time.Minute)
	if cost != 3948 {
		t.Errorf("expected half hour cost to be %d, was %d", 3948, cost)
	}
	cost = GetEffectiveCost(2 * time.Hour)
	if cost != 15793 {
		t.Errorf("expected half hour cost to be %d, was %d", 15793, cost)
	}
}

package styles

import (
	"testing"
)

func TestGetBreathingBorder(t *testing.T) {
	c0 := GetBreathingBorder(0)
	c1 := GetBreathingBorder(1)
	if c0 == "" || c1 == "" {
		t.Errorf("expected non-empty border color, got %v, %v", c0, c1)
	}

	// Test cycling modulo
	count := len(BreathingColors)
	if count > 0 {
		cCycle := GetBreathingBorder(count)
		if cCycle != c0 {
			t.Errorf("expected cyclical color at index %d to match index 0, got %v vs %v", count, cCycle, c0)
		}
	}
}

func TestGetTelemetryWave(t *testing.T) {
	w0 := GetTelemetryWave(0)
	w1 := GetTelemetryWave(1)
	if w0 == "" || w1 == "" {
		t.Errorf("expected non-empty telemetry wave frames, got %q, %q", w0, w1)
	}

	// Test cycling modulo
	count := len(TelemetryWaves)
	if count > 0 {
		wCycle := GetTelemetryWave(count)
		if wCycle != w0 {
			t.Errorf("expected cyclical wave at index %d to match index 0, got %q vs %q", count, wCycle, w0)
		}
	}
}

func TestGetActivePanelBox(t *testing.T) {
	boxStatic := GetActivePanelBox(0, false)
	boxAnim := GetActivePanelBox(0, true)

	renderedStatic := boxStatic.Render("test")
	renderedAnim := boxAnim.Render("test")

	if len(renderedStatic) == 0 || len(renderedAnim) == 0 {
		t.Errorf("expected rendered boxes to contain content")
	}
}

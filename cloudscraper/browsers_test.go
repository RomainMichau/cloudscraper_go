package cloudscraper

import (
	"testing"
)

func TestUserAgentBrowserConf(t *testing.T) {
	err := fetchJA3()
	if err != nil {
		t.Errorf("TESTS failed: %s", err.Error())
	}
	type test struct {
		N      int
		Device Device
	}

	devices := []Device{
		Chrome,
		Firefox,
		Default,
		Safari,
		Mobile,
	}

	tests := make([]test, 0, len(devices))
	for i, device := range devices {
		tests = append(tests, test{
			N:      i + 1,
			Device: device,
		})
	}

	for _, tc := range tests {
		newCl, err := getUserAgents(tc.Device)
		if err != nil {
			t.Errorf("TEST [%d] with %s failed: %s", tc.N, tc.Device, err.Error())
			continue
		}
		if newCl.UserAgent == "" || newCl.Ja3 == "" {
			t.Errorf("TEST [%d] with %s returned empty fields: %+v", tc.N, tc.Device, newCl)
		}
		t.Logf("TEST [%d] with %s succeeded. BrowserConf struct: %+v", tc.N, tc.Device, newCl)
	}
}

func TestFetchJA3(t *testing.T) {
	err := fetchJA3()
	if err != nil {
		t.Errorf("TEST failed: %s", err.Error())
	}

	t.Logf("TEST succeed: %s", recourseJA3)
}

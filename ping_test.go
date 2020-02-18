package ping

import (
	"testing"
)

func TestPing(t *testing.T) {
	host := "www.google.com"
	actual := Ping(host)
	expected := true

	if actual != expected {
		t.Fatal("Expected successful ping for "+host+" but failed", actual, expected)
	}

	host = "172.16.88.88"
	actual = Ping(host)
	expected = false

	if actual != expected {
		t.Fatal("Expected failed ping for "+host+" but succeeded", actual, expected)
	}
}

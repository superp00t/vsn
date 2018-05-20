package vsn

import "testing"

func TestVersion(t *testing.T) {
	v1 := []Version{
		"2.33.5",
		"1.0.0",
		"7.0.0",
		"7.0.1",
		"7.2.0",
		"6.0.0",
	}

	v1s := SortVersions(v1)
	if v1s[0] != "7.2.0" {
		t.Fatal("failed")
	}

	if Version("3.4").IsLessThan("3.5") == false {
		t.Fatal("failed")
	}

	if Version("3.4.5").IsGreaterThan("3.4") == false {
		t.Fatal("failed")
	}
}

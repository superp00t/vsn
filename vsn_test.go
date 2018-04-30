package vsn

import "fmt"
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

	fmt.Println(SortVersions(v1))
}

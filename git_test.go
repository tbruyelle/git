package git

import (
	"testing"
)

func TestLog(t *testing.T) {

	commits, err := Log("f3c2fb", "HEAD")

	if err != nil {
		t.Fatalf("Error occured during Log %v", err)
	}
	t.Errorf("Found commits %v\n", commits)
}

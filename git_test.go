package git

import (
	"reflect"
	"testing"
)

func TestLog(t *testing.T) {

	commits, err := Log("f3c2fb", "HEAD")

	if err != nil {
		t.Fatalf("Error occured during Log %v", err)
	}
	if size := len(commits); size != 1 {
		t.Fatalf("Log returns %d commits, want 1", size)
	}
	want := Commit{Ref: "e83f98a", Message: "Start testing"}
	if !reflect.DeepEqual(commits[0], want) {
		t.Errorf("Log returned %v, want %v", commits[0], want)
	}
}

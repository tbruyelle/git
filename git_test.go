package git

import (
	"reflect"
	"testing"
)

func TestBranch(t *testing.T) {
	want := "master"

	branch, err := Branch()

	if err != nil {
		t.Fatalf("Error occured during Branch %v", err)
	}
	if branch != want {
		t.Errorf("Branch returns %v, want %v", branch, want)
	}
}

func TestRevParse(t *testing.T) {
	want := "428e743ec249d734bdcb1fcbb1b603dd752f9687"

	ref, err := RevParse("428e743ec")

	if err != nil {
		t.Fatalf("Error occurred during RevParse %v", err)
	}
	if ref != want {
		t.Errorf("RevParse returns %v, want %v", ref, want)
	}
}

func TestHasLocalDiff(t *testing.T) {
	_, err := HasLocalDiff()

	if err != nil {
		t.Fatalf("Error occurred during HasLocalDiff %v", err)
	}
}

func TestRefExists(t *testing.T) {
	ref := "moncul"
	exist, err := RefExists(ref)
	if err != nil {
		t.Fatalf("Error occured during RefExists %v", err)
	}
	if exist {
		t.Errorf("Ref %s shoudn't exist", ref)
	}

	ref = "f3c2fb"
	exist, err = RefExists(ref)
	if err != nil {
		t.Fatalf("Error occured during RefExists %v", err)
	}
	if !exist {
		t.Errorf("Ref %s shoud exist", ref)
	}
}

func TestLog(t *testing.T) {

	commits, err := Log("f3c2fb", "6cbe88b")

	if err != nil {
		t.Fatalf("Error occured during Log %v", err)
	}
	if size := len(commits); size != 2 {
		t.Fatalf("Log returns %d commits, want 2", size)
	}
	want := []Commit{
		Commit{Ref: "6cbe88b", Message: "Test Log"},
		Commit{Ref: "e83f98a", Message: "Start testing"},
	}
	if !reflect.DeepEqual(commits, want) {
		t.Errorf("Log returned %v, want %v", commits, want)
	}
}

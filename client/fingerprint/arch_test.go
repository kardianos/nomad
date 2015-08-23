package fingerprint

import (
	"testing"

	"github.com/hashicorp/nomad/nomad/structs"
)

func TestArchFingerprint(t *testing.T) {
	f := NewArchFingerprint(testLogger())
	node := &structs.Node{
		Attributes: make(map[string]string),
	}
	ok, err := f.Fingerprint(node)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if !ok {
		t.Fatalf("should apply")
	}
	if node.Attributes["arch"] == "" {
		t.Fatalf("missing arch")
	}
}
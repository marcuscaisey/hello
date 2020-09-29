package bye

import "testing"

func TestHello(t *testing.T) {
	want := "Bye, marcus."
	name := "marcus"
	if got := Bye(name); got != want {
		t.Errorf("Bye(%q) = %q, got %q", name, want, got)
	}
}

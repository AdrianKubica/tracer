package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Funkcja New nie moze zwraca nil")
	} else {
		tracer.Trace("Witamy w pakiecie trace")
		if buf.String() != "Witamy w pakiecie trace\n" {
			t.Errorf("Metoda Trace nie powinna wyświetlić: %s", buf.String())
		}
	}
}

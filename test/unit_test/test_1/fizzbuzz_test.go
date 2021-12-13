package fizzbuzz // --> (1)
import (
	"testing" // --> (2)
)

func TestInput1ShouldBeDisplay1(t *testing.T) { // --> (3)
	v := FizzBuzz(1) // --> (4)
	if "1" != v {    // --> 5
		t.Error("fizzbuzz of 1 should be '1' but have", v) // --> (6)
	}
}

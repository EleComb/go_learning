package __loop

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1== 1; a {
		t.Log("1==1")
	}

	//if v, err := someFunc(); err == nil {
	//	t.Log("1==1")
	//} else {
	//	t.Log("2==2")
	//}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even1")
		case i%2 == 1:
			t.Log("Odd1")
		default:
			t.Log("unknown")
		}
	}
}

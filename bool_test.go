package as

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// OnOff is a simple fmt.Stringer implementation
type OnOff struct{ value bool }

func (o OnOff) String() (repr string) {
	switch o.value {
	case true:
		repr = "ON"
	case false:
		repr = "OFF"
	}
	return
}

func TestBool(t *testing.T) {
	var test = func(inp interface{}, expValue, expOk bool) {
		var actValue, actOk = Bool(inp)
		var msg = fmt.Sprintf("testing %#v, expect (%v, %v)", inp, expValue, expOk)
		assert.Equal(t, expValue, actValue, msg)
		assert.Equal(t, expOk, actOk, msg)
	}
	const OK = true
	const bad = false

	test(true, true, OK)
	test(false, false, OK)
	test(1, true, OK)
	test(0, false, OK)
	test(123, bad, bad)
	test(1.0, true, OK)
	test(0.0, false, OK)
	test(0.5, bad, bad)
	test(" True ", true, OK)
	test("  faLsE", false, OK)
	test("YES", true, OK)
	test(" NO ", false, OK)
	test("On", true, OK)
	test("OFF", false, OK)
	test("ENABLE", true, OK)
	test("disable", false, OK)
	test("enabled", true, OK)
	test("DISABLED", false, OK)
	test("unknown", bad, bad)
	test("1", true, OK)
	test("0  ", false, OK)
	test("-1", bad, bad)
	test(OnOff{true}, true, OK)
	test(OnOff{false}, false, OK)
}

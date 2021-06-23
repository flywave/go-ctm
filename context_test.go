package ctm

import "testing"

func TestCTMReader(t *testing.T) {
	m := &Mesh{ctx: NewContext(CTM_IMPORT)}
	m.ctx.Load("./testdata/test.ctm")

	if m.GetVertCount() == 0 {
		t.FailNow()
	}
}

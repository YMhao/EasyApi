package swagger

import "testing"

func TestAvoidRe(t *testing.T) {
	t.Log(AvoidRepeatMap.GetTypeName("xx", "abc"))
	t.Log(AvoidRepeatMap.GetTypeName("yy", "abc"))
	t.Log(AvoidRepeatMap.GetTypeName("yy", "abc"))
	t.Log(AvoidRepeatMap.GetTypeName("xx", "abc"))
	t.Log(AvoidRepeatMap.GetTypeName("zz", "abc"))
	t.Log(AvoidRepeatMap.GetTypeName("zz", "abc"))
}

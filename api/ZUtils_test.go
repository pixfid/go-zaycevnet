package api

import "testing"

func TestMD5Hash(t *testing.T) {
	s := "test"
	c := "098f6bcd4621d373cade4e832627b4f6"
	h := MD5Hash(s)

	if h != c {
		t.Errorf("Hash of string %#v should be %#v, not %#v.", s, c, h)
	}
}

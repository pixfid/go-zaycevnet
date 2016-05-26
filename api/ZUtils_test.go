/*
	Copyright (C) 2016  <Semchenko Aleksandr>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.If not, see <http://www.gnu.org/licenses/>.
*/

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

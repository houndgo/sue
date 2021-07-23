package utils

import "testing"

func TestRandomString(t *testing.T) {

 	str := RandomString(32)
 	t.Log(str)
}
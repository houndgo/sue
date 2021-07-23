package validator

import (
	"regexp"	"testing"
)


func TestIsMail(t *testing)  {
	var mail = "suihuashneg@gmail.com"

	is := IsIsEmailValid(mail)
	
	t.Log(is)
}
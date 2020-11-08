package web

import "testing"

func TestIsValidModuleName(t *testing.T) {
	t.Log(IsValidModuleName(""))
	t.Log(IsValidModuleName("a"))
	t.Log(IsValidModuleName("abc"))
	t.Log(IsValidModuleName("0abc"))
	t.Log(IsValidModuleName("abc."))
	t.Log(IsValidModuleName("abc.a"))
	t.Log(IsValidModuleName("abc.b//ss"))
	t.Log(IsValidModuleName("abc..b.//"))
}

func TestIsValidAppName(t *testing.T) {
	t.Log(IsValidAppName(""))
	t.Log(IsValidAppName("a"))
	t.Log(IsValidAppName("abc"))
	t.Log(IsValidAppName("0abc"))
	t.Log(IsValidAppName("abc."))
	t.Log(IsValidAppName("abc.a"))
	t.Log(IsValidAppName("abc.b/ss"))
	t.Log(IsValidAppName("abc..b.//"))
	t.Log(IsValidAppName("abc4_s"))
}

func TestIsValidDockerName(t *testing.T) {
	t.Log(IsValidDockerName(""))
	t.Log(IsValidDockerName("a"))
	t.Log(IsValidDockerName("abc"))
	t.Log(IsValidDockerName("0abc"))
	t.Log(IsValidDockerName("abc."))
	t.Log(IsValidDockerName("abc4_s"))
	t.Log(IsValidDockerName("abC4_s"))
}

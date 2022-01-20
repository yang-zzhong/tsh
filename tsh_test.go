package tsh

import "testing"

func TestMD5(t *testing.T) {
	if MD5("helloworld") != "fc5e038d38a57032085441e7fe7010b0" {
		t.Fatalf("md5 error")
	}
}

package main

import "testing"

func TestConctUrl(t *testing.T) {
	cases := []struct {
		in struct { user, pass, senha string}, 
		want string
	}{
		{{"root", "root", "localhost.com"}, "http://root:root@localhost.com"},
		{{"admin", "admin", "localhost.com"}, "http://admin:admin@localhost.com"},
	}

	for _, c := range cases {

	}

}

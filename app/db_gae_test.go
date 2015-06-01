// +build appengine

package oauthinator

import (
	"testing"

	"appengine/aetest"
)

func TestPutUser(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	u := &User{
		Login: "wolfeidau",
		Email: "mark@wolfe.id.au",
		Name:  "Mark Wolfe",
	}

	k := u.Key(c)

	if k.String() != "/User,wolfeidau" {
		t.Fatal("unexpected key", k)
	}

	err = putUser(c, u)

	if err != nil {
		t.Fatal(err)
	}

}

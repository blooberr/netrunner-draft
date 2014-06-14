package draft

import (
  "testing"
)

func TestNewName(t *testing.T) {
  t.Logf("prefix - %#v \n", NamePrefix)
  t.Logf("suffix - %#v \n", NameSuffix)

  t.Logf("random name - %s \n", CreateName())
}



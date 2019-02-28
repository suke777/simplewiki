package wiki

import (
	"bytes"
	"testing"
)

func TestNl2Br(t *testing.T) {
	expect := []byte("abc<br>テスト1<br>テスト２")
	actual := nl2Br([]byte("abc\nテスト1\nテスト２"))
	if !bytes.Equal(expect, actual) {
		t.Errorf("%s != %s", expect, actual)
	}
}

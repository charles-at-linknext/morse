package morse_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alwindoss/morse"
)

func TestEncode(t *testing.T) {
	h := morse.NewHacker()
	r := strings.NewReader("Name")
	data, err := h.Encode(r)
	if err != nil {
		t.Fail()
	}
	expected := "-. .- -- ."
	if string(data) != expected {
		t.Errorf("Expected the value to be %s but was %s", expected, string(data))
		t.Fail()
	}
}

func TestEncodeWithSentence(t *testing.T) {
	h := morse.NewHacker()
	r := strings.NewReader("My Name is Alwin Doss")
	data, err := h.Encode(r)
	if err != nil {
		t.Fail()
	}
	fmt.Println(string(data))
	expected := "-- -.-- / -. .- -- . / .. ... / .- .-.. .-- .. -. / -.. --- ... ..."
	if string(data) != expected {
		t.Errorf("Expected the value to be %s but was %s", expected, string(data))
		t.Fail()
	}
}

func TestDecode(t *testing.T) {
	h := morse.NewHacker()
	r := strings.NewReader("-- -.-- / -. .- -- . / .. ... / .- .-.. .-- .. -. / -.. --- ... ...")
	alphaNum, err := h.Decode(r)
	if err != nil {
		t.Fail()
	}
	expected := "MY NAME IS ALWIN DOSS"
	if string(alphaNum) != expected {
		t.Errorf("Expected the value to be %s but was %s", expected, string(alphaNum))
		t.Fail()
	}
}

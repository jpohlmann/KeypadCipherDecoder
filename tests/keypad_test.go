package tests

import (
	"github.com/jpohlmann/KeypadCipherDecoder/model"
	"testing"
)

func TestDecoder(t *testing.T) {
	var code = "DR2UL2UL3DR1 DC2SC3SL3SR2DR1 DR1SC3 DC2UC2SR3SR3"
	var answer = "THIS WORKS SO WELL"
	decoderKeypad := model.GetKeyPad(0)
	testAnswer := model.DecodeMessage(decoderKeypad, code)
	if (testAnswer != answer) {
		t.Errorf("%s does not meet the expected value %s", testAnswer, answer)
	}
}
func TestGetBestDecoding(t *testing.T) {
	var code = "DR2UL2UL3DR1 DC2SC3SL3SR2DR1 DR1SC3 DC2UC2SR3SR3\n"
	var answer = "THIS WORKS SO WELL"
	testAnswer := model.GetBestDecoding(code)
	if (testAnswer != answer) {
		t.Errorf("%s does not meet the expected value %s", testAnswer, answer)
	}
}
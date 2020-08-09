package tests

import (
	"github.com/jpohlmann/KeypadCipherDecoder/model"
	"testing"
)

func TestDecoder(t *testing.T) {
	var code = "DL2UR2UR3DL1 DC2CC3CR3CL2DL1 DL1CC3 DC2UC2CL3CL3"
	var answer = "THIS WORKS SO WELL"
	decoderKeypad := model.GetKeyPad(0)
	testAnswer := model.DecodeMessage(decoderKeypad, code)
	if (testAnswer != answer) {
		t.Errorf("%s does not meet the expected value %s", testAnswer, answer)
	}
}
func TestGetBestDecoding(t *testing.T) {
	var code = "DL2UR2UR3DL1 DC2CC3CR3CL2DL1 DL1CC3 DC2UC2CL3CL3"
	var answer = "THIS WORKS SO WELL"
	testAnswer := model.GetBestDecoding(code)
	if (testAnswer != answer) {
		t.Errorf("%s does not meet the expected value %s", testAnswer, answer)
	}
}
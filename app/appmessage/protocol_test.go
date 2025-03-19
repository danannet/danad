// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package appmessage

import "testing"

// TestServiceFlagStringer tests the stringized output for service flag types.
func TestServiceFlagStringer(t *testing.T) {
	tests := []struct {
		in   ServiceFlag
		want string
	}{
		{0, "0x0"},
		{SFNodeNetwork, "SFNodeNetwork"},
		{SFNodeGetUTXO, "SFNodeGetUTXO"},
		{SFNodeBloom, "SFNodeBloom"},
		{SFNodeXthin, "SFNodeXthin"},
		{SFNodeBit5, "SFNodeBit5"},
		{SFNodeCF, "SFNodeCF"},
		{0xffffffff, "SFNodeNetwork|SFNodeGetUTXO|SFNodeBloom|SFNodeXthin|SFNodeBit5|SFNodeCF|0xffffffc0"},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestDanaNetStringer tests the stringized output for dana net types.
func TestDanaNetStringer(t *testing.T) {
	tests := []struct {
		in   DanaNet
		want string
	}{
		{Mainnet, "Mainnet"},
		{Testnet, "Testnet"},
		{Simnet, "Simnet"},
		{0xffffffff, "Unknown DanaNet (4294967295)"},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// DanaMainnetPrivate is the version that is used for
// dana mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var DanaMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// DanaMainnetPublic is the version that is used for
// dana mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var DanaMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// DanaTestnetPrivate is the version that is used for
// dana testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var DanaTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// DanaTestnetPublic is the version that is used for
// dana testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var DanaTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// DanaDevnetPrivate is the version that is used for
// dana devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var DanaDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// DanaDevnetPublic is the version that is used for
// dana devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var DanaDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// DanaSimnetPrivate is the version that is used for
// dana simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var DanaSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// DanaSimnetPublic is the version that is used for
// dana simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var DanaSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case DanaMainnetPrivate:
		return DanaMainnetPublic, nil
	case DanaTestnetPrivate:
		return DanaTestnetPublic, nil
	case DanaDevnetPrivate:
		return DanaDevnetPublic, nil
	case DanaSimnetPrivate:
		return DanaSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case DanaMainnetPrivate:
		return true
	case DanaTestnetPrivate:
		return true
	case DanaDevnetPrivate:
		return true
	case DanaSimnetPrivate:
		return true
	}

	return false
}

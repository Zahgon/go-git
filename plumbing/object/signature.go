package object

import (
	"io"

	"github.com/go-git/go-git/v6/plumbing"
)

const (
	signatureTypeUnknown signatureType = iota
	signatureTypeOpenPGP
	signatureTypeX509
	signatureTypeSSH
)

var (
	openPGPSignatureFormat = signatureFormat{
		[]byte("-----BEGIN PGP SIGNATURE-----"),
		[]byte("-----BEGIN PGP MESSAGE-----"),
	}

	x509SignatureFormat = signatureFormat{
		[]byte("-----BEGIN SIGNED MESSAGE-----"),
	}

	sshSignatureFormat = signatureFormat{
		[]byte("-----BEGIN SSH SIGNATURE-----"),
	}
)

var knownSignatureFormats = map[signatureType]signatureFormat{
	signatureTypeOpenPGP: openPGPSignatureFormat,
	signatureTypeX509:    x509SignatureFormat,
	signatureTypeSSH:     sshSignatureFormat,
}

type signatureType int8

type signatureFormat [][]byte

func typeForSignature(b []byte) signatureType {
	_ = "STUB: not implemented"
	return *new(signatureType)
}

func parseSignedBytes(b []byte) (int, signatureType) {
	_ = "STUB: not implemented"
	return 0, *new(signatureType)
}

func countSignatureBlocks(b []byte) int { _ = "STUB: not implemented"; return 0 }

func isSignatureHeader(line []byte) bool { _ = "STUB: not implemented"; return false }

func stripObjectSignatures(dst, src plumbing.EncodedObject, objType plumbing.ObjectType) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func stripHeaderSignatures(w io.Writer, r io.Reader) error { _ = "STUB: not implemented"; return nil }

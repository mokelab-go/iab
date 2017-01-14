package android

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	rj "github.com/fkmhrk-go/rawjson"
)

// Verifier for Android in-app billing receipt
type Verifier struct {
	publicKey *rsa.PublicKey
}

// NewVerifier creates Verifier. Pass Base64 public key that
// you can get in Developer console.
func NewVerifier(publicKey string) *Verifier {
	key, err := readPublicKey("-----BEGIN PUBLIC KEY-----\n" + publicKey + "\n-----END PUBLIC KEY-----")
	if err != nil {
		return nil
	}
	return &Verifier{
		publicKey: key,
	}
}

// Verify receipt JSON string
func (o *Verifier) Verify(src, signature string) (Receipt, error) {
	// src to digest
	h := sha1.New()
	h.Write([]byte(src))
	digest := h.Sum(nil)
	// decode signature
	rawSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return Receipt{}, newVerifyError(SignatureDecodeFailed, "Failed to decode signature", err)
	}
	// verify
	err = rsa.VerifyPKCS1v15(o.publicKey, crypto.SHA1, digest, rawSignature)
	if err != nil {
		return Receipt{}, newVerifyError(VerificationFailed, "Verification failed", err)
	}
	// parse
	receipt, err := o.parse(src)
	if err != nil {
		return Receipt{}, newVerifyError(ParseFailed, "Failed to parse receipt JSON", err)
	}
	return receipt, nil
}

func (o *Verifier) parse(src string) (Receipt, error) {
	json, err := rj.ObjectFromString(src)
	if err != nil {
		return Receipt{}, err
	}
	orderID, err := json.String("orderId")
	if err != nil {
		return Receipt{}, err
	}
	packageName, err := json.String("packageName")
	if err != nil {
		return Receipt{}, err
	}
	productID, err := json.String("productId")
	if err != nil {
		return Receipt{}, err
	}
	developerPayload := ""
	if _, ok := json["developerPayload"]; ok {
		developerPayload, err = json.String("developerPayload")
		if err != nil {
			return Receipt{}, err
		}
	}
	purchaseState, err := json.Int("purchaseState")
	if err != nil {
		return Receipt{}, err
	}
	purchaseTime, err := json.Long("purchaseTime")
	if err != nil {
		return Receipt{}, err
	}
	purchaseToken, err := json.String("purchaseToken")
	if err != nil {
		return Receipt{}, err
	}
	return Receipt{
		OrderID:          orderID,
		PackageName:      packageName,
		ProductID:        productID,
		DeveloperPayload: developerPayload,
		PurchaseState:    purchaseState,
		PurchaseTime:     purchaseTime,
		PurchaseToken:    purchaseToken,
	}, nil
}

func readPublicKey(key string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, errors.New("bad block")
	}
	if block.Type != "PUBLIC KEY" {
		return nil, errors.New("wrong block type")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("key is not RSA public key")
	}

	return rsaKey, nil
}

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

type Verifier struct {
	publicKey *rsa.PublicKey
}

func NewVerifier(publicKey string) *Verifier {
	key, err := readPublicKey("-----BEGIN PUBLIC KEY-----\n" + publicKey + "\n-----END PUBLIC KEY-----")
	if err != nil {
		return nil
	}
	return &Verifier{
		publicKey: key,
	}
}

func (o *Verifier) Verify(src, signature string) (Receipt, error) {
	// src to digest
	h := sha1.New()
	h.Write([]byte(src))
	digest := h.Sum(nil)
	// decode signature
	rawSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return Receipt{}, newVerifyError(ERR_SignatureDecodeFailed, "Failed to decode signature", err)
	}
	// verify
	err = rsa.VerifyPKCS1v15(o.publicKey, crypto.SHA1, digest, rawSignature)
	if err != nil {
		return Receipt{}, newVerifyError(ERR_VerificationFailed, "Verification failed", err)
	}
	// parse
	receipt, err := o.parse(src)
	if err != nil {
		return Receipt{}, newVerifyError(ERR_ParseFailed, "Failed to parse receipt JSON", err)
	}
	return receipt, nil
}

func (o *Verifier) parse(src string) (Receipt, error) {
	json, err := rj.ObjectFromString(src)
	if err != nil {
		return Receipt{}, err
	}
	orderId, err := json.String("orderId")
	if err != nil {
		return Receipt{}, err
	}
	packageName, err := json.String("packageName")
	if err != nil {
		return Receipt{}, err
	}
	productId, err := json.String("productId")
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
		OrderId:          orderId,
		PackageName:      packageName,
		ProductId:        productId,
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

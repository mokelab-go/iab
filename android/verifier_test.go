package android

import (
	"testing"
)

func TestVerifier_OK(t *testing.T) {
	publicKey := `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDcgE5K4uCUycylqjl047n3wmo2
S4BRrNc4uK8PbWv8IwG+mIjr5ydmPuZUNnbNS9CJd4RjL45USlx20Ne4qzEbGoJ0
lCRh69QwdLrvfg5rK7MLiO5ppGcGBvil4Uqpd8sE3FOow+3HwJkvdz+aYbESfG+x
xSI3tJaufX5qePejTwIDAQAB`
	v := NewVerifier(publicKey)
	input := `{"orderId":"12999763169054705758.1371073745894165","packageName":"com.mokelab.dummyApp","productId":"dummy_app_item","purchaseTime":1367655044246,"purchaseState":0,"purchaseToken":"epprpkcryioqkjdpbmflohvh.AO-J1OwKL0OEHgT_H1hezDIw8pBZuT_6JfZHWiwBLtAXCPLlkjUwfszTqt59mBEulqp4WAsqPsJlG4T6nD-1Er53w9LicLloOTVOOyzwX0U02gLBH2ZS_WxQXlfLGvDDIdZoFkDsZwvN"}
`
	signature := `ko28oNiLOOJ1FOeZTJSj4I3U6t125X1OAz/IFPzrOLMDj7FXF3y+TZcY38VK58ZAIWHypgS0pKLisOYpeR+KPEtFNvEevNiUbsc/a6NnNfI+LyJ3FrB1weqOiYUgU3C0B03SRwXmcWB1cN/eac1fKNOsjxIW07CrRHDailN1lxM=`
	receipt, err := v.Verify(input, signature)
	if err != nil {
		t.Errorf("verification error : %s", err)
	}
	if receipt.OrderId != "12999763169054705758.1371073745894165" {
		t.Errorf("Wrong orderId : %s", receipt.OrderId)
	}
	if receipt.PackageName != "com.mokelab.dummyApp" {
		t.Errorf("Wrong packageName : %s", receipt.PackageName)
	}
	if receipt.ProductId != "dummy_app_item" {
		t.Errorf("Wrong productId : %s", receipt.ProductId)
	}
	if receipt.PurchaseState != 0 {
		t.Errorf("Wrong purchaseState : %d", receipt.PurchaseState)
	}
	if receipt.PurchaseTime != 1367655044246 {
		t.Errorf("Wrong purchaseTime : %d", receipt.PurchaseTime)
	}
	if receipt.PurchaseToken != "epprpkcryioqkjdpbmflohvh.AO-J1OwKL0OEHgT_H1hezDIw8pBZuT_6JfZHWiwBLtAXCPLlkjUwfszTqt59mBEulqp4WAsqPsJlG4T6nD-1Er53w9LicLloOTVOOyzwX0U02gLBH2ZS_WxQXlfLGvDDIdZoFkDsZwvN" {
		t.Errorf("Wrong purchaseToken : %s", receipt.PurchaseToken)
	}
}

func TestVerifier_Failed(t *testing.T) {
	publicKey := `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDcgE5K4uCUycylqjl047n3wmo2
S4BRrNc4uK8PbWv8IwG+mIjr5ydmPuZUNnbNS9CJd4RjL45USlx20Ne4qzEbGoJ0
lCRh69QwdLrvfg5rK7MLiO5ppGcGBvil4Uqpd8sE3FOow+3HwJkvdz+aYbESfG+x
xSI3tJaufX5qePejTwIDAQAB`
	v := NewVerifier(publicKey)
	input := `{"orderId":"12999763169054705758.1371073745894165","packageName":"com.mokelab.dummyApp","productId":"dummy_app_item","purchaseTime":1367655044246,"purchaseState":0,"purchaseToken":"epprpkcryioqkjdpbmflohvh.AO-J1OwKL0OEHgT_H1hezDIw8pBZuT_6JfZHWiwBLtAXCPLlkjUwfszTqt59mBEulqp4WAsqPsJlG4T6nD-1Er53w9LicLloOTVOOyzwX0U02gLBH2ZS_WxQXlfLGvDDIdZoFkDsZwvN_BROKEN_INPUT"}`
	signature := `ko28oNiLOOJ1FOeZTJSj4I3U6t125X1OAz/IFPzrOLMDj7FXF3y+TZcY38VK58ZAIWHypgS0pKLisOYpeR+KPEtFNvEevNiUbsc/a6NnNfI+LyJ3FrB1weqOiYUgU3C0B03SRwXmcWB1cN/eac1fKNOsjxIW07CrRHDailN1lxM=`
	_, err := v.Verify(input, signature)
	if err == nil {
		t.Errorf("err must not be nil")
		return
	}
	err2, ok := err.(*VerifyError)
	if !ok {
		t.Errorf("err must be VerifyError")
		return
	}
	if err2.Code != ERR_VerificationFailed {
		t.Errorf("Error code must be VerificationFailed but %d", err2.Code)
		return
	}
}

func TestVerifier_Payload(t *testing.T) {
	publicKey := `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDcgE5K4uCUycylqjl047n3wmo2
S4BRrNc4uK8PbWv8IwG+mIjr5ydmPuZUNnbNS9CJd4RjL45USlx20Ne4qzEbGoJ0
lCRh69QwdLrvfg5rK7MLiO5ppGcGBvil4Uqpd8sE3FOow+3HwJkvdz+aYbESfG+x
xSI3tJaufX5qePejTwIDAQAB`
	v := NewVerifier(publicKey)
	input := `{"orderId":"GPA.1111-2222-3333-44444","packageName":"com.mokelab.demoapp","productId":"demo_product_id","purchaseTime":1454561642505,"purchaseState":0,"developerPayload":"paylord","purchaseToken":"aaaaaaaapflgcdgpbbbbbbbb.AO-AAAABBBBBBRmJs7xFa5fyRS203c5pc_gu-djJL9cmMB5ZPq474O0kdVyVkqWCAanYDd6s6wMr1wGcBYEh0pg-qgnNp1L3qI8qjerBG_cW-AAAAAAAAvcKGJcEUKG3uAnEjeQq9JPg"}
`
	signature := `DQ5iGUYeUBNzby23mWVzxSJ275x8PkIF5icBw5OLPqBxZQhxE7lNwHbkLdUGKK73U6HVshdK0jnb8YfkpWYwVDAAQ+IYJVeaE7DFmONTRSFakg6aMNsU5P0T+3Ta38+BcNejVQtnMulCelDnR7ukLnLO3kuczgGuqDGWXiN7kt8=`
	receipt, err := v.Verify(input, signature)
	if err != nil {
		t.Errorf("err must be nil : %s", err)
		return
	}
	if receipt.DeveloperPayload != "paylord" {
		t.Errorf("developerPayload must be %s but %s", "paylord", receipt.DeveloperPayload)
		return
	}
}

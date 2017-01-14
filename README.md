# iab
This library provides a receipt verification feature. 

# Android

```go
package main

import (
        "fmt"
        "github.com/mokelab-go/iab/android"
)

func main() {
        // Receipt JSON string given by Android In-app billing response
        input := `{"orderId":"12999763169054705758.1371073745894165","packageName":"com.mokelab.dummyApp","productId":"dummy_app_item","purcha\
seTime":1367655044246,"purchaseState":0,"purchaseToken":"epprpkcryioqkjdpbmflohvh.AO-J1OwKL0OEHgT_H1hezDIw8pBZuT_6JfZHWiwBLtAXCPLlkjUwfszTqt59\
mBEulqp4WAsqPsJlG4T6nD-1Er53w9LicLloOTVOOyzwX0U02gLBH2ZS_WxQXlfLGvDDIdZoFkDsZwvN"}`

        // Receipt signature given by Android In-app billing response
        signature := `ko28oNiLOOJ1FOeZTJSj4I3U6t125X1OAz/IFPzrOLMDj7FXF3y+TZcY38VK58ZAIWHypgS0pKLisOYpeR+KPEtFNvEevNiUbsc/a6NnNfI+LyJ3FrB1weqO\
iYUgU3C0B03SRwXmcWB1cN/eac1fKNOsjxIW07CrRHDailN1lxM=`

        // You can get your public key on Developer console.
        publicKey := `__PUT_YOUR_PUBLIC_KEY__`
        v := android.NewVerifier(publicKey)
        receipt, err := v.Verify(input, signature)
        if err != nil {
                fmt.Printf("Failed to verify : %s", err)
                return
        }
        // Receipt is parsed as android.Receipt object.
        fmt.Printf("order id=%s\n", receipt.OrderId)
}

```

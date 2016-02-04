package android

type Receipt struct {
	OrderId          string
	PackageName      string
	ProductId        string
	DeveloperPayload string
	PurchaseState    int
	PurchaseTime     int64
	PurchaseToken    string
}

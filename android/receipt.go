package android

// Receipt of a purchase
type Receipt struct {
	OrderID          string
	PackageName      string
	ProductID        string
	DeveloperPayload string
	PurchaseState    int
	PurchaseTime     int64
	PurchaseToken    string
}

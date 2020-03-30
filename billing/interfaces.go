package billing

// CreditCard 결제에 사용할 카드 정보입니다.
type CreditCard struct {
	CardNo string
}

// ChargeResult 결제 결과 입니다.
type ChargeResult struct {
	WasSuccessful  bool
	Amount         int
	DeclineMessage string
}

// TransactionLog 결제 로그를 기록하는 interface입니다.
type TransactionLog interface {
	LogChargeResult(result *ChargeResult)
	LogConnectException(e error)
}

// CreditCardProcessor 카드 결제를 담당하는 interface 입니다.
type CreditCardProcessor interface {
	Charge(card CreditCard, amount int) (*ChargeResult, error)
}

// PizzaOrder Pizza 를 주문하기 위한  request 입니다.
type PizzaOrder struct {
	Amount int
}

// Receipt 결제후 손님에게 줄 영수증 interface 입니다.
type Receipt interface {
}

// BillingService pizza 를 card로 결제하고 영수증을 return 하는 서비스 입니다.
type BillingService interface {
	ChargeOrder(order PizzaOrder, card CreditCard) Receipt
}

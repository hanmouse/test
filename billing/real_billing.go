package billing

import (
	"net/http"

	"camel.uangel.com/ua5g/ulib.git/errcode"
	"camel.uangel.com/ua5g/ulib.git/testhelper"
	"camel.uangel.com/ua5g/ulib.git/uconf"
	"camel.uangel.com/ua5g/ulib.git/ulog"
)

// 결제에 성공 했을 때의 영수증 입니다.
type ReceiptSuccess struct {
	Amount int
}

// 결제가 거절 되었을 때의 영수증 입니다.
type ReceiptDeclined struct {
	DeclineMessage string
}

// 시스템 오류로 결제를 하지 못했을 때의 영수증 입니다.
type ReceiptSystemFailure struct {
	failureMessage string
}

// 실제 BillingService 구현 struct 입니다.
type RealBillingService struct {
	processor      CreditCardProcessor
	transactionLog TransactionLog
}

// pizza 결제 구현입니다.
func (r *RealBillingService) ChargeOrder(order PizzaOrder, card CreditCard) Receipt {

	// CreditCardProcessor 의 Charge를 호출 합니다.
	result, err := r.processor.Charge(card, order.Amount)
	if err != nil {
		// 카드사에 연결하지 못했으면 오류 로그를 남기고
		r.transactionLog.LogConnectException(err)
		// 실패 영수증을 리턴합니다.
		return &ReceiptSystemFailure{err.Error()}
	}

	// 성공하면 성공 로그를 남기고
	r.transactionLog.LogChargeResult(result)

	// 결제가 되었으면 결제 영수증을 리턴합니다.
	if result.WasSuccessful {
		return &ReceiptSuccess{order.Amount}
	}

	// 결제가 되지 않았다면  거절 영수증을 리턴합니다.
	return &ReceiptDeclined{result.DeclineMessage}
}

// RealBillingService의 생성자 입니다.  Constructor Binding을 할 예정입니다.
func NewRealBillingService(processor CreditCardProcessor, transactionLog TransactionLog) *RealBillingService {
	return &RealBillingService{processor, transactionLog}
}

// 시험에 사용할 mockup ( fake )  카드 결제기 입니다.
type FakeCreditCardProcessor struct {
	// TestBehavior 를 embed 합니다.
	testhelper.TestBehavior
}

func (r *FakeCreditCardProcessor) Charge(card CreditCard, amount int) (*ChargeResult, error) {

	// connect-error 상황이면
	if r.BehaviorEnabled("connect-error") {
		// 503 에러를 리턴합니다.
		return nil, errcode.ServiceUnavailable("can't connect to card service")
	}

	// 잔액이 없는 상황이면
	if r.BehaviorEnabled("not-enough-credit") {
		// 거절 사유를 리턴합니다.
		return &ChargeResult{
			WasSuccessful:  false,
			DeclineMessage: "sorry. not enough credit",
		}, nil
	}

	// 결제가 성공되었습니다.
	return &ChargeResult{
		WasSuccessful: true,
		Amount:        amount,
	}, nil

}

// FakeCreditCardProcessor의 생성자입니다. constructor binding을 할 예정입니다.
func NewFakeCreditCardProcessor() *FakeCreditCardProcessor {
	return &FakeCreditCardProcessor{}
}

// TransactionLog를 logger 를 사용해서 남기는 구현입니다.
type LoggerTransactionLog struct {
	chargeLogger ulog.Logger
	errorLogger  ulog.Logger
}

// chargeLogger에 info로 로그를 남깁니다.
func (r *LoggerTransactionLog) LogChargeResult(result *ChargeResult) {
	r.chargeLogger.Info("Charge result = %v", *result)
}

// 카드사에 접속 하지 못했을 경우 사용되는 함수입니다.
func (r *LoggerTransactionLog) LogConnectException(e error) {
	// error에서  http code를 가져 옵니다.  code가 없는경우는 500을 사용합니다.
	code := errcode.GetCode(e, http.StatusInternalServerError)

	// 에러 유형을 가져옵니다. 에러 유형이 없는 경우 http의 기본 text를 사용합니다.
	title := errcode.GetTitle(e, http.StatusText(code))

	// errorLogger에 로그를 남깁니다.
	r.errorLogger.Warn("credit service error. code : %d, type : %s , reason : %s", code, title, e.Error())
}

// LoggerTransactionLog의 생성자입니다.  constructor binding을 할 예정입니다.
func NewLoggerTransactionLog(conf uconf.Config, lf ulog.LoggerFactory) *LoggerTransactionLog {
	chargeLoggerName := conf.GetString("transaction-log.charge-logger", "com.uangel.billing.charge")
	chargeLogger := lf.GetLogger(chargeLoggerName)

	errLoggerName := conf.GetString("transaction-log.error-logger", "com.uangel.billing.error")
	errLogger := lf.GetLogger(errLoggerName)

	return &LoggerTransactionLog{chargeLogger, errLogger}
}

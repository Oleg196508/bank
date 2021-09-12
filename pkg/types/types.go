package types

// Money представляет собой денежную сумму в минимальных единицах
type Money int64

//  Currency  представляет код валюты
type Currency string 

 const (
  	TJS Currency = "TJS"
  	RUB Currency = "RUB"
  	USD Currency = "USD"
  )

// PAN представляет номер карты
type PAN string

//Category представляет собой категорию, в которой был совершён
//платёж(авто, аптеки, рестораны и т.д.)
type Category string

//Status представляет собойтстатус платежа.
type Status string

//Предопределённые статусы платежей
const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

type Payment struct {
	ID int
	Amount Money
	Category Category
	Status Status
}

// Card представляет информацию о платёжной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}
//PaymentSource выбираются источники оплаты
type PaymentSource struct {
	Type      string
	Number    string
	Balance   Money
}
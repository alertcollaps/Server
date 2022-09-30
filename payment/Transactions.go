package payment

import (
	"Server/sql"
	"fmt"
)

func GetBalance(id int) (cash int, err error) {
	return sql.GetBalance(id)
}

func ExchangeBetweenUsers(idSender, idRecipient, sum int) (err error) {
	if sum < 0 {
		return fmt.Errorf("sum incorrect")
	}

	us1, _ := sql.Check(idSender)
	us2, _ := sql.Check(idRecipient)

	if !(us1 && us2) {
		return fmt.Errorf("no user or users exists")
	}

	if idRecipient == idSender {
		return
	}

	balanceSender, err := sql.GetBalance(idSender)
	if err != nil {
		return fmt.Errorf("error get balance")
	}

	if balanceSender-sum < 0 {
		return fmt.Errorf("not enough money")
	}

	err = SubtractionSum(idSender, sum)
	if err != nil {
		return
	}

	err = AddSum(idRecipient, sum)
	if err != nil {
		return
	}
	return
}

func AddSum(id int, sum int) (err error) {
	if sum < 0 {
		return fmt.Errorf("sum incorrect")
	}

	var curBalance int

	if b, _ := sql.Check(id); b {
		curBalance, err = sql.GetBalance(id)
		if err != nil {
			return
		}
	}

	err = sql.InsertBalance(id, curBalance+sum)
	if err != nil {
		return fmt.Errorf("no insert")
	}
	return
}

func SubtractionSum(id int, sum int) (err error) {
	if sum < 0 {
		return fmt.Errorf("sum incorrect")
	}
	if b, _ := sql.Check(id); !b {
		return fmt.Errorf("no user exists")
	}

	currBal, err := sql.GetBalance(id)
	if err != nil {
		return fmt.Errorf("no getBalance")
	}
	if currBal-sum < 0 {
		return fmt.Errorf("not enough money")
	}

	err = sql.InsertBalance(id, currBal-sum)
	if err != nil {
		return fmt.Errorf("no insert")
	}
	return

}

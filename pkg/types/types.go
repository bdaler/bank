package types

const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

type Money int64

type Category string

type Status string

type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

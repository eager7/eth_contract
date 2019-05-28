package tables

type DbMessage interface {
	GetInstance() interface{}
	SqlCommand() string
}

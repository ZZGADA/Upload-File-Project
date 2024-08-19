package bo

type TransactionStudy struct {
	Id      int64  `mapstructure:"id" gorm:"primarykey"`
	Name    string `mapstructure:"name"`
	Age     int32  `mapstructure:"age"`
	Address string `mapstructure:"address"`
}

func NewTransactionStudy(id int64, name string, age int32, address string) *TransactionStudy {
	return &TransactionStudy{Id: id, Name: name, Age: age, Address: address}
}

// TableName 指定表名
func (TransactionStudy) TableName() string {
	return "transaction_study"
}

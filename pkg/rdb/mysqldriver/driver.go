package mysqldriver

type (
	Driver struct {}
)

func New() *Driver {
	driver := &Driver{}
	return driver
}
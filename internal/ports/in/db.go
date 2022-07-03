package in

type DatabasePort interface {
	CloseConnection()
	RestartConnection()
}

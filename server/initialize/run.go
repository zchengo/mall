package initialize

func Run()  {
	LoadConfig()
	Mysql()
	Redis()
	Router()
}

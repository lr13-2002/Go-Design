package singleton

func init() {
	singleton = new(Singleton)
}

func GetHungryInstance() *Singleton {
	return singleton
}

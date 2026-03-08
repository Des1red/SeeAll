package devmode

var dev = false

func setDev(status bool) {
	dev = status
}

func isDev() bool { return dev }

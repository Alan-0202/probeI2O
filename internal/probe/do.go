package probe

import "I2Oprobe/internal/list"

func Do() {
	d := new(list.Dolist).Run()
	appClient := NewAppProbe()
	NewServer(d, appClient).Run()
}

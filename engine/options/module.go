package options

const (
	STATUS_READY	= iota
	STATUS_PLAYING
	STATUS_ENDING
)

func Path(name string) string {
	return "/var/gogame/logo/" + name + ".png"
}

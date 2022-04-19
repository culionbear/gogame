package options

const (
	STATUS_READY	= iota
	STATUS_PLAYING
	STATUS_ENDING
)

type Value interface {
	~ int | ~ string | ~ float64 | ~ float32 | ~ bool
}

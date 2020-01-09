package person

type Male interface {
	Get(id int64) int
	GetName(id int64) string
}

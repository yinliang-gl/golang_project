package person

type Male interface {
	GetId(id int64) int
	GetName(id int64) string
}

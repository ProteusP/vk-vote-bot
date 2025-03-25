package tarantool

type Vote struct {
	ID       string
	Creator  string
	Question string
	Options  map[string]string
	Votes    map[string]string
	Status   string
}

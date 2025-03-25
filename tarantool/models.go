package tarantool

type Vote struct {
	ID       string
	Creator  string
	Question string
	Options  map[string]int    // {"option": votes_count}
	Votes    map[string]string // {"user_id": "option"}
	Status   string            // "active" | "ended"
}

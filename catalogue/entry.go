package catalogue

type Entry struct {
	Name         string
	Owner        string
	Repo         string
	Role         string
	Protocol     string
	Format       string
	Dependencies struct {
		Critical    []interface{}
		Noncritical []interface{} `json:"non-critical"`
	}
	Environment []string
	Mievents    []string
	Metrics     []string
}

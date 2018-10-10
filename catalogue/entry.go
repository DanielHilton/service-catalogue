package catalogue

type Entry struct {
	Name         string
	Owner        string
	Repo         string
	Role         string
	Protocol     string
	Format       string
	Dependencies struct {
		Critical    []Dependency `json:"critical"`
		Noncritical []Dependency `json:"non-critical"`
	} `json:"dependencies"`
	Environment []string
	Mievents    []string
	Metrics     []string
}

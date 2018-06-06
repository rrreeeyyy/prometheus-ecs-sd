package discovery

type TargetGroup struct {
	Targets []string `json:"targets,omitempty"`
	Labels map[string][string] `json:"labels,omitempty"`
}

func main() {
	fmt.Println("vim-go")
}

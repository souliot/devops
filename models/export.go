package models

type Export struct {
	Targets []string `json:"targets"`
}

func (e *Export) Node() (exs []*Export) {
	exs = make([]*Export, 0)
	ex := &Export{
		Targets: []string{"192.168.0.252:8080"},
	}
	exs = append(exs, ex)
	return
}

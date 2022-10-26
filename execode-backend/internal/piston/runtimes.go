package piston

type Runtime struct {
	Language string   `json:"language"`
	Version  string   `json:"version"`
	Aliases  []string `json:"aliases"`
	Runtime  string   `json:"node"`
}

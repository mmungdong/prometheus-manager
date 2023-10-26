package strategy

type Strategy struct {
	Filename string   `json:"filename,omitempty"`
	Groups   []*Group `json:"groups,omitempty"`
}

type Group struct {
	Name  string  `json:"name,omitempty"`
	Rules []*Rule `json:"rules,omitempty"`
}

type Rule struct {
	Alert       string            `json:"alert,omitempty"`
	Expr        string            `json:"expr,omitempty"`
	For         string            `json:"for,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type Dir struct {
	Dir        string      `json:"dir,omitempty"`
	Strategies []*Strategy `json:"strategies,omitempty"`
}

// GetDir ...
func (l *Dir) GetDir() string {
	if l == nil {
		return ""
	}
	return l.Dir
}

// GetStrategies ...
func (l *Dir) GetStrategies() []*Strategy {
	if l == nil {
		return nil
	}
	return l.Strategies
}

// GetFilename ...
func (l *Strategy) GetFilename() string {
	if l == nil {
		return ""
	}
	return l.Filename
}

// GetGroups ...
func (l *Strategy) GetGroups() []*Group {
	if l == nil {
		return nil
	}
	return l.Groups
}

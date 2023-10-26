package alert

type (
	// Alert ...
	Alert struct {
	}

	// Option ...
	Option func(*Alert)
)

// NewAlert new an Alert instance
func NewAlert(opts ...Option) *Alert {
	alert := &Alert{}
	for _, o := range opts {
		o(alert)
	}

	return alert
}

package ready

import "encoding/json"

type Status int

const (
	Ready Status = iota
	NotReady
)

type StatusReport struct {
	Overall    Status            `json:"overall"`
	Indicators map[string]Status `json:"indicators"`
}

type Indicator interface {
	Ready() Status
}

type Monitor struct {
	indicators map[string]Indicator
}

func NewMonitor(indicators map[string]Indicator) Monitor {
	return Monitor{indicators: indicators}
}

func (m Monitor) ReadyCheck() StatusReport {
	report := StatusReport{
		Overall: Ready,
		Indicators: make(map[string]Status, len(m.indicators)),
	}
	for name, indicator := range m.indicators {
		readyStatus := indicator.Ready()
		report.Indicators[name] = readyStatus
		if readyStatus == NotReady {
			report.Overall = NotReady
		}
	}
	return report
}

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Status) String() string {
	switch s {
	case Ready:
		return "READY"
	case NotReady:
		return "NOT_READY"
	default:
		return "INVALID"
	}
}

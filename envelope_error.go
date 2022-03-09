package rambda

import "encoding/json"

type EnvelopeError struct {
	Title    string `json:"title,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Type     string `json:"type,omitempty"`
	Envelope `json:"-"`
}

func (e EnvelopeError) MarshalJSON() ([]byte, error) {
	type envelope EnvelopeError // prevent recursion
	data, err := json.Marshal(envelope(e))
	if err != nil {
		return data, err
	}

	var payload map[string]json.RawMessage
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return data, err
	}

	for k, v := range e.Envelope {
		if _, ok := payload[k]; ok {
			continue
		}
		tmpData, err := json.Marshal(v)
		if err != nil {
			return tmpData, err
		}
		payload[k] = tmpData
	}

	return json.Marshal(payload)
}

package main

type Consumption struct {
}

type Provider interface {
	DoRequest(string) (bool, error)
}

type RiskServiceImpl struct {
	client Provider
}

func (r RiskServiceImpl) Validate(data string) error {

	valid, err := r.client.DoRequest(data)
	if err != nil {
		return NewFraudInfrastructureError(err)
	}

	if !valid {
		return NewFraudSuspectError()
	}

	return nil
}

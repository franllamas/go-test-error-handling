package main

type RiskService interface {
	Validate(string) (Consumption, error)
}

type UseCaseImpl struct {
	risk RiskService
}

func (u UseCaseImpl) Create() (Consumption, error) {
	authorizationData := "authorization authorizationData..."
	consumptionId := "eeeee"
	processorId := "zzzz"

	consumption, err := u.risk.Validate(authorizationData)
	if err != nil {
		u.logError(err, consumptionId, processorId)
		if _, ok := err.(FraudSuspectError); ok {
			return Consumption{}, err
		}
	}

	return consumption, nil
}

func (u UseCaseImpl) logError(err error, consumptionId, processorId string) {
	/* tags := [
	  request_id: 8736178361287361837926
		entity_type: consumption
		entity_id: eeeeee
		related_id: zzzzzz (id que le pone Bari)
		action: transactional_creation
	]
	// logger.Log(err.Error(), tags)
	*/
}

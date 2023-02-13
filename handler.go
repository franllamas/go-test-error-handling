package main

type UseCase interface {
	Create() (Consumption, error)
}

type Handler struct {
	useCase UseCase
}

func (h Handler) Handle() string {

	/*
		if badRequest {
			logError()...
		  return "invalid_message"
		}
	*/

	_, err := h.useCase.Create()
	switch err.(type) {
	case FraudSuspectError:
		return "fraud_suspect"
	default:
		return "unknown_error"
	}

}

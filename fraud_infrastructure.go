package main

type FraudInfrastructureError struct {
	message string
	cause   error
}

func NewFraudInfrastructureError(cause error) FraudInfrastructureError {
	return FraudInfrastructureError{
		message: "infrastructure error validating fraud",
		cause:   cause,
	}
}

func (e FraudInfrastructureError) Error() string {
	return e.message + " - " + e.cause.Error()
}

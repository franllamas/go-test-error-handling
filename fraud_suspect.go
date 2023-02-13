package main

type FraudSuspectError struct {
	message string
}

func NewFraudSuspectError() FraudSuspectError {
	return FraudSuspectError{
		message: "there is a suspicion of fraud",
	}
}

func (e FraudSuspectError) Error() string {
	return e.message
}

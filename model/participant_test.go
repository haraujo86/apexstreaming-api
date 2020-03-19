package model

import "testing"

const errorMessage = "Valor esperado %v, mas o resultado encontrado foi %v."
const participantID = "1234"

func TestInsertParticipant(t *testing.T) {
	t.Parallel()
	var fistValueExpected error = nil
	secondValueExpected := participantID
	value, err := InsertParticipant(mountParticipant())

	if value != secondValueExpected || err != fistValueExpected {
		t.Errorf(errorMessage, fistValueExpected, err)
	}
}

func mountParticipant() (p Participant) {
	p.ID = participantID
	p.Name = "Test"
	p.Content = "Testing"

	return p
}

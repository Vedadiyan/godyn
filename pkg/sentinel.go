package godyn

import (
	"fmt"
)

type Error string

const (
	INVALID_NUMBER_OF_ARGUMENTS Error = "requires %d arguments but recieved %d"
	ARGUMENT_MISMATCH           Error = "%d argument should be of type %s but recieved %s"
)

func InvalidNumberOfArgumentsError(require int, recieved int) error {
	return fmt.Errorf(fmt.Sprintf(string(INVALID_NUMBER_OF_ARGUMENTS), require, recieved))
}

func ArgumentMismatchError(argNumber int, requireType string, recievedType string) error {
	return fmt.Errorf(fmt.Sprintf(string(ARGUMENT_MISMATCH), argNumber, requireType, recievedType))
}

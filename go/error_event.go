package discoveryextensionevents

import (
	"fmt"
)

// / <summary>
// / Initializes a new instance of the <see cref="ErrorEvent"/> class.
// / </summary>
// / <param name="errorCode">The error code to be associated with the exception.</param>
// / <param name="message">The message associated with the exception.</param>
// / <param name="possibleCauses">The possibleCauses associated with the exception.</param>
// / <param name="recommendedAction">The recommendedAction associated with the exception.</param>
// / <param name="errorMsgParams">Parameters required to construct the error message.
func getErrorEvent(
	errorCode AgentErrorCode,
	message string,
	possibleCauses string,
	recommendedAction string,
	errorMsgParams map[string]string) ErrorEvent {
	a := ErrorEvent{}
	a.ErrorCode = errorCode
	a.Message = message
	a.PossibleCauses = possibleCauses
	a.RecommendedAction = recommendedAction
	a.MessageParameters = errorMsgParams

	return a
}

// / <summary>
// / This class represents exceptions thrown by the discovery agent.
// / </summary>
type ErrorEvent struct {
	/// <summary>
	/// Gets the error code associated with the exception.
	/// </summary>
	ErrorCode AgentErrorCode

	/// <summary>
	/// The message associated with the error.
	/// </summary>
	Message string

	/// <summary>
	/// The possible causes for the error.
	/// </summary>
	PossibleCauses string

	/// <summary>
	/// The recommended action for resolving the error.
	/// </summary>
	RecommendedAction string
  
	/// <summary>
	/// Gets the message parameters needed to form the error message.
	/// </summary>
	MessageParameters map[string]string
}

// define Error() method on the struct
func (ae ErrorEvent) Error() string {
	return fmt.Sprint(ae.ErrorCode)
}

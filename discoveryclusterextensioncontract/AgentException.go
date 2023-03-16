package discoveryclusterextensioncontract

import (
	"fmt"
)

// / <summary>
// / Initializes a new instance of the <see cref="AgentException"/> class.
// / </summary>
// / <param name="errorCode">The error code to be associated with the exception.</param>
// / <param name="errorMsgParams">Parameters required to construct the error message.
func getAgentException(errorCode AgentErrorCode, errorMsgParams map[string]string) AgentException {
	a := AgentException{}
	a.ErrorCode = errorCode
	a.MessageParameters = errorMsgParams

	return a
}

// / <summary>
// / This class represents exceptions thrown by the discovery agent.
// / </summary>
type AgentException struct {
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
func (ae AgentException) Error() string {
	return fmt.Sprint(ae.ErrorCode)
}

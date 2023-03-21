package discoveryclusterextensioncontract

// / <summary>
// / Initializes a new instance of the <see cref="AgentException"/> class.
// / </summary>
// / <param name="errorCode">The error code to be associated with the exception.</param>
// / <param name="message">The message associated with the exception.</param>
// / <param name="possibleCauses">The possibleCauses associated with the exception.</param>
// / <param name="recommendedAction">The recommendedAction associated with the exception.</param>
// / <param name="errorMsgParams">Parameters required to construct the error message.
func getAgentException(
	errorCode AgentErrorCode,
	message string,
	possibleCauses string,
	recommendedAction string,
	errorMsgParams map[string]string) AgentException {
	a := AgentException{}
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

package discoveryextensionevents

// / <summary>
// / Initializes a new instance of the <see cref="TelemetryEvent"/> class.
// / </summary>
// / <param name="errorCode">The error code to be associated with the exception.</param>
// / <param name="errorMsgParams">Parameters required to construct the error message.
func getTelemetryEvent(telemetryEventId TelemetryEventId, dimensions map[string]string) TelemetryEvent {
	t := TelemetryEvent{}
	t.TelemetryEventId = telemetryEventId
	t.Dimensions = dimensions

	return t
}

// / <summary>
// / This class represents exceptions thrown by the discovery agent.
// / </summary>
type TelemetryEvent struct {
	/// <summary>
	/// Gets the error code associated with the exception.
	/// </summary>
	TelemetryEventId TelemetryEventId

	/// <summary>
	/// Gets the message parameters needed to form the error message.
	/// </summary>
	Dimensions map[string]string
}

package discoveryextensionevents

// / <summary>
// / Initializes a new instance of the <see cref="MetrciEvent"/> class.
// / </summary>
// / <param name="errorCode">The error code to be associated with the exception.</param>
// / <param name="errorMsgParams">Parameters required to construct the error message.
func getMetricEvent(metricName MetricEventId, dimensions map[string]string, annotations map[string]string) MetricEvent {
	t := MetricEvent{}
	t.MetricName = metricName
	t.Dimensions = dimensions
	t.Annotations = annotations

	return t
}

// / <summary>
// / This class represents metric event by discovery agent.
// / </summary>
type MetricEvent struct {
	/// <summary>
	/// Gets the MetricName that represent the metric.
	/// </summary>
	MetricName MetricEventId

	/// <summary>
	/// Gets the dimension dictionary with the parameter needed for the metric.
	/// </summary>
	Dimensions map[string]string

	/// <summary>
	/// Gets the annotation dictionary with the parameter needed for the metric.
	/// </summary>
	Annotations map[string]string
}

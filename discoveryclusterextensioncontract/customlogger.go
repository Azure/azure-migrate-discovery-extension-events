package discoveryclusterextensioncontract

import (
	"context"
	"strings"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type CustomLogger struct {
	logr logr.Logger
}

func GetAzureLogger(ctx context.Context, annotationsMap map[string]string) CustomLogger {
	c := CustomLogger{logr: log.FromContext(ctx)}
	c.addAzureAnnotations(annotationsMap)

	return c
}

func (c *CustomLogger) GetLogr() logr.Logger {
	return c.logr
}

// Adding annotation in custom logger.
func (c *CustomLogger) addAzureAnnotations(annotationsMap map[string]string) {
	for annotationkey, annotationvalue := range annotationsMap {
		// only adding azure annotations
		if strings.HasPrefix(annotationkey, "management.azure.com/") {
			scrubbedKey := strings.TrimPrefix(annotationkey, "management.azure.com/")
			// strcase.ToCamel() and serialize the system data.
			c.logr = c.logr.WithValues(scrubbedKey, annotationvalue)
		}
	}
}

// Logging Exception Event
func (c *CustomLogger) LogAgentException(ae discoveryclusterextensioncontract.AgentException) {

	// todo: add message parameter into message, posssiblecauses, recommended action.
	logger := c.logr.WithValues("MessageParameters", ae.MessageParameters)
	logger = logger.WithValues("Message", ae.Message)
	logger = logger.WithValues("PossibleCauses", ae.PossibleCauses)
	logger = logger.WithValues("RecommendedAction", ae.RecommendedAction)
	logger.Error(ae, "Agent Exception Logged.")
}

// Logging Telemetry Event
func (c *CustomLogger) LogTelemetryEvent(te discoveryclusterextensioncontract.TelemetryEvent) {
	logger := c.logr.WithValues("ServiceEventName", "TelemetryEvent")
	logger = logger.WithValues("TelemetryEventId", te.TelemetryEventId)
	logger = logger.WithValues("Dimensions", te.Dimensions)

	logger.Info("TelemetryEvent")
}

// Logging Metric Event
func (c *CustomLogger) LogMetricEvent(metricName string, metricValue int, metricDimensions map[string]string) {
	logger := c.logr.WithValues("ServiceEventName", "MetricEvent")
	logger = logger.WithValues("MetricValue", metricValue)

	for metricDimensionKey, metricDimensionValue := range metricDimensions {
		logger = logger.WithValues(metricDimensionKey, metricDimensionValue)
	}

	logger.Info(metricName)
}

package discoveryextensionevents

import (
	"context"
	"strings"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type CustomLogger struct {
	logr logr.Logger
}

func AzureLogger(ctx context.Context, annotationsMap map[string]string) CustomLogger {
	c := CustomLogger{logr: log.FromContext(ctx)}
	c.addAzureAnnotations(annotationsMap)

	return c
}

func (c *CustomLogger) Logr() logr.Logger {
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
func (c *CustomLogger) LogErrorEvent(ee ErrorEvent) {

	// todo: add message parameter into message, posssiblecauses, recommended action.
	logger := c.logr.WithValues("MessageParameters", ee.MessageParameters)
	logger = logger.WithValues("Message", ee.Message)
	logger = logger.WithValues("PossibleCauses", ee.PossibleCauses)
	logger = logger.WithValues("RecommendedAction", ee.RecommendedAction)
	logger.Error(ee, "Agent Exception Logged.")
}

// Logging Telemetry Event
func (c *CustomLogger) LogTelemetryEvent(te TelemetryEvent) {
	logger := c.logr.WithValues("ServiceEventName", "TelemetryEvent")
	logger = logger.WithValues("TelemetryEventId", te.TelemetryEventId)
	logger = logger.WithValues("Dimensions", te.Dimensions)

	logger.Info("TelemetryEvent")
}

// Logging Metric Event
func (c *CustomLogger) LogMetricEvent(me MetricEvent, metricValue int) {
	logger := c.logr.WithValues("ServiceEventName", "MetricEvent")
	logger = logger.WithValues("MetricValue", metricValue)

	for metricDimensionKey, metricDimensionValue := range me.Dimensions {
		logger = logger.WithValues(metricDimensionKey, metricDimensionValue)
	}

	for metricAnnotationKey, metricAnnotationValue := range me.Annotations {
		logger = logger.WithValues(metricAnnotationKey, metricAnnotationValue)
	}
	logger = logger.WithValues("MetricName", string(me.MetricName))

	logger.Info("MetricEvent")
}
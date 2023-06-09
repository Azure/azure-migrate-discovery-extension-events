﻿
// Code generated by controller-gen. DO NOT EDIT.
// <copyright file="MetricEvents.go" company="Microsoft Corporation">
// Copyright (c) Microsoft Corporation. All rights reserved.
// </copyright>
// <autogenerated />

package discoveryextensionevents

type MetricEventId string

    const (
        SASURINotValid MetricEventId = "SASURINotValid"
        RoleStart MetricEventId = "RoleStart"
        RoleStop MetricEventId = "RoleStop"
        IISApplicationStart MetricEventId = "IISApplicationStart"
        IISApplicationStop MetricEventId = "IISApplicationStop"
    )

//Annotation class
    type SASURINotValidType struct {
            AnnClientRequestId string
        }

//Method to get AnnotationType
    func SASURINotValidAnnotation(annClientRequestId string) SASURINotValidType {
        return SASURINotValidType{AnnClientRequestId: annClientRequestId}
    }

// Event methods
    func SASURINotValidMetricEvent(subscriptionId string, clientRequestId string, errorCode string, annotation SASURINotValidType) MetricEvent {
        dimensionsDataCollection := map[string]string{}
        dimensionsDataCollection["SubscriptionId"] = subscriptionId
        dimensionsDataCollection["ClientRequestId"] = clientRequestId
        dimensionsDataCollection["ErrorCode"] = errorCode
        
        annotationsDataCollection := map[string]string{}
        annotationsDataCollection["AnnClientRequestId"] = annotation.AnnClientRequestId
 
        return getMetricEvent(SASURINotValid, dimensionsDataCollection, annotationsDataCollection)
    }
    func RoleStartMetricEvent() MetricEvent {
        dimensionsDataCollection := map[string]string{}
        
        annotationsDataCollection := map[string]string{}
 
        return getMetricEvent(RoleStart, dimensionsDataCollection, annotationsDataCollection)
    }
    func RoleStopMetricEvent() MetricEvent {
        dimensionsDataCollection := map[string]string{}
        
        annotationsDataCollection := map[string]string{}
 
        return getMetricEvent(RoleStop, dimensionsDataCollection, annotationsDataCollection)
    }
    func IISApplicationStartMetricEvent() MetricEvent {
        dimensionsDataCollection := map[string]string{}
        
        annotationsDataCollection := map[string]string{}
 
        return getMetricEvent(IISApplicationStart, dimensionsDataCollection, annotationsDataCollection)
    }
    func IISApplicationStopMetricEvent(hostingEnvShutdownReason string) MetricEvent {
        dimensionsDataCollection := map[string]string{}
        dimensionsDataCollection["HostingEnvShutdownReason"] = hostingEnvShutdownReason
        
        annotationsDataCollection := map[string]string{}
 
        return getMetricEvent(IISApplicationStop, dimensionsDataCollection, annotationsDataCollection)
    }

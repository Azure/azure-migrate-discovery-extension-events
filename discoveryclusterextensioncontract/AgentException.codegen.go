﻿
// <copyright file="AgentException.codegen.go" company="Microsoft Corporation">
// Copyright (c) Microsoft Corporation. All rights reserved.
// </copyright>
// <autogenerated />

package discoveryclusterextensioncontract
func GetUnhandledException(errorCode string, errorMessage string, errortype string) AgentException {
    return getAgentException(
        UnhandledException,
        map[string]string{
                    "ErrorCode": errorCode ,
                    "ErrorMessage": errorMessage ,
                    "Errortype": errortype ,
                });
        }
func GetNullRefException(reason string) AgentException {
    return getAgentException(
        NullRefException,
        map[string]string{
                    "Reason": reason ,
                });
        }
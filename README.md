**To Consume the Go module**
1. Use go get to fetch the module with latest verison ```go get github.com/Azure/azure-migrate-discovery-extension-events```.
2. For module with specific commit/version use ```go get go get github.com/Azure/azure-migrate-discovery-extension-events@<commit-id>```.
3. In your controller reconciler.

You can use funcitons provided in https://github.com/Azure/azure-migrate-discovery-extension-events/blob/main/go/zapr_logger.go

a. **To Log Error Events**
````
  azureLogger := AzureLogger(ctx, cr.GetAnnotations())
  ne := discoveryextensionevents.NullReferenceErrorEvent("null ref while accessing x variable")
  discoveryextensionevents.azureLogger.LogErrorEvent(ne)
````
b. **To Log Metric Events**
````
  azureLogger := AzureLogger(ctx, cr.GetAnnotations())
  me := discoveryextensionevents.RoleStartMetricEvent()
  discoveryextensionevents.azureLogger.LogMetricEvent(me);
````

c. **To Log Telemetry Events**
````
  azureLogger := AzureLogger(ctx, cr.GetAnnotations())
  te := discoveryextensionevents.EntityDiscoveredTelemetryEvent()
  discoveryextensionevents.azureLogger.LogTelemetryEvent(te)
````


d. **To Log Error Messages**
````
azureLogger.logr.Error(err, "Failed to get Seed Discovery Entity, lets requeue for reconcilation.")
````

e. **To Log Info Messages**
````
 azureLogger.logr.V(2).Info("Seed Discovery Entity has already been reconciled, skip...",
			"spec.Generation", cr.Generation,
			"status.Generation", cr.Status.ObservedGeneration);
````


**To Add New Events**
1. Add new events in the Telemetry.xml, Error.xml or Metrics.xml in  AzureMigrate-ClusterExtensionCommon repo (https://msazure.visualstudio.com/One/_git/AzureMigrate-ClusterExtensionCommon?path=/src/DiscoveryClusterExtension/EventXml).

2. Kick off official build after PR merge.
3. Get the artifiact of name Gomodule from drop_build_main\codegen and download it in your local machine.
4. Raise the PR in this Git Repo.
5. Consume the Go module in respective discovery cluster extension.

**North Star Goal**
- The step 3 adn 4 would be automated in future.
- The go module should have new events by running the build pipeline on this repo
https://msazure.visualstudio.com/One/_git/AzureMigrate-ClusterExtensionCommon

**To add the tags**

- Add the Tag
````
git add .
git commit -m "Updating read me"

git tag v1.5.0 HEAD -m "Version 1.5.0 released"
git push origin --tags
````

- In your source code
````
 go get github.com/Azure/azure-migrate-discovery-extension-events@v1.5.0
````
Reference Link
https://phoenixnap.com/kb/git-tag-commit

**To Consume the Go module**
1. Use go get to fetch the module with latest verison ```go get github.com/Azure/AzureMigrate-discoveryclusterextension-event```.
2. For module with specific commit/version use ```go get github.com/Azure/AzureMigrate-discoveryclusterextension-event@<commit-id>```.

**To add new Events**
1. Add new events in the Telemetry.xml, Error.xml or Metrics.xml in  AzureMigrate-ClusterExtensionCommon repo (https://msazure.visualstudio.com/One/_git/AzureMigrate-ClusterExtensionCommon?path=/src/DiscoveryClusterExtension/EventXml).

2. Kick off official build after PR merge.
3. Get the artifiact of name Gomodule from drop_build_main\codegen and download it in your local machine.
4. Raise the PR in this Git Repo.
5. Consume the Go module in respective discovery cluster extension.

**North Star Goal**
- The step 3 adn 4 would be automated in future.
- The go module should have new events by running the build pipeline on this repo
https://msazure.visualstudio.com/One/_git/AzureMigrate-ClusterExtensionCommon

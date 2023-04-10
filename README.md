**To Consume the Go module**
1. Use go get to fetch the module with latest verison ```go get github.com/Azure/azure-migrate-discovery-extension-events```.
2. For module with specific commit/version use ```go get go get github.com/Azure/azure-migrate-discovery-extension-events@<commit-id>```.

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

** To make changes and to add tags**
````
git add .
git commit -m "Updating read me"

git tag v1.5.0 HEAD -m "Version 1.5.0 released"
git push origin --tags
````

https://phoenixnap.com/kb/git-tag-commit

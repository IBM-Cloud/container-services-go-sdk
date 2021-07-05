# IBM Cloud Container Services Go SDK Version 1.0 

Go client library for [IBM Cloud Kubernetes service](https://cloud.ibm.com/apidocs/kubernetes) and Satellite services.

**NOTE:**

The [IBM Cloud Provider for Terraform](https://github.com/IBM-Cloud/terraform-provider-ibm) uses this client library, alongside the [bluemix-go](https://github.com/IBM-Cloud/bluemix-go) SDK.  No additional support is provided for this SDK.

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  * [Go modules](#go-modules)
  * [`go get` command](#go-get-command)
- [Using the SDK](#using-the-sdk)

<!-- tocstop -->

## Overview

The IBM Cloud Container Services Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[Kubernetes](https://cloud.ibm.com/apidocs/kubernetes) | kubernetesserviceapiv1
[Satellite Link](https://cloud.ibm.com/docs/satellite?topic=satellite-link-location-cloud) | satellitelinkv1

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one
[here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.12 or above.

## Installation
The current version of this SDK: 1.0

### Go modules  
If your application uses Go modules for dependency management (recommended), just add an import for each service 
that you will use in your application.  
Here is an example:

```go
import (
	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
)
```
Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's
`go.mod` file.  

In the example above, the `kubernetesserviceapiv1` part of the import path is the package name
associated with the kubernetes service.
See the service table above to find the approprate package name for the services used by your application.

### `go get` command  
Alternatively, you can use the `go get` command to download and install the appropriate packages needed by your application:
```
go get -u github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1
```
Be sure to use the appropriate package name from the service table above for the services used by your application.


## License

The IBM Cloud Container Services Go SDK is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).

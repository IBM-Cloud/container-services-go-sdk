// +build integration

/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package satellitelinkv1_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

/**
 * This file contains an integration test for the satellitelinkv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`SatelliteLinkV1 Integration Tests`, func() {

	const externalConfigFile = "../satellite_link_v1.env"

	var (
		err                  error
		satelliteLinkService *satellitelinkv1.SatelliteLinkV1
		serviceURL           string
		config               map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(satellitelinkv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			satelliteLinkServiceOptions := &satellitelinkv1.SatelliteLinkV1Options{}

			satelliteLinkService, err = satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(satelliteLinkServiceOptions)

			Expect(err).To(BeNil())
			Expect(satelliteLinkService).ToNot(BeNil())
			Expect(satelliteLinkService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`CreateLink - create link [Administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateLink(createLinkOptions *CreateLinkOptions)`, func() {

			createLinkOptions := &satellitelinkv1.CreateLinkOptions{
				Crn:        core.StringPtr("crn:v1:staging:public:satellite:us-south:a/1ae4eb57181a46ceade4846519678888::location:brbats7009sqna3dtest"),
				LocationID: core.StringPtr("brbats7009sqna3dtest"),
			}

			location, response, err := satelliteLinkService.CreateLink(createLinkOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(location).ToNot(BeNil())

		})
	})

	Describe(`GetLink - read link [Administrator, Editor, Operator, Viewer, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLink(getLinkOptions *GetLinkOptions)`, func() {

			getLinkOptions := &satellitelinkv1.GetLinkOptions{
				LocationID: core.StringPtr("testString"),
			}

			location, response, err := satelliteLinkService.GetLink(getLinkOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(location).ToNot(BeNil())

		})
	})

	Describe(`UpdateLink - update link [Administrator, Editor, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateLink(updateLinkOptions *UpdateLinkOptions)`, func() {

			updateLinkOptions := &satellitelinkv1.UpdateLinkOptions{
				LocationID: core.StringPtr("testString"),
				WsEndpoint: core.StringPtr("{satellite-link-tunnel-server}"),
			}

			location, response, err := satelliteLinkService.UpdateLink(updateLinkOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(location).ToNot(BeNil())

		})
	})

	Describe(`ListEndpoints - list endpoints [Administrator, Editor, Viewer, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListEndpoints(listEndpointsOptions *ListEndpointsOptions)`, func() {

			listEndpointsOptions := &satellitelinkv1.ListEndpointsOptions{
				LocationID: core.StringPtr("testString"),
				Type:       core.StringPtr("enabled"),
			}

			endpoints, response, err := satelliteLinkService.ListEndpoints(listEndpointsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpoints).ToNot(BeNil())

		})
	})

	Describe(`CreateEndpoints - create endpoint [Administrator, Editor, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateEndpoints(createEndpointsOptions *CreateEndpointsOptions)`, func() {

			additionalNewEndpointRequestCertsClientCertModel := &satellitelinkv1.AdditionalNewEndpointRequestCertsClientCert{
				Filename:     core.StringPtr("clientEndpointCert.pem"),
				FileContents: core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n"),
			}

			additionalNewEndpointRequestCertsClientModel := &satellitelinkv1.AdditionalNewEndpointRequestCertsClient{
				Cert: additionalNewEndpointRequestCertsClientCertModel,
			}

			additionalNewEndpointRequestCertsServerCertModel := &satellitelinkv1.AdditionalNewEndpointRequestCertsServerCert{
				Filename:     core.StringPtr("serverEndpointCert.pem"),
				FileContents: core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n"),
			}

			additionalNewEndpointRequestCertsServerModel := &satellitelinkv1.AdditionalNewEndpointRequestCertsServer{
				Cert: additionalNewEndpointRequestCertsServerCertModel,
			}

			additionalNewEndpointRequestCertsConnectorCertModel := &satellitelinkv1.AdditionalNewEndpointRequestCertsConnectorCert{
				Filename:     core.StringPtr("ConnectorCert.pem"),
				FileContents: core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n"),
			}

			additionalNewEndpointRequestCertsConnectorKeyModel := &satellitelinkv1.AdditionalNewEndpointRequestCertsConnectorKey{
				Filename:     core.StringPtr("ConnectorPrivateKey.pem"),
				FileContents: core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n"),
			}

			additionalNewEndpointRequestCertsConnectorModel := &satellitelinkv1.AdditionalNewEndpointRequestCertsConnector{
				Cert: additionalNewEndpointRequestCertsConnectorCertModel,
				Key:  additionalNewEndpointRequestCertsConnectorKeyModel,
			}

			additionalNewEndpointRequestCertsModel := &satellitelinkv1.AdditionalNewEndpointRequestCerts{
				Client:    additionalNewEndpointRequestCertsClientModel,
				Server:    additionalNewEndpointRequestCertsServerModel,
				Connector: additionalNewEndpointRequestCertsConnectorModel,
			}

			createEndpointsOptions := &satellitelinkv1.CreateEndpointsOptions{
				LocationID:       core.StringPtr("testString"),
				ConnType:         core.StringPtr("cloud"),
				DisplayName:      core.StringPtr("My endpoint"),
				ServerHost:       core.StringPtr("example.com"),
				ServerPort:       core.Int64Ptr(int64(443)),
				Sni:              core.StringPtr("example.com"),
				ClientProtocol:   core.StringPtr("https"),
				ClientMutualAuth: core.BoolPtr(true),
				ServerProtocol:   core.StringPtr("tls"),
				ServerMutualAuth: core.BoolPtr(true),
				RejectUnauth:     core.BoolPtr(true),
				Timeout:          core.Int64Ptr(int64(60)),
				CreatedBy:        core.StringPtr("My service"),
				Certs:            additionalNewEndpointRequestCertsModel,
			}

			endpoint, response, err := satelliteLinkService.CreateEndpoints(createEndpointsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpoint).ToNot(BeNil())

		})
	})

	Describe(`ImportEndpoints - import an endpoint [Administrator, Editor, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ImportEndpoints(importEndpointsOptions *ImportEndpointsOptions)`, func() {

			importEndpointsOptions := &satellitelinkv1.ImportEndpointsOptions{
				LocationID:       core.StringPtr("testString"),
				State:            CreateMockReader("This is a mock file."),
				StateContentType: core.StringPtr("testString"),
			}

			endpoint, response, err := satelliteLinkService.ImportEndpoints(importEndpointsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpoint).ToNot(BeNil())

		})
	})

	Describe(`ExportEndpoints - export an endpoint [Administrator, Editor, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ExportEndpoints(exportEndpointsOptions *ExportEndpointsOptions)`, func() {

			exportEndpointsOptions := &satellitelinkv1.ExportEndpointsOptions{
				LocationID: core.StringPtr("testString"),
				EndpointID: core.StringPtr("testString"),
			}

			exportEndpointsResponse, response, err := satelliteLinkService.ExportEndpoints(exportEndpointsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(exportEndpointsResponse).ToNot(BeNil())

		})
	})

	Describe(`GetEndpoints - read endpoint [Administrator, Editor, Viewer, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEndpoints(getEndpointsOptions *GetEndpointsOptions)`, func() {

			getEndpointsOptions := &satellitelinkv1.GetEndpointsOptions{
				LocationID: core.StringPtr("testString"),
				EndpointID: core.StringPtr("testString"),
			}

			endpoint, response, err := satelliteLinkService.GetEndpoints(getEndpointsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpoint).ToNot(BeNil())

		})
	})

	Describe(`UpdateEndpoints - update endpoint [Administrator, Editor, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateEndpoints(updateEndpointsOptions *UpdateEndpointsOptions)`, func() {

			updatedEndpointRequestCertsClientCertModel := &satellitelinkv1.UpdatedEndpointRequestCertsClientCert{
				Filename:     core.StringPtr("clientEndpointCert.pem"),
				FileContents: core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n"),
			}

			updatedEndpointRequestCertsClientModel := &satellitelinkv1.UpdatedEndpointRequestCertsClient{
				Cert: updatedEndpointRequestCertsClientCertModel,
			}

			updatedEndpointRequestCertsServerCertModel := &satellitelinkv1.UpdatedEndpointRequestCertsServerCert{
				Filename:     core.StringPtr("serverEndpointCert.pem"),
				FileContents: core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n"),
			}

			updatedEndpointRequestCertsServerModel := &satellitelinkv1.UpdatedEndpointRequestCertsServer{
				Cert: updatedEndpointRequestCertsServerCertModel,
			}

			updatedEndpointRequestCertsConnectorCertModel := &satellitelinkv1.UpdatedEndpointRequestCertsConnectorCert{
				Filename:     core.StringPtr("ConnectorCert.pem"),
				FileContents: core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n"),
			}

			updatedEndpointRequestCertsConnectorKeyModel := &satellitelinkv1.UpdatedEndpointRequestCertsConnectorKey{
				Filename:     core.StringPtr("ConnectorPrivateKey.pem"),
				FileContents: core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n"),
			}

			updatedEndpointRequestCertsConnectorModel := &satellitelinkv1.UpdatedEndpointRequestCertsConnector{
				Cert: updatedEndpointRequestCertsConnectorCertModel,
				Key:  updatedEndpointRequestCertsConnectorKeyModel,
			}

			updatedEndpointRequestCertsModel := &satellitelinkv1.UpdatedEndpointRequestCerts{
				Client:    updatedEndpointRequestCertsClientModel,
				Server:    updatedEndpointRequestCertsServerModel,
				Connector: updatedEndpointRequestCertsConnectorModel,
			}

			updateEndpointsOptions := &satellitelinkv1.UpdateEndpointsOptions{
				LocationID:       core.StringPtr("testString"),
				EndpointID:       core.StringPtr("testString"),
				DisplayName:      core.StringPtr("My endpoint"),
				ServerHost:       core.StringPtr("example.com"),
				ServerPort:       core.Int64Ptr(int64(443)),
				Sni:              core.StringPtr("example.com"),
				ClientProtocol:   core.StringPtr("https"),
				ClientMutualAuth: core.BoolPtr(true),
				ServerProtocol:   core.StringPtr("tls"),
				ServerMutualAuth: core.BoolPtr(true),
				RejectUnauth:     core.BoolPtr(true),
				Timeout:          core.Int64Ptr(int64(60)),
				CreatedBy:        core.StringPtr("My service"),
				Certs:            updatedEndpointRequestCertsModel,
				Enabled:          core.BoolPtr(true),
			}

			endpoint, response, err := satelliteLinkService.UpdateEndpoints(updateEndpointsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpoint).ToNot(BeNil())

		})
	})

	Describe(`GetEndpointCerts - read endpoint certs [Administrator, Editor, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEndpointCerts(getEndpointCertsOptions *GetEndpointCertsOptions)`, func() {

			getEndpointCertsOptions := &satellitelinkv1.GetEndpointCertsOptions{
				LocationID: core.StringPtr("testString"),
				EndpointID: core.StringPtr("testString"),
				NoZip:      core.BoolPtr(true),
			}

			downloadedCerts, response, err := satelliteLinkService.GetEndpointCerts(getEndpointCertsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(downloadedCerts).ToNot(BeNil())

		})
	})

	Describe(`UploadEndpointCerts - upload endpoint certs [Administrator, Editor, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UploadEndpointCerts(uploadEndpointCertsOptions *UploadEndpointCertsOptions)`, func() {

			uploadEndpointCertsOptions := &satellitelinkv1.UploadEndpointCertsOptions{
				LocationID:               core.StringPtr("testString"),
				EndpointID:               core.StringPtr("testString"),
				ClientCert:               CreateMockReader("This is a mock file."),
				ClientCertContentType:    core.StringPtr("testString"),
				ServerCert:               CreateMockReader("This is a mock file."),
				ServerCertContentType:    core.StringPtr("testString"),
				ConnectorCert:            CreateMockReader("This is a mock file."),
				ConnectorCertContentType: core.StringPtr("testString"),
				ConnectorKey:             CreateMockReader("This is a mock file."),
				ConnectorKeyContentType:  core.StringPtr("testString"),
			}

			endpoint, response, err := satelliteLinkService.UploadEndpointCerts(uploadEndpointCertsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpoint).ToNot(BeNil())

		})
	})

	Describe(`ListEndpointSources - list endpoint sources [Administrator, Editor, Viewer, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListEndpointSources(listEndpointSourcesOptions *ListEndpointSourcesOptions)`, func() {

			listEndpointSourcesOptions := &satellitelinkv1.ListEndpointSourcesOptions{
				LocationID: core.StringPtr("testString"),
				EndpointID: core.StringPtr("testString"),
			}

			sourceStatus, response, err := satelliteLinkService.ListEndpointSources(listEndpointSourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sourceStatus).ToNot(BeNil())

		})
	})

	Describe(`UpdateEndpointSources - update endpoint sources [Administrator, Editor, Operator, satellite-link-source-access-controller]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateEndpointSources(updateEndpointSourcesOptions *UpdateEndpointSourcesOptions)`, func() {

			sourceStatusRequestObjectModel := &satellitelinkv1.SourceStatusRequestObject{
				SourceID: core.StringPtr("us-south--K9kQEVFmqNpP-Source-Q87fe"),
				Enabled:  core.BoolPtr(true),
			}

			updateEndpointSourcesOptions := &satellitelinkv1.UpdateEndpointSourcesOptions{
				LocationID: core.StringPtr("testString"),
				EndpointID: core.StringPtr("testString"),
				Sources:    []satellitelinkv1.SourceStatusRequestObject{*sourceStatusRequestObjectModel},
			}

			sourceStatus, response, err := satelliteLinkService.UpdateEndpointSources(updateEndpointSourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sourceStatus).ToNot(BeNil())

		})
	})

	Describe(`ListSources - list sources [Administrator, Editor, Viewer, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSources(listSourcesOptions *ListSourcesOptions)`, func() {

			listSourcesOptions := &satellitelinkv1.ListSourcesOptions{
				LocationID: core.StringPtr("testString"),
				Type:       core.StringPtr("user"),
			}

			sources, response, err := satelliteLinkService.ListSources(listSourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(sources).ToNot(BeNil())

		})
	})

	Describe(`CreateSources - create source [Administrator, Editor, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSources(createSourcesOptions *CreateSourcesOptions)`, func() {

			createSourcesOptions := &satellitelinkv1.CreateSourcesOptions{
				LocationID: core.StringPtr("testString"),
				Type:       core.StringPtr("user"),
				SourceName: core.StringPtr("DB2"),
				Addresses:  []string{"192.168.20.1/24"},
			}

			source, response, err := satelliteLinkService.CreateSources(createSourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(source).ToNot(BeNil())

		})
	})

	Describe(`UpdateSources - update source [Administrator, Editor, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSources(updateSourcesOptions *UpdateSourcesOptions)`, func() {

			updateSourcesOptions := &satellitelinkv1.UpdateSourcesOptions{
				LocationID: core.StringPtr("testString"),
				SourceID:   core.StringPtr("testString"),
				SourceName: core.StringPtr("DB2"),
				Addresses:  []string{"192.168.20.1/24"},
			}

			source, response, err := satelliteLinkService.UpdateSources(updateSourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(source).ToNot(BeNil())

		})
	})

	Describe(`ListSourceEndpoints - list source status for all endpoints [Administrator, Editor, Viewer, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSourceEndpoints(listSourceEndpointsOptions *ListSourceEndpointsOptions)`, func() {

			listSourceEndpointsOptions := &satellitelinkv1.ListSourceEndpointsOptions{
				LocationID: core.StringPtr("testString"),
				SourceID:   core.StringPtr("testString"),
			}

			endpointSourceStatus, response, err := satelliteLinkService.ListSourceEndpoints(listSourceEndpointsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointSourceStatus).ToNot(BeNil())

		})
	})

	Describe(`UpdateSourceEndpoints - update source status for listed endpoints [Administrator, Editor, Operator, satellite-link-source-access-controller]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSourceEndpoints(updateSourceEndpointsOptions *UpdateSourceEndpointsOptions)`, func() {

			endpointSourceStatusEndpointsItemModel := &satellitelinkv1.EndpointSourceStatusEndpointsItem{
				EndpointID: core.StringPtr("us-south--K9kQEVFmqNpP_Q87fe"),
				Enabled:    core.BoolPtr(true),
			}

			updateSourceEndpointsOptions := &satellitelinkv1.UpdateSourceEndpointsOptions{
				LocationID: core.StringPtr("testString"),
				SourceID:   core.StringPtr("testString"),
				Endpoints:  []satellitelinkv1.EndpointSourceStatusEndpointsItem{*endpointSourceStatusEndpointsItemModel},
			}

			executionResult, response, err := satelliteLinkService.UpdateSourceEndpoints(updateSourceEndpointsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(executionResult).ToNot(BeNil())

		})
	})

	Describe(`DeleteSources - delete source [Administrator, Editor, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSources(deleteSourcesOptions *DeleteSourcesOptions)`, func() {

			deleteSourcesOptions := &satellitelinkv1.DeleteSourcesOptions{
				LocationID: core.StringPtr("testString"),
				SourceID:   core.StringPtr("testString"),
			}

			executionResult, response, err := satelliteLinkService.DeleteSources(deleteSourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(executionResult).ToNot(BeNil())

		})
	})

	Describe(`DeleteLink - delete link [Administrator, Operator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLink(deleteLinkOptions *DeleteLinkOptions)`, func() {

			deleteLinkOptions := &satellitelinkv1.DeleteLinkOptions{
				LocationID: core.StringPtr("testString"),
			}

			executionResult, response, err := satelliteLinkService.DeleteLink(deleteLinkOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(executionResult).ToNot(BeNil())

		})
	})

	Describe(`DeleteEndpoints - delete endpoint [Administrator, Editor, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteEndpoints(deleteEndpointsOptions *DeleteEndpointsOptions)`, func() {

			deleteEndpointsOptions := &satellitelinkv1.DeleteEndpointsOptions{
				LocationID: core.StringPtr("testString"),
				EndpointID: core.StringPtr("testString"),
			}

			executionResult, response, err := satelliteLinkService.DeleteEndpoints(deleteEndpointsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(executionResult).ToNot(BeNil())

		})
	})

	Describe(`DeleteEndpointCerts - delete endpoint certs [Administrator, Editor, Operator, satellite-link-administrator]`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteEndpointCerts(deleteEndpointCertsOptions *DeleteEndpointCertsOptions)`, func() {

			deleteEndpointCertsOptions := &satellitelinkv1.DeleteEndpointCertsOptions{
				LocationID: core.StringPtr("testString"),
				EndpointID: core.StringPtr("testString"),
			}

			endpoint, response, err := satelliteLinkService.DeleteEndpointCerts(deleteEndpointCertsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpoint).ToNot(BeNil())

		})
	})
})

//
// Utility functions are declared in the unit test file
//

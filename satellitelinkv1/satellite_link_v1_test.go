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
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM-Cloud/container-services-go-sdk/satellitelinkv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`SatelliteLinkV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(satelliteLinkService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(satelliteLinkService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
				URL: "https://satellitelinkv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(satelliteLinkService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SATELLITE_LINK_URL":       "https://satellitelinkv1/api",
				"SATELLITE_LINK_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{})
				Expect(satelliteLinkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := satelliteLinkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != satelliteLinkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(satelliteLinkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(satelliteLinkService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{
					URL: "https://testService/api",
				})
				Expect(satelliteLinkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := satelliteLinkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != satelliteLinkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(satelliteLinkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(satelliteLinkService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{})
				err := satelliteLinkService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := satelliteLinkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != satelliteLinkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(satelliteLinkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(satelliteLinkService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SATELLITE_LINK_URL":       "https://satellitelinkv1/api",
				"SATELLITE_LINK_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(satelliteLinkService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SATELLITE_LINK_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(satelliteLinkService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = satellitelinkv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateLink(createLinkOptions *CreateLinkOptions) - Operation response error`, func() {
		createLinkPath := "/v1/locations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLinkPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLink with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the CreateLinkOptions model
				createLinkOptionsModel := new(satellitelinkv1.CreateLinkOptions)
				createLinkOptionsModel.Crn = core.StringPtr("crn:v1:staging:public:satellite:us-south:a/1ae4eb57181a46ceade4846519678888::location:brbats7009sqna3dtest")
				createLinkOptionsModel.LocationID = core.StringPtr("brbats7009sqna3dtest")
				createLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.CreateLink(createLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.CreateLink(createLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateLink(createLinkOptions *CreateLinkOptions)`, func() {
		createLinkPath := "/v1/locations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLinkPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ws_endpoint": "{satellite-link-tunnel-server}", "location_id": "brbats7009sqna3dtest", "crn": "{crn}", "desc": "My Location", "satellite_link_host": "{satellite-link-tunnel-server}", "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"tunnels": 3, "healthStatus": "Up", "avg_latency": 4, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "latency": 4, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke CreateLink successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the CreateLinkOptions model
				createLinkOptionsModel := new(satellitelinkv1.CreateLinkOptions)
				createLinkOptionsModel.Crn = core.StringPtr("crn:v1:staging:public:satellite:us-south:a/1ae4eb57181a46ceade4846519678888::location:brbats7009sqna3dtest")
				createLinkOptionsModel.LocationID = core.StringPtr("brbats7009sqna3dtest")
				createLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.CreateLinkWithContext(ctx, createLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.CreateLink(createLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.CreateLinkWithContext(ctx, createLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLinkPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ws_endpoint": "{satellite-link-tunnel-server}", "location_id": "brbats7009sqna3dtest", "crn": "{crn}", "desc": "My Location", "satellite_link_host": "{satellite-link-tunnel-server}", "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"tunnels": 3, "healthStatus": "Up", "avg_latency": 4, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "latency": 4, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke CreateLink successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.CreateLink(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateLinkOptions model
				createLinkOptionsModel := new(satellitelinkv1.CreateLinkOptions)
				createLinkOptionsModel.Crn = core.StringPtr("crn:v1:staging:public:satellite:us-south:a/1ae4eb57181a46ceade4846519678888::location:brbats7009sqna3dtest")
				createLinkOptionsModel.LocationID = core.StringPtr("brbats7009sqna3dtest")
				createLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.CreateLink(createLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateLink with error: Operation request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the CreateLinkOptions model
				createLinkOptionsModel := new(satellitelinkv1.CreateLinkOptions)
				createLinkOptionsModel.Crn = core.StringPtr("crn:v1:staging:public:satellite:us-south:a/1ae4eb57181a46ceade4846519678888::location:brbats7009sqna3dtest")
				createLinkOptionsModel.LocationID = core.StringPtr("brbats7009sqna3dtest")
				createLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.CreateLink(createLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLink(getLinkOptions *GetLinkOptions) - Operation response error`, func() {
		getLinkPath := "/v1/locations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLinkPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLink with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the GetLinkOptions model
				getLinkOptionsModel := new(satellitelinkv1.GetLinkOptions)
				getLinkOptionsModel.LocationID = core.StringPtr("testString")
				getLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.GetLink(getLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.GetLink(getLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLink(getLinkOptions *GetLinkOptions)`, func() {
		getLinkPath := "/v1/locations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLinkPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ws_endpoint": "{satellite-link-tunnel-server}", "location_id": "brbats7009sqna3dtest", "crn": "{crn}", "desc": "My Location", "satellite_link_host": "{satellite-link-tunnel-server}", "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"tunnels": 3, "healthStatus": "Up", "avg_latency": 4, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "latency": 4, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke GetLink successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetLinkOptions model
				getLinkOptionsModel := new(satellitelinkv1.GetLinkOptions)
				getLinkOptionsModel.LocationID = core.StringPtr("testString")
				getLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.GetLinkWithContext(ctx, getLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.GetLink(getLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.GetLinkWithContext(ctx, getLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLinkPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ws_endpoint": "{satellite-link-tunnel-server}", "location_id": "brbats7009sqna3dtest", "crn": "{crn}", "desc": "My Location", "satellite_link_host": "{satellite-link-tunnel-server}", "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"tunnels": 3, "healthStatus": "Up", "avg_latency": 4, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "latency": 4, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke GetLink successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.GetLink(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLinkOptions model
				getLinkOptionsModel := new(satellitelinkv1.GetLinkOptions)
				getLinkOptionsModel.LocationID = core.StringPtr("testString")
				getLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.GetLink(getLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLink with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the GetLinkOptions model
				getLinkOptionsModel := new(satellitelinkv1.GetLinkOptions)
				getLinkOptionsModel.LocationID = core.StringPtr("testString")
				getLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.GetLink(getLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLinkOptions model with no property values
				getLinkOptionsModelNew := new(satellitelinkv1.GetLinkOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.GetLink(getLinkOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateLink(updateLinkOptions *UpdateLinkOptions) - Operation response error`, func() {
		updateLinkPath := "/v1/locations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLinkPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateLink with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the UpdateLinkOptions model
				updateLinkOptionsModel := new(satellitelinkv1.UpdateLinkOptions)
				updateLinkOptionsModel.LocationID = core.StringPtr("testString")
				updateLinkOptionsModel.WsEndpoint = core.StringPtr("{satellite-link-tunnel-server}")
				updateLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.UpdateLink(updateLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.UpdateLink(updateLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateLink(updateLinkOptions *UpdateLinkOptions)`, func() {
		updateLinkPath := "/v1/locations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLinkPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ws_endpoint": "{satellite-link-tunnel-server}", "location_id": "brbats7009sqna3dtest", "crn": "{crn}", "desc": "My Location", "satellite_link_host": "{satellite-link-tunnel-server}", "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"tunnels": 3, "healthStatus": "Up", "avg_latency": 4, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "latency": 4, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke UpdateLink successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the UpdateLinkOptions model
				updateLinkOptionsModel := new(satellitelinkv1.UpdateLinkOptions)
				updateLinkOptionsModel.LocationID = core.StringPtr("testString")
				updateLinkOptionsModel.WsEndpoint = core.StringPtr("{satellite-link-tunnel-server}")
				updateLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.UpdateLinkWithContext(ctx, updateLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.UpdateLink(updateLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.UpdateLinkWithContext(ctx, updateLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLinkPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"ws_endpoint": "{satellite-link-tunnel-server}", "location_id": "brbats7009sqna3dtest", "crn": "{crn}", "desc": "My Location", "satellite_link_host": "{satellite-link-tunnel-server}", "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"tunnels": 3, "healthStatus": "Up", "avg_latency": 4, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "latency": 4, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke UpdateLink successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.UpdateLink(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateLinkOptions model
				updateLinkOptionsModel := new(satellitelinkv1.UpdateLinkOptions)
				updateLinkOptionsModel.LocationID = core.StringPtr("testString")
				updateLinkOptionsModel.WsEndpoint = core.StringPtr("{satellite-link-tunnel-server}")
				updateLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.UpdateLink(updateLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateLink with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the UpdateLinkOptions model
				updateLinkOptionsModel := new(satellitelinkv1.UpdateLinkOptions)
				updateLinkOptionsModel.LocationID = core.StringPtr("testString")
				updateLinkOptionsModel.WsEndpoint = core.StringPtr("{satellite-link-tunnel-server}")
				updateLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.UpdateLink(updateLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateLinkOptions model with no property values
				updateLinkOptionsModelNew := new(satellitelinkv1.UpdateLinkOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.UpdateLink(updateLinkOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteLink(deleteLinkOptions *DeleteLinkOptions) - Operation response error`, func() {
		deleteLinkPath := "/v1/locations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLinkPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteLink with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteLinkOptions model
				deleteLinkOptionsModel := new(satellitelinkv1.DeleteLinkOptions)
				deleteLinkOptionsModel.LocationID = core.StringPtr("testString")
				deleteLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.DeleteLink(deleteLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.DeleteLink(deleteLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteLink(deleteLinkOptions *DeleteLinkOptions)`, func() {
		deleteLinkPath := "/v1/locations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLinkPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "Success"}`)
				}))
			})
			It(`Invoke DeleteLink successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the DeleteLinkOptions model
				deleteLinkOptionsModel := new(satellitelinkv1.DeleteLinkOptions)
				deleteLinkOptionsModel.LocationID = core.StringPtr("testString")
				deleteLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.DeleteLinkWithContext(ctx, deleteLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.DeleteLink(deleteLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.DeleteLinkWithContext(ctx, deleteLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLinkPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "Success"}`)
				}))
			})
			It(`Invoke DeleteLink successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.DeleteLink(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteLinkOptions model
				deleteLinkOptionsModel := new(satellitelinkv1.DeleteLinkOptions)
				deleteLinkOptionsModel.LocationID = core.StringPtr("testString")
				deleteLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.DeleteLink(deleteLinkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteLink with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteLinkOptions model
				deleteLinkOptionsModel := new(satellitelinkv1.DeleteLinkOptions)
				deleteLinkOptionsModel.LocationID = core.StringPtr("testString")
				deleteLinkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.DeleteLink(deleteLinkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteLinkOptions model with no property values
				deleteLinkOptionsModelNew := new(satellitelinkv1.DeleteLinkOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.DeleteLink(deleteLinkOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(satelliteLinkService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(satelliteLinkService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
				URL: "https://satellitelinkv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(satelliteLinkService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SATELLITE_LINK_URL":       "https://satellitelinkv1/api",
				"SATELLITE_LINK_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{})
				Expect(satelliteLinkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := satelliteLinkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != satelliteLinkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(satelliteLinkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(satelliteLinkService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{
					URL: "https://testService/api",
				})
				Expect(satelliteLinkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := satelliteLinkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != satelliteLinkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(satelliteLinkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(satelliteLinkService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{})
				err := satelliteLinkService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := satelliteLinkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != satelliteLinkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(satelliteLinkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(satelliteLinkService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SATELLITE_LINK_URL":       "https://satellitelinkv1/api",
				"SATELLITE_LINK_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(satelliteLinkService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SATELLITE_LINK_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(satelliteLinkService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = satellitelinkv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListEndpoints(listEndpointsOptions *ListEndpointsOptions) - Operation response error`, func() {
		listEndpointsPath := "/v1/locations/testString/endpoints"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEndpointsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"enabled"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEndpoints with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the ListEndpointsOptions model
				listEndpointsOptionsModel := new(satellitelinkv1.ListEndpointsOptions)
				listEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				listEndpointsOptionsModel.Type = core.StringPtr("enabled")
				listEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.ListEndpoints(listEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.ListEndpoints(listEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListEndpoints(listEndpointsOptions *ListEndpointsOptions)`, func() {
		listEndpointsPath := "/v1/locations/testString/endpoints"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEndpointsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"enabled"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"endpoints": [{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}]}`)
				}))
			})
			It(`Invoke ListEndpoints successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListEndpointsOptions model
				listEndpointsOptionsModel := new(satellitelinkv1.ListEndpointsOptions)
				listEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				listEndpointsOptionsModel.Type = core.StringPtr("enabled")
				listEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.ListEndpointsWithContext(ctx, listEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.ListEndpoints(listEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.ListEndpointsWithContext(ctx, listEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEndpointsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"enabled"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"endpoints": [{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}]}`)
				}))
			})
			It(`Invoke ListEndpoints successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.ListEndpoints(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEndpointsOptions model
				listEndpointsOptionsModel := new(satellitelinkv1.ListEndpointsOptions)
				listEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				listEndpointsOptionsModel.Type = core.StringPtr("enabled")
				listEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.ListEndpoints(listEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListEndpoints with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the ListEndpointsOptions model
				listEndpointsOptionsModel := new(satellitelinkv1.ListEndpointsOptions)
				listEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				listEndpointsOptionsModel.Type = core.StringPtr("enabled")
				listEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.ListEndpoints(listEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListEndpointsOptions model with no property values
				listEndpointsOptionsModelNew := new(satellitelinkv1.ListEndpointsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.ListEndpoints(listEndpointsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateEndpoints(createEndpointsOptions *CreateEndpointsOptions) - Operation response error`, func() {
		createEndpointsPath := "/v1/locations/testString/endpoints"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEndpointsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateEndpoints with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the AdditionalNewEndpointRequestCertsClientCert model
				additionalNewEndpointRequestCertsClientCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsClientCert)
				additionalNewEndpointRequestCertsClientCertModel.Filename = core.StringPtr("clientEndpointCert.pem")
				additionalNewEndpointRequestCertsClientCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsClient model
				additionalNewEndpointRequestCertsClientModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsClient)
				additionalNewEndpointRequestCertsClientModel.Cert = additionalNewEndpointRequestCertsClientCertModel

				// Construct an instance of the AdditionalNewEndpointRequestCertsServerCert model
				additionalNewEndpointRequestCertsServerCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsServerCert)
				additionalNewEndpointRequestCertsServerCertModel.Filename = core.StringPtr("serverEndpointCert.pem")
				additionalNewEndpointRequestCertsServerCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsServer model
				additionalNewEndpointRequestCertsServerModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsServer)
				additionalNewEndpointRequestCertsServerModel.Cert = additionalNewEndpointRequestCertsServerCertModel

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnectorCert model
				additionalNewEndpointRequestCertsConnectorCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnectorCert)
				additionalNewEndpointRequestCertsConnectorCertModel.Filename = core.StringPtr("ConnectorCert.pem")
				additionalNewEndpointRequestCertsConnectorCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnectorKey model
				additionalNewEndpointRequestCertsConnectorKeyModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnectorKey)
				additionalNewEndpointRequestCertsConnectorKeyModel.Filename = core.StringPtr("ConnectorPrivateKey.pem")
				additionalNewEndpointRequestCertsConnectorKeyModel.FileContents = core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnector model
				additionalNewEndpointRequestCertsConnectorModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnector)
				additionalNewEndpointRequestCertsConnectorModel.Cert = additionalNewEndpointRequestCertsConnectorCertModel
				additionalNewEndpointRequestCertsConnectorModel.Key = additionalNewEndpointRequestCertsConnectorKeyModel

				// Construct an instance of the AdditionalNewEndpointRequestCerts model
				additionalNewEndpointRequestCertsModel := new(satellitelinkv1.AdditionalNewEndpointRequestCerts)
				additionalNewEndpointRequestCertsModel.Client = additionalNewEndpointRequestCertsClientModel
				additionalNewEndpointRequestCertsModel.Server = additionalNewEndpointRequestCertsServerModel
				additionalNewEndpointRequestCertsModel.Connector = additionalNewEndpointRequestCertsConnectorModel

				// Construct an instance of the CreateEndpointsOptions model
				createEndpointsOptionsModel := new(satellitelinkv1.CreateEndpointsOptions)
				createEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				createEndpointsOptionsModel.ConnType = core.StringPtr("cloud")
				createEndpointsOptionsModel.DisplayName = core.StringPtr("My endpoint")
				createEndpointsOptionsModel.ServerHost = core.StringPtr("example.com")
				createEndpointsOptionsModel.ServerPort = core.Int64Ptr(int64(443))
				createEndpointsOptionsModel.Sni = core.StringPtr("example.com")
				createEndpointsOptionsModel.ClientProtocol = core.StringPtr("https")
				createEndpointsOptionsModel.ClientMutualAuth = core.BoolPtr(true)
				createEndpointsOptionsModel.ServerProtocol = core.StringPtr("tls")
				createEndpointsOptionsModel.ServerMutualAuth = core.BoolPtr(true)
				createEndpointsOptionsModel.RejectUnauth = core.BoolPtr(true)
				createEndpointsOptionsModel.Timeout = core.Int64Ptr(int64(60))
				createEndpointsOptionsModel.CreatedBy = core.StringPtr("My service")
				createEndpointsOptionsModel.Certs = additionalNewEndpointRequestCertsModel
				createEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.CreateEndpoints(createEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.CreateEndpoints(createEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateEndpoints(createEndpointsOptions *CreateEndpointsOptions)`, func() {
		createEndpointsPath := "/v1/locations/testString/endpoints"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEndpointsPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke CreateEndpoints successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the AdditionalNewEndpointRequestCertsClientCert model
				additionalNewEndpointRequestCertsClientCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsClientCert)
				additionalNewEndpointRequestCertsClientCertModel.Filename = core.StringPtr("clientEndpointCert.pem")
				additionalNewEndpointRequestCertsClientCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsClient model
				additionalNewEndpointRequestCertsClientModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsClient)
				additionalNewEndpointRequestCertsClientModel.Cert = additionalNewEndpointRequestCertsClientCertModel

				// Construct an instance of the AdditionalNewEndpointRequestCertsServerCert model
				additionalNewEndpointRequestCertsServerCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsServerCert)
				additionalNewEndpointRequestCertsServerCertModel.Filename = core.StringPtr("serverEndpointCert.pem")
				additionalNewEndpointRequestCertsServerCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsServer model
				additionalNewEndpointRequestCertsServerModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsServer)
				additionalNewEndpointRequestCertsServerModel.Cert = additionalNewEndpointRequestCertsServerCertModel

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnectorCert model
				additionalNewEndpointRequestCertsConnectorCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnectorCert)
				additionalNewEndpointRequestCertsConnectorCertModel.Filename = core.StringPtr("ConnectorCert.pem")
				additionalNewEndpointRequestCertsConnectorCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnectorKey model
				additionalNewEndpointRequestCertsConnectorKeyModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnectorKey)
				additionalNewEndpointRequestCertsConnectorKeyModel.Filename = core.StringPtr("ConnectorPrivateKey.pem")
				additionalNewEndpointRequestCertsConnectorKeyModel.FileContents = core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnector model
				additionalNewEndpointRequestCertsConnectorModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnector)
				additionalNewEndpointRequestCertsConnectorModel.Cert = additionalNewEndpointRequestCertsConnectorCertModel
				additionalNewEndpointRequestCertsConnectorModel.Key = additionalNewEndpointRequestCertsConnectorKeyModel

				// Construct an instance of the AdditionalNewEndpointRequestCerts model
				additionalNewEndpointRequestCertsModel := new(satellitelinkv1.AdditionalNewEndpointRequestCerts)
				additionalNewEndpointRequestCertsModel.Client = additionalNewEndpointRequestCertsClientModel
				additionalNewEndpointRequestCertsModel.Server = additionalNewEndpointRequestCertsServerModel
				additionalNewEndpointRequestCertsModel.Connector = additionalNewEndpointRequestCertsConnectorModel

				// Construct an instance of the CreateEndpointsOptions model
				createEndpointsOptionsModel := new(satellitelinkv1.CreateEndpointsOptions)
				createEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				createEndpointsOptionsModel.ConnType = core.StringPtr("cloud")
				createEndpointsOptionsModel.DisplayName = core.StringPtr("My endpoint")
				createEndpointsOptionsModel.ServerHost = core.StringPtr("example.com")
				createEndpointsOptionsModel.ServerPort = core.Int64Ptr(int64(443))
				createEndpointsOptionsModel.Sni = core.StringPtr("example.com")
				createEndpointsOptionsModel.ClientProtocol = core.StringPtr("https")
				createEndpointsOptionsModel.ClientMutualAuth = core.BoolPtr(true)
				createEndpointsOptionsModel.ServerProtocol = core.StringPtr("tls")
				createEndpointsOptionsModel.ServerMutualAuth = core.BoolPtr(true)
				createEndpointsOptionsModel.RejectUnauth = core.BoolPtr(true)
				createEndpointsOptionsModel.Timeout = core.Int64Ptr(int64(60))
				createEndpointsOptionsModel.CreatedBy = core.StringPtr("My service")
				createEndpointsOptionsModel.Certs = additionalNewEndpointRequestCertsModel
				createEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.CreateEndpointsWithContext(ctx, createEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.CreateEndpoints(createEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.CreateEndpointsWithContext(ctx, createEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createEndpointsPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke CreateEndpoints successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.CreateEndpoints(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AdditionalNewEndpointRequestCertsClientCert model
				additionalNewEndpointRequestCertsClientCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsClientCert)
				additionalNewEndpointRequestCertsClientCertModel.Filename = core.StringPtr("clientEndpointCert.pem")
				additionalNewEndpointRequestCertsClientCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsClient model
				additionalNewEndpointRequestCertsClientModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsClient)
				additionalNewEndpointRequestCertsClientModel.Cert = additionalNewEndpointRequestCertsClientCertModel

				// Construct an instance of the AdditionalNewEndpointRequestCertsServerCert model
				additionalNewEndpointRequestCertsServerCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsServerCert)
				additionalNewEndpointRequestCertsServerCertModel.Filename = core.StringPtr("serverEndpointCert.pem")
				additionalNewEndpointRequestCertsServerCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsServer model
				additionalNewEndpointRequestCertsServerModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsServer)
				additionalNewEndpointRequestCertsServerModel.Cert = additionalNewEndpointRequestCertsServerCertModel

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnectorCert model
				additionalNewEndpointRequestCertsConnectorCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnectorCert)
				additionalNewEndpointRequestCertsConnectorCertModel.Filename = core.StringPtr("ConnectorCert.pem")
				additionalNewEndpointRequestCertsConnectorCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnectorKey model
				additionalNewEndpointRequestCertsConnectorKeyModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnectorKey)
				additionalNewEndpointRequestCertsConnectorKeyModel.Filename = core.StringPtr("ConnectorPrivateKey.pem")
				additionalNewEndpointRequestCertsConnectorKeyModel.FileContents = core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnector model
				additionalNewEndpointRequestCertsConnectorModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnector)
				additionalNewEndpointRequestCertsConnectorModel.Cert = additionalNewEndpointRequestCertsConnectorCertModel
				additionalNewEndpointRequestCertsConnectorModel.Key = additionalNewEndpointRequestCertsConnectorKeyModel

				// Construct an instance of the AdditionalNewEndpointRequestCerts model
				additionalNewEndpointRequestCertsModel := new(satellitelinkv1.AdditionalNewEndpointRequestCerts)
				additionalNewEndpointRequestCertsModel.Client = additionalNewEndpointRequestCertsClientModel
				additionalNewEndpointRequestCertsModel.Server = additionalNewEndpointRequestCertsServerModel
				additionalNewEndpointRequestCertsModel.Connector = additionalNewEndpointRequestCertsConnectorModel

				// Construct an instance of the CreateEndpointsOptions model
				createEndpointsOptionsModel := new(satellitelinkv1.CreateEndpointsOptions)
				createEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				createEndpointsOptionsModel.ConnType = core.StringPtr("cloud")
				createEndpointsOptionsModel.DisplayName = core.StringPtr("My endpoint")
				createEndpointsOptionsModel.ServerHost = core.StringPtr("example.com")
				createEndpointsOptionsModel.ServerPort = core.Int64Ptr(int64(443))
				createEndpointsOptionsModel.Sni = core.StringPtr("example.com")
				createEndpointsOptionsModel.ClientProtocol = core.StringPtr("https")
				createEndpointsOptionsModel.ClientMutualAuth = core.BoolPtr(true)
				createEndpointsOptionsModel.ServerProtocol = core.StringPtr("tls")
				createEndpointsOptionsModel.ServerMutualAuth = core.BoolPtr(true)
				createEndpointsOptionsModel.RejectUnauth = core.BoolPtr(true)
				createEndpointsOptionsModel.Timeout = core.Int64Ptr(int64(60))
				createEndpointsOptionsModel.CreatedBy = core.StringPtr("My service")
				createEndpointsOptionsModel.Certs = additionalNewEndpointRequestCertsModel
				createEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.CreateEndpoints(createEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateEndpoints with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the AdditionalNewEndpointRequestCertsClientCert model
				additionalNewEndpointRequestCertsClientCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsClientCert)
				additionalNewEndpointRequestCertsClientCertModel.Filename = core.StringPtr("clientEndpointCert.pem")
				additionalNewEndpointRequestCertsClientCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsClient model
				additionalNewEndpointRequestCertsClientModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsClient)
				additionalNewEndpointRequestCertsClientModel.Cert = additionalNewEndpointRequestCertsClientCertModel

				// Construct an instance of the AdditionalNewEndpointRequestCertsServerCert model
				additionalNewEndpointRequestCertsServerCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsServerCert)
				additionalNewEndpointRequestCertsServerCertModel.Filename = core.StringPtr("serverEndpointCert.pem")
				additionalNewEndpointRequestCertsServerCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsServer model
				additionalNewEndpointRequestCertsServerModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsServer)
				additionalNewEndpointRequestCertsServerModel.Cert = additionalNewEndpointRequestCertsServerCertModel

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnectorCert model
				additionalNewEndpointRequestCertsConnectorCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnectorCert)
				additionalNewEndpointRequestCertsConnectorCertModel.Filename = core.StringPtr("ConnectorCert.pem")
				additionalNewEndpointRequestCertsConnectorCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnectorKey model
				additionalNewEndpointRequestCertsConnectorKeyModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnectorKey)
				additionalNewEndpointRequestCertsConnectorKeyModel.Filename = core.StringPtr("ConnectorPrivateKey.pem")
				additionalNewEndpointRequestCertsConnectorKeyModel.FileContents = core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n")

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnector model
				additionalNewEndpointRequestCertsConnectorModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnector)
				additionalNewEndpointRequestCertsConnectorModel.Cert = additionalNewEndpointRequestCertsConnectorCertModel
				additionalNewEndpointRequestCertsConnectorModel.Key = additionalNewEndpointRequestCertsConnectorKeyModel

				// Construct an instance of the AdditionalNewEndpointRequestCerts model
				additionalNewEndpointRequestCertsModel := new(satellitelinkv1.AdditionalNewEndpointRequestCerts)
				additionalNewEndpointRequestCertsModel.Client = additionalNewEndpointRequestCertsClientModel
				additionalNewEndpointRequestCertsModel.Server = additionalNewEndpointRequestCertsServerModel
				additionalNewEndpointRequestCertsModel.Connector = additionalNewEndpointRequestCertsConnectorModel

				// Construct an instance of the CreateEndpointsOptions model
				createEndpointsOptionsModel := new(satellitelinkv1.CreateEndpointsOptions)
				createEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				createEndpointsOptionsModel.ConnType = core.StringPtr("cloud")
				createEndpointsOptionsModel.DisplayName = core.StringPtr("My endpoint")
				createEndpointsOptionsModel.ServerHost = core.StringPtr("example.com")
				createEndpointsOptionsModel.ServerPort = core.Int64Ptr(int64(443))
				createEndpointsOptionsModel.Sni = core.StringPtr("example.com")
				createEndpointsOptionsModel.ClientProtocol = core.StringPtr("https")
				createEndpointsOptionsModel.ClientMutualAuth = core.BoolPtr(true)
				createEndpointsOptionsModel.ServerProtocol = core.StringPtr("tls")
				createEndpointsOptionsModel.ServerMutualAuth = core.BoolPtr(true)
				createEndpointsOptionsModel.RejectUnauth = core.BoolPtr(true)
				createEndpointsOptionsModel.Timeout = core.Int64Ptr(int64(60))
				createEndpointsOptionsModel.CreatedBy = core.StringPtr("My service")
				createEndpointsOptionsModel.Certs = additionalNewEndpointRequestCertsModel
				createEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.CreateEndpoints(createEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateEndpointsOptions model with no property values
				createEndpointsOptionsModelNew := new(satellitelinkv1.CreateEndpointsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.CreateEndpoints(createEndpointsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ImportEndpoints(importEndpointsOptions *ImportEndpointsOptions) - Operation response error`, func() {
		importEndpointsPath := "/v1/locations/testString/endpoints/import"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importEndpointsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ImportEndpoints with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the ImportEndpointsOptions model
				importEndpointsOptionsModel := new(satellitelinkv1.ImportEndpointsOptions)
				importEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				importEndpointsOptionsModel.State = CreateMockReader("This is a mock file.")
				importEndpointsOptionsModel.StateContentType = core.StringPtr("testString")
				importEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.ImportEndpoints(importEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.ImportEndpoints(importEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ImportEndpoints(importEndpointsOptions *ImportEndpointsOptions)`, func() {
		importEndpointsPath := "/v1/locations/testString/endpoints/import"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importEndpointsPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke ImportEndpoints successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the ImportEndpointsOptions model
				importEndpointsOptionsModel := new(satellitelinkv1.ImportEndpointsOptions)
				importEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				importEndpointsOptionsModel.State = CreateMockReader("This is a mock file.")
				importEndpointsOptionsModel.StateContentType = core.StringPtr("testString")
				importEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.ImportEndpointsWithContext(ctx, importEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.ImportEndpoints(importEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.ImportEndpointsWithContext(ctx, importEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(importEndpointsPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke ImportEndpoints successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.ImportEndpoints(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ImportEndpointsOptions model
				importEndpointsOptionsModel := new(satellitelinkv1.ImportEndpointsOptions)
				importEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				importEndpointsOptionsModel.State = CreateMockReader("This is a mock file.")
				importEndpointsOptionsModel.StateContentType = core.StringPtr("testString")
				importEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.ImportEndpoints(importEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ImportEndpoints with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the ImportEndpointsOptions model
				importEndpointsOptionsModel := new(satellitelinkv1.ImportEndpointsOptions)
				importEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				importEndpointsOptionsModel.State = CreateMockReader("This is a mock file.")
				importEndpointsOptionsModel.StateContentType = core.StringPtr("testString")
				importEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.ImportEndpoints(importEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ImportEndpointsOptions model with no property values
				importEndpointsOptionsModelNew := new(satellitelinkv1.ImportEndpointsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.ImportEndpoints(importEndpointsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ExportEndpoints(exportEndpointsOptions *ExportEndpointsOptions) - Operation response error`, func() {
		exportEndpointsPath := "/v1/locations/testString/endpoints/testString/export"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(exportEndpointsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ExportEndpoints with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the ExportEndpointsOptions model
				exportEndpointsOptionsModel := new(satellitelinkv1.ExportEndpointsOptions)
				exportEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				exportEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				exportEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.ExportEndpoints(exportEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.ExportEndpoints(exportEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ExportEndpoints(exportEndpointsOptions *ExportEndpointsOptions)`, func() {
		exportEndpointsPath := "/v1/locations/testString/endpoints/testString/export"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(exportEndpointsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem", "file_contents": "-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n"}}, "server": {"cert": {"filename": "serverEndpointCert.pem", "file_contents": "-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n"}}, "connector": {"cert": {"filename": "ConnectorCert.pem", "file_contents": "-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n"}, "key": {"filename": "ConnectorPrivateKey.pem", "file_contents": "-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n"}}}}`)
				}))
			})
			It(`Invoke ExportEndpoints successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the ExportEndpointsOptions model
				exportEndpointsOptionsModel := new(satellitelinkv1.ExportEndpointsOptions)
				exportEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				exportEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				exportEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.ExportEndpointsWithContext(ctx, exportEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.ExportEndpoints(exportEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.ExportEndpointsWithContext(ctx, exportEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(exportEndpointsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem", "file_contents": "-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n"}}, "server": {"cert": {"filename": "serverEndpointCert.pem", "file_contents": "-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n"}}, "connector": {"cert": {"filename": "ConnectorCert.pem", "file_contents": "-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n"}, "key": {"filename": "ConnectorPrivateKey.pem", "file_contents": "-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n"}}}}`)
				}))
			})
			It(`Invoke ExportEndpoints successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.ExportEndpoints(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ExportEndpointsOptions model
				exportEndpointsOptionsModel := new(satellitelinkv1.ExportEndpointsOptions)
				exportEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				exportEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				exportEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.ExportEndpoints(exportEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ExportEndpoints with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the ExportEndpointsOptions model
				exportEndpointsOptionsModel := new(satellitelinkv1.ExportEndpointsOptions)
				exportEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				exportEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				exportEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.ExportEndpoints(exportEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ExportEndpointsOptions model with no property values
				exportEndpointsOptionsModelNew := new(satellitelinkv1.ExportEndpointsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.ExportEndpoints(exportEndpointsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEndpoints(getEndpointsOptions *GetEndpointsOptions) - Operation response error`, func() {
		getEndpointsPath := "/v1/locations/testString/endpoints/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEndpointsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEndpoints with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the GetEndpointsOptions model
				getEndpointsOptionsModel := new(satellitelinkv1.GetEndpointsOptions)
				getEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				getEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				getEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.GetEndpoints(getEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.GetEndpoints(getEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetEndpoints(getEndpointsOptions *GetEndpointsOptions)`, func() {
		getEndpointsPath := "/v1/locations/testString/endpoints/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEndpointsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke GetEndpoints successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetEndpointsOptions model
				getEndpointsOptionsModel := new(satellitelinkv1.GetEndpointsOptions)
				getEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				getEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				getEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.GetEndpointsWithContext(ctx, getEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.GetEndpoints(getEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.GetEndpointsWithContext(ctx, getEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEndpointsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke GetEndpoints successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.GetEndpoints(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEndpointsOptions model
				getEndpointsOptionsModel := new(satellitelinkv1.GetEndpointsOptions)
				getEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				getEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				getEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.GetEndpoints(getEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEndpoints with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the GetEndpointsOptions model
				getEndpointsOptionsModel := new(satellitelinkv1.GetEndpointsOptions)
				getEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				getEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				getEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.GetEndpoints(getEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEndpointsOptions model with no property values
				getEndpointsOptionsModelNew := new(satellitelinkv1.GetEndpointsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.GetEndpoints(getEndpointsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEndpoints(updateEndpointsOptions *UpdateEndpointsOptions) - Operation response error`, func() {
		updateEndpointsPath := "/v1/locations/testString/endpoints/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEndpointsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateEndpoints with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the UpdatedEndpointRequestCertsClientCert model
				updatedEndpointRequestCertsClientCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsClientCert)
				updatedEndpointRequestCertsClientCertModel.Filename = core.StringPtr("clientEndpointCert.pem")
				updatedEndpointRequestCertsClientCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the UpdatedEndpointRequestCertsClient model
				updatedEndpointRequestCertsClientModel := new(satellitelinkv1.UpdatedEndpointRequestCertsClient)
				updatedEndpointRequestCertsClientModel.Cert = updatedEndpointRequestCertsClientCertModel

				// Construct an instance of the UpdatedEndpointRequestCertsServerCert model
				updatedEndpointRequestCertsServerCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsServerCert)
				updatedEndpointRequestCertsServerCertModel.Filename = core.StringPtr("serverEndpointCert.pem")
				updatedEndpointRequestCertsServerCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the UpdatedEndpointRequestCertsServer model
				updatedEndpointRequestCertsServerModel := new(satellitelinkv1.UpdatedEndpointRequestCertsServer)
				updatedEndpointRequestCertsServerModel.Cert = updatedEndpointRequestCertsServerCertModel

				// Construct an instance of the UpdatedEndpointRequestCertsConnectorCert model
				updatedEndpointRequestCertsConnectorCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnectorCert)
				updatedEndpointRequestCertsConnectorCertModel.Filename = core.StringPtr("ConnectorCert.pem")
				updatedEndpointRequestCertsConnectorCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the UpdatedEndpointRequestCertsConnectorKey model
				updatedEndpointRequestCertsConnectorKeyModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnectorKey)
				updatedEndpointRequestCertsConnectorKeyModel.Filename = core.StringPtr("ConnectorPrivateKey.pem")
				updatedEndpointRequestCertsConnectorKeyModel.FileContents = core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n")

				// Construct an instance of the UpdatedEndpointRequestCertsConnector model
				updatedEndpointRequestCertsConnectorModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnector)
				updatedEndpointRequestCertsConnectorModel.Cert = updatedEndpointRequestCertsConnectorCertModel
				updatedEndpointRequestCertsConnectorModel.Key = updatedEndpointRequestCertsConnectorKeyModel

				// Construct an instance of the UpdatedEndpointRequestCerts model
				updatedEndpointRequestCertsModel := new(satellitelinkv1.UpdatedEndpointRequestCerts)
				updatedEndpointRequestCertsModel.Client = updatedEndpointRequestCertsClientModel
				updatedEndpointRequestCertsModel.Server = updatedEndpointRequestCertsServerModel
				updatedEndpointRequestCertsModel.Connector = updatedEndpointRequestCertsConnectorModel

				// Construct an instance of the UpdateEndpointsOptions model
				updateEndpointsOptionsModel := new(satellitelinkv1.UpdateEndpointsOptions)
				updateEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				updateEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				updateEndpointsOptionsModel.DisplayName = core.StringPtr("My endpoint")
				updateEndpointsOptionsModel.ServerHost = core.StringPtr("example.com")
				updateEndpointsOptionsModel.ServerPort = core.Int64Ptr(int64(443))
				updateEndpointsOptionsModel.Sni = core.StringPtr("example.com")
				updateEndpointsOptionsModel.ClientProtocol = core.StringPtr("https")
				updateEndpointsOptionsModel.ClientMutualAuth = core.BoolPtr(true)
				updateEndpointsOptionsModel.ServerProtocol = core.StringPtr("tls")
				updateEndpointsOptionsModel.ServerMutualAuth = core.BoolPtr(true)
				updateEndpointsOptionsModel.RejectUnauth = core.BoolPtr(true)
				updateEndpointsOptionsModel.Timeout = core.Int64Ptr(int64(60))
				updateEndpointsOptionsModel.CreatedBy = core.StringPtr("My service")
				updateEndpointsOptionsModel.Certs = updatedEndpointRequestCertsModel
				updateEndpointsOptionsModel.Enabled = core.BoolPtr(true)
				updateEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.UpdateEndpoints(updateEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.UpdateEndpoints(updateEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateEndpoints(updateEndpointsOptions *UpdateEndpointsOptions)`, func() {
		updateEndpointsPath := "/v1/locations/testString/endpoints/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEndpointsPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke UpdateEndpoints successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the UpdatedEndpointRequestCertsClientCert model
				updatedEndpointRequestCertsClientCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsClientCert)
				updatedEndpointRequestCertsClientCertModel.Filename = core.StringPtr("clientEndpointCert.pem")
				updatedEndpointRequestCertsClientCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the UpdatedEndpointRequestCertsClient model
				updatedEndpointRequestCertsClientModel := new(satellitelinkv1.UpdatedEndpointRequestCertsClient)
				updatedEndpointRequestCertsClientModel.Cert = updatedEndpointRequestCertsClientCertModel

				// Construct an instance of the UpdatedEndpointRequestCertsServerCert model
				updatedEndpointRequestCertsServerCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsServerCert)
				updatedEndpointRequestCertsServerCertModel.Filename = core.StringPtr("serverEndpointCert.pem")
				updatedEndpointRequestCertsServerCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the UpdatedEndpointRequestCertsServer model
				updatedEndpointRequestCertsServerModel := new(satellitelinkv1.UpdatedEndpointRequestCertsServer)
				updatedEndpointRequestCertsServerModel.Cert = updatedEndpointRequestCertsServerCertModel

				// Construct an instance of the UpdatedEndpointRequestCertsConnectorCert model
				updatedEndpointRequestCertsConnectorCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnectorCert)
				updatedEndpointRequestCertsConnectorCertModel.Filename = core.StringPtr("ConnectorCert.pem")
				updatedEndpointRequestCertsConnectorCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the UpdatedEndpointRequestCertsConnectorKey model
				updatedEndpointRequestCertsConnectorKeyModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnectorKey)
				updatedEndpointRequestCertsConnectorKeyModel.Filename = core.StringPtr("ConnectorPrivateKey.pem")
				updatedEndpointRequestCertsConnectorKeyModel.FileContents = core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n")

				// Construct an instance of the UpdatedEndpointRequestCertsConnector model
				updatedEndpointRequestCertsConnectorModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnector)
				updatedEndpointRequestCertsConnectorModel.Cert = updatedEndpointRequestCertsConnectorCertModel
				updatedEndpointRequestCertsConnectorModel.Key = updatedEndpointRequestCertsConnectorKeyModel

				// Construct an instance of the UpdatedEndpointRequestCerts model
				updatedEndpointRequestCertsModel := new(satellitelinkv1.UpdatedEndpointRequestCerts)
				updatedEndpointRequestCertsModel.Client = updatedEndpointRequestCertsClientModel
				updatedEndpointRequestCertsModel.Server = updatedEndpointRequestCertsServerModel
				updatedEndpointRequestCertsModel.Connector = updatedEndpointRequestCertsConnectorModel

				// Construct an instance of the UpdateEndpointsOptions model
				updateEndpointsOptionsModel := new(satellitelinkv1.UpdateEndpointsOptions)
				updateEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				updateEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				updateEndpointsOptionsModel.DisplayName = core.StringPtr("My endpoint")
				updateEndpointsOptionsModel.ServerHost = core.StringPtr("example.com")
				updateEndpointsOptionsModel.ServerPort = core.Int64Ptr(int64(443))
				updateEndpointsOptionsModel.Sni = core.StringPtr("example.com")
				updateEndpointsOptionsModel.ClientProtocol = core.StringPtr("https")
				updateEndpointsOptionsModel.ClientMutualAuth = core.BoolPtr(true)
				updateEndpointsOptionsModel.ServerProtocol = core.StringPtr("tls")
				updateEndpointsOptionsModel.ServerMutualAuth = core.BoolPtr(true)
				updateEndpointsOptionsModel.RejectUnauth = core.BoolPtr(true)
				updateEndpointsOptionsModel.Timeout = core.Int64Ptr(int64(60))
				updateEndpointsOptionsModel.CreatedBy = core.StringPtr("My service")
				updateEndpointsOptionsModel.Certs = updatedEndpointRequestCertsModel
				updateEndpointsOptionsModel.Enabled = core.BoolPtr(true)
				updateEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.UpdateEndpointsWithContext(ctx, updateEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.UpdateEndpoints(updateEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.UpdateEndpointsWithContext(ctx, updateEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEndpointsPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke UpdateEndpoints successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.UpdateEndpoints(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdatedEndpointRequestCertsClientCert model
				updatedEndpointRequestCertsClientCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsClientCert)
				updatedEndpointRequestCertsClientCertModel.Filename = core.StringPtr("clientEndpointCert.pem")
				updatedEndpointRequestCertsClientCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the UpdatedEndpointRequestCertsClient model
				updatedEndpointRequestCertsClientModel := new(satellitelinkv1.UpdatedEndpointRequestCertsClient)
				updatedEndpointRequestCertsClientModel.Cert = updatedEndpointRequestCertsClientCertModel

				// Construct an instance of the UpdatedEndpointRequestCertsServerCert model
				updatedEndpointRequestCertsServerCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsServerCert)
				updatedEndpointRequestCertsServerCertModel.Filename = core.StringPtr("serverEndpointCert.pem")
				updatedEndpointRequestCertsServerCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the UpdatedEndpointRequestCertsServer model
				updatedEndpointRequestCertsServerModel := new(satellitelinkv1.UpdatedEndpointRequestCertsServer)
				updatedEndpointRequestCertsServerModel.Cert = updatedEndpointRequestCertsServerCertModel

				// Construct an instance of the UpdatedEndpointRequestCertsConnectorCert model
				updatedEndpointRequestCertsConnectorCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnectorCert)
				updatedEndpointRequestCertsConnectorCertModel.Filename = core.StringPtr("ConnectorCert.pem")
				updatedEndpointRequestCertsConnectorCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the UpdatedEndpointRequestCertsConnectorKey model
				updatedEndpointRequestCertsConnectorKeyModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnectorKey)
				updatedEndpointRequestCertsConnectorKeyModel.Filename = core.StringPtr("ConnectorPrivateKey.pem")
				updatedEndpointRequestCertsConnectorKeyModel.FileContents = core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n")

				// Construct an instance of the UpdatedEndpointRequestCertsConnector model
				updatedEndpointRequestCertsConnectorModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnector)
				updatedEndpointRequestCertsConnectorModel.Cert = updatedEndpointRequestCertsConnectorCertModel
				updatedEndpointRequestCertsConnectorModel.Key = updatedEndpointRequestCertsConnectorKeyModel

				// Construct an instance of the UpdatedEndpointRequestCerts model
				updatedEndpointRequestCertsModel := new(satellitelinkv1.UpdatedEndpointRequestCerts)
				updatedEndpointRequestCertsModel.Client = updatedEndpointRequestCertsClientModel
				updatedEndpointRequestCertsModel.Server = updatedEndpointRequestCertsServerModel
				updatedEndpointRequestCertsModel.Connector = updatedEndpointRequestCertsConnectorModel

				// Construct an instance of the UpdateEndpointsOptions model
				updateEndpointsOptionsModel := new(satellitelinkv1.UpdateEndpointsOptions)
				updateEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				updateEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				updateEndpointsOptionsModel.DisplayName = core.StringPtr("My endpoint")
				updateEndpointsOptionsModel.ServerHost = core.StringPtr("example.com")
				updateEndpointsOptionsModel.ServerPort = core.Int64Ptr(int64(443))
				updateEndpointsOptionsModel.Sni = core.StringPtr("example.com")
				updateEndpointsOptionsModel.ClientProtocol = core.StringPtr("https")
				updateEndpointsOptionsModel.ClientMutualAuth = core.BoolPtr(true)
				updateEndpointsOptionsModel.ServerProtocol = core.StringPtr("tls")
				updateEndpointsOptionsModel.ServerMutualAuth = core.BoolPtr(true)
				updateEndpointsOptionsModel.RejectUnauth = core.BoolPtr(true)
				updateEndpointsOptionsModel.Timeout = core.Int64Ptr(int64(60))
				updateEndpointsOptionsModel.CreatedBy = core.StringPtr("My service")
				updateEndpointsOptionsModel.Certs = updatedEndpointRequestCertsModel
				updateEndpointsOptionsModel.Enabled = core.BoolPtr(true)
				updateEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.UpdateEndpoints(updateEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateEndpoints with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the UpdatedEndpointRequestCertsClientCert model
				updatedEndpointRequestCertsClientCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsClientCert)
				updatedEndpointRequestCertsClientCertModel.Filename = core.StringPtr("clientEndpointCert.pem")
				updatedEndpointRequestCertsClientCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the UpdatedEndpointRequestCertsClient model
				updatedEndpointRequestCertsClientModel := new(satellitelinkv1.UpdatedEndpointRequestCertsClient)
				updatedEndpointRequestCertsClientModel.Cert = updatedEndpointRequestCertsClientCertModel

				// Construct an instance of the UpdatedEndpointRequestCertsServerCert model
				updatedEndpointRequestCertsServerCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsServerCert)
				updatedEndpointRequestCertsServerCertModel.Filename = core.StringPtr("serverEndpointCert.pem")
				updatedEndpointRequestCertsServerCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the UpdatedEndpointRequestCertsServer model
				updatedEndpointRequestCertsServerModel := new(satellitelinkv1.UpdatedEndpointRequestCertsServer)
				updatedEndpointRequestCertsServerModel.Cert = updatedEndpointRequestCertsServerCertModel

				// Construct an instance of the UpdatedEndpointRequestCertsConnectorCert model
				updatedEndpointRequestCertsConnectorCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnectorCert)
				updatedEndpointRequestCertsConnectorCertModel.Filename = core.StringPtr("ConnectorCert.pem")
				updatedEndpointRequestCertsConnectorCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")

				// Construct an instance of the UpdatedEndpointRequestCertsConnectorKey model
				updatedEndpointRequestCertsConnectorKeyModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnectorKey)
				updatedEndpointRequestCertsConnectorKeyModel.Filename = core.StringPtr("ConnectorPrivateKey.pem")
				updatedEndpointRequestCertsConnectorKeyModel.FileContents = core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n")

				// Construct an instance of the UpdatedEndpointRequestCertsConnector model
				updatedEndpointRequestCertsConnectorModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnector)
				updatedEndpointRequestCertsConnectorModel.Cert = updatedEndpointRequestCertsConnectorCertModel
				updatedEndpointRequestCertsConnectorModel.Key = updatedEndpointRequestCertsConnectorKeyModel

				// Construct an instance of the UpdatedEndpointRequestCerts model
				updatedEndpointRequestCertsModel := new(satellitelinkv1.UpdatedEndpointRequestCerts)
				updatedEndpointRequestCertsModel.Client = updatedEndpointRequestCertsClientModel
				updatedEndpointRequestCertsModel.Server = updatedEndpointRequestCertsServerModel
				updatedEndpointRequestCertsModel.Connector = updatedEndpointRequestCertsConnectorModel

				// Construct an instance of the UpdateEndpointsOptions model
				updateEndpointsOptionsModel := new(satellitelinkv1.UpdateEndpointsOptions)
				updateEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				updateEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				updateEndpointsOptionsModel.DisplayName = core.StringPtr("My endpoint")
				updateEndpointsOptionsModel.ServerHost = core.StringPtr("example.com")
				updateEndpointsOptionsModel.ServerPort = core.Int64Ptr(int64(443))
				updateEndpointsOptionsModel.Sni = core.StringPtr("example.com")
				updateEndpointsOptionsModel.ClientProtocol = core.StringPtr("https")
				updateEndpointsOptionsModel.ClientMutualAuth = core.BoolPtr(true)
				updateEndpointsOptionsModel.ServerProtocol = core.StringPtr("tls")
				updateEndpointsOptionsModel.ServerMutualAuth = core.BoolPtr(true)
				updateEndpointsOptionsModel.RejectUnauth = core.BoolPtr(true)
				updateEndpointsOptionsModel.Timeout = core.Int64Ptr(int64(60))
				updateEndpointsOptionsModel.CreatedBy = core.StringPtr("My service")
				updateEndpointsOptionsModel.Certs = updatedEndpointRequestCertsModel
				updateEndpointsOptionsModel.Enabled = core.BoolPtr(true)
				updateEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.UpdateEndpoints(updateEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateEndpointsOptions model with no property values
				updateEndpointsOptionsModelNew := new(satellitelinkv1.UpdateEndpointsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.UpdateEndpoints(updateEndpointsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteEndpoints(deleteEndpointsOptions *DeleteEndpointsOptions) - Operation response error`, func() {
		deleteEndpointsPath := "/v1/locations/testString/endpoints/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEndpointsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteEndpoints with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteEndpointsOptions model
				deleteEndpointsOptionsModel := new(satellitelinkv1.DeleteEndpointsOptions)
				deleteEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				deleteEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				deleteEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.DeleteEndpoints(deleteEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.DeleteEndpoints(deleteEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteEndpoints(deleteEndpointsOptions *DeleteEndpointsOptions)`, func() {
		deleteEndpointsPath := "/v1/locations/testString/endpoints/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEndpointsPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "Success"}`)
				}))
			})
			It(`Invoke DeleteEndpoints successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the DeleteEndpointsOptions model
				deleteEndpointsOptionsModel := new(satellitelinkv1.DeleteEndpointsOptions)
				deleteEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				deleteEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				deleteEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.DeleteEndpointsWithContext(ctx, deleteEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.DeleteEndpoints(deleteEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.DeleteEndpointsWithContext(ctx, deleteEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEndpointsPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "Success"}`)
				}))
			})
			It(`Invoke DeleteEndpoints successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.DeleteEndpoints(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteEndpointsOptions model
				deleteEndpointsOptionsModel := new(satellitelinkv1.DeleteEndpointsOptions)
				deleteEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				deleteEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				deleteEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.DeleteEndpoints(deleteEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteEndpoints with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteEndpointsOptions model
				deleteEndpointsOptionsModel := new(satellitelinkv1.DeleteEndpointsOptions)
				deleteEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				deleteEndpointsOptionsModel.EndpointID = core.StringPtr("testString")
				deleteEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.DeleteEndpoints(deleteEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteEndpointsOptions model with no property values
				deleteEndpointsOptionsModelNew := new(satellitelinkv1.DeleteEndpointsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.DeleteEndpoints(deleteEndpointsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEndpointCerts(getEndpointCertsOptions *GetEndpointCertsOptions) - Operation response error`, func() {
		getEndpointCertsPath := "/v1/locations/testString/endpoints/testString/cert"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEndpointCertsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for no_zip query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEndpointCerts with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the GetEndpointCertsOptions model
				getEndpointCertsOptionsModel := new(satellitelinkv1.GetEndpointCertsOptions)
				getEndpointCertsOptionsModel.LocationID = core.StringPtr("testString")
				getEndpointCertsOptionsModel.EndpointID = core.StringPtr("testString")
				getEndpointCertsOptionsModel.NoZip = core.BoolPtr(true)
				getEndpointCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.GetEndpointCerts(getEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.GetEndpointCerts(getEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetEndpointCerts(getEndpointCertsOptions *GetEndpointCertsOptions)`, func() {
		getEndpointCertsPath := "/v1/locations/testString/endpoints/testString/cert"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEndpointCertsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for no_zip query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certs": [{"name": "DigiCertCA2.pem", "content": "-----BEGIN CERTIFICATE-----\r\n<The-content-of-the-cert>\r\n-----END CERTIFICATE-----\r\n"}]}`)
				}))
			})
			It(`Invoke GetEndpointCerts successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the GetEndpointCertsOptions model
				getEndpointCertsOptionsModel := new(satellitelinkv1.GetEndpointCertsOptions)
				getEndpointCertsOptionsModel.LocationID = core.StringPtr("testString")
				getEndpointCertsOptionsModel.EndpointID = core.StringPtr("testString")
				getEndpointCertsOptionsModel.NoZip = core.BoolPtr(true)
				getEndpointCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.GetEndpointCertsWithContext(ctx, getEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.GetEndpointCerts(getEndpointCertsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.GetEndpointCertsWithContext(ctx, getEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEndpointCertsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for no_zip query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certs": [{"name": "DigiCertCA2.pem", "content": "-----BEGIN CERTIFICATE-----\r\n<The-content-of-the-cert>\r\n-----END CERTIFICATE-----\r\n"}]}`)
				}))
			})
			It(`Invoke GetEndpointCerts successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.GetEndpointCerts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEndpointCertsOptions model
				getEndpointCertsOptionsModel := new(satellitelinkv1.GetEndpointCertsOptions)
				getEndpointCertsOptionsModel.LocationID = core.StringPtr("testString")
				getEndpointCertsOptionsModel.EndpointID = core.StringPtr("testString")
				getEndpointCertsOptionsModel.NoZip = core.BoolPtr(true)
				getEndpointCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.GetEndpointCerts(getEndpointCertsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEndpointCerts with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the GetEndpointCertsOptions model
				getEndpointCertsOptionsModel := new(satellitelinkv1.GetEndpointCertsOptions)
				getEndpointCertsOptionsModel.LocationID = core.StringPtr("testString")
				getEndpointCertsOptionsModel.EndpointID = core.StringPtr("testString")
				getEndpointCertsOptionsModel.NoZip = core.BoolPtr(true)
				getEndpointCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.GetEndpointCerts(getEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetEndpointCertsOptions model with no property values
				getEndpointCertsOptionsModelNew := new(satellitelinkv1.GetEndpointCertsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.GetEndpointCerts(getEndpointCertsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UploadEndpointCerts(uploadEndpointCertsOptions *UploadEndpointCertsOptions) - Operation response error`, func() {
		uploadEndpointCertsPath := "/v1/locations/testString/endpoints/testString/cert"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadEndpointCertsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UploadEndpointCerts with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the UploadEndpointCertsOptions model
				uploadEndpointCertsOptionsModel := new(satellitelinkv1.UploadEndpointCertsOptions)
				uploadEndpointCertsOptionsModel.LocationID = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.EndpointID = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ClientCert = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ClientCertContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ServerCert = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ServerCertContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ConnectorCert = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ConnectorCertContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ConnectorKey = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ConnectorKeyContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.UploadEndpointCerts(uploadEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.UploadEndpointCerts(uploadEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UploadEndpointCerts(uploadEndpointCertsOptions *UploadEndpointCertsOptions)`, func() {
		uploadEndpointCertsPath := "/v1/locations/testString/endpoints/testString/cert"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadEndpointCertsPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke UploadEndpointCerts successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the UploadEndpointCertsOptions model
				uploadEndpointCertsOptionsModel := new(satellitelinkv1.UploadEndpointCertsOptions)
				uploadEndpointCertsOptionsModel.LocationID = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.EndpointID = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ClientCert = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ClientCertContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ServerCert = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ServerCertContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ConnectorCert = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ConnectorCertContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ConnectorKey = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ConnectorKeyContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.UploadEndpointCertsWithContext(ctx, uploadEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.UploadEndpointCerts(uploadEndpointCertsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.UploadEndpointCertsWithContext(ctx, uploadEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uploadEndpointCertsPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke UploadEndpointCerts successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.UploadEndpointCerts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UploadEndpointCertsOptions model
				uploadEndpointCertsOptionsModel := new(satellitelinkv1.UploadEndpointCertsOptions)
				uploadEndpointCertsOptionsModel.LocationID = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.EndpointID = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ClientCert = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ClientCertContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ServerCert = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ServerCertContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ConnectorCert = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ConnectorCertContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ConnectorKey = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ConnectorKeyContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.UploadEndpointCerts(uploadEndpointCertsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UploadEndpointCerts with error: Param validation error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the UploadEndpointCertsOptions model
				uploadEndpointCertsOptionsModel := new(satellitelinkv1.UploadEndpointCertsOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := satelliteLinkService.UploadEndpointCerts(uploadEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke UploadEndpointCerts with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the UploadEndpointCertsOptions model
				uploadEndpointCertsOptionsModel := new(satellitelinkv1.UploadEndpointCertsOptions)
				uploadEndpointCertsOptionsModel.LocationID = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.EndpointID = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ClientCert = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ClientCertContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ServerCert = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ServerCertContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ConnectorCert = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ConnectorCertContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.ConnectorKey = CreateMockReader("This is a mock file.")
				uploadEndpointCertsOptionsModel.ConnectorKeyContentType = core.StringPtr("testString")
				uploadEndpointCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.UploadEndpointCerts(uploadEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UploadEndpointCertsOptions model with no property values
				uploadEndpointCertsOptionsModelNew := new(satellitelinkv1.UploadEndpointCertsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.UploadEndpointCerts(uploadEndpointCertsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteEndpointCerts(deleteEndpointCertsOptions *DeleteEndpointCertsOptions) - Operation response error`, func() {
		deleteEndpointCertsPath := "/v1/locations/testString/endpoints/testString/cert"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEndpointCertsPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteEndpointCerts with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteEndpointCertsOptions model
				deleteEndpointCertsOptionsModel := new(satellitelinkv1.DeleteEndpointCertsOptions)
				deleteEndpointCertsOptionsModel.LocationID = core.StringPtr("testString")
				deleteEndpointCertsOptionsModel.EndpointID = core.StringPtr("testString")
				deleteEndpointCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.DeleteEndpointCerts(deleteEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.DeleteEndpointCerts(deleteEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteEndpointCerts(deleteEndpointCertsOptions *DeleteEndpointCertsOptions)`, func() {
		deleteEndpointCertsPath := "/v1/locations/testString/endpoints/testString/cert"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEndpointCertsPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke DeleteEndpointCerts successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the DeleteEndpointCertsOptions model
				deleteEndpointCertsOptionsModel := new(satellitelinkv1.DeleteEndpointCertsOptions)
				deleteEndpointCertsOptionsModel.LocationID = core.StringPtr("testString")
				deleteEndpointCertsOptionsModel.EndpointID = core.StringPtr("testString")
				deleteEndpointCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.DeleteEndpointCertsWithContext(ctx, deleteEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.DeleteEndpointCerts(deleteEndpointCertsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.DeleteEndpointCertsWithContext(ctx, deleteEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteEndpointCertsPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"conn_type": "cloud", "display_name": "My endpoint", "server_host": "example.com", "server_port": 443, "sni": "example.com", "client_protocol": "https", "client_mutual_auth": true, "server_protocol": "tls", "server_mutual_auth": true, "reject_unauth": true, "timeout": 60, "created_by": "My service", "sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}], "connector_port": 29999, "crn": "{crn}", "endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "service_name": "myendpoint", "location_id": "us-south--K9kQEVFmqNpP", "client_host": "{satellite-link-tunnel-server}", "client_port": 15001, "certs": {"client": {"cert": {"filename": "clientEndpointCert.pem"}}, "server": {"cert": {"filename": "serverEndpointCert.pem"}}, "connector": {"cert": {"filename": "ConnectorCert.pem"}, "key": {"filename": "ConnectorPrivateKey.pem"}}}, "status": "enabled", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z", "performance": {"connection": 5, "rx_bandwidth": 5, "tx_bandwidth": 5, "bandwidth": 10, "connectors": [{"connector": "satellite-link-connector-9487bf46c-4sp9z", "connections": 5, "rxBW": 5, "txBW": 5}]}}`)
				}))
			})
			It(`Invoke DeleteEndpointCerts successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.DeleteEndpointCerts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteEndpointCertsOptions model
				deleteEndpointCertsOptionsModel := new(satellitelinkv1.DeleteEndpointCertsOptions)
				deleteEndpointCertsOptionsModel.LocationID = core.StringPtr("testString")
				deleteEndpointCertsOptionsModel.EndpointID = core.StringPtr("testString")
				deleteEndpointCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.DeleteEndpointCerts(deleteEndpointCertsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteEndpointCerts with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteEndpointCertsOptions model
				deleteEndpointCertsOptionsModel := new(satellitelinkv1.DeleteEndpointCertsOptions)
				deleteEndpointCertsOptionsModel.LocationID = core.StringPtr("testString")
				deleteEndpointCertsOptionsModel.EndpointID = core.StringPtr("testString")
				deleteEndpointCertsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.DeleteEndpointCerts(deleteEndpointCertsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteEndpointCertsOptions model with no property values
				deleteEndpointCertsOptionsModelNew := new(satellitelinkv1.DeleteEndpointCertsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.DeleteEndpointCerts(deleteEndpointCertsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEndpointSources(listEndpointSourcesOptions *ListEndpointSourcesOptions) - Operation response error`, func() {
		listEndpointSourcesPath := "/v1/locations/testString/endpoints/testString/sources"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEndpointSourcesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEndpointSources with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the ListEndpointSourcesOptions model
				listEndpointSourcesOptionsModel := new(satellitelinkv1.ListEndpointSourcesOptions)
				listEndpointSourcesOptionsModel.LocationID = core.StringPtr("testString")
				listEndpointSourcesOptionsModel.EndpointID = core.StringPtr("testString")
				listEndpointSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.ListEndpointSources(listEndpointSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.ListEndpointSources(listEndpointSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListEndpointSources(listEndpointSourcesOptions *ListEndpointSourcesOptions)`, func() {
		listEndpointSourcesPath := "/v1/locations/testString/endpoints/testString/sources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEndpointSourcesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}]}`)
				}))
			})
			It(`Invoke ListEndpointSources successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListEndpointSourcesOptions model
				listEndpointSourcesOptionsModel := new(satellitelinkv1.ListEndpointSourcesOptions)
				listEndpointSourcesOptionsModel.LocationID = core.StringPtr("testString")
				listEndpointSourcesOptionsModel.EndpointID = core.StringPtr("testString")
				listEndpointSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.ListEndpointSourcesWithContext(ctx, listEndpointSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.ListEndpointSources(listEndpointSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.ListEndpointSourcesWithContext(ctx, listEndpointSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEndpointSourcesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}]}`)
				}))
			})
			It(`Invoke ListEndpointSources successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.ListEndpointSources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEndpointSourcesOptions model
				listEndpointSourcesOptionsModel := new(satellitelinkv1.ListEndpointSourcesOptions)
				listEndpointSourcesOptionsModel.LocationID = core.StringPtr("testString")
				listEndpointSourcesOptionsModel.EndpointID = core.StringPtr("testString")
				listEndpointSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.ListEndpointSources(listEndpointSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListEndpointSources with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the ListEndpointSourcesOptions model
				listEndpointSourcesOptionsModel := new(satellitelinkv1.ListEndpointSourcesOptions)
				listEndpointSourcesOptionsModel.LocationID = core.StringPtr("testString")
				listEndpointSourcesOptionsModel.EndpointID = core.StringPtr("testString")
				listEndpointSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.ListEndpointSources(listEndpointSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListEndpointSourcesOptions model with no property values
				listEndpointSourcesOptionsModelNew := new(satellitelinkv1.ListEndpointSourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.ListEndpointSources(listEndpointSourcesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateEndpointSources(updateEndpointSourcesOptions *UpdateEndpointSourcesOptions) - Operation response error`, func() {
		updateEndpointSourcesPath := "/v1/locations/testString/endpoints/testString/sources"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEndpointSourcesPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateEndpointSources with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the SourceStatusRequestObject model
				sourceStatusRequestObjectModel := new(satellitelinkv1.SourceStatusRequestObject)
				sourceStatusRequestObjectModel.SourceID = core.StringPtr("us-south--K9kQEVFmqNpP-Source-Q87fe")
				sourceStatusRequestObjectModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateEndpointSourcesOptions model
				updateEndpointSourcesOptionsModel := new(satellitelinkv1.UpdateEndpointSourcesOptions)
				updateEndpointSourcesOptionsModel.LocationID = core.StringPtr("testString")
				updateEndpointSourcesOptionsModel.EndpointID = core.StringPtr("testString")
				updateEndpointSourcesOptionsModel.Sources = []satellitelinkv1.SourceStatusRequestObject{*sourceStatusRequestObjectModel}
				updateEndpointSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.UpdateEndpointSources(updateEndpointSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.UpdateEndpointSources(updateEndpointSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateEndpointSources(updateEndpointSourcesOptions *UpdateEndpointSourcesOptions)`, func() {
		updateEndpointSourcesPath := "/v1/locations/testString/endpoints/testString/sources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEndpointSourcesPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}]}`)
				}))
			})
			It(`Invoke UpdateEndpointSources successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the SourceStatusRequestObject model
				sourceStatusRequestObjectModel := new(satellitelinkv1.SourceStatusRequestObject)
				sourceStatusRequestObjectModel.SourceID = core.StringPtr("us-south--K9kQEVFmqNpP-Source-Q87fe")
				sourceStatusRequestObjectModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateEndpointSourcesOptions model
				updateEndpointSourcesOptionsModel := new(satellitelinkv1.UpdateEndpointSourcesOptions)
				updateEndpointSourcesOptionsModel.LocationID = core.StringPtr("testString")
				updateEndpointSourcesOptionsModel.EndpointID = core.StringPtr("testString")
				updateEndpointSourcesOptionsModel.Sources = []satellitelinkv1.SourceStatusRequestObject{*sourceStatusRequestObjectModel}
				updateEndpointSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.UpdateEndpointSourcesWithContext(ctx, updateEndpointSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.UpdateEndpointSources(updateEndpointSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.UpdateEndpointSourcesWithContext(ctx, updateEndpointSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateEndpointSourcesPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"sources": [{"source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "enabled": true, "last_change": "2019-11-27T09:07:27.245Z", "pending": false}]}`)
				}))
			})
			It(`Invoke UpdateEndpointSources successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.UpdateEndpointSources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SourceStatusRequestObject model
				sourceStatusRequestObjectModel := new(satellitelinkv1.SourceStatusRequestObject)
				sourceStatusRequestObjectModel.SourceID = core.StringPtr("us-south--K9kQEVFmqNpP-Source-Q87fe")
				sourceStatusRequestObjectModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateEndpointSourcesOptions model
				updateEndpointSourcesOptionsModel := new(satellitelinkv1.UpdateEndpointSourcesOptions)
				updateEndpointSourcesOptionsModel.LocationID = core.StringPtr("testString")
				updateEndpointSourcesOptionsModel.EndpointID = core.StringPtr("testString")
				updateEndpointSourcesOptionsModel.Sources = []satellitelinkv1.SourceStatusRequestObject{*sourceStatusRequestObjectModel}
				updateEndpointSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.UpdateEndpointSources(updateEndpointSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateEndpointSources with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the SourceStatusRequestObject model
				sourceStatusRequestObjectModel := new(satellitelinkv1.SourceStatusRequestObject)
				sourceStatusRequestObjectModel.SourceID = core.StringPtr("us-south--K9kQEVFmqNpP-Source-Q87fe")
				sourceStatusRequestObjectModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateEndpointSourcesOptions model
				updateEndpointSourcesOptionsModel := new(satellitelinkv1.UpdateEndpointSourcesOptions)
				updateEndpointSourcesOptionsModel.LocationID = core.StringPtr("testString")
				updateEndpointSourcesOptionsModel.EndpointID = core.StringPtr("testString")
				updateEndpointSourcesOptionsModel.Sources = []satellitelinkv1.SourceStatusRequestObject{*sourceStatusRequestObjectModel}
				updateEndpointSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.UpdateEndpointSources(updateEndpointSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateEndpointSourcesOptions model with no property values
				updateEndpointSourcesOptionsModelNew := new(satellitelinkv1.UpdateEndpointSourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.UpdateEndpointSources(updateEndpointSourcesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(satelliteLinkService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(satelliteLinkService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
				URL: "https://satellitelinkv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(satelliteLinkService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SATELLITE_LINK_URL":       "https://satellitelinkv1/api",
				"SATELLITE_LINK_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{})
				Expect(satelliteLinkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := satelliteLinkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != satelliteLinkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(satelliteLinkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(satelliteLinkService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{
					URL: "https://testService/api",
				})
				Expect(satelliteLinkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := satelliteLinkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != satelliteLinkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(satelliteLinkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(satelliteLinkService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{})
				err := satelliteLinkService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := satelliteLinkService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != satelliteLinkService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(satelliteLinkService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(satelliteLinkService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SATELLITE_LINK_URL":       "https://satellitelinkv1/api",
				"SATELLITE_LINK_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(satelliteLinkService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SATELLITE_LINK_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1UsingExternalConfig(&satellitelinkv1.SatelliteLinkV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(satelliteLinkService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = satellitelinkv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListSources(listSourcesOptions *ListSourcesOptions) - Operation response error`, func() {
		listSourcesPath := "/v1/locations/testString/sources"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSourcesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"user"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSources with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the ListSourcesOptions model
				listSourcesOptionsModel := new(satellitelinkv1.ListSourcesOptions)
				listSourcesOptionsModel.LocationID = core.StringPtr("testString")
				listSourcesOptionsModel.Type = core.StringPtr("user")
				listSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.ListSources(listSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.ListSources(listSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListSources(listSourcesOptions *ListSourcesOptions)`, func() {
		listSourcesPath := "/v1/locations/testString/sources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSourcesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"user"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"sources": [{"type": "user", "source_name": "DB2", "addresses": ["192.168.20.1/24"], "source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "location_id": "us-south--K9kQEVFmqNpP", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z"}]}`)
				}))
			})
			It(`Invoke ListSources successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListSourcesOptions model
				listSourcesOptionsModel := new(satellitelinkv1.ListSourcesOptions)
				listSourcesOptionsModel.LocationID = core.StringPtr("testString")
				listSourcesOptionsModel.Type = core.StringPtr("user")
				listSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.ListSourcesWithContext(ctx, listSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.ListSources(listSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.ListSourcesWithContext(ctx, listSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSourcesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"user"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"sources": [{"type": "user", "source_name": "DB2", "addresses": ["192.168.20.1/24"], "source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "location_id": "us-south--K9kQEVFmqNpP", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z"}]}`)
				}))
			})
			It(`Invoke ListSources successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.ListSources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSourcesOptions model
				listSourcesOptionsModel := new(satellitelinkv1.ListSourcesOptions)
				listSourcesOptionsModel.LocationID = core.StringPtr("testString")
				listSourcesOptionsModel.Type = core.StringPtr("user")
				listSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.ListSources(listSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSources with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the ListSourcesOptions model
				listSourcesOptionsModel := new(satellitelinkv1.ListSourcesOptions)
				listSourcesOptionsModel.LocationID = core.StringPtr("testString")
				listSourcesOptionsModel.Type = core.StringPtr("user")
				listSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.ListSources(listSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListSourcesOptions model with no property values
				listSourcesOptionsModelNew := new(satellitelinkv1.ListSourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.ListSources(listSourcesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSources(createSourcesOptions *CreateSourcesOptions) - Operation response error`, func() {
		createSourcesPath := "/v1/locations/testString/sources"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSourcesPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSources with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the CreateSourcesOptions model
				createSourcesOptionsModel := new(satellitelinkv1.CreateSourcesOptions)
				createSourcesOptionsModel.LocationID = core.StringPtr("testString")
				createSourcesOptionsModel.Type = core.StringPtr("user")
				createSourcesOptionsModel.SourceName = core.StringPtr("DB2")
				createSourcesOptionsModel.Addresses = []string{"192.168.20.1/24"}
				createSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.CreateSources(createSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.CreateSources(createSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateSources(createSourcesOptions *CreateSourcesOptions)`, func() {
		createSourcesPath := "/v1/locations/testString/sources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSourcesPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "user", "source_name": "DB2", "addresses": ["192.168.20.1/24"], "source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "location_id": "us-south--K9kQEVFmqNpP", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z"}`)
				}))
			})
			It(`Invoke CreateSources successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the CreateSourcesOptions model
				createSourcesOptionsModel := new(satellitelinkv1.CreateSourcesOptions)
				createSourcesOptionsModel.LocationID = core.StringPtr("testString")
				createSourcesOptionsModel.Type = core.StringPtr("user")
				createSourcesOptionsModel.SourceName = core.StringPtr("DB2")
				createSourcesOptionsModel.Addresses = []string{"192.168.20.1/24"}
				createSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.CreateSourcesWithContext(ctx, createSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.CreateSources(createSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.CreateSourcesWithContext(ctx, createSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSourcesPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "user", "source_name": "DB2", "addresses": ["192.168.20.1/24"], "source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "location_id": "us-south--K9kQEVFmqNpP", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z"}`)
				}))
			})
			It(`Invoke CreateSources successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.CreateSources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSourcesOptions model
				createSourcesOptionsModel := new(satellitelinkv1.CreateSourcesOptions)
				createSourcesOptionsModel.LocationID = core.StringPtr("testString")
				createSourcesOptionsModel.Type = core.StringPtr("user")
				createSourcesOptionsModel.SourceName = core.StringPtr("DB2")
				createSourcesOptionsModel.Addresses = []string{"192.168.20.1/24"}
				createSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.CreateSources(createSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSources with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the CreateSourcesOptions model
				createSourcesOptionsModel := new(satellitelinkv1.CreateSourcesOptions)
				createSourcesOptionsModel.LocationID = core.StringPtr("testString")
				createSourcesOptionsModel.Type = core.StringPtr("user")
				createSourcesOptionsModel.SourceName = core.StringPtr("DB2")
				createSourcesOptionsModel.Addresses = []string{"192.168.20.1/24"}
				createSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.CreateSources(createSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSourcesOptions model with no property values
				createSourcesOptionsModelNew := new(satellitelinkv1.CreateSourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.CreateSources(createSourcesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSources(updateSourcesOptions *UpdateSourcesOptions) - Operation response error`, func() {
		updateSourcesPath := "/v1/locations/testString/sources/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSourcesPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSources with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the UpdateSourcesOptions model
				updateSourcesOptionsModel := new(satellitelinkv1.UpdateSourcesOptions)
				updateSourcesOptionsModel.LocationID = core.StringPtr("testString")
				updateSourcesOptionsModel.SourceID = core.StringPtr("testString")
				updateSourcesOptionsModel.SourceName = core.StringPtr("DB2")
				updateSourcesOptionsModel.Addresses = []string{"192.168.20.1/24"}
				updateSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.UpdateSources(updateSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.UpdateSources(updateSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateSources(updateSourcesOptions *UpdateSourcesOptions)`, func() {
		updateSourcesPath := "/v1/locations/testString/sources/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSourcesPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "user", "source_name": "DB2", "addresses": ["192.168.20.1/24"], "source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "location_id": "us-south--K9kQEVFmqNpP", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z"}`)
				}))
			})
			It(`Invoke UpdateSources successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the UpdateSourcesOptions model
				updateSourcesOptionsModel := new(satellitelinkv1.UpdateSourcesOptions)
				updateSourcesOptionsModel.LocationID = core.StringPtr("testString")
				updateSourcesOptionsModel.SourceID = core.StringPtr("testString")
				updateSourcesOptionsModel.SourceName = core.StringPtr("DB2")
				updateSourcesOptionsModel.Addresses = []string{"192.168.20.1/24"}
				updateSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.UpdateSourcesWithContext(ctx, updateSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.UpdateSources(updateSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.UpdateSourcesWithContext(ctx, updateSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSourcesPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "user", "source_name": "DB2", "addresses": ["192.168.20.1/24"], "source_id": "us-south--K9kQEVFmqNpP-Source-Q87fe", "location_id": "us-south--K9kQEVFmqNpP", "created_at": "2019-11-27T09:07:27.245Z", "last_change": "2019-11-27T09:07:27.245Z"}`)
				}))
			})
			It(`Invoke UpdateSources successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.UpdateSources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateSourcesOptions model
				updateSourcesOptionsModel := new(satellitelinkv1.UpdateSourcesOptions)
				updateSourcesOptionsModel.LocationID = core.StringPtr("testString")
				updateSourcesOptionsModel.SourceID = core.StringPtr("testString")
				updateSourcesOptionsModel.SourceName = core.StringPtr("DB2")
				updateSourcesOptionsModel.Addresses = []string{"192.168.20.1/24"}
				updateSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.UpdateSources(updateSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSources with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the UpdateSourcesOptions model
				updateSourcesOptionsModel := new(satellitelinkv1.UpdateSourcesOptions)
				updateSourcesOptionsModel.LocationID = core.StringPtr("testString")
				updateSourcesOptionsModel.SourceID = core.StringPtr("testString")
				updateSourcesOptionsModel.SourceName = core.StringPtr("DB2")
				updateSourcesOptionsModel.Addresses = []string{"192.168.20.1/24"}
				updateSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.UpdateSources(updateSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSourcesOptions model with no property values
				updateSourcesOptionsModelNew := new(satellitelinkv1.UpdateSourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.UpdateSources(updateSourcesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSources(deleteSourcesOptions *DeleteSourcesOptions) - Operation response error`, func() {
		deleteSourcesPath := "/v1/locations/testString/sources/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSourcesPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteSources with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteSourcesOptions model
				deleteSourcesOptionsModel := new(satellitelinkv1.DeleteSourcesOptions)
				deleteSourcesOptionsModel.LocationID = core.StringPtr("testString")
				deleteSourcesOptionsModel.SourceID = core.StringPtr("testString")
				deleteSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.DeleteSources(deleteSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.DeleteSources(deleteSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteSources(deleteSourcesOptions *DeleteSourcesOptions)`, func() {
		deleteSourcesPath := "/v1/locations/testString/sources/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSourcesPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "Success"}`)
				}))
			})
			It(`Invoke DeleteSources successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the DeleteSourcesOptions model
				deleteSourcesOptionsModel := new(satellitelinkv1.DeleteSourcesOptions)
				deleteSourcesOptionsModel.LocationID = core.StringPtr("testString")
				deleteSourcesOptionsModel.SourceID = core.StringPtr("testString")
				deleteSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.DeleteSourcesWithContext(ctx, deleteSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.DeleteSources(deleteSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.DeleteSourcesWithContext(ctx, deleteSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSourcesPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "Success"}`)
				}))
			})
			It(`Invoke DeleteSources successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.DeleteSources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteSourcesOptions model
				deleteSourcesOptionsModel := new(satellitelinkv1.DeleteSourcesOptions)
				deleteSourcesOptionsModel.LocationID = core.StringPtr("testString")
				deleteSourcesOptionsModel.SourceID = core.StringPtr("testString")
				deleteSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.DeleteSources(deleteSourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteSources with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the DeleteSourcesOptions model
				deleteSourcesOptionsModel := new(satellitelinkv1.DeleteSourcesOptions)
				deleteSourcesOptionsModel.LocationID = core.StringPtr("testString")
				deleteSourcesOptionsModel.SourceID = core.StringPtr("testString")
				deleteSourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.DeleteSources(deleteSourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteSourcesOptions model with no property values
				deleteSourcesOptionsModelNew := new(satellitelinkv1.DeleteSourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.DeleteSources(deleteSourcesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSourceEndpoints(listSourceEndpointsOptions *ListSourceEndpointsOptions) - Operation response error`, func() {
		listSourceEndpointsPath := "/v1/locations/testString/sources/testString/endpoints"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSourceEndpointsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSourceEndpoints with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the ListSourceEndpointsOptions model
				listSourceEndpointsOptionsModel := new(satellitelinkv1.ListSourceEndpointsOptions)
				listSourceEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				listSourceEndpointsOptionsModel.SourceID = core.StringPtr("testString")
				listSourceEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.ListSourceEndpoints(listSourceEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.ListSourceEndpoints(listSourceEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListSourceEndpoints(listSourceEndpointsOptions *ListSourceEndpointsOptions)`, func() {
		listSourceEndpointsPath := "/v1/locations/testString/sources/testString/endpoints"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSourceEndpointsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"endpoints": [{"endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "enabled": true}]}`)
				}))
			})
			It(`Invoke ListSourceEndpoints successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the ListSourceEndpointsOptions model
				listSourceEndpointsOptionsModel := new(satellitelinkv1.ListSourceEndpointsOptions)
				listSourceEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				listSourceEndpointsOptionsModel.SourceID = core.StringPtr("testString")
				listSourceEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.ListSourceEndpointsWithContext(ctx, listSourceEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.ListSourceEndpoints(listSourceEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.ListSourceEndpointsWithContext(ctx, listSourceEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSourceEndpointsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"endpoints": [{"endpoint_id": "us-south--K9kQEVFmqNpP_Q87fe", "enabled": true}]}`)
				}))
			})
			It(`Invoke ListSourceEndpoints successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.ListSourceEndpoints(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSourceEndpointsOptions model
				listSourceEndpointsOptionsModel := new(satellitelinkv1.ListSourceEndpointsOptions)
				listSourceEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				listSourceEndpointsOptionsModel.SourceID = core.StringPtr("testString")
				listSourceEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.ListSourceEndpoints(listSourceEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSourceEndpoints with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the ListSourceEndpointsOptions model
				listSourceEndpointsOptionsModel := new(satellitelinkv1.ListSourceEndpointsOptions)
				listSourceEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				listSourceEndpointsOptionsModel.SourceID = core.StringPtr("testString")
				listSourceEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.ListSourceEndpoints(listSourceEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListSourceEndpointsOptions model with no property values
				listSourceEndpointsOptionsModelNew := new(satellitelinkv1.ListSourceEndpointsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.ListSourceEndpoints(listSourceEndpointsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSourceEndpoints(updateSourceEndpointsOptions *UpdateSourceEndpointsOptions) - Operation response error`, func() {
		updateSourceEndpointsPath := "/v1/locations/testString/sources/testString/endpoints"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSourceEndpointsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSourceEndpoints with error: Operation response processing error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the EndpointSourceStatusEndpointsItem model
				endpointSourceStatusEndpointsItemModel := new(satellitelinkv1.EndpointSourceStatusEndpointsItem)
				endpointSourceStatusEndpointsItemModel.EndpointID = core.StringPtr("us-south--K9kQEVFmqNpP_Q87fe")
				endpointSourceStatusEndpointsItemModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateSourceEndpointsOptions model
				updateSourceEndpointsOptionsModel := new(satellitelinkv1.UpdateSourceEndpointsOptions)
				updateSourceEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				updateSourceEndpointsOptionsModel.SourceID = core.StringPtr("testString")
				updateSourceEndpointsOptionsModel.Endpoints = []satellitelinkv1.EndpointSourceStatusEndpointsItem{*endpointSourceStatusEndpointsItemModel}
				updateSourceEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := satelliteLinkService.UpdateSourceEndpoints(updateSourceEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				satelliteLinkService.EnableRetries(0, 0)
				result, response, operationErr = satelliteLinkService.UpdateSourceEndpoints(updateSourceEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateSourceEndpoints(updateSourceEndpointsOptions *UpdateSourceEndpointsOptions)`, func() {
		updateSourceEndpointsPath := "/v1/locations/testString/sources/testString/endpoints"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSourceEndpointsPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "Success"}`)
				}))
			})
			It(`Invoke UpdateSourceEndpoints successfully with retries`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())
				satelliteLinkService.EnableRetries(0, 0)

				// Construct an instance of the EndpointSourceStatusEndpointsItem model
				endpointSourceStatusEndpointsItemModel := new(satellitelinkv1.EndpointSourceStatusEndpointsItem)
				endpointSourceStatusEndpointsItemModel.EndpointID = core.StringPtr("us-south--K9kQEVFmqNpP_Q87fe")
				endpointSourceStatusEndpointsItemModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateSourceEndpointsOptions model
				updateSourceEndpointsOptionsModel := new(satellitelinkv1.UpdateSourceEndpointsOptions)
				updateSourceEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				updateSourceEndpointsOptionsModel.SourceID = core.StringPtr("testString")
				updateSourceEndpointsOptionsModel.Endpoints = []satellitelinkv1.EndpointSourceStatusEndpointsItem{*endpointSourceStatusEndpointsItemModel}
				updateSourceEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := satelliteLinkService.UpdateSourceEndpointsWithContext(ctx, updateSourceEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				satelliteLinkService.DisableRetries()
				result, response, operationErr := satelliteLinkService.UpdateSourceEndpoints(updateSourceEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = satelliteLinkService.UpdateSourceEndpointsWithContext(ctx, updateSourceEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSourceEndpointsPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "Success"}`)
				}))
			})
			It(`Invoke UpdateSourceEndpoints successfully`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := satelliteLinkService.UpdateSourceEndpoints(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EndpointSourceStatusEndpointsItem model
				endpointSourceStatusEndpointsItemModel := new(satellitelinkv1.EndpointSourceStatusEndpointsItem)
				endpointSourceStatusEndpointsItemModel.EndpointID = core.StringPtr("us-south--K9kQEVFmqNpP_Q87fe")
				endpointSourceStatusEndpointsItemModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateSourceEndpointsOptions model
				updateSourceEndpointsOptionsModel := new(satellitelinkv1.UpdateSourceEndpointsOptions)
				updateSourceEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				updateSourceEndpointsOptionsModel.SourceID = core.StringPtr("testString")
				updateSourceEndpointsOptionsModel.Endpoints = []satellitelinkv1.EndpointSourceStatusEndpointsItem{*endpointSourceStatusEndpointsItemModel}
				updateSourceEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = satelliteLinkService.UpdateSourceEndpoints(updateSourceEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSourceEndpoints with error: Operation validation and request error`, func() {
				satelliteLinkService, serviceErr := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(satelliteLinkService).ToNot(BeNil())

				// Construct an instance of the EndpointSourceStatusEndpointsItem model
				endpointSourceStatusEndpointsItemModel := new(satellitelinkv1.EndpointSourceStatusEndpointsItem)
				endpointSourceStatusEndpointsItemModel.EndpointID = core.StringPtr("us-south--K9kQEVFmqNpP_Q87fe")
				endpointSourceStatusEndpointsItemModel.Enabled = core.BoolPtr(true)

				// Construct an instance of the UpdateSourceEndpointsOptions model
				updateSourceEndpointsOptionsModel := new(satellitelinkv1.UpdateSourceEndpointsOptions)
				updateSourceEndpointsOptionsModel.LocationID = core.StringPtr("testString")
				updateSourceEndpointsOptionsModel.SourceID = core.StringPtr("testString")
				updateSourceEndpointsOptionsModel.Endpoints = []satellitelinkv1.EndpointSourceStatusEndpointsItem{*endpointSourceStatusEndpointsItemModel}
				updateSourceEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := satelliteLinkService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := satelliteLinkService.UpdateSourceEndpoints(updateSourceEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSourceEndpointsOptions model with no property values
				updateSourceEndpointsOptionsModelNew := new(satellitelinkv1.UpdateSourceEndpointsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = satelliteLinkService.UpdateSourceEndpoints(updateSourceEndpointsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			satelliteLinkService, _ := satellitelinkv1.NewSatelliteLinkV1(&satellitelinkv1.SatelliteLinkV1Options{
				URL:           "http://satellitelinkv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateEndpointsOptions successfully`, func() {
				// Construct an instance of the AdditionalNewEndpointRequestCertsClientCert model
				additionalNewEndpointRequestCertsClientCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsClientCert)
				Expect(additionalNewEndpointRequestCertsClientCertModel).ToNot(BeNil())
				additionalNewEndpointRequestCertsClientCertModel.Filename = core.StringPtr("clientEndpointCert.pem")
				additionalNewEndpointRequestCertsClientCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")
				Expect(additionalNewEndpointRequestCertsClientCertModel.Filename).To(Equal(core.StringPtr("clientEndpointCert.pem")))
				Expect(additionalNewEndpointRequestCertsClientCertModel.FileContents).To(Equal(core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")))

				// Construct an instance of the AdditionalNewEndpointRequestCertsClient model
				additionalNewEndpointRequestCertsClientModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsClient)
				Expect(additionalNewEndpointRequestCertsClientModel).ToNot(BeNil())
				additionalNewEndpointRequestCertsClientModel.Cert = additionalNewEndpointRequestCertsClientCertModel
				Expect(additionalNewEndpointRequestCertsClientModel.Cert).To(Equal(additionalNewEndpointRequestCertsClientCertModel))

				// Construct an instance of the AdditionalNewEndpointRequestCertsServerCert model
				additionalNewEndpointRequestCertsServerCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsServerCert)
				Expect(additionalNewEndpointRequestCertsServerCertModel).ToNot(BeNil())
				additionalNewEndpointRequestCertsServerCertModel.Filename = core.StringPtr("serverEndpointCert.pem")
				additionalNewEndpointRequestCertsServerCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")
				Expect(additionalNewEndpointRequestCertsServerCertModel.Filename).To(Equal(core.StringPtr("serverEndpointCert.pem")))
				Expect(additionalNewEndpointRequestCertsServerCertModel.FileContents).To(Equal(core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")))

				// Construct an instance of the AdditionalNewEndpointRequestCertsServer model
				additionalNewEndpointRequestCertsServerModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsServer)
				Expect(additionalNewEndpointRequestCertsServerModel).ToNot(BeNil())
				additionalNewEndpointRequestCertsServerModel.Cert = additionalNewEndpointRequestCertsServerCertModel
				Expect(additionalNewEndpointRequestCertsServerModel.Cert).To(Equal(additionalNewEndpointRequestCertsServerCertModel))

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnectorCert model
				additionalNewEndpointRequestCertsConnectorCertModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnectorCert)
				Expect(additionalNewEndpointRequestCertsConnectorCertModel).ToNot(BeNil())
				additionalNewEndpointRequestCertsConnectorCertModel.Filename = core.StringPtr("ConnectorCert.pem")
				additionalNewEndpointRequestCertsConnectorCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")
				Expect(additionalNewEndpointRequestCertsConnectorCertModel.Filename).To(Equal(core.StringPtr("ConnectorCert.pem")))
				Expect(additionalNewEndpointRequestCertsConnectorCertModel.FileContents).To(Equal(core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")))

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnectorKey model
				additionalNewEndpointRequestCertsConnectorKeyModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnectorKey)
				Expect(additionalNewEndpointRequestCertsConnectorKeyModel).ToNot(BeNil())
				additionalNewEndpointRequestCertsConnectorKeyModel.Filename = core.StringPtr("ConnectorPrivateKey.pem")
				additionalNewEndpointRequestCertsConnectorKeyModel.FileContents = core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n")
				Expect(additionalNewEndpointRequestCertsConnectorKeyModel.Filename).To(Equal(core.StringPtr("ConnectorPrivateKey.pem")))
				Expect(additionalNewEndpointRequestCertsConnectorKeyModel.FileContents).To(Equal(core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n")))

				// Construct an instance of the AdditionalNewEndpointRequestCertsConnector model
				additionalNewEndpointRequestCertsConnectorModel := new(satellitelinkv1.AdditionalNewEndpointRequestCertsConnector)
				Expect(additionalNewEndpointRequestCertsConnectorModel).ToNot(BeNil())
				additionalNewEndpointRequestCertsConnectorModel.Cert = additionalNewEndpointRequestCertsConnectorCertModel
				additionalNewEndpointRequestCertsConnectorModel.Key = additionalNewEndpointRequestCertsConnectorKeyModel
				Expect(additionalNewEndpointRequestCertsConnectorModel.Cert).To(Equal(additionalNewEndpointRequestCertsConnectorCertModel))
				Expect(additionalNewEndpointRequestCertsConnectorModel.Key).To(Equal(additionalNewEndpointRequestCertsConnectorKeyModel))

				// Construct an instance of the AdditionalNewEndpointRequestCerts model
				additionalNewEndpointRequestCertsModel := new(satellitelinkv1.AdditionalNewEndpointRequestCerts)
				Expect(additionalNewEndpointRequestCertsModel).ToNot(BeNil())
				additionalNewEndpointRequestCertsModel.Client = additionalNewEndpointRequestCertsClientModel
				additionalNewEndpointRequestCertsModel.Server = additionalNewEndpointRequestCertsServerModel
				additionalNewEndpointRequestCertsModel.Connector = additionalNewEndpointRequestCertsConnectorModel
				Expect(additionalNewEndpointRequestCertsModel.Client).To(Equal(additionalNewEndpointRequestCertsClientModel))
				Expect(additionalNewEndpointRequestCertsModel.Server).To(Equal(additionalNewEndpointRequestCertsServerModel))
				Expect(additionalNewEndpointRequestCertsModel.Connector).To(Equal(additionalNewEndpointRequestCertsConnectorModel))

				// Construct an instance of the CreateEndpointsOptions model
				locationID := "testString"
				createEndpointsOptionsModel := satelliteLinkService.NewCreateEndpointsOptions(locationID)
				createEndpointsOptionsModel.SetLocationID("testString")
				createEndpointsOptionsModel.SetConnType("cloud")
				createEndpointsOptionsModel.SetDisplayName("My endpoint")
				createEndpointsOptionsModel.SetServerHost("example.com")
				createEndpointsOptionsModel.SetServerPort(int64(443))
				createEndpointsOptionsModel.SetSni("example.com")
				createEndpointsOptionsModel.SetClientProtocol("https")
				createEndpointsOptionsModel.SetClientMutualAuth(true)
				createEndpointsOptionsModel.SetServerProtocol("tls")
				createEndpointsOptionsModel.SetServerMutualAuth(true)
				createEndpointsOptionsModel.SetRejectUnauth(true)
				createEndpointsOptionsModel.SetTimeout(int64(60))
				createEndpointsOptionsModel.SetCreatedBy("My service")
				createEndpointsOptionsModel.SetCerts(additionalNewEndpointRequestCertsModel)
				createEndpointsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createEndpointsOptionsModel).ToNot(BeNil())
				Expect(createEndpointsOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(createEndpointsOptionsModel.ConnType).To(Equal(core.StringPtr("cloud")))
				Expect(createEndpointsOptionsModel.DisplayName).To(Equal(core.StringPtr("My endpoint")))
				Expect(createEndpointsOptionsModel.ServerHost).To(Equal(core.StringPtr("example.com")))
				Expect(createEndpointsOptionsModel.ServerPort).To(Equal(core.Int64Ptr(int64(443))))
				Expect(createEndpointsOptionsModel.Sni).To(Equal(core.StringPtr("example.com")))
				Expect(createEndpointsOptionsModel.ClientProtocol).To(Equal(core.StringPtr("https")))
				Expect(createEndpointsOptionsModel.ClientMutualAuth).To(Equal(core.BoolPtr(true)))
				Expect(createEndpointsOptionsModel.ServerProtocol).To(Equal(core.StringPtr("tls")))
				Expect(createEndpointsOptionsModel.ServerMutualAuth).To(Equal(core.BoolPtr(true)))
				Expect(createEndpointsOptionsModel.RejectUnauth).To(Equal(core.BoolPtr(true)))
				Expect(createEndpointsOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(60))))
				Expect(createEndpointsOptionsModel.CreatedBy).To(Equal(core.StringPtr("My service")))
				Expect(createEndpointsOptionsModel.Certs).To(Equal(additionalNewEndpointRequestCertsModel))
				Expect(createEndpointsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateLinkOptions successfully`, func() {
				// Construct an instance of the CreateLinkOptions model
				createLinkOptionsModel := satelliteLinkService.NewCreateLinkOptions()
				createLinkOptionsModel.SetCrn("crn:v1:staging:public:satellite:us-south:a/1ae4eb57181a46ceade4846519678888::location:brbats7009sqna3dtest")
				createLinkOptionsModel.SetLocationID("brbats7009sqna3dtest")
				createLinkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLinkOptionsModel).ToNot(BeNil())
				Expect(createLinkOptionsModel.Crn).To(Equal(core.StringPtr("crn:v1:staging:public:satellite:us-south:a/1ae4eb57181a46ceade4846519678888::location:brbats7009sqna3dtest")))
				Expect(createLinkOptionsModel.LocationID).To(Equal(core.StringPtr("brbats7009sqna3dtest")))
				Expect(createLinkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSourcesOptions successfully`, func() {
				// Construct an instance of the CreateSourcesOptions model
				locationID := "testString"
				createSourcesOptionsModel := satelliteLinkService.NewCreateSourcesOptions(locationID)
				createSourcesOptionsModel.SetLocationID("testString")
				createSourcesOptionsModel.SetType("user")
				createSourcesOptionsModel.SetSourceName("DB2")
				createSourcesOptionsModel.SetAddresses([]string{"192.168.20.1/24"})
				createSourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSourcesOptionsModel).ToNot(BeNil())
				Expect(createSourcesOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(createSourcesOptionsModel.Type).To(Equal(core.StringPtr("user")))
				Expect(createSourcesOptionsModel.SourceName).To(Equal(core.StringPtr("DB2")))
				Expect(createSourcesOptionsModel.Addresses).To(Equal([]string{"192.168.20.1/24"}))
				Expect(createSourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEndpointCertsOptions successfully`, func() {
				// Construct an instance of the DeleteEndpointCertsOptions model
				locationID := "testString"
				endpointID := "testString"
				deleteEndpointCertsOptionsModel := satelliteLinkService.NewDeleteEndpointCertsOptions(locationID, endpointID)
				deleteEndpointCertsOptionsModel.SetLocationID("testString")
				deleteEndpointCertsOptionsModel.SetEndpointID("testString")
				deleteEndpointCertsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEndpointCertsOptionsModel).ToNot(BeNil())
				Expect(deleteEndpointCertsOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEndpointCertsOptionsModel.EndpointID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEndpointCertsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteEndpointsOptions successfully`, func() {
				// Construct an instance of the DeleteEndpointsOptions model
				locationID := "testString"
				endpointID := "testString"
				deleteEndpointsOptionsModel := satelliteLinkService.NewDeleteEndpointsOptions(locationID, endpointID)
				deleteEndpointsOptionsModel.SetLocationID("testString")
				deleteEndpointsOptionsModel.SetEndpointID("testString")
				deleteEndpointsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteEndpointsOptionsModel).ToNot(BeNil())
				Expect(deleteEndpointsOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEndpointsOptionsModel.EndpointID).To(Equal(core.StringPtr("testString")))
				Expect(deleteEndpointsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLinkOptions successfully`, func() {
				// Construct an instance of the DeleteLinkOptions model
				locationID := "testString"
				deleteLinkOptionsModel := satelliteLinkService.NewDeleteLinkOptions(locationID)
				deleteLinkOptionsModel.SetLocationID("testString")
				deleteLinkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLinkOptionsModel).ToNot(BeNil())
				Expect(deleteLinkOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteLinkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSourcesOptions successfully`, func() {
				// Construct an instance of the DeleteSourcesOptions model
				locationID := "testString"
				sourceID := "testString"
				deleteSourcesOptionsModel := satelliteLinkService.NewDeleteSourcesOptions(locationID, sourceID)
				deleteSourcesOptionsModel.SetLocationID("testString")
				deleteSourcesOptionsModel.SetSourceID("testString")
				deleteSourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSourcesOptionsModel).ToNot(BeNil())
				Expect(deleteSourcesOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSourcesOptionsModel.SourceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewExportEndpointsOptions successfully`, func() {
				// Construct an instance of the ExportEndpointsOptions model
				locationID := "testString"
				endpointID := "testString"
				exportEndpointsOptionsModel := satelliteLinkService.NewExportEndpointsOptions(locationID, endpointID)
				exportEndpointsOptionsModel.SetLocationID("testString")
				exportEndpointsOptionsModel.SetEndpointID("testString")
				exportEndpointsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(exportEndpointsOptionsModel).ToNot(BeNil())
				Expect(exportEndpointsOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(exportEndpointsOptionsModel.EndpointID).To(Equal(core.StringPtr("testString")))
				Expect(exportEndpointsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEndpointCertsOptions successfully`, func() {
				// Construct an instance of the GetEndpointCertsOptions model
				locationID := "testString"
				endpointID := "testString"
				getEndpointCertsOptionsModel := satelliteLinkService.NewGetEndpointCertsOptions(locationID, endpointID)
				getEndpointCertsOptionsModel.SetLocationID("testString")
				getEndpointCertsOptionsModel.SetEndpointID("testString")
				getEndpointCertsOptionsModel.SetNoZip(true)
				getEndpointCertsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEndpointCertsOptionsModel).ToNot(BeNil())
				Expect(getEndpointCertsOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(getEndpointCertsOptionsModel.EndpointID).To(Equal(core.StringPtr("testString")))
				Expect(getEndpointCertsOptionsModel.NoZip).To(Equal(core.BoolPtr(true)))
				Expect(getEndpointCertsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEndpointsOptions successfully`, func() {
				// Construct an instance of the GetEndpointsOptions model
				locationID := "testString"
				endpointID := "testString"
				getEndpointsOptionsModel := satelliteLinkService.NewGetEndpointsOptions(locationID, endpointID)
				getEndpointsOptionsModel.SetLocationID("testString")
				getEndpointsOptionsModel.SetEndpointID("testString")
				getEndpointsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEndpointsOptionsModel).ToNot(BeNil())
				Expect(getEndpointsOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(getEndpointsOptionsModel.EndpointID).To(Equal(core.StringPtr("testString")))
				Expect(getEndpointsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLinkOptions successfully`, func() {
				// Construct an instance of the GetLinkOptions model
				locationID := "testString"
				getLinkOptionsModel := satelliteLinkService.NewGetLinkOptions(locationID)
				getLinkOptionsModel.SetLocationID("testString")
				getLinkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLinkOptionsModel).ToNot(BeNil())
				Expect(getLinkOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(getLinkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImportEndpointsOptions successfully`, func() {
				// Construct an instance of the ImportEndpointsOptions model
				locationID := "testString"
				state := CreateMockReader("This is a mock file.")
				importEndpointsOptionsModel := satelliteLinkService.NewImportEndpointsOptions(locationID, state)
				importEndpointsOptionsModel.SetLocationID("testString")
				importEndpointsOptionsModel.SetState(CreateMockReader("This is a mock file."))
				importEndpointsOptionsModel.SetStateContentType("testString")
				importEndpointsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(importEndpointsOptionsModel).ToNot(BeNil())
				Expect(importEndpointsOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(importEndpointsOptionsModel.State).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(importEndpointsOptionsModel.StateContentType).To(Equal(core.StringPtr("testString")))
				Expect(importEndpointsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEndpointSourcesOptions successfully`, func() {
				// Construct an instance of the ListEndpointSourcesOptions model
				locationID := "testString"
				endpointID := "testString"
				listEndpointSourcesOptionsModel := satelliteLinkService.NewListEndpointSourcesOptions(locationID, endpointID)
				listEndpointSourcesOptionsModel.SetLocationID("testString")
				listEndpointSourcesOptionsModel.SetEndpointID("testString")
				listEndpointSourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEndpointSourcesOptionsModel).ToNot(BeNil())
				Expect(listEndpointSourcesOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(listEndpointSourcesOptionsModel.EndpointID).To(Equal(core.StringPtr("testString")))
				Expect(listEndpointSourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEndpointsOptions successfully`, func() {
				// Construct an instance of the ListEndpointsOptions model
				locationID := "testString"
				listEndpointsOptionsModel := satelliteLinkService.NewListEndpointsOptions(locationID)
				listEndpointsOptionsModel.SetLocationID("testString")
				listEndpointsOptionsModel.SetType("enabled")
				listEndpointsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEndpointsOptionsModel).ToNot(BeNil())
				Expect(listEndpointsOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(listEndpointsOptionsModel.Type).To(Equal(core.StringPtr("enabled")))
				Expect(listEndpointsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSourceEndpointsOptions successfully`, func() {
				// Construct an instance of the ListSourceEndpointsOptions model
				locationID := "testString"
				sourceID := "testString"
				listSourceEndpointsOptionsModel := satelliteLinkService.NewListSourceEndpointsOptions(locationID, sourceID)
				listSourceEndpointsOptionsModel.SetLocationID("testString")
				listSourceEndpointsOptionsModel.SetSourceID("testString")
				listSourceEndpointsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSourceEndpointsOptionsModel).ToNot(BeNil())
				Expect(listSourceEndpointsOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(listSourceEndpointsOptionsModel.SourceID).To(Equal(core.StringPtr("testString")))
				Expect(listSourceEndpointsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSourcesOptions successfully`, func() {
				// Construct an instance of the ListSourcesOptions model
				locationID := "testString"
				listSourcesOptionsModel := satelliteLinkService.NewListSourcesOptions(locationID)
				listSourcesOptionsModel.SetLocationID("testString")
				listSourcesOptionsModel.SetType("user")
				listSourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSourcesOptionsModel).ToNot(BeNil())
				Expect(listSourcesOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(listSourcesOptionsModel.Type).To(Equal(core.StringPtr("user")))
				Expect(listSourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEndpointSourcesOptions successfully`, func() {
				// Construct an instance of the SourceStatusRequestObject model
				sourceStatusRequestObjectModel := new(satellitelinkv1.SourceStatusRequestObject)
				Expect(sourceStatusRequestObjectModel).ToNot(BeNil())
				sourceStatusRequestObjectModel.SourceID = core.StringPtr("us-south--K9kQEVFmqNpP-Source-Q87fe")
				sourceStatusRequestObjectModel.Enabled = core.BoolPtr(true)
				Expect(sourceStatusRequestObjectModel.SourceID).To(Equal(core.StringPtr("us-south--K9kQEVFmqNpP-Source-Q87fe")))
				Expect(sourceStatusRequestObjectModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the UpdateEndpointSourcesOptions model
				locationID := "testString"
				endpointID := "testString"
				updateEndpointSourcesOptionsModel := satelliteLinkService.NewUpdateEndpointSourcesOptions(locationID, endpointID)
				updateEndpointSourcesOptionsModel.SetLocationID("testString")
				updateEndpointSourcesOptionsModel.SetEndpointID("testString")
				updateEndpointSourcesOptionsModel.SetSources([]satellitelinkv1.SourceStatusRequestObject{*sourceStatusRequestObjectModel})
				updateEndpointSourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEndpointSourcesOptionsModel).ToNot(BeNil())
				Expect(updateEndpointSourcesOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(updateEndpointSourcesOptionsModel.EndpointID).To(Equal(core.StringPtr("testString")))
				Expect(updateEndpointSourcesOptionsModel.Sources).To(Equal([]satellitelinkv1.SourceStatusRequestObject{*sourceStatusRequestObjectModel}))
				Expect(updateEndpointSourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateEndpointsOptions successfully`, func() {
				// Construct an instance of the UpdatedEndpointRequestCertsClientCert model
				updatedEndpointRequestCertsClientCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsClientCert)
				Expect(updatedEndpointRequestCertsClientCertModel).ToNot(BeNil())
				updatedEndpointRequestCertsClientCertModel.Filename = core.StringPtr("clientEndpointCert.pem")
				updatedEndpointRequestCertsClientCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")
				Expect(updatedEndpointRequestCertsClientCertModel.Filename).To(Equal(core.StringPtr("clientEndpointCert.pem")))
				Expect(updatedEndpointRequestCertsClientCertModel.FileContents).To(Equal(core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")))

				// Construct an instance of the UpdatedEndpointRequestCertsClient model
				updatedEndpointRequestCertsClientModel := new(satellitelinkv1.UpdatedEndpointRequestCertsClient)
				Expect(updatedEndpointRequestCertsClientModel).ToNot(BeNil())
				updatedEndpointRequestCertsClientModel.Cert = updatedEndpointRequestCertsClientCertModel
				Expect(updatedEndpointRequestCertsClientModel.Cert).To(Equal(updatedEndpointRequestCertsClientCertModel))

				// Construct an instance of the UpdatedEndpointRequestCertsServerCert model
				updatedEndpointRequestCertsServerCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsServerCert)
				Expect(updatedEndpointRequestCertsServerCertModel).ToNot(BeNil())
				updatedEndpointRequestCertsServerCertModel.Filename = core.StringPtr("serverEndpointCert.pem")
				updatedEndpointRequestCertsServerCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")
				Expect(updatedEndpointRequestCertsServerCertModel.Filename).To(Equal(core.StringPtr("serverEndpointCert.pem")))
				Expect(updatedEndpointRequestCertsServerCertModel.FileContents).To(Equal(core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")))

				// Construct an instance of the UpdatedEndpointRequestCertsServer model
				updatedEndpointRequestCertsServerModel := new(satellitelinkv1.UpdatedEndpointRequestCertsServer)
				Expect(updatedEndpointRequestCertsServerModel).ToNot(BeNil())
				updatedEndpointRequestCertsServerModel.Cert = updatedEndpointRequestCertsServerCertModel
				Expect(updatedEndpointRequestCertsServerModel.Cert).To(Equal(updatedEndpointRequestCertsServerCertModel))

				// Construct an instance of the UpdatedEndpointRequestCertsConnectorCert model
				updatedEndpointRequestCertsConnectorCertModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnectorCert)
				Expect(updatedEndpointRequestCertsConnectorCertModel).ToNot(BeNil())
				updatedEndpointRequestCertsConnectorCertModel.Filename = core.StringPtr("ConnectorCert.pem")
				updatedEndpointRequestCertsConnectorCertModel.FileContents = core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")
				Expect(updatedEndpointRequestCertsConnectorCertModel.Filename).To(Equal(core.StringPtr("ConnectorCert.pem")))
				Expect(updatedEndpointRequestCertsConnectorCertModel.FileContents).To(Equal(core.StringPtr("-----BEGIN CERTIFICATE-----\r\n<the-content-of-the-cert>-----END CERTIFICATE-----\r\n")))

				// Construct an instance of the UpdatedEndpointRequestCertsConnectorKey model
				updatedEndpointRequestCertsConnectorKeyModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnectorKey)
				Expect(updatedEndpointRequestCertsConnectorKeyModel).ToNot(BeNil())
				updatedEndpointRequestCertsConnectorKeyModel.Filename = core.StringPtr("ConnectorPrivateKey.pem")
				updatedEndpointRequestCertsConnectorKeyModel.FileContents = core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n")
				Expect(updatedEndpointRequestCertsConnectorKeyModel.Filename).To(Equal(core.StringPtr("ConnectorPrivateKey.pem")))
				Expect(updatedEndpointRequestCertsConnectorKeyModel.FileContents).To(Equal(core.StringPtr("-----BEGIN PRIVATE KEY-----\n<the-content-of-the-key>\n-----END PRIVATE KEY-----\n")))

				// Construct an instance of the UpdatedEndpointRequestCertsConnector model
				updatedEndpointRequestCertsConnectorModel := new(satellitelinkv1.UpdatedEndpointRequestCertsConnector)
				Expect(updatedEndpointRequestCertsConnectorModel).ToNot(BeNil())
				updatedEndpointRequestCertsConnectorModel.Cert = updatedEndpointRequestCertsConnectorCertModel
				updatedEndpointRequestCertsConnectorModel.Key = updatedEndpointRequestCertsConnectorKeyModel
				Expect(updatedEndpointRequestCertsConnectorModel.Cert).To(Equal(updatedEndpointRequestCertsConnectorCertModel))
				Expect(updatedEndpointRequestCertsConnectorModel.Key).To(Equal(updatedEndpointRequestCertsConnectorKeyModel))

				// Construct an instance of the UpdatedEndpointRequestCerts model
				updatedEndpointRequestCertsModel := new(satellitelinkv1.UpdatedEndpointRequestCerts)
				Expect(updatedEndpointRequestCertsModel).ToNot(BeNil())
				updatedEndpointRequestCertsModel.Client = updatedEndpointRequestCertsClientModel
				updatedEndpointRequestCertsModel.Server = updatedEndpointRequestCertsServerModel
				updatedEndpointRequestCertsModel.Connector = updatedEndpointRequestCertsConnectorModel
				Expect(updatedEndpointRequestCertsModel.Client).To(Equal(updatedEndpointRequestCertsClientModel))
				Expect(updatedEndpointRequestCertsModel.Server).To(Equal(updatedEndpointRequestCertsServerModel))
				Expect(updatedEndpointRequestCertsModel.Connector).To(Equal(updatedEndpointRequestCertsConnectorModel))

				// Construct an instance of the UpdateEndpointsOptions model
				locationID := "testString"
				endpointID := "testString"
				updateEndpointsOptionsModel := satelliteLinkService.NewUpdateEndpointsOptions(locationID, endpointID)
				updateEndpointsOptionsModel.SetLocationID("testString")
				updateEndpointsOptionsModel.SetEndpointID("testString")
				updateEndpointsOptionsModel.SetDisplayName("My endpoint")
				updateEndpointsOptionsModel.SetServerHost("example.com")
				updateEndpointsOptionsModel.SetServerPort(int64(443))
				updateEndpointsOptionsModel.SetSni("example.com")
				updateEndpointsOptionsModel.SetClientProtocol("https")
				updateEndpointsOptionsModel.SetClientMutualAuth(true)
				updateEndpointsOptionsModel.SetServerProtocol("tls")
				updateEndpointsOptionsModel.SetServerMutualAuth(true)
				updateEndpointsOptionsModel.SetRejectUnauth(true)
				updateEndpointsOptionsModel.SetTimeout(int64(60))
				updateEndpointsOptionsModel.SetCreatedBy("My service")
				updateEndpointsOptionsModel.SetCerts(updatedEndpointRequestCertsModel)
				updateEndpointsOptionsModel.SetEnabled(true)
				updateEndpointsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateEndpointsOptionsModel).ToNot(BeNil())
				Expect(updateEndpointsOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(updateEndpointsOptionsModel.EndpointID).To(Equal(core.StringPtr("testString")))
				Expect(updateEndpointsOptionsModel.DisplayName).To(Equal(core.StringPtr("My endpoint")))
				Expect(updateEndpointsOptionsModel.ServerHost).To(Equal(core.StringPtr("example.com")))
				Expect(updateEndpointsOptionsModel.ServerPort).To(Equal(core.Int64Ptr(int64(443))))
				Expect(updateEndpointsOptionsModel.Sni).To(Equal(core.StringPtr("example.com")))
				Expect(updateEndpointsOptionsModel.ClientProtocol).To(Equal(core.StringPtr("https")))
				Expect(updateEndpointsOptionsModel.ClientMutualAuth).To(Equal(core.BoolPtr(true)))
				Expect(updateEndpointsOptionsModel.ServerProtocol).To(Equal(core.StringPtr("tls")))
				Expect(updateEndpointsOptionsModel.ServerMutualAuth).To(Equal(core.BoolPtr(true)))
				Expect(updateEndpointsOptionsModel.RejectUnauth).To(Equal(core.BoolPtr(true)))
				Expect(updateEndpointsOptionsModel.Timeout).To(Equal(core.Int64Ptr(int64(60))))
				Expect(updateEndpointsOptionsModel.CreatedBy).To(Equal(core.StringPtr("My service")))
				Expect(updateEndpointsOptionsModel.Certs).To(Equal(updatedEndpointRequestCertsModel))
				Expect(updateEndpointsOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(updateEndpointsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateLinkOptions successfully`, func() {
				// Construct an instance of the UpdateLinkOptions model
				locationID := "testString"
				updateLinkOptionsModel := satelliteLinkService.NewUpdateLinkOptions(locationID)
				updateLinkOptionsModel.SetLocationID("testString")
				updateLinkOptionsModel.SetWsEndpoint("{satellite-link-tunnel-server}")
				updateLinkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateLinkOptionsModel).ToNot(BeNil())
				Expect(updateLinkOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(updateLinkOptionsModel.WsEndpoint).To(Equal(core.StringPtr("{satellite-link-tunnel-server}")))
				Expect(updateLinkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSourceEndpointsOptions successfully`, func() {
				// Construct an instance of the EndpointSourceStatusEndpointsItem model
				endpointSourceStatusEndpointsItemModel := new(satellitelinkv1.EndpointSourceStatusEndpointsItem)
				Expect(endpointSourceStatusEndpointsItemModel).ToNot(BeNil())
				endpointSourceStatusEndpointsItemModel.EndpointID = core.StringPtr("us-south--K9kQEVFmqNpP_Q87fe")
				endpointSourceStatusEndpointsItemModel.Enabled = core.BoolPtr(true)
				Expect(endpointSourceStatusEndpointsItemModel.EndpointID).To(Equal(core.StringPtr("us-south--K9kQEVFmqNpP_Q87fe")))
				Expect(endpointSourceStatusEndpointsItemModel.Enabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the UpdateSourceEndpointsOptions model
				locationID := "testString"
				sourceID := "testString"
				updateSourceEndpointsOptionsModel := satelliteLinkService.NewUpdateSourceEndpointsOptions(locationID, sourceID)
				updateSourceEndpointsOptionsModel.SetLocationID("testString")
				updateSourceEndpointsOptionsModel.SetSourceID("testString")
				updateSourceEndpointsOptionsModel.SetEndpoints([]satellitelinkv1.EndpointSourceStatusEndpointsItem{*endpointSourceStatusEndpointsItemModel})
				updateSourceEndpointsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSourceEndpointsOptionsModel).ToNot(BeNil())
				Expect(updateSourceEndpointsOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(updateSourceEndpointsOptionsModel.SourceID).To(Equal(core.StringPtr("testString")))
				Expect(updateSourceEndpointsOptionsModel.Endpoints).To(Equal([]satellitelinkv1.EndpointSourceStatusEndpointsItem{*endpointSourceStatusEndpointsItemModel}))
				Expect(updateSourceEndpointsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSourcesOptions successfully`, func() {
				// Construct an instance of the UpdateSourcesOptions model
				locationID := "testString"
				sourceID := "testString"
				updateSourcesOptionsModel := satelliteLinkService.NewUpdateSourcesOptions(locationID, sourceID)
				updateSourcesOptionsModel.SetLocationID("testString")
				updateSourcesOptionsModel.SetSourceID("testString")
				updateSourcesOptionsModel.SetSourceName("DB2")
				updateSourcesOptionsModel.SetAddresses([]string{"192.168.20.1/24"})
				updateSourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSourcesOptionsModel).ToNot(BeNil())
				Expect(updateSourcesOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(updateSourcesOptionsModel.SourceID).To(Equal(core.StringPtr("testString")))
				Expect(updateSourcesOptionsModel.SourceName).To(Equal(core.StringPtr("DB2")))
				Expect(updateSourcesOptionsModel.Addresses).To(Equal([]string{"192.168.20.1/24"}))
				Expect(updateSourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUploadEndpointCertsOptions successfully`, func() {
				// Construct an instance of the UploadEndpointCertsOptions model
				locationID := "testString"
				endpointID := "testString"
				uploadEndpointCertsOptionsModel := satelliteLinkService.NewUploadEndpointCertsOptions(locationID, endpointID)
				uploadEndpointCertsOptionsModel.SetLocationID("testString")
				uploadEndpointCertsOptionsModel.SetEndpointID("testString")
				uploadEndpointCertsOptionsModel.SetClientCert(CreateMockReader("This is a mock file."))
				uploadEndpointCertsOptionsModel.SetClientCertContentType("testString")
				uploadEndpointCertsOptionsModel.SetServerCert(CreateMockReader("This is a mock file."))
				uploadEndpointCertsOptionsModel.SetServerCertContentType("testString")
				uploadEndpointCertsOptionsModel.SetConnectorCert(CreateMockReader("This is a mock file."))
				uploadEndpointCertsOptionsModel.SetConnectorCertContentType("testString")
				uploadEndpointCertsOptionsModel.SetConnectorKey(CreateMockReader("This is a mock file."))
				uploadEndpointCertsOptionsModel.SetConnectorKeyContentType("testString")
				uploadEndpointCertsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uploadEndpointCertsOptionsModel).ToNot(BeNil())
				Expect(uploadEndpointCertsOptionsModel.LocationID).To(Equal(core.StringPtr("testString")))
				Expect(uploadEndpointCertsOptionsModel.EndpointID).To(Equal(core.StringPtr("testString")))
				Expect(uploadEndpointCertsOptionsModel.ClientCert).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(uploadEndpointCertsOptionsModel.ClientCertContentType).To(Equal(core.StringPtr("testString")))
				Expect(uploadEndpointCertsOptionsModel.ServerCert).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(uploadEndpointCertsOptionsModel.ServerCertContentType).To(Equal(core.StringPtr("testString")))
				Expect(uploadEndpointCertsOptionsModel.ConnectorCert).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(uploadEndpointCertsOptionsModel.ConnectorCertContentType).To(Equal(core.StringPtr("testString")))
				Expect(uploadEndpointCertsOptionsModel.ConnectorKey).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(uploadEndpointCertsOptionsModel.ConnectorKeyContentType).To(Equal(core.StringPtr("testString")))
				Expect(uploadEndpointCertsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}

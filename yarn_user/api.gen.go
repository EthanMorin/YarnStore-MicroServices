// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new user
	// (POST /user)
	PostUser(c *gin.Context)
	// Check if the service is running
	// (GET /user/check)
	GetUserCheck(c *gin.Context)
	// Login a user
	// (POST /user/login)
	PostUserLogin(c *gin.Context)
	// Delete user by ID
	// (DELETE /user/{user_id})
	DeleteUserUserId(c *gin.Context, userId string)
	// Get user by ID
	// (GET /user/{user_id})
	GetUserUserId(c *gin.Context, userId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostUser operation middleware
func (siw *ServerInterfaceWrapper) PostUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUser(c)
}

// GetUserCheck operation middleware
func (siw *ServerInterfaceWrapper) GetUserCheck(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserCheck(c)
}

// PostUserLogin operation middleware
func (siw *ServerInterfaceWrapper) PostUserLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUserLogin(c)
}

// DeleteUserUserId operation middleware
func (siw *ServerInterfaceWrapper) DeleteUserUserId(c *gin.Context) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", c.Param("user_id"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteUserUserId(c, userId)
}

// GetUserUserId operation middleware
func (siw *ServerInterfaceWrapper) GetUserUserId(c *gin.Context) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", c.Param("user_id"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserUserId(c, userId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/user", wrapper.PostUser)
	router.GET(options.BaseURL+"/user/check", wrapper.GetUserCheck)
	router.POST(options.BaseURL+"/user/login", wrapper.PostUserLogin)
	router.DELETE(options.BaseURL+"/user/:user_id", wrapper.DeleteUserUserId)
	router.GET(options.BaseURL+"/user/:user_id", wrapper.GetUserUserId)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8SVUWvbMBDHv4q47dGJ061PfttaKIFBC6MPo5SiWBdHXaTTpHPaEPzdh85J2iXuxgrr",
	"HoyF7nT63+9O0gZqcoE8ek5QbSDVC3RahtcJY/6HSAEjW5RZdNou84DXAaGCxNH6BroCvHY4aAg6pQeK",
	"JhvnFJ1mqJ4mi+MFbcJ4Z81xsAIeRw2NrAsUWbRpXkAFDY0d+YbMbEyxKWU8MtGuMJazRL4M0TrLdoU5",
	"vMTYBt4bxpeze6x5ei6bkA52VJPBBv0IHznqEetGAOR4UMGdNQU5y+gCr6HrugIi/mhtRAPVTQ+j2NJ6",
	"huB2ny7Jfv1K6+ck2VpeZts3Hb3K/NWnq6kaqcuAPo8+jifjCRSwwpisyDiRma4ACuh1sFDBzimzEcVl",
	"u6skJaGW66nZkp8aqOCKEkut+www8Wcy6+xXk2f0skSHsLS1LCrvhcCuV/5Hi3SDFJ/4c2xRJlIgn3pV",
	"HyYn+Wcw1dEG7vEJ4zqiZjRZx+lkcuw09Su9tEZZH1qWzVPrnI5rqOBM1iqtPD4o4ZztQrysF1h/z+Ea",
	"HMB+gUL9TJyOtA7I+IpxZWtUNqnYer8D8UxLDqXsXPECVRr27pUtqbH+zx3xRdz+fVu8SfUnL1R/SU2D",
	"ubivqr8gUvqw9JvtFdb10ZbIeIz5XOaziPxNjZzZqB0yxgTVzQZykfo7bnd29nfjYcLFM+6HvG6PYJy+",
	"AKOXuj0KLzl5YjWn1psDFn0+gkLN1mp6nsP8rvffPO3JX7Xv+4hzqOBd+fREltv3sbzu690VQ4C2cF7D",
	"8AL5F4DS7PLy3AzGcdrrBl3OZ88qGxJ0t93PAAAA//8VazlJ3QcAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

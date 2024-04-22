// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

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
	// returns the catalog of all Yarns
	// (GET /catalog)
	GetCatalog(c *gin.Context)
	// Adds a Yarn to the catalog
	// (POST /catalog)
	PostCatalog(c *gin.Context)
	// Delete one item in the catalog
	// (DELETE /catalog/{product_id})
	DeleteCatalogProductId(c *gin.Context, productId string)
	// gets one yarn from the catalog
	// (GET /catalog/{product_id})
	GetCatalogProductId(c *gin.Context, productId string)
	// Update the Yarn availability
	// (PATCH /catalog/{product_id})
	PatchCatalogProductId(c *gin.Context, productId string)
	// Check if the service is running
	// (GET /check)
	GetCheck(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetCatalog operation middleware
func (siw *ServerInterfaceWrapper) GetCatalog(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCatalog(c)
}

// PostCatalog operation middleware
func (siw *ServerInterfaceWrapper) PostCatalog(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostCatalog(c)
}

// DeleteCatalogProductId operation middleware
func (siw *ServerInterfaceWrapper) DeleteCatalogProductId(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteCatalogProductId(c, productId)
}

// GetCatalogProductId operation middleware
func (siw *ServerInterfaceWrapper) GetCatalogProductId(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCatalogProductId(c, productId)
}

// PatchCatalogProductId operation middleware
func (siw *ServerInterfaceWrapper) PatchCatalogProductId(c *gin.Context) {

	var err error

	// ------------- Path parameter "product_id" -------------
	var productId string

	err = runtime.BindStyledParameterWithOptions("simple", "product_id", c.Param("product_id"), &productId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchCatalogProductId(c, productId)
}

// GetCheck operation middleware
func (siw *ServerInterfaceWrapper) GetCheck(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCheck(c)
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

	router.GET(options.BaseURL+"/catalog", wrapper.GetCatalog)
	router.POST(options.BaseURL+"/catalog", wrapper.PostCatalog)
	router.DELETE(options.BaseURL+"/catalog/:product_id", wrapper.DeleteCatalogProductId)
	router.GET(options.BaseURL+"/catalog/:product_id", wrapper.GetCatalogProductId)
	router.PATCH(options.BaseURL+"/catalog/:product_id", wrapper.PatchCatalogProductId)
	router.GET(options.BaseURL+"/check", wrapper.GetCheck)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xWTW/jNhD9KwTbo2x52z3ptu0ChXtoAix6CIogoMmxzFTisMOREcPQfy848vdHkxSb",
	"ICfT5HA4896bB621xTZigMBJV2ud7AJaI8tfDZsG67yMhBGIPciBZ2iPFz8SzHWlfyj3ycpNpvLOUNB9",
	"oXkVQVfaEJmV7vcbOHsEyzlCIs9eM0vjGzNrIP/Z3JkhNmAkbSR0neUH7w7OE5MPtS7006jGkW8jEktm",
	"wwtd6RrHLYYa3WyMVJeyHjnyS6ByljCUkXzr2S8hvyA5Nol3B+MbKXv6VR5BE/3IooMawgiemMyITS3l",
	"53y60g/eFdhmwCIP7XfB84PFBum88O1xMC1cP43k7eFx6NoZ0CVs85YPc5RgzxlLQVttOFZfbqdqpG4i",
	"hLz6eTwZT3Shl0DJS/mfZKcvNEYIJnpd6W1QxlQ6Le1eMDUI3plHwx7D1OlK/wa81VShCVLEkAaOf5pM",
	"8o/FwBDkpomx8Vbulo8C4Vabz+lt+4Q07SBZ8pGHLn7/dvOHEgEqnKuVoZAErdS1raGVrjQBdxSS4gWo",
	"TTs51DSNupPwQg/E/rUbj/ssQkwX+r3FdNTwPx0k/gXd6lW9vmoa3ltUuSlP4HTF1EF/Ruun70brYCPn",
	"nIqMjXPgVOqshZTmXdOscj+fB1Udx0/D0jTeKR9ixyfsf3EuKSNMK8ZDEVzkvS92mi/Xex/qh0cbYDjX",
	"xFfZ3+S4He5MnYwRmRYYKD9yWvQmUHmX1ZjryuLVeaizKWVXK/TA66EhnvJTHGB9wn5/f8bd53PwBBmC",
	"FpcX4b52IyCrOXbBneA9gKEwgMrOqHx4FvPiOWv52JhO3nwexOOGEd2a3FVyphn0a+TUwEmoySnUnLB9",
	"ATnRsF1ccMK8/VEJenNb/n/WOXkf6+yiM3wyzeMXjvP4RDJ/Si6harDlARXfeF5dd9AF2L//85tBAi7j",
	"c1zdN6Clt6B8UtSFMHB8WKCkUn5QU7oU3e/KXG+1tfuguO//DQAA///dFBedLgsAAA==",
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

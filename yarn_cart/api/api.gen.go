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
	// Check if the service is running
	// (GET /cart/check)
	GetCartCheck(c *gin.Context)
	// Creates a new cart
	// (POST /cart/new)
	PostCartNew(c *gin.Context)
	// Clear the cart
	// (DELETE /cart/{cart_id})
	DeleteCartCartId(c *gin.Context, cartId string)
	// Returns current Yarns in cart
	// (GET /cart/{cart_id})
	GetCartCartId(c *gin.Context, cartId string)
	// Adds a Yarn to the cart
	// (PATCH /cart/{cart_id})
	PatchCartCartId(c *gin.Context, cartId string)
	// Remove one item in the cart
	// (DELETE /cart/{cart_id}/{product_id})
	DeleteCartCartIdProductId(c *gin.Context, cartId string, productId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetCartCheck operation middleware
func (siw *ServerInterfaceWrapper) GetCartCheck(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCartCheck(c)
}

// PostCartNew operation middleware
func (siw *ServerInterfaceWrapper) PostCartNew(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostCartNew(c)
}

// DeleteCartCartId operation middleware
func (siw *ServerInterfaceWrapper) DeleteCartCartId(c *gin.Context) {

	var err error

	// ------------- Path parameter "cart_id" -------------
	var cartId string

	err = runtime.BindStyledParameterWithOptions("simple", "cart_id", c.Param("cart_id"), &cartId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter cart_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteCartCartId(c, cartId)
}

// GetCartCartId operation middleware
func (siw *ServerInterfaceWrapper) GetCartCartId(c *gin.Context) {

	var err error

	// ------------- Path parameter "cart_id" -------------
	var cartId string

	err = runtime.BindStyledParameterWithOptions("simple", "cart_id", c.Param("cart_id"), &cartId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter cart_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCartCartId(c, cartId)
}

// PatchCartCartId operation middleware
func (siw *ServerInterfaceWrapper) PatchCartCartId(c *gin.Context) {

	var err error

	// ------------- Path parameter "cart_id" -------------
	var cartId string

	err = runtime.BindStyledParameterWithOptions("simple", "cart_id", c.Param("cart_id"), &cartId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter cart_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchCartCartId(c, cartId)
}

// DeleteCartCartIdProductId operation middleware
func (siw *ServerInterfaceWrapper) DeleteCartCartIdProductId(c *gin.Context) {

	var err error

	// ------------- Path parameter "cart_id" -------------
	var cartId string

	err = runtime.BindStyledParameterWithOptions("simple", "cart_id", c.Param("cart_id"), &cartId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter cart_id: %w", err), http.StatusBadRequest)
		return
	}

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

	siw.Handler.DeleteCartCartIdProductId(c, cartId, productId)
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

	router.GET(options.BaseURL+"/cart/check", wrapper.GetCartCheck)
	router.POST(options.BaseURL+"/cart/new", wrapper.PostCartNew)
	router.DELETE(options.BaseURL+"/cart/:cart_id", wrapper.DeleteCartCartId)
	router.GET(options.BaseURL+"/cart/:cart_id", wrapper.GetCartCartId)
	router.PATCH(options.BaseURL+"/cart/:cart_id", wrapper.PatchCartCartId)
	router.DELETE(options.BaseURL+"/cart/:cart_id/:product_id", wrapper.DeleteCartCartIdProductId)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RWT0/jOhD/Kta8d8yj5T1OufFYadU9AFpOaIWQ60xbs8nY2OOiqsp3X41TaGlSEIgV",
	"2lOc8fz9zW8mWYNxjXeExBHKNUSzwEbn45kOLE8fnMfAFrPU6MC3tpLjzIVGM5SQkq2gAF55hBIiB0tz",
	"aAuwjE02+jvgDEr4a7QNNtpEGkmYSVZs2ycfbnqHhsXH9rpcbx0+z+k+aWLLKzlvHFhinGMQDysd6LUk",
	"rkWn3UlAh6BXYn29sX4eUS+1rfW0xp2QU+dq1CRGPrgqmUecergksnxrXO3C4WvSDR6+9cGa3WtKzVSq",
	"7SMoIkszl5UtS8a5JiXAqtPLifpHXXgkOf13ND4aQwFLDNE6ghKOs6QtwHkk7S2U8KjkNS8yFiOhxMgs",
	"0PyU1zlm1ghamq2jSQUlfEWWeGdZqYCA0TuKHZb/jsfyqDCaYD13ga8wLK1BZaMKiSjX3hYQU9PosIIS",
	"sitlZ4oXqOKwdpcZ4UNuoIsDiV26mDM7x4ec133CyP+7KnPJOGKkbKW9r63JdqO76Gg7K2/hd9uFsAEr",
	"KDkkbHtYHL8p8Htnc5gnz1uQCWICasZKxWQMxjhLdb3a70RWiUorwgclOUg8PY9Q/uiWyM22GetNjm3X",
	"8hoZ+035kuWZMAJeldkWdIOMQbyuwUqCwkAooJuUp+L3IS524NpH4aYH/0mfih0ONerwCg6ikuk4jEHx",
	"8mgcqHQgl1zkb0Ng/KHMHyLWt6uLc5UXrHIzJes57kH5HTkFisqkEJBYycKKytJBZL1msxiYbhF/Nrp/",
	"2Ep5T0fzB0VXVW9ACjgZ2u4TWuraVsqST7zX+9Oqkl2SXbJ7YZ76O2W03n5437RgLjuzT2BIMbjLdv4f",
	"PnydZWADNm452K1DFuRYzVyiqjeq4kk5QiV/ZzKkL7SsfRKtH2vtOHXT/goAAP//Ar84MooKAAA=",
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

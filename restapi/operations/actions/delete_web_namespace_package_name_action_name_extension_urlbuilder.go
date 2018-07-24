// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// DeleteWebNamespacePackageNameActionNameExtensionURL generates an URL for the delete web namespace package name action name extension operation
type DeleteWebNamespacePackageNameActionNameExtensionURL struct {
	ActionName  string
	Extension   string
	Namespace   string
	PackageName string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteWebNamespacePackageNameActionNameExtensionURL) WithBasePath(bp string) *DeleteWebNamespacePackageNameActionNameExtensionURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteWebNamespacePackageNameActionNameExtensionURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteWebNamespacePackageNameActionNameExtensionURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/web/{namespace}/{packageName}/{actionName}.{extension}"

	actionName := o.ActionName
	if actionName != "" {
		_path = strings.Replace(_path, "{actionName}", actionName, -1)
	} else {
		return nil, errors.New("ActionName is required on DeleteWebNamespacePackageNameActionNameExtensionURL")
	}

	extension := o.Extension
	if extension != "" {
		_path = strings.Replace(_path, "{extension}", extension, -1)
	} else {
		return nil, errors.New("Extension is required on DeleteWebNamespacePackageNameActionNameExtensionURL")
	}

	namespace := o.Namespace
	if namespace != "" {
		_path = strings.Replace(_path, "{namespace}", namespace, -1)
	} else {
		return nil, errors.New("Namespace is required on DeleteWebNamespacePackageNameActionNameExtensionURL")
	}

	packageName := o.PackageName
	if packageName != "" {
		_path = strings.Replace(_path, "{packageName}", packageName, -1)
	} else {
		return nil, errors.New("PackageName is required on DeleteWebNamespacePackageNameActionNameExtensionURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteWebNamespacePackageNameActionNameExtensionURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteWebNamespacePackageNameActionNameExtensionURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteWebNamespacePackageNameActionNameExtensionURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteWebNamespacePackageNameActionNameExtensionURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteWebNamespacePackageNameActionNameExtensionURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *DeleteWebNamespacePackageNameActionNameExtensionURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

package framework

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-zoo/bone"
)

// Request struct basically adds a context to the http.Request so that
// authenticator or any other middleware could push out the data
// to main request handler
type Request struct {
	*http.Request
	context map[string]interface{}
}

// Flash contains session flash message structure
type Flash struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// GetBoneValue wrapper for getting url values
func (r *Request) GetBoneValue(key string) string {
	return bone.GetValue(r.Request, key)
}

// Push push a value to request context
func (r *Request) Push(key string, value interface{}) {
	if r.context == nil {
		r.context = map[string]interface{}{}
	}
	r.context[key] = value
}

// Value get the value from request context
func (r *Request) Value(key string) interface{} {
	return r.context[key]
}

// QueryParam get values from GET params
func (r *Request) QueryParam(key string) string {
	return r.URL.Query().Get(key)
}

// ReadBody reads request body
func (r *Request) ReadBody() (map[string]interface{}, error) {
	return ReadBody(r.Request)
}

// ReadBody reads request body
func ReadBody(r *http.Request) (map[string]interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	bodyMap := make(map[string]interface{})
	err := decoder.Decode(&bodyMap)
	if err != nil {
		return bodyMap, err
	}
	return bodyMap, nil
}

// Bind binds request body to certain structure
func (r *Request) Bind(v interface{}) error {
	return Bind(r.Request.Body, v)
}

// Bind binds request body to certain structure
func Bind(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	err := json.NewDecoder(body).Decode(v)
	return err
}

// GetFlash returns flash if set by response
func (r *Request) GetFlash(w http.ResponseWriter) *Flash {
	c, err := r.Cookie("FLAMSGC")
	if err != nil {
		return nil
	}

	v, err := base64.URLEncoding.DecodeString(c.Value)
	if err != nil {
		return nil
	}

	f := &Flash{}
	if err := json.Unmarshal(v, f); err != nil {
		return nil
	}

	c.MaxAge = -1
	http.SetCookie(w, c)

	return f
}

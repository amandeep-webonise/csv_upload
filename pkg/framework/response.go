package framework

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

var (

	// DefaultErrorCode specifying default error code as 409 conflict
	DefaultErrorCode = http.StatusConflict
)

// Response contains http response
type Response struct {
	http.ResponseWriter

	msg        string
	data       JSONResponse
	statusCode int
	err        error
	written    bool
	success    bool
	minifier   *minify.M
}

// NewResponse creates new response for a request
func NewResponse(w http.ResponseWriter) Response {
	r := Response{}
	r.ResponseWriter = w
	r.data = JSONResponse{}
	r.statusCode = -1
	r.msg = "success"
	r.written = false
	r.success = true
	r.minifier = minify.New()
	r.minifier.AddFunc("text/html", html.Minify)

	return r
}

// Data sends data in response
func (r *Response) Data(data map[string]interface{}) {
	r.data = data
}

// PutInData sends data in response
func (r *Response) PutInData(k string, v interface{}) {
	r.data[k] = v
}

// Written set response written to true
func (r *Response) Written() {
	r.written = true
}

// StatusCode set status code for a response
func (r *Response) StatusCode(code int) {
	r.statusCode = code
}

// SetSuccess sends success false but http status as 200
func (r *Response) SetSuccess(flag bool) {
	r.success = flag
}

// Error writes error in response
func (r *Response) Error(err error) {
	r.err = err
}

// BadRequest throws bad request 400 response
func (r *Response) BadRequest(err ...error) {
	r.statusCode = http.StatusBadRequest
	if len(err) > 0 {
		r.err = err[0]
	}
}

// NotFound throws not found 404 response
func (r *Response) NotFound(err ...error) {
	r.statusCode = http.StatusNotFound
	if len(err) > 0 {
		r.err = err[0]
	}

}

// Unauthorised throws unauthorized 403 response
func (r *Response) Unauthorised(err ...error) {
	r.statusCode = http.StatusUnauthorized
	if len(err) > 0 {
		r.err = err[0]
	}
	return
}

// InternalError throws internal server 500 response
func (r *Response) InternalError(err ...error) {
	r.statusCode = http.StatusInternalServerError
	if len(err) > 0 {
		r.err = err[0]
	}
}

// Conflict throws conflict 409 response
func (r *Response) Conflict(err ...error) {
	r.statusCode = DefaultErrorCode
	if len(err) > 0 {
		r.err = err[0]
	}
}

// Message writes message in response
func (r *Response) Message(msg string) {
	r.msg = msg

}

// Write writes response
func (r *Response) Write() {
	if r.written {
		return
	}
	if r.err != nil || r.statusCode > 399 {
		r.writeErrorResponse()
		return
	}
	if r.statusCode == http.StatusFound {
		return
	}
	r.writeResponse()
}

func (r *Response) writeResponse() {
	r.ResponseWriter.Header().Add("Content-Type", "application/json")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Auth-Key, Session-Key")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
	if r.statusCode == -1 {
		r.ResponseWriter.WriteHeader(http.StatusOK)
	} else {
		r.ResponseWriter.WriteHeader(r.statusCode)
	}
	res := JSONResponse{
		"message": r.msg,
		"success": r.success,
		"data":    r.data,
	}
	r.ResponseWriter.Write(res.ByteArray())
}

func (r *Response) writeErrorResponse() {
	r.ResponseWriter.Header().Add("Content-Type", "application/json")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Auth-Key")
	if r.statusCode == -1 {
		r.ResponseWriter.WriteHeader(DefaultErrorCode)
	} else {
		r.ResponseWriter.WriteHeader(r.statusCode)
	}
	if r.err == nil {
		switch r.statusCode {
		case http.StatusUnauthorized:
			r.err = errors.New("Unauthorized access")
		default:
			r.err = errors.New("Illegal request")
		}

	}

	res := JSONResponse{
		"message": r.err.Error(),
		"success": false,
		"data":    r.data,
	}
	r.ResponseWriter.Write(res.ByteArray())

}

// Redirect redirects user to given URL
func (r *Response) Redirect(url string, req *http.Request) {
	r.StatusCode(http.StatusFound)
	r.ResponseWriter.Header().Add("Content-Type", "application/json")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Auth-Key, Session-Key")
	r.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
	http.Redirect(r.ResponseWriter, req, url, r.statusCode)
}

// MinifyHTML minifies html data by removing bytes from a file(such as whitespace)
// without changing its output
func (r *Response) MinifyHTML(res string) (string, error) {
	return r.minifier.String("text/html", res)
}

//RenderHTML renders HTML template
func (r *Response) RenderHTML(res string) {
	minRes, err := r.MinifyHTML(res)
	if err != nil {
		io.WriteString(r.ResponseWriter, res)
		return
	}
	io.WriteString(r.ResponseWriter, minRes)
}

// SetFlash sets flash structure containing flash message and type to cookie value
func (r *Response) SetFlash(msg, msgType string) error {
	f := &Flash{}
	f.Message = msg
	f.Type = msgType

	mr, err := json.Marshal(f)
	if err != nil {
		return err
	}

	c := &http.Cookie{}
	c.Name = "FLAMSGC"
	c.Value = base64.URLEncoding.EncodeToString(mr)
	c.HttpOnly = true
	http.SetCookie(r.ResponseWriter, c)

	return nil
}

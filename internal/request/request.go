package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

var ERROR_BAD_START_LINE = fmt.Errorf("bad start line")
var ERROR_WRONG_START_LINE = fmt.Errorf("malformed start line")
var ERROR_UNSUPPORTED_HTTP_VERSION = fmt.Errorf("IS NOT A VALID HTTP TYPE\n")

// var BAD_START_LINE = fmt.Errorf("bad start line")
var Breaker = "\r\n"

func (r *RequestLine) ValidHTTP() bool {
	return r.HttpVersion == "HTTP/1.1"
}

func parseRequestLine(b string) (*RequestLine, string, error) {
	idx := strings.Index(b, Breaker)

	if idx == -1 {
		return nil, b, nil
	}

	startline := b[:idx]

	remainingdata := b[idx+len(Breaker):]

	//RFC implements SP i,e single space ! >DSP is a malfunction
	partsofStartLine := strings.Split(startline, " ")

	httpParts := strings.Split(partsofStartLine[2], "/")

	if len(httpParts) != 2 || httpParts[0] != "HTTP" || httpParts[1] != "1.1" {
		return nil, remainingdata, ERROR_UNSUPPORTED_HTTP_VERSION
	}

	if len(partsofStartLine) != 3 {
		return nil, remainingdata, ERROR_WRONG_START_LINE
	}

	rl := &RequestLine{
		Method:        partsofStartLine[0],
		RequestTarget: partsofStartLine[1],
		HttpVersion:   httpParts[1],
	}

	return rl, remainingdata, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	data, err := io.ReadAll(reader)

	if err != nil {
		return nil, fmt.Errorf("READER TERMINATED, %w\n", err)
	}
	str := string(data)

	rl, _, err := parseRequestLine(str)

	if err !=nil{
		return  nil, err
	}

	//after parseReqeust line and validation we recicve the Request

	return &Request{
		RequestLine: *rl,
	}, err

}

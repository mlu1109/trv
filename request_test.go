package trafikverket

import (
	"encoding/xml"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	request := newRequest("test")
	out, err := xml.Marshal(request)
	assert.Nil(t, err)
	log.Print(string(out))
}

package conv

import (
	"testing"
	"time"

	"github.com/mindworker/go-openapi-without-unsafe/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestDateTimeValue(t *testing.T) {
	assert.Equal(t, strfmt.DateTime{}, DateTimeValue(nil))
	time := strfmt.DateTime(time.Now())
	assert.Equal(t, time, DateTimeValue(&time))
}

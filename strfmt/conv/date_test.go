package conv

import (
	"testing"
	"time"

	"github.com/mindworker/go-openapi-without-unsafe/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestDateValue(t *testing.T) {
	assert.Equal(t, strfmt.Date{}, DateValue(nil))
	date := strfmt.Date(time.Now())
	assert.Equal(t, date, DateValue(&date))
}

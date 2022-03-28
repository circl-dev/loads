package loads_test

import (
	"testing"

	"github.com/circl-dev/loads"
	"github.com/circl-dev/validate"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPathLoaderIssue(t *testing.T) {
	swaggerFile := "fixtures/json/resources/pathLoaderIssue.json"
	document, err := loads.Spec(swaggerFile)
	require.NoError(t, err)
	require.NotNil(t, document)
	validationErrs := validate.Spec(document, strfmt.Default)
	assert.NoError(t, validationErrs)
}

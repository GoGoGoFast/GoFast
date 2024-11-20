package urlutil_test

import (
	"net/url"
	"testing"

	"GoFast/pkg/util/urlutil"
	"github.com/stretchr/testify/assert"
)

func TestNewURL(t *testing.T) {
	u := &urlutil.URLUtil{}
	testURL := "https://example.com"
	parsedURL, err := u.NewURL(testURL)
	assert.NoError(t, err)
	assert.Equal(t, testURL, parsedURL.String())
}

func TestGetURL(t *testing.T) {
	u := &urlutil.URLUtil{}
	resourceName := "example/resource"
	expectedURL := "classpath://" + resourceName
	parsedURL, err := u.GetURL(resourceName)
	assert.NoError(t, err)
	assert.Equal(t, expectedURL, parsedURL.String())
}

func TestNormalize(t *testing.T) {
	u := &urlutil.URLUtil{}
	rawURL := "https://example.com//foo//bar"
	expectedURL := "https://example.com/foo/bar"
	normalizedURL, err := u.Normalize(rawURL)
	assert.NoError(t, err)
	assert.Equal(t, expectedURL, normalizedURL)
}
func TestEncode(t *testing.T) {
	u := &urlutil.URLUtil{}
	content := "foo bar"
	expectedEncoded := "foo+bar"
	encoded := u.Encode(content)
	assert.Equal(t, expectedEncoded, encoded)
}

func TestDecode(t *testing.T) {
	u := &urlutil.URLUtil{}
	content := "foo+bar"
	expectedDecoded := "foo bar"
	decoded, err := u.Decode(content)
	assert.NoError(t, err)
	assert.Equal(t, expectedDecoded, decoded)
}

func TestGetPath(t *testing.T) {
	u := &urlutil.URLUtil{}
	rawURL := "https://example.com/foo/bar"
	expectedPath := "/foo/bar"
	path, err := u.GetPath(rawURL)
	assert.NoError(t, err)
	assert.Equal(t, expectedPath, path)
}

func TestToURI(t *testing.T) {
	u := &urlutil.URLUtil{}
	rawURL := "https://example.com"
	expectedURI, err := url.Parse(rawURL)
	assert.NoError(t, err)
	uri, err := u.ToURI(rawURL)
	assert.NoError(t, err)
	assert.Equal(t, expectedURI, uri)
}

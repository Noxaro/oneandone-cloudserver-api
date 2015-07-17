/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package oneandone_cloudserver_api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUrl_1(t *testing.T) {
	api := New("token", "http://test.de/v1")

	result := createUrl(api)
	assert.Equal(t, "http://test.de/v1", result)
}

func TestCreateUrl_2(t *testing.T) {
	api := New("token", "http://test.de/v1")

	result := createUrl(api, "servers")
	assert.Equal(t, "http://test.de/v1/servers", result)
}

func TestCreateUrl_3(t *testing.T) {
	api := New("token", "http://test.de/v1")

	result := createUrl(api, "servers", 1)
	assert.Equal(t, "http://test.de/v1/servers/1", result)
}

func TestAppendQueryParams_1(t *testing.T) {
	params := map[string]interface{}{
		"foo": "bar",
	}
	result := appendQueryParams("http://test/", params)
	assert.Equal(t, "http://test/?foo=bar", result)
}

func TestAppendQueryParams_2(t *testing.T) {
	params := map[string]interface{}{
		"foo":  "bar",
		"size": 5,
	}
	result := appendQueryParams("http://test/", params)
	assert.Equal(t, "http://test/?foo=bar&size=5", result)
}

func TestAppendQueryParams_3(t *testing.T) {
	params := map[string]interface{}{}
	result := appendQueryParams("http://test/", params)
	assert.Equal(t, "http://test/", result)
}

func TestAppendQueryParams_UrlEncode_1(t *testing.T) {
	params := map[string]interface{}{
		"test": "1&2=3",
	}
	result := appendQueryParams("http://test/", params)
	assert.Equal(t, "http://test/?test=1%262%3D3", result)
}

func TestAppendQueryParams_UrlEncode_2(t *testing.T) {
	params := map[string]interface{}{
		"test!": "1&2=3",
	}
	result := appendQueryParams("http://test/", params)
	assert.Equal(t, "http://test/?test%21=1%262%3D3", result)
}

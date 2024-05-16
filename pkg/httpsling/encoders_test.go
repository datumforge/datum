package httpsling_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/pkg/httpsling"
)

func TestJSONEncoder_Encode(t *testing.T) {
	encoder := &httpsling.JSONEncoder{}

	// Test encoding a struct
	data := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "John Doe",
		Age:  30,
	}

	reader, err := encoder.Encode(data)
	require.NoError(t, err)

	encodedData, err := ioutil.ReadAll(reader)
	require.NoError(t, err)

	expectedData := `{"name":"John Doe","age":30}`
	require.Equal(t, expectedData, string(encodedData))
}

func TestJSONDecoder_Decode(t *testing.T) {
	decoder := &httpsling.JSONDecoder{}

	// Test decoding JSON data into a struct
	jsonData := `{"name":"John Doe","age":30}`
	var data struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	err := decoder.Decode(bytes.NewReader([]byte(jsonData)), &data)
	require.NoError(t, err)

	expectedData := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "John Doe",
		Age:  30,
	}
	require.Equal(t, expectedData, data)
}

func TestXMLEncoder_Encode(t *testing.T) {
	encoder := &httpsling.XMLEncoder{}

	// Test encoding a struct
	data := struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}{
		Name: "John Doe",
		Age:  30,
	}

	reader, err := encoder.Encode(data)
	require.NoError(t, err)

	encodedData, err := ioutil.ReadAll(reader)
	require.NoError(t, err)

	expectedData := `<root><name>John Doe</name><age>30</age></root>`
	require.Equal(t, expectedData, string(encodedData))
}

func TestXMLDecoder_Decode(t *testing.T) {
	decoder := &httpsling.XMLDecoder{}

	// Test decoding XML data into a struct
	xmlData := `<root><name>John Doe</name><age>30</age></root>`
	var data struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}

	err := decoder.Decode(bytes.NewReader([]byte(xmlData)), &data)
	require.NoError(t, err)

	expectedData := struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}{
		Name: "John Doe",
		Age:  30,
	}
	require.Equal(t, expectedData, data)
}

func TestYAMLEncoder_Encode(t *testing.T) {
	encoder := &httpsling.YAMLEncoder{}

	// Test encoding a struct
	data := struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
	}{
		Name: "John Doe",
		Age:  30,
	}

	reader, err := encoder.Encode(data)
	require.NoError(t, err)

	encodedData, err := ioutil.ReadAll(reader)
	require.NoError(t, err)

	expectedData := "name: John Doe\nage: 30\n"
	require.Equal(t, expectedData, string(encodedData))
}

func TestYAMLDecoder_Decode(t *testing.T) {
	decoder := &httpsling.YAMLDecoder{}

	// Test decoding YAML data into a struct
	yamlData := "name: John Doe\nage: 30\n"
	var data struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
	}

	err := decoder.Decode(bytes.NewReader([]byte(yamlData)), &data)
	require.NoError(t, err)

	expectedData := struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
	}{
		Name: "John Doe",
		Age:  30,
	}
	require.Equal(t, expectedData, data)
}

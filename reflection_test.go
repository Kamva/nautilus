package nautilus

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type inner struct{}

type Embed struct{}

type sampleStruct struct {
	inner
	Exposed      int `tag:"tag_value"`
	unExposed    string
	ExposedEmbed Embed
}

var expectedFieldData = []FieldData{
	{Name: "inner", Type: reflect.TypeOf(inner{}), Tags: "", Exported: false, Anonymous: true},
	{Name: "Exposed", Type: reflect.TypeOf(1), Tags: `tag:"tag_value"`, Exported: true, Anonymous: false},
	{Name: "unExposed", Type: reflect.TypeOf(""), Tags: "", Exported: false, Anonymous: false},
	{Name: "ExposedEmbed", Type: reflect.TypeOf(Embed{}), Tags: "", Exported: true, Anonymous: false},
}

func TestGetType(t *testing.T) {
	s := sampleStruct{}
	res := GetType(s)
	res2 := GetType(&s)

	assert.Equal(t, "sampleStruct", res)
	assert.Equal(t, "sampleStruct", res2)

	assert.Panics(t, func() {
		_ = GetType(1)
	})
}

func TestGetStructFieldsData(t *testing.T) {
	t.Run("direct", func(t *testing.T) {
		s := sampleStruct{Exposed: 10, unExposed: "str", ExposedEmbed: Embed{}}
		fieldsData, err := GetStructFieldsData(s)

		assert.Nil(t, err)
		assert.Equal(t, 4, len(fieldsData))

		for i, fieldData := range fieldsData {
			assert.Equal(t, expectedFieldData[i].Name, fieldData.Name)
			assert.Equal(t, expectedFieldData[i].Type, fieldData.Type)
			assert.Equal(t, expectedFieldData[i].Tags, fieldData.Tags)
			assert.Equal(t, expectedFieldData[i].Exported, fieldData.Exported)
			assert.Equal(t, expectedFieldData[i].Anonymous, fieldData.Anonymous)
		}
	})
	t.Run("pointer", func(t *testing.T) {
		s := &sampleStruct{Exposed: 10, unExposed: "str", ExposedEmbed: Embed{}}
		fieldsData, err := GetStructFieldsData(s)

		assert.Nil(t, err)
		assert.Equal(t, 4, len(fieldsData))

		for i, fieldData := range fieldsData {
			assert.Equal(t, expectedFieldData[i].Name, fieldData.Name)
			assert.Equal(t, expectedFieldData[i].Type, fieldData.Type)
			assert.Equal(t, expectedFieldData[i].Tags, fieldData.Tags)
			assert.Equal(t, expectedFieldData[i].Exported, fieldData.Exported)
			assert.Equal(t, expectedFieldData[i].Anonymous, fieldData.Anonymous)
		}
	})
	t.Run("panic", func(t *testing.T) {
		fieldsData, err := GetStructFieldsData(42)

		assert.NotNil(t, err)
		assert.Nil(t, fieldsData)
	})
}

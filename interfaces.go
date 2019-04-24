package nautilus

import "reflect"

// Taggable is an interface for struct that have tagging
type Taggable interface {
	GetTag(caller interface{}, field string, tag string) string
}

// BaseTaggable is a base struct for helper for struct needed to implement
// Taggable interface
type BaseTaggable struct{}

// GetTag Get `tag` value on the `field` of `caller`
func (r BaseTaggable) GetTag(caller interface{}, field string, tag string) string {
	t := reflect.TypeOf(caller)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	taggedField, _ := t.FieldByName(field)
	return taggedField.Tag.Get(tag)
}

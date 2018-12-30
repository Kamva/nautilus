package nautilus

// Taggable is an interface for struct that have tagging
type Taggable interface {
	GetTag(caller interface{}, field string, tag string) string
}

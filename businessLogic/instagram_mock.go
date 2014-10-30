package businessLogic

import (
	"instagram/structs"
)

type Endpoint struct {
}

func (e *Endpoint) GetTag(tagName string) (tag structs.Tag) {
	return
}

func (e *Endpoint) GetTaggedMedia(
	tag string,
	count int,
	min_tag_id string,
	max_tag_id string,
) (media []structs.Media) {
	if tag != "" {
		return append(media, structs.Media{Type: "image"})
	}
	return
}

func (e *Endpoint) SearchTag(query string) (tags []structs.Tag) {
	return
}

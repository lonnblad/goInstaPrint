package businessLogic

import (
	"instagram/endpoints"
	"instagram/structs"
	"strings"
)

type Logic struct {
	instagram endpoints.Instagram
	tag       string
}

type picture struct {
	caption string
	id      string
	images  images
}

type images struct {
	normal    image
	thumbnail image
}

type image struct {
	url    string
	width  int
	height int
}

func tagExistsInMediaCaption(tag string, caption string) bool {
	if tag == "" || caption == "" {
		return false
	}
	return strings.Contains(caption, "#"+tag)
}

func (l *Logic) getNewPictures() (pics []picture) {
	media := l.instagram.GetTaggedMedia(l.tag, 1, "", "")
	for _, elem := range media {
		if filterMediaType("image", elem.Type) {
			pics = append(pics, mapInstagramMediaToPicture(elem))
		}
	}
	return pics
}

func printPicture(pic picture) {}

func filterMediaType(filter, mediaType string) bool {
	return filter == mediaType || filter == ""
}

func mapInstagramMediaToPicture(media structs.Media) (pic picture) {
	pic.caption = media.Caption.Text
	pic.id = media.Id
	copyNormalImageFromMediaToPicture(media, &pic)
	copyThumbImageFromMediaToPicture(media, &pic)
	return
}

func copyNormalImageFromMediaToPicture(media structs.Media, pic *picture) {
	pic.images.normal.height = media.Images.Normal.Height
	pic.images.normal.width = media.Images.Normal.Width
	pic.images.normal.url = media.Images.Normal.URL
}

func copyThumbImageFromMediaToPicture(media structs.Media, pic *picture) {
	pic.images.thumbnail.height = media.Images.Thumb.Height
	pic.images.thumbnail.width = media.Images.Thumb.Width
	pic.images.thumbnail.url = media.Images.Thumb.URL
}

func (l *Logic) Run() {
	pics := l.getNewPictures()
	if len(pics) > 0 {
		for _, pic := range pics {
			if tagExistsInMediaCaption(l.tag, pic.caption) {
				printPicture(pic)
			}
		}
	}
}

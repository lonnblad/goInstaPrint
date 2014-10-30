package businessLogic

import (
	"instagram/structs"
	"testing"
)

func Test_getNewPictures_1(t *testing.T) {
	l := Logic{instagram: &Endpoint{}}
	pics := l.getNewPictures()
	if len(pics) != 0 {
		t.Errorf("got pictures with empty tag, pics: %+v", pics)
	}
}

func Test_getNewPictures_2(t *testing.T) {
	l := Logic{instagram: &Endpoint{}, tag: "lol"}
	pics := l.getNewPictures()
	if len(pics) != 1 {
		t.Errorf("got pictures with empty tag, pics: %+v", pics)
	}
}

func Test_mapInstagramMediaToPicture_caption(t *testing.T) {
	media := structs.Media{Caption: structs.Caption{"lol"}}
	expected := picture{caption: "lol"}
	pic := mapInstagramMediaToPicture(media)
	if pic != expected {
		t.Errorf("pics not equal: [%+v] != [%+v]", pic, expected)
	}
}

func Test_mapInstagramMediaToPicture_id(t *testing.T) {
	media := structs.Media{Id: "lol"}
	expected := picture{id: "lol"}
	pic := mapInstagramMediaToPicture(media)
	if pic != expected {
		t.Errorf("pics not equal: [%+v] != [%+v]", pic, expected)
	}
}

func Test_mapInstagramMediaToPicture_normalImage(t *testing.T) {
	media := structs.Media{}
	media.Images.Normal = structs.Image{"lol", 1, 1}
	expected := picture{}
	expected.images.normal = image{"lol", 1, 1}
	pic := mapInstagramMediaToPicture(media)
	if pic != expected {
		t.Errorf("pics not equal: [%+v] != [%+v]", pic, expected)
	}
}

func Test_mapInstagramMediaToPicture_thumbImage(t *testing.T) {
	media := structs.Media{}
	media.Images.Thumb = structs.Image{"lol", 1, 1}
	expected := picture{}
	expected.images.thumbnail = image{"lol", 1, 1}
	pic := mapInstagramMediaToPicture(media)
	if pic != expected {
		t.Errorf("pics not equal: [%+v] != [%+v]", pic, expected)
	}
}

func Test_filterMediaType_1(t *testing.T) {
	if !filterMediaType("", "") {
		t.Errorf("filter: [%s] == mediaType [%s]", "", "")
	}
}

func Test_filterMediaType_2(t *testing.T) {
	filter := "a"
	mediaType := ""
	if filterMediaType(filter, mediaType) {
		t.Errorf("filter: [%s] == mediaType [%s]", filter, mediaType)
	}
}

func Test_filterMediaType_3(t *testing.T) {
	filter := "a"
	mediaType := "a"
	if !filterMediaType(filter, mediaType) {
		t.Errorf("filter: [%s] == mediaType [%s]", filter, mediaType)
	}
}

func Test_filterMediaType_4(t *testing.T) {
	filter := ""
	mediaType := "a"
	if !filterMediaType(filter, mediaType) {
		t.Errorf("filter: [%s] == mediaType [%s]", filter, mediaType)
	}
}

func Test_tagExistsInMediaCaption_1(t *testing.T) {
	if tagExistsInMediaCaption("", "") {
		t.Errorf("Tag: [%s] did exist in caption: [%s]", "", "")
	}
}

func Test_tagExistsInMediaCaption_2(t *testing.T) {
	tag := "LOL"
	caption := "#LOL"
	if !tagExistsInMediaCaption(tag, caption) {
		t.Errorf("Tag: [%s] didn't exist in caption: [%s]", tag, caption)
	}
}

func Test_tagExistsInMediaCaption_3(t *testing.T) {
	tag := "LOL"
	caption := ""
	if tagExistsInMediaCaption(tag, caption) {
		t.Errorf("Tag: [%s] did exist in caption: [%s]", tag, caption)
	}
}

func Test_tagExistsInMediaCaption_4(t *testing.T) {
	tag := "LOL"
	caption := "NAH"
	if tagExistsInMediaCaption(tag, caption) {
		t.Errorf("Tag: [%s] did exist in caption: [%s]", tag, caption)
	}
}

func Test_tagExistsInMediaCaption_5(t *testing.T) {
	tag := "LOL"
	caption := "LOL"
	if tagExistsInMediaCaption(tag, caption) {
		t.Errorf("Tag: [%s] did exist in caption: [%s]", tag, caption)
	}
}

// func Test_Run(t *testing.T) {
// 	l := &Logic{}
// 	l.Run()
// }

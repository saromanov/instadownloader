package instadownloader

import (
	"image"
	"io/ioutil"
	"os"
	"testing"
)

//client id took from http://tutorialzine.com/2014/07/5-practical-examples-for-learning-facebooks-react-framework/

func readImage(path string) (image.Image, error) {
	img, _ := os.Open(path)
	return decodeImage(img)
}

func TestDownloadLinks(t *testing.T) {
	links, err := GetLinksToPopularPhotos("642176ece1e7445e99244cec26f4de1f")
	if err != nil {
		t.Error(err)
	}

	if len(links) == 0 {
		t.Error("Error with downloading links from Instagram")
	}
}

func TestDownloadLinksByTag(t *testing.T) {
	links, err := GetLinksByTag("642176ece1e7445e99244cec26f4de1f", "bird")
	if err != nil {
		t.Error(err)
	}

	if len(links) == 0 {
		t.Error("Error with downloading links by tags from Instagram")
	}
}

func TestDownloadAndResize(t *testing.T) {
	links, err := GetLinksToPopularPhotos("642176ece1e7445e99244cec26f4de1f")
	if err != nil {
		t.Error(err)
	}
	newwidth := 50
	newheight := 50
	dirname := "./example1"

	SaveWithNewSize(links, dirname, newwidth, newheight)

	files, _ := ioutil.ReadDir(dirname)
	img, err := readImage(dirname + "/" + files[0].Name())
	if err != nil {
		t.Error(err)
	}
	bounds := img.Bounds()
	if bounds.Max.X != newwidth {
		t.Error("Unexpected width of the images")
	}

	if bounds.Max.Y != newheight {
		t.Error("Unexpected height of the images")
	}

}

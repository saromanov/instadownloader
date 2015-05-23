package instadownloader

import (
	"testing"
)

//client id took from http://tutorialzine.com/2014/07/5-practical-examples-for-learning-facebooks-react-framework/

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

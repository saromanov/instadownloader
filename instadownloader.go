package instadownloader

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"math/rand"
)


type Image struct {
	Width  int64
	Height int64
	Url    string
}

type InstagramData struct {
	Images      struct {
		Standard Image `json:"standard_resolution"`
		Low      Image `json:"low_resolution"`
	}
	Link string
	Tags []string
	Type string
}

type Instagram struct {
	Data []InstagramData
}


//Return list of urls to photos
func GetLinksToPopularPhotos(clientid string) ([]string, error) {
	log.Print("Getting popular photos from Instagram...")
	path := fmt.Sprintf("https://api.instagram.com/v1/media/popular?client_id=%s&callback=?", clientid)
	data, err := getInstagramData(path)
	if err != nil {
		return nil, err
	}
	return getLinks(data), nil
}

//This method provides search instagram photos by tag. Need for tile images
//Return link to photos
func GetLinksByTag(clientid string, tag string) ([]string, error) {
	path := fmt.Sprintf("https://api.instagram.com/v1/tags/%s/media/recent?client_id=%s&callback=", tag, clientid)
	data, err := getInstagramData(path)
	if err != nil {
		return nil, err
	}
	return getLinks(data), nil
}

//This method provides download saving data from output of GetLinksToPopularPhotos
func DownloadAndSave(links []string, outdir string) {
	log.Println("Start to download data")
	for _, title := range links {
		resp, err := http.Get(title)
		if err != nil {
			log.Print("Can't download picture")
		}
		img, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Print("Can't decode picture")
		}
		resp.Body.Close()

		if _, err := os.Stat(outdir); os.IsNotExist(err) {
			os.Mkdir(outdir, 0777)
		}
		path := fmt.Sprintf("%s/%s.jpg", outdir, randName(5))
		toimg, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		jpeg.Encode(toimg, img, &jpeg.Options{jpeg.DefaultQuality})
		toimg.Close()
	}
	log.Println("Finished download data")
}

func randName(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

//This private method provide download photos from instagram
func getInstagramData(path string) (Instagram, error) {
	response, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
		return Instagram{}, err
	}
	preresult, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return Instagram{}, err
	}
	var data Instagram
	json.Unmarshal(preresult, &data)
	defer response.Body.Close()
	return data, nil
}

//This private method provides getting links to photos
func getLinks(inst Instagram) []string {
	links := []string{}
	log.Print("Getting links from Loaded Instagram Data")
	for _, data := range inst.Data {
		links = append(links, data.Images.Standard.Url)
	}
	log.Print("Finished to getting links from Instagram Data")
	return links
}

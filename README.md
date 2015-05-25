# instadownloader

Simple package for downloading photos from Instagram

**Usage**

Download popular photos
```go
import (
 "instadownloader"
)

result, err := instadownloader.GetLinksToPopularPhotos("Your client id")
if err != nil {
	panic(err)
}
instadownloader.DownloadAndSave(result, "./outdir")
```

or download photos by tags
```go
import (
 "instadownloader"
)

links, err := GetLinksByTag("Your client id", "bird")
if err != nil {
	panic(err)
}
instadownloader.DownloadAndSave(result, "./outdir")
```

After downloading links, you can resize images
```go
import (
 "instadownloader"
)

links, err := GetLinksToPopularPhotos("Your client id", "bird")
if err != nil {
	panic(err)
}
instadownloader.SaveWithNewSize(result, "./outdir", 50,50)
```

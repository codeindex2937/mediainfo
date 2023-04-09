= Intro
For read media info on Windows

= Usage
1. download MediaInfo.dll from https://mediaarea.net/en/MediaInfo/Download/Windows
2. run
```go
func main() {
	h := mediainfo.New()
	defer h.Delete()

	h.Option("CharSet", "UTF-8")
	h.Option("Inform", "JSON")
	h.Option("Complete", "1")
	h.Option("ParseSpeed", "0.5")

	if !h.Open(os.Args[1]) {
		h.Close()
		return
	}
	defer h.Close()

	info, err := h.Inform()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v\n", info.Audio()[0])
}
```

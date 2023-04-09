package mediainfo

import (
	"encoding/json"
	"log"
)

type Track = map[string]interface{}

type MediaInfo struct {
	CreatingLibrary struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		URL     string `json:"url"`
	} `json:"creatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []Track
	} `json:"media"`
}

func (i MediaInfo) OtherTrack() []Track {
	tracks := []Track{}
	for _, t := range i.Media.Track {
		if t["@type"].(string) == "Other" {
			tracks = append(tracks, t)
		}
	}
	return tracks
}

func (i MediaInfo) General() []GeneralTrack {
	tracks := []GeneralTrack{}
	for _, t := range i.Media.Track {
		if t["@type"].(string) == "General" {
			var track GeneralTrack
			jObj, _ := json.Marshal(t)
			if err := json.Unmarshal(jObj, &track); err != nil {
				log.Printf("convert failed: %v\n", err)
				continue
			}
			tracks = append(tracks, track)
		}
	}
	return tracks
}

func (i MediaInfo) Audio() []AudioTrack {
	tracks := []AudioTrack{}
	for _, t := range i.Media.Track {
		if t["@type"].(string) == "Audio" {
			var track AudioTrack
			jObj, _ := json.Marshal(t)
			if err := json.Unmarshal(jObj, &track); err != nil {
				log.Printf("convert failed: %v\n", err)
				continue
			}
			tracks = append(tracks, track)
		}
	}
	return tracks
}

func (i MediaInfo) Video() []VideoTrack {
	tracks := []VideoTrack{}
	for _, t := range i.Media.Track {
		if t["@type"].(string) == "Video" {
			var track VideoTrack
			jObj, _ := json.Marshal(t)
			if err := json.Unmarshal(jObj, &track); err != nil {
				log.Printf("convert failed: %v\n", err)
				continue
			}
			tracks = append(tracks, track)
		}
	}
	return tracks
}

func (i MediaInfo) Text() []Track {
	tracks := []Track{}
	for _, t := range i.Media.Track {
		if t["@type"].(string) == "Text" {
			tracks = append(tracks, t)
		}
	}
	return tracks
}

func (i MediaInfo) Image() []Track {
	tracks := []Track{}
	for _, t := range i.Media.Track {
		if t["@type"].(string) == "Image" {
			tracks = append(tracks, t)
		}
	}
	return tracks
}

func (i MediaInfo) Menu() []Track {
	tracks := []Track{}
	for _, t := range i.Media.Track {
		if t["@type"].(string) == "Menu" {
			tracks = append(tracks, t)
		}
	}
	return tracks
}

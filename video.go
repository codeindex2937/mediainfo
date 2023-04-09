package mediainfo

type VideoTrack struct {
	StreamKind               string  `json:"StreamKind"`
	StreamKindID             int     `json:"StreamKindID,string"`
	StreamOrder              int     `json:"StreamOrder,string"`
	ID                       int     `json:"ID,string"`
	Format                   string  `json:"Format"`
	FormatInfo               string  `json:"Format_Info"`
	FormatProfile            string  `json:"Format_Profile"`
	FormatLevel              string  `json:"Format_Level"`
	CodecID                  string  `json:"CodecID"`
	CodecIDCompatible        string  `json:"CodecID_Compatible"`
	CodecIDInfo              string  `json:"CodecID_Info"`
	CodecIDHint              string  `json:"CodecID_Hint"`
	Duration                 float64 `json:"Duration,string"`
	DurationString           string  `json:"Duration/String"`
	DurationString1          string  `json:"Duration/String1"`
	DurationString2          string  `json:"Duration/String2"`
	DurationString3          string  `json:"Duration/String3"`
	DurationString4          string  `json:"Duration/String4"`
	DurationString5          string  `json:"Duration/String5"`
	Width                    int     `json:"Width,string"`
	Height                   int     `json:"Height,string"`
	StoredWidth              int     `json:"Stored_Width,string"`
	StoredHeight             int     `json:"Stored_Height,string"`
	SampleAspectRatio        float64 `json:"Sampled_Width,string"`
	DisplayAspectRatioString string  `json:"DisplayAspectRatio/String"`
	DisplayAspectRatioFloat  float64 `json:"DisplayAspectRatio,string"`
	PixelAspectRatioString   string  `json:"PixelAspectRatio/String"`
	PixelAspectRatioFloat    float64 `json:"PixelAspectRatio,string"`
	FrameRateMode            string  `json:"FrameRate_Mode"`
	FrameRate                float64 `json:"FrameRate,string"`
	FrameRateString          string  `json:"FrameRate/String"`
	FrameCount               int     `json:"FrameCount,string"`
	CompressionMode          string  `json:"Compression_Mode"`
	BitsPerPixel             float64 `json:"Bits-(Pixel*Frame),string"`
	StreamSize               int64   `json:"StreamSize,string"`
	StreamSizeString         string  `json:"StreamSize/String"`
	StreamSizeString1        string  `json:"StreamSize/String1"`
	StreamSizeString2        string  `json:"StreamSize/String2"`
	StreamSizeString3        string  `json:"StreamSize/String3"`
	StreamSizeString4        string  `json:"StreamSize/String4"`
	StreamSizeString5        string  `json:"StreamSize/String5"`
	EncodedDate              string  `json:"Encoded_Date"`
	TaggedDate               string  `json:"Tagged_Date"`
	ColorSpace               string  `json:"ColorSpace"`
	ChromaSubsampling        string  `json:"ChromaSubsampling"`
	BitDepth                 int     `json:"BitDepth,string"`
	ScanType                 string  `json:"ScanType"`
	CompressionType          string  `json:"Compression_Type"`
	Language                 string  `json:"Language"`
	Default                  bool    `json:"Default"`
	Forced                   bool    `json:"Forced"`
	Extra                    map[string]string
}

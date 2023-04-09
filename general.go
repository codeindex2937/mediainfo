package mediainfo

type GeneralTrack struct {
	Type                     string  `json:"@type"`
	AudioCount               int64   `json:"AudioCount,string"`
	AudioCodecList           string  `json:"Audio_Codec_List"`
	AudioFormatList          string  `json:"Audio_Format_List"`
	AudioFormatWithHintList  string  `json:"Audio_Format_WithHint_List"`
	VideoCount               int64   `json:"VideoCount,string"`
	VideoCodecList           string  `json:"Video_Codec_List"`
	VideoFormatList          string  `json:"Video_Format_List"`
	VideoFormatWithHintList  string  `json:"Video_Format_WithHint_List"`
	CodecID                  string  `json:"CodecID"`
	CodecIDCompatible        string  `json:"CodecID_Compatible"`
	CodecIDString            string  `json:"CodecID_String"`
	CodecIDUrl               string  `json:"CodecID_Url"`
	CompleteName             string  `json:"CompleteName"`
	Count                    int64   `json:"Count,string"`
	DataSize                 int64   `json:"DataSize,string"`
	Duration                 float64 `json:"Duration,string"`
	DurationString           string  `json:"Duration_String"`
	DurationString1          string  `json:"Duration_String1"`
	DurationString2          string  `json:"Duration_String2"`
	DurationString3          string  `json:"Duration_String3"`
	DurationString4          string  `json:"Duration_String4"`
	DurationString5          string  `json:"Duration_String5"`
	EncodedDate              string  `json:"Encoded_Date"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FileNameExtension        string  `json:"FileNameExtension"`
	FileSize                 int64   `json:"FileSize,string"`
	FileSizeString           string  `json:"FileSize_String"`
	FileSizeString1          string  `json:"FileSize_String1"`
	FileSizeString2          string  `json:"FileSize_String2"`
	FileSizeString3          string  `json:"FileSize_String3"`
	FileSizeString4          string  `json:"FileSize_String4"`
	FileCreatedDate          string  `json:"File_Created_Date"`
	FileCreatedDateLocal     string  `json:"File_Created_Date_Local"`
	FileModifiedDate         string  `json:"File_Modified_Date"`
	FileModifiedDateLocal    string  `json:"File_Modified_Date_Local"`
	FooterSize               int64   `json:"FooterSize,string"`
	Format                   string  `json:"Format"`
	FormatCommercial         string  `json:"Format_Commercial"`
	FormatExtensions         string  `json:"Format_Extensions"`
	FormatProfile            string  `json:"Format_Profile"`
	FormatString             string  `json:"Format_String"`
	FrameCount               int64   `json:"FrameCount,string"`
	FrameRate                string  `json:"FrameRate"`
	FrameRateString          string  `json:"FrameRate_String"`
	HeaderSize               int64   `json:"HeaderSize,string"`
	InternetMediaType        string  `json:"InternetMediaType"`
	IsStreamable             string  `json:"IsStreamable"`
	OverallBitRate           int64   `json:"OverallBitRate,string"`
	OverallBitRateMode       string  `json:"OverallBitRate_Mode"`
	OverallBitRateModeString string  `json:"OverallBitRate_Mode_String"`
	OverallBitRateString     string  `json:"OverallBitRate_String"`
	StreamCount              int64   `json:"StreamCount,string"`
	StreamKind               string  `json:"StreamKind"`
	StreamKindID             int64   `json:"StreamKindID,string"`
	StreamKindString         string  `json:"StreamKind_String"`
	StreamSize               int64   `json:"StreamSize,string"`
	StreamSizeProportion     string  `json:"StreamSize_Proportion"`
	StreamSizeString         string  `json:"StreamSize_String"`
	StreamSizeString1        string  `json:"StreamSize_String1"`
	StreamSizeString2        string  `json:"StreamSize_String2"`
	StreamSizeString3        string  `json:"StreamSize_String3"`
	StreamSizeString4        string  `json:"StreamSize_String4"`
	StreamSizeString5        string  `json:"StreamSize_String5"`
	Extra                    map[string]string
}

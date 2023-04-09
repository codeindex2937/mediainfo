package mediainfo

import (
	"encoding/json"
	"log"
	"regexp"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	libMediaInfo        = windows.NewLazyDLL("MediaInfo.dll")
	procMediaInfoNew    = libMediaInfo.NewProc("MediaInfo_New")
	procMediaInfoDelete = libMediaInfo.NewProc("MediaInfo_Delete")
	procMediaInfoOption = libMediaInfo.NewProc("MediaInfo_Option")
	procMediaInfoOpen   = libMediaInfo.NewProc("MediaInfo_Open")
	procMediaInfoClose  = libMediaInfo.NewProc("MediaInfo_Close")
	procMediaInfoInform = libMediaInfo.NewProc("MediaInfo_Inform")
	versionParser       = regexp.MustCompile(`^MediaInfoLib - v(\S+)`)
)

type MediaInfoHelper struct {
	h uintptr
}

func New() *MediaInfoHelper {
	h, _, err := procMediaInfoNew.Call()
	if err != windows.ERROR_SUCCESS {
		log.Fatal(err)
	}

	return &MediaInfoHelper{
		h: h,
	}
}

func (m *MediaInfoHelper) Delete() {
	if m.h == uintptr(windows.InvalidHandle) {
		return
	}
	procMediaInfoDelete.Call(m.h)
	m.h = uintptr(windows.InvalidHandle)
}

func (m MediaInfoHelper) Option(name, value string) string {
	nameUtf16, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return ""
	}
	valueUtf16, err := windows.UTF16PtrFromString(value)
	if err != nil {
		return ""
	}
	originValue, _, err := procMediaInfoOption.Call(m.h, uintptr(unsafe.Pointer(nameUtf16)), uintptr(unsafe.Pointer(valueUtf16)))
	if err != windows.ERROR_SUCCESS {
		log.Fatal(err)
	}

	return windows.UTF16PtrToString((*uint16)(unsafe.Pointer(originValue)))
}

func (m MediaInfoHelper) Open(path string) bool {
	pathUtf16, err := windows.UTF16PtrFromString(path)
	if err != nil {
		log.Println(err)
		return false
	}
	success, _, err := procMediaInfoOpen.Call(m.h, uintptr(unsafe.Pointer(pathUtf16)))
	if err != windows.ERROR_SUCCESS {
		log.Fatal(err)
	}

	return success != 0
}

func (m MediaInfoHelper) Close() {
	_, _, err := procMediaInfoClose.Call(m.h)
	if err != windows.ERROR_SUCCESS {
		log.Fatal(err)
	}
}

func (m MediaInfoHelper) Version() string {
	versionStr := m.Option("Info_Version", "")
	matched := versionParser.FindStringSubmatch(versionStr)
	if len(matched) == 0 {
		return ""
	}
	return matched[1]
}

func (m MediaInfoHelper) Inform() (MediaInfo, error) {
	ptr, _, err := procMediaInfoInform.Call(m.h)
	if err != windows.ERROR_SUCCESS {
		log.Fatal(err)
	}

	infoStr := windows.UTF16PtrToString((*uint16)(unsafe.Pointer(ptr)))
	var info MediaInfo
	if err := json.Unmarshal([]byte(infoStr), &info); err != nil {
		return MediaInfo{}, err
	}
	return info, nil
}

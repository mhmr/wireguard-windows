// Copyright 2010 The go-winapi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

import (
	"syscall"
	"unsafe"
)

const MAX_PATH = 260

// Error codes
const (
	ERROR_SUCCESS             = 0
	ERROR_FILE_NOT_FOUND      = 2
	ERROR_INVALID_PARAMETER   = 87
	ERROR_INSUFFICIENT_BUFFER = 122
	ERROR_MORE_DATA           = 234
)

// GlobalAlloc flags
const (
	GHND          = 0x0042
	GMEM_FIXED    = 0x0000
	GMEM_MOVEABLE = 0x0002
	GMEM_ZEROINIT = 0x0040
	GPTR          = 0x004
)

// Predefined locale ids
const (
	LOCALE_INVARIANT      = 0x007f
	LOCALE_USER_DEFAULT   = 0x0400
	LOCALE_SYSTEM_DEFAULT = 0x0800
)

var (
	// Library
	libkernel32 uintptr

	// Functions
	closeHandle            uintptr
	fileTimeToSystemTime   uintptr
	getLastError           uintptr
	getLogicalDriveStrings uintptr
	getModuleHandle        uintptr
	getNumberFormat        uintptr
	getThreadLocale        uintptr
	getVersion             uintptr
	globalAlloc            uintptr
	globalFree             uintptr
	globalLock             uintptr
	globalUnlock           uintptr
	moveMemory             uintptr
	mulDiv                 uintptr
	setLastError           uintptr
	systemTimeToFileTime   uintptr
	getProfileString       uintptr
)

type (
	ATOM      uint16
	HANDLE    uintptr
	HGLOBAL   HANDLE
	HINSTANCE HANDLE
	LCID      uint32
)

type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

type NUMBERFMT struct {
	NumDigits     uint32
	LeadingZero   uint32
	Grouping      uint32
	LpDecimalSep  *uint16
	LpThousandSep *uint16
	NegativeOrder uint32
}

type SYSTEMTIME struct {
	WYear         uint16
	WMonth        uint16
	WDayOfWeek    uint16
	WDay          uint16
	WHour         uint16
	WMinute       uint16
	WSecond       uint16
	WMilliseconds uint16
}

func init() {
	// Library
	libkernel32 = MustLoadLibrary("kernel32.dll")

	// Functions
	closeHandle = MustGetProcAddress(libkernel32, "CloseHandle")
	fileTimeToSystemTime = MustGetProcAddress(libkernel32, "FileTimeToSystemTime")
	getLastError = MustGetProcAddress(libkernel32, "GetLastError")
	getLogicalDriveStrings = MustGetProcAddress(libkernel32, "GetLogicalDriveStringsW")
	getModuleHandle = MustGetProcAddress(libkernel32, "GetModuleHandleW")
	getNumberFormat = MustGetProcAddress(libkernel32, "GetNumberFormatW")
	getThreadLocale = MustGetProcAddress(libkernel32, "GetThreadLocale")
	getVersion = MustGetProcAddress(libkernel32, "GetVersion")
	globalAlloc = MustGetProcAddress(libkernel32, "GlobalAlloc")
	globalFree = MustGetProcAddress(libkernel32, "GlobalFree")
	globalLock = MustGetProcAddress(libkernel32, "GlobalLock")
	globalUnlock = MustGetProcAddress(libkernel32, "GlobalUnlock")
	moveMemory = MustGetProcAddress(libkernel32, "RtlMoveMemory")
	mulDiv = MustGetProcAddress(libkernel32, "MulDiv")
	setLastError = MustGetProcAddress(libkernel32, "SetLastError")
	systemTimeToFileTime = MustGetProcAddress(libkernel32, "SystemTimeToFileTime")
	getProfileString = MustGetProcAddress(libkernel32, "GetProfileStringW")
}

func CloseHandle(hObject HANDLE) bool {
	ret, _, _ := syscall.Syscall(closeHandle, 1,
		uintptr(hObject),
		0,
		0)

	return ret != 0
}

func FileTimeToSystemTime(lpFileTime *FILETIME, lpSystemTime *SYSTEMTIME) bool {
	ret, _, _ := syscall.Syscall(fileTimeToSystemTime, 2,
		uintptr(unsafe.Pointer(lpFileTime)),
		uintptr(unsafe.Pointer(lpSystemTime)),
		0)

	return ret != 0
}

func GetLastError() uint32 {
	ret, _, _ := syscall.Syscall(getLastError, 0,
		0,
		0,
		0)

	return uint32(ret)
}

func GetLogicalDriveStrings(nBufferLength uint32, lpBuffer *uint16) uint32 {
	ret, _, _ := syscall.Syscall(getLogicalDriveStrings, 2,
		uintptr(nBufferLength),
		uintptr(unsafe.Pointer(lpBuffer)),
		0)

	return uint32(ret)
}

func GetModuleHandle(lpModuleName *uint16) HINSTANCE {
	ret, _, _ := syscall.Syscall(getModuleHandle, 1,
		uintptr(unsafe.Pointer(lpModuleName)),
		0,
		0)

	return HINSTANCE(ret)
}

func GetNumberFormat(Locale LCID, dwFlags uint32, lpValue *uint16, lpFormat *NUMBERFMT, lpNumberStr *uint16, cchNumber int32) int32 {
	ret, _, _ := syscall.Syscall6(getNumberFormat, 6,
		uintptr(Locale),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpValue)),
		uintptr(unsafe.Pointer(lpFormat)),
		uintptr(unsafe.Pointer(lpNumberStr)),
		uintptr(cchNumber))

	return int32(ret)
}

func GetProfileString(lpAppName, lpKeyName, lpDefault *uint16, lpReturnedString uintptr, nSize uint32) int32 {
	ret, _, _ := syscall.Syscall6(getProfileString, 5,
		uintptr(unsafe.Pointer(lpAppName)),
		uintptr(unsafe.Pointer(lpKeyName)),
		uintptr(unsafe.Pointer(lpDefault)),
		lpReturnedString,
		uintptr(nSize),
		0)
	return int32(ret)
}

func GetThreadLocale() LCID {
	ret, _, _ := syscall.Syscall(getThreadLocale, 0,
		0,
		0,
		0)

	return LCID(ret)
}

func GetVersion() int64 {
	ret, _, _ := syscall.Syscall(getVersion, 0,
		0,
		0,
		0)
	return int64(ret)
}

func GlobalAlloc(uFlags uint32, dwBytes uintptr) HGLOBAL {
	ret, _, _ := syscall.Syscall(globalAlloc, 2,
		uintptr(uFlags),
		dwBytes,
		0)

	return HGLOBAL(ret)
}

func GlobalFree(hMem HGLOBAL) HGLOBAL {
	ret, _, _ := syscall.Syscall(globalFree, 1,
		uintptr(hMem),
		0,
		0)

	return HGLOBAL(ret)
}

func GlobalLock(hMem HGLOBAL) unsafe.Pointer {
	ret, _, _ := syscall.Syscall(globalLock, 1,
		uintptr(hMem),
		0,
		0)

	return unsafe.Pointer(ret)
}

func GlobalUnlock(hMem HGLOBAL) bool {
	ret, _, _ := syscall.Syscall(globalUnlock, 1,
		uintptr(hMem),
		0,
		0)

	return ret != 0
}

func MoveMemory(destination, source unsafe.Pointer, length uintptr) {
	syscall.Syscall(moveMemory, 3,
		uintptr(unsafe.Pointer(destination)),
		uintptr(source),
		uintptr(length))
}

func MulDiv(nNumber, nNumerator, nDenominator int32) int32 {
	ret, _, _ := syscall.Syscall(mulDiv, 3,
		uintptr(nNumber),
		uintptr(nNumerator),
		uintptr(nDenominator))

	return int32(ret)
}

func SetLastError(dwErrorCode uint32) {
	syscall.Syscall(setLastError, 1,
		uintptr(dwErrorCode),
		0,
		0)
}

func SystemTimeToFileTime(lpSystemTime *SYSTEMTIME, lpFileTime *FILETIME) bool {
	ret, _, _ := syscall.Syscall(systemTimeToFileTime, 2,
		uintptr(unsafe.Pointer(lpSystemTime)),
		uintptr(unsafe.Pointer(lpFileTime)),
		0)

	return ret != 0
}

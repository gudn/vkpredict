package store

import "unsafe"

// Конвертировать слайс ID в слайс строк без дополнительной памяти
func IdsToStrigns(ids []ID) []string {
	if len(ids) == 0 {
		return nil
	}
	p := unsafe.Pointer(&(ids[0]))
	return unsafe.Slice((*string)(p), len(ids))
}

// Конвертировать слайс строк в слайс ID без дополнительной памяти
func StringsToIds(strings []string) []ID {
	if len(strings) == 0 {
		return nil
	}
	p := unsafe.Pointer(&(strings[0]))
	return unsafe.Slice((*ID)(p), len(strings))
}

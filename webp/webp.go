package webp

type WEBP struct {
	// the file header

	// the header of WebP. Should always be 'RIFF' for a valid webp file.
	Header [4]byte
	// the total size in bytes of the file after this field and the 4 byte FourCC.
	FileSize uint32
	// the unique character code for WebP. Should always be 'WEBP' for a valid webp file.
	FourCC [4]byte

	// the content of the file

}

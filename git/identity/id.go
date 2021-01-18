package identity

// HID 表示一个 HASH-ID
type HID interface {
	ToString() string
	ToBytes() []byte
	Equals(other HID) bool
}

// ObjectId 表示一个git 对象的ID
type ObjectId interface {
	HID
}

// PackId 表示一个 git 对象包的ID
type PackId interface {
	HID
}

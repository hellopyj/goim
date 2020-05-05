package msg

type Content interface {
	String()string
	Bytes()[]byte
}

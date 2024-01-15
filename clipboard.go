package clipboard

import "golang.design/x/clipboard"

func Push(str string) {
	PushBytes([]byte(str))
}

func Top() string {
	return string(TopBytes())
}

func PushBytes(data []byte) {
	initialize()

	clipboard.Write(clipboard.FmtText, data)
}

func TopBytes() []byte {
	return clipboard.Read(clipboard.FmtText)
}

func initialize() {
	err := clipboard.Init()

	if err != nil {
		panic(err)
	}
}

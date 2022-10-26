package lang

var LangMessage = map[string]string{}

func NewLang(language string) map[string]string {
	var StubStorage = map[string]func(){
		"ID": ID,
		"EN": EN,
	}
	CallLang(language, StubStorage[language])
	return Id
}

func CallLang(function string, f func()) {
	f()
}

func ID() {
	LangMessage = Id
}

func EN() {
	LangMessage = En
}

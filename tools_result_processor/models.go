package tools_result_processor

const TscanCodeName = "TscanCode"
const FlawFinderName = "FlawFinder"

type CPPCheckResult []struct {
	File     string `csv:"file"`
	Line     int    `csv:"line"`
	Severity string `csv:"severity"`
	ID       string `csv:"id"`      // 故障名称
	Message  string `csv:"message"` // 故障描述
}
type TScancodeResult struct {
	Name   string `xml:"name,attr"` // 绝对路径
	Errors []struct {
		ID       string `xml:"id,attr"`
		SubID    string `xml:"subid,attr"`
		Line     int    `xml:"line,attr"`
		File     string `xml:"file,attr"`
		Severity string `xml:"severity,attr"`
		Msg      string `xml:"msg,attr"`
	} `xml:"error"`
}

type FlawFinderResult []struct {
	File     string `csv:"File"`
	Line     int    `csv:"Line"`
	Column   int    `csv:"Column"`
	Level    string `csv:"Level"`
	Category string `csv:"Category"`
	Name     string `csv:"Name"`
	Warning  string `csv:"Warning"` // 故障描述
	CWEs     string `csv:"CWEs"`
}

type StandardError struct {
	FileName    string
	Start       int
	End         int
	SourcePath  string
	Description string
	Type        string
	Priority    string
}

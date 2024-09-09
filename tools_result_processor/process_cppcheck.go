package tools_result_processor

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/gocarina/gocsv"
)

// 标准化数据
func standardizeCPPCheckResult(projectDir string, result CPPCheckResult) []StandardError {
	standardized := make([]StandardError, 0)
	for i := 0; i < len(result); i++ {
		standardized = append(standardized, StandardError{
			FileName:    path.Base(result[i].File),
			Start:       result[i].Line,
			End:         result[i].Line,
			SourcePath:  GetRelPath(result[i].File, projectDir),
			Description: result[i].Message,
			Type:        result[i].ID,
			Priority:    result[i].Severity,
		})
	}
	return standardized
}

// 解析CPPCheck的数据
func ParseCppCheckResult(projectDir, resultFile, output string) error {

	formattedResultFile := resultFile + ".formatted"
	err := stripCSVLast(resultFile, formattedResultFile)
	if err != nil {
		log.Fatal(err)
	}
	content, err := os.ReadFile(formattedResultFile)
	if err != nil {
		log.Fatal(err)
	}
	v := CPPCheckResult{}
	err = gocsv.UnmarshalBytes([]byte(content), &v)
	if err != nil {

		fmt.Println("Output: \n", string(content))
		return fmt.Errorf("error: %v", err)
	}
	standardized := standardizeCPPCheckResult(projectDir, v)
	FormatStandardResultJson(standardized, output)
	return nil
}

package tools_result_processor

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gocarina/gocsv"
)

// 标准化数据
func standardizeFlawFinderResult(projectDir string, result FlawFinderResult) []StandardError {
	standardized := make([]StandardError, 0)

	for i := 0; i < len(result); i++ {
		result[i].File = strings.TrimLeft(result[i].File, "./")
		standardized = append(standardized, StandardError{
			FileName:    path.Base(result[i].File),
			Start:       result[i].Line,
			End:         result[i].Line,
			SourcePath:  GetRelPath(result[i].File, projectDir),
			Description: result[i].Warning,
			Type:        result[i].Category,
			Priority:    result[i].Level,
		})
	}
	return standardized
}

// 解析CPPCheck的数据
func ParseFlawFinderResult(projectDir, csvFile, output string) error {
	content, err := os.ReadFile(csvFile)

	if err != nil {
		log.Fatal(err)
	}
	v := FlawFinderResult{}
	err = gocsv.UnmarshalBytes([]byte(content), &v)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	standardized := standardizeFlawFinderResult(projectDir, v)
	FormatStandardResultJson(standardized, output)
	return nil
}

package tools_result_processor

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// 标准化数据
func standardizeTscancodeResult(resultStruct TScancodeResult, sourceCodeDir string) []StandardError {
	standardized := make([]StandardError, 0)
	violations := resultStruct.Errors
	for i := 0; i < len(violations); i++ {
		violation := violations[i]

		if !strings.HasPrefix(violation.File, sourceCodeDir) {
			violation.File = strings.TrimPrefix(violation.File, "/private")
			if !strings.HasPrefix(violation.File, sourceCodeDir) {
				panic(fmt.Sprintf("violation file path %s does not match source code directory %s", violation.File, sourceCodeDir))
			}
		}

		relpath, err := filepath.Rel(sourceCodeDir, violation.File)
		if err != nil {
			log.Println(err.Error())
		}
		standardized = append(standardized, StandardError{
			FileName:    path.Base(violation.File),
			Start:       violation.Line,
			End:         violation.Line,
			SourcePath:  relpath,
			Description: violation.Msg,
			Type:        violation.ID,
			Priority:    violation.Severity,
		})

	}

	return standardized
}

func ParseTscancodeResult(sourceCodeDir, xmlFileName, output string) error {
	content, err := os.ReadFile(xmlFileName)

	if err != nil {
		log.Fatal(err)
	}
	v := TScancodeResult{}
	err = xml.Unmarshal([]byte(content), &v)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	standardized := standardizeTscancodeResult(v, sourceCodeDir)
	FormatStandardResultJson(standardized, output)
	return nil
}

// func tscancodeCallBack(ctx context.Context, testcase parameter.Set, m iactor.IMailbox) error {
// 	// sourceCode := static_analysis_tool_utils.SourceCodeArg.ToMeta()
// 	// err := testcase.Decode(sourceCode.Name, &sourceCode)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// formatCommand := func(analysisDir string, resultFileDir string) string {
// 	// 	return fmt.Sprintf(`/usr/bin/tscancode --xml %s 2> %s .`, analysisDir, resultFileDir)
// 	// }
// 	postProcess := func(projectDir, rawResultFileDir, formattedResultFileDir string) error {
// 		return ParseTscancodeResult(projectDir, rawResultFileDir, formattedResultFileDir)
// 	}
// 	config := static_analysis_tool_utils.RoutineConfig{
// 		Toolname:         "tscancode",
// 		ResultFileFormat: "xml",
// 		SuccessCode:      []int{0},
// 		IgnoreError:      false,
// 	}
// 	static_analysis_tool_utils.StaticToolAnalysisRoutine(string(sourceCode.Value), formatCommand, postProcess, m, config)
// 	return nil
// }

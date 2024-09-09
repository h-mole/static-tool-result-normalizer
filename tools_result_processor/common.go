package tools_result_processor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func stripCSVLast(csvFileName string, targetCSVFileName string) error {
	fTarget, err := os.Create(targetCSVFileName)
	if err != nil {
		return fmt.Errorf("create map file error: %v\n", err)
	}
	defer fTarget.Close()
	fTemp, err := os.Open(csvFileName)
	if err != nil {
		return err
	}
	defer fTemp.Close()

	rd := bufio.NewReader(fTemp)
	w := bufio.NewWriter(fTarget)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		strs := strings.SplitN(line, ",", 5)
		if len(strs) >= 5 {
			s := strings.ReplaceAll(strs[4], "'", "")
			s = strings.ReplaceAll(s, "\"", "")
			s = strings.ReplaceAll(s, ",", "")
			s = strings.TrimSpace(s)
			strs[4] = "\"" + s + "\""
		}
		fmt.Fprintln(w, strings.Join(strs, ","))
	}
	w.Flush()
	return nil
}

func FormatStandardResultJson(standardErrors []StandardError, fileName string) error {

	m := make(map[string]interface{})
	for i := 0; i < len(standardErrors); i++ {

		standardError := standardErrors[i]
		if standardError.FileName == "" {
			continue
		}
		if _, stat := m[standardError.SourcePath]; !stat {
			list := make([]interface{}, 0)
			list = append(list, path.Dir(standardError.SourcePath))
			list = append(list, []string{
				strconv.Itoa(standardError.Start),
				standardError.Priority,
				standardError.Description,
			})
			m[standardError.SourcePath] = list
		} else {
			list, ok := m[standardError.SourcePath].([]interface{})
			if !ok {
				fmt.Println("Get Elem failed")
			}
			list = append(list, []string{
				strconv.Itoa(standardError.Start),
				standardError.Priority,
				standardError.Description,
			})

			m[standardError.SourcePath] = list
		}

	}

	filePtr, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("文件创建失败: %s", err.Error())
	}
	defer filePtr.Close()

	jsonWithIndent, _ := json.MarshalIndent(m, "", "    ")
	filePtr.Write(jsonWithIndent)
	return nil
}

// Get Relative Path
// For mac, remove prefix /private
func GetRelPath(absPath, projectRoot string) string {
	sourceFilePath := strings.TrimPrefix(absPath, "/private")
	sourceFileRelPath, err := filepath.Rel(projectRoot, sourceFilePath)
	if err != nil {
		fmt.Println("Err occured in GetRelPath:", err)
	}
	return sourceFileRelPath
}

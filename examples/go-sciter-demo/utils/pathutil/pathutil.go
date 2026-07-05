package pathutil

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// 返回路径父路径,带斜杠
func GetParentPath(pathStr string) string {
	if !strings.Contains(pathStr, "/") {
		return pathStr
	}
	return pathStr[:strings.LastIndex(pathStr, "/")] + "/"
}

// 获取指定路径的文件名称
func GetFileName(pathStr string) string {
	_, filename := path.Split(pathStr)
	return filename
}

// 返回路径最后一级
func GetPathEndName(path string) string {
	if !strings.Contains(path, "/") {
		return path
	}
	return path[strings.LastIndex(path, "/")+1:]
}

// 返回路径最后是否带有斜杠，若最后不带斜杠则添加斜杠
func GetPathWithEnd(path string) string {
	if strings.HasSuffix(path, "/") {
		return path
	}
	return path + "/"
}

// 返回当前执行程序目录，最后不带斜杠
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err == nil {
		//将\替换成/
		return strings.Replace(dir, "\\", "/", -1)
	}
	return ""
}

// 判断文件或者文件夹是否存在（linux下同一目录不能同时存在名称相同的文件和文件夹）
func IsPathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func IsFileExist(filePath string) bool {
	_, err := os.Lstat(filePath)
	return !os.IsNotExist(err)
}

func MakeDir(filePath string) error {
	if !IsPathExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

// 判断所给路径是否为文件夹
func IsDir(filePath string) bool {
	s, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(filePath string) bool {
	return !IsDir(filePath)
}

func RemoveSpecial(filePath string, ext string) error {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return err
	}
	for _, file := range files {
		fileName := path.Join(filePath, file.Name())
		if file.IsDir() {
			err = RemoveSpecial(fileName, ext)
		} else {
			if path.Ext(file.Name()) == ext {
				err = os.Remove(fileName)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// 查找目录下的指定特征文件：Find("/app","*.apk")
func Find(filePath string, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(filePath, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(p)); err != nil {
			return err
		} else if matched {
			matches = append(matches, p)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return matches, nil
}

func ListFiles(sourceDir string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(sourceDir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过目录和非常规文件
		if info.IsDir() || !info.Mode().IsRegular() {
			return nil
		}
		files = append(files, p)
		return nil
	})
	return files, err
}

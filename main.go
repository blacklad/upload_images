package main

import (
	"blacklad.com/upload_images/conf"
	"blacklad.com/upload_images/yun"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	args := flag.Args()

	ossUrls := make([]string, 0, len(args))
	for _, filePath := range args {
		if isExist(filePath) {
			fileUrl, err := uploadFile(filePath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			ossUrls = append(ossUrls, fileUrl)
		}
	}

	fmt.Println("Upload Success:")
	for _, url := range ossUrls {
		fmt.Println(url)
	}
}

func uploadFile(filePath string) (string, error) {
	userDir, err := user.Current()
	if err != nil {
		return  "", err
	}

	config, err := conf.GetConf(filepath.Join(userDir.HomeDir, "/.upload.conf"))
	if err != nil {
		return  "", err
	}
	ossClient, err := yun.NewOssAli(config)
	if err != nil {
		return  "", err
	}
	fileName := uploadFileName()
	err = ossClient.UploadFile(fileName, filePath)
	if err != nil {
		return  "", err
	}

	fileUrl := getOssFileUrl(fileName, config)
	return fileUrl, nil
}

func getOssFileUrl(fileName string, conf *conf.Config) string {
	return fmt.Sprintf("%s%s", conf.OssDomain, fileName)
}

func uploadFileName() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(9000)+1000
	return time.Now().Format("20060102150405") + strconv.Itoa(n) + ".png"
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
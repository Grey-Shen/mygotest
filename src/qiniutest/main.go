package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

var (
	domain = "ofmufa5nc.bkt.clouddn.com"
	// keys         = []string{"test", "test1", "gopl-zh.pdf", "ssss"}
	keys         = []string{"test", "test1"}
	ch           chan string
	allfilenanes []string
	client       = http.Client{}
)

const donwloadPath = "/tmp/mydata"

var childPath string

func main() {

	conf.ACCESS_KEY = "C2QcWBwdgB5XIPgexk9LoAJp4uJoW0yaHywpqP0q"
	conf.SECRET_KEY = "v-RAMVOIuLKBGCFemR1omfSAGlc8NgGxqTudEnPQ"
	ch = make(chan string, 2)
	// var chnum = len(keys)
	fmt.Println("home is ", os.Getenv("HOME"))
	// var filename string
	tmpPath := path.Join(os.Getenv("HOME"), donwloadPath)
	os.MkdirAll(tmpPath, 0777)
	go downloadUrl(domain, keys)

	time.Sleep(10 * time.Second)
}

func downloadUrl(domain string, keys []string) {
	fmt.Println("begin to download")

	for num, key := range keys {

		baseUrl := kodo.MakeBaseUrl(domain, key)
		fmt.Println("baseurl is-------", baseUrl)
		policy := kodo.GetPolicy{}
		c := kodo.New(0, nil)
		privateUrl := c.MakePrivateUrl(baseUrl, &policy)
		fmt.Println("privateurl is---", privateUrl)
		fmt.Println("begin to mkdir")
		childPath = path.Join(donwloadPath, key)
		if err := os.MkdirAll(childPath, os.ModePerm); err != nil {
			fmt.Println("mkdir err: ", err)
			return
			// }
		}
		fmt.Println("url is ", privateUrl)
		if err := downloadFile(privateUrl, key, num); err != nil {
			fmt.Println("downfile faild :", err)
		}
	}
	return
}

func writeIntoFile(filename string, content []byte) error {
	// filename = filepath.Join(childPath, filename)
	fmt.Println("begin write file :", filename)
	fi, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fi.Close()
	fi.Write(content)
	return nil

}

func downloadFile(url string, key string, num int) error {
	var (
		err         error
		content     []byte
		downloadURL string
	)

	request, err := http.NewRequest("GET", url, nil)
	// fmt.Println("url is ", downloadURL)

	resp, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("faild to download url: %s:%s", downloadURL, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("http err code is ", resp.StatusCode)
		return fmt.Errorf("http error code is %d", resp.StatusCode)
	}
	content, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("read body faild : %s", err)
	}

	fileName := filepath.Join(childPath, key+"_"+strconv.Itoa(num))

	fmt.Println("filename is :", fileName)

	if err = writeIntoFile(fileName, content); err != nil {
		return err
	}
	fmt.Println("finsh to down file :", key)
	return nil
}

//创建压缩文件
func CreateTar(filename string, files []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var fileWrite io.WriteCloser = file

	if strings.HasSuffix(filename, ".gz") {
		fileWrite = gzip.NewWriter(file)
		defer fileWrite.Close()
	}

	writer := tar.NewWriter(fileWrite)
	defer writer.Close()

	for _, name := range files {
		if err := writeFileToTar(writer, name); err != nil {
			return err
		}
	}
	return nil
}

//把文件写入压缩包
func writeFileToTar(writer *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil
	}
	header := &tar.Header{
		Name:    filename,
		Mode:    int64(stat.Mode()),
		Uid:     os.Getuid(),
		Gid:     os.Getegid(),
		ModTime: stat.ModTime(),
		Size:    stat.Size(),
	}
	if err := writer.WriteHeader(header); err != nil {
		return err
	}
	var size int64
	if size, err = io.Copy(writer, file); err != nil {
		fmt.Println("copy err:", err)
		return err
	} else {
		fmt.Println("copy size is:", size)
		return nil
	}
}

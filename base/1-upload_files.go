package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"

	"github.com/kaleidodev/dify-sdk-go/types"
)

func (c *AppClient) UploadFile(filePath string, f *os.File, user string) (info types.FileInfo, err error) {
	if user == "" {
		user = c.GetUser()
	}

	var file *os.File

	// 确定使用哪个文件对象
	if f != nil {
		file = f
	} else {
		file, err = os.Open(filePath)
		if err != nil {
			err = fmt.Errorf("failed to open file: %w", err)
			return
		}
		defer file.Close()
	}

	// 获取文件名
	var fileName string
	if f != nil {
		fileInfo, err := f.Stat()
		if err != nil {
			return info, fmt.Errorf("failed to get file info: %w", err)
		}
		fileName = fileInfo.Name()
	} else {
		fileName = filepath.Base(filePath)
	}

	// 检测文件的 MIME 类型
	// 先通过文件扩展名判断
	ext := filepath.Ext(fileName)
	mimeType := mime.TypeByExtension(ext)

	// 如果通过扩展名无法判断，则读取文件头部进行判断
	if mimeType == "" {
		// 创建buffer用于存储文件头部数据
		buffer := make([]byte, 512)
		_, err := file.Read(buffer)
		if err != nil {
			return info, fmt.Errorf("failed to read file header: %w", err)
		}

		// 重置文件指针到开始位置
		_, err = file.Seek(0, 0)
		if err != nil {
			return info, fmt.Errorf("failed to reset file pointer: %w", err)
		}

		// 通过文件头部数据检测 MIME 类型
		mimeType = http.DetectContentType(buffer)
	}

	if !strings.Contains(fileName, ".") {
		ext := strings.Split(mimeType, "/")
		if len(ext) == 2 {
			fileName += "." + ext[1]
		}
	}

	// 创建multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加userId字段
	err = writer.WriteField("user", user)
	if err != nil {
		err = fmt.Errorf("failed to write user field: %w", err)
		return
	}

	// 创建form file部分，设置正确的Content-Type
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileName))
	header.Set("Content-Type", mimeType)

	part, err := writer.CreatePart(header)
	if err != nil {
		err = fmt.Errorf("failed to create form file: %w", err)
		return
	}

	// 复制文件内容
	_, err = io.Copy(part, file)
	if err != nil {
		err = fmt.Errorf("failed to copy file content: %w", err)
		return
	}

	// 关闭writer
	err = writer.Close()
	if err != nil {
		err = fmt.Errorf("failed to close writer: %w", err)
		return
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", c.apiServer+"/files/upload", body)
	if err != nil {
		err = fmt.Errorf("failed to create request: %w", err)
		return
	}

	// 设置Content-Type header
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	// 发送请求
	resp, err := (*Client)(c).HttpClient().SendRequest(req)
	if err != nil {
		err = fmt.Errorf("failed to send request: %w", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		var errbody errBody
		err = json.NewDecoder(resp.Body).Decode(&errbody)
		if err != nil {
			return info, err
		}
		return info, fmt.Errorf("ERROR: %d [%v] %v", errbody.Status, errbody.Code, errbody.Message)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed read body: %w", err)
		return
	}

	err = json.Unmarshal(respBody, &info)

	return
}

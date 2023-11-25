package service

import (
	"TikTokLite/config"
	"TikTokLite/log"
	"TikTokLite/minioStore"
	message "TikTokLite/proto/pkg"
	"TikTokLite/repository"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func PublishVideo(saveFile, filename string) (*message.DouyinPublishActionResponse, error) {
	client := minioStore.GetMinio()
	/*	videourl, err := client.UploadFile("video", saveFile, strconv.FormatInt(userId, 10))
		if err != nil {
			return nil, err
	}*/
	ctx := context.Background()
	fileInfo, err := os.Stat(saveFile)
	if err == os.ErrNotExist { //这这是在对错误进行处理
		log.Info("%s目标文件不存在", saveFile)
	}
	f, err := os.Open(saveFile)
	contentType := "multipart/form-data"
	videourl, err := client.MinioClient.PutObject(ctx, "video", filename, f, fileInfo.Size(), minio.PutObjectOptions{
		ContentType: contentType})
	logrus.Info(videourl)
	//imageFile, err := GetImageFile(saveFile) //用户获取封面图像的路径
	//logrus.Info(imageFile)
	if err != nil {
		return nil, err
	}

	//log.Debugf("imageFile %v\n", imageFile)

	/*	picurl, err := client.UploadFile("pic", imageFile, strconv.FormatInt(userId, 10))
		if err != nil {
			picurl = "https://github.com/Diligenthader"
		}
		err = repository.InsertVideo(userId, videourl, picurl, title)
		if err != nil {
			return nil, err
		}*/
	return &message.DouyinPublishActionResponse{}, nil
}

func PublishList(tokenUserId, userId int64) (*message.DouyinPublishListResponse, error) {
	videos, err := repository.GetVideoList(userId)
	if err != nil {
		return nil, err
	}
	list := &message.DouyinPublishListResponse{
		VideoList: VideoList(videos, tokenUserId),
	}

	return list, nil
}

func GetImageFile(videoPath string) (string, error) {
	temp := strings.Split(videoPath, "/")
	videoName := temp[len(temp)-1]
	b := []byte(videoName)
	videoName = string(b[:len(b)-3]) + "png"
	picpath := config.GetConfig().Path.Picfile
	//然后，函数使用config.GetConfig().Path.Picfile获取图片文件的存储路径(picpath)，并将其与图片文件名拼接起来，得到图片文件的完整路径(picName)。

	picName := filepath.Join(picpath, videoName)                                                            //这段代码是将图片路径和视频路径合成一个返回路径.
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "1", "-f", "image2", "-t", "0.01", "-y", picName) //函数使用ffmpeg命令从视频文件中提取一帧作为图片。
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	log.Debugf(picName)
	return picName, nil
}

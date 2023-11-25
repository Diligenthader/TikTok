package minioStore

import (
	"TikTokLite/config"
	"TikTokLite/log"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	MinioClient  *minio.Client
	endpoint     string
	port         string
	VideoBuckets string
	PicBuckets   string
}

var client Minio

func GetMinio() Minio {
	return client //返回一个Minio对象.
}

func InitMinio() {
	conf := config.GetConfig()
	endpoint := conf.Minio.Host
	port := conf.Minio.Port
	endpoint = endpoint + ":" + port
	accessKeyID := conf.Minio.AccessKeyID
	secretAccessKey := conf.Minio.SecretAccessKey
	videoBucket := conf.Minio.Videobuckets
	picBucket := conf.Minio.Picbuckets

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Error(err)
	}
	//创建存储桶
	//ToDo:无法进行视频的传输和和在本地创建桶以及修改密码等操作.
	/*	creatBucket(minioClient, videoBucket)
		creatBucket(minioClient, picBucket)*/
	ctx := context.Background()
	//ctx2,err:=context.WithDeadline(context.Background(),time.Now().Add(3*time.Hour))用于设置过期时间
	// 创建这个bucket
	err = minioClient.MakeBucket(ctx, "video", minio.MakeBucketOptions{})
	if err != nil {
		// 检测这个bucket是否已经存在
		exists, errBucketExists := minioClient.BucketExists(ctx, "video")
		if errBucketExists == nil && exists {
			log.Info("We already own %s\n", "video")
		} else {
			log.Info(err)
		}
	} else {
		log.Info("Successfully created %s\n", "video")
	}
	err = minioClient.MakeBucket(ctx, "pic", minio.MakeBucketOptions{})
	if err != nil {
		// 检测这个bucket是否已经存在
		exists, errBucketExists := minioClient.BucketExists(ctx, "pic")
		if errBucketExists == nil && exists {
			log.Info("We already own %s\n", "pic")
		} else {
			log.Info(err)
		}
	} else {
		log.Info("Successfully created %s\n", "pic")
	}
	client = Minio{minioClient, endpoint, port, videoBucket, picBucket}
	//对定义的的结构体进行赋值
}

/*func creatBucket(m *minio.Client, bucket string) {
	// log.Debug("bucketname", bucket)
	found, err := m.BucketExists(bucket)
	if err != nil {
		log.Errorf("check %s bucketExists err:%s", bucket, err.Error())
	}
	if !found {
		m.MakeBucket(bucket, "us-east-1")
	}
	//设置桶策略
	policy := `{"Version": "2012-10-17",
				"Statement":
					[{
						"Action":["s3:GetObject"],
						"Effect": "Allow",
						"Principal": {"AWS": ["*"]},
						"Resource": ["arn:aws:s3:::` + bucket + `/*"],
						"Sid": ""
					}]
				}`
	err = m.SetBucketPolicy(bucket, policy)
	if err != nil {
		log.Errorf("SetBucketPolicy %s  err:%s", bucket, err.Error())
	}
}*/

/*func (m *Minio) UploadFile(filetype, file, userID string) (string, error) {
	var fileName strings.Builder
	var contentType, bucket string
	if filetype == "video" {
		contentType = "multipart/form-data"
		//Suffix = ".mp4"
		bucket = m.VideoBuckets
	} else {
		contentType = "multipart/form-data"
		//Suffix = ".png"
		bucket = m.PicBuckets
	}
	fileName.WriteString(userID)
	fileName.WriteString("_")
	fileName.WriteString(strconv.FormatInt(util.GetCurrentTime(), 10))
	//进制转换为10进制形式
	//fileName.WriteString(Suffix)
	//ToDo :这里是用于将解析文件的格式，并且用于返回
	n, err := m.MinioClient.FPutObject(bucket, fileName.String(), file, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Errorf("upload file error:%s", err.Error())
		return "", err
	}
	log.Infof("upload file %dbyte success,fileName:%s", n, fileName)
	url := "http://" + m.endpoint + "/" + bucket + "/" + fileName.String()
	return url, nil
}*/

package constants

import (
	"github.com/spf13/viper"
	"net"
	"os"
	"strings"
)

var (
	MySQLDSN             string
	RedisAddr            string
	RedisPassword        string
	EtcdAddress          string
	MinioEndPoint        string
	MinioAccessKeyID     string
	MinioAccessSecretKey string

	SecretKey   string
	IdentityKey string

	VideoFeedCount int
	UserNameMaxLen int
	PassWordMaxLen int
	PasswordMinLen int

	UserTableName     string
	RelationTableName string
	MessageTableName  string
	PublishTableName  string
	FavoriteTableName string
	CommentTableName  string

	MinioVideoBucketName      string
	MinioImageBucketName      string
	MinioAvatarBucketName     string
	MinioBackgroundBucketName string

	UserServiceName     string
	RelationServiceName string
	PublishServiceName  string
	MessageServiceName  string
	FavoriteServiceName string
	CommentServiceName  string

	LimitsPerSecond int

	TestAva        string
	TestBackground string

	RabbitMqURI string
)

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // 配置文件所在的路径

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	MySQLDSN = os.Getenv("MYSQL_DSN")
	if MySQLDSN == "" {
		MySQLDSN = viper.GetString("MySQLDSN")
	}

	RedisAddr = os.Getenv("REDIS_ADDR")
	if RedisAddr == "" {
		RedisAddr = viper.GetString("RedisAddr")
	}
	RedisPassword = viper.GetString("RedisPassword")

	EtcdAddress = os.Getenv("ETCD_ADDR")
	if EtcdAddress == "" {
		EtcdAddress = viper.GetString("EtcdAddress")
	}

	MinioEndPoint = os.Getenv("MINIO_ENDPOINT")
	if MinioEndPoint == "" {
		MinioEndPoint = viper.GetString("MinioEndPoint")
	}
	MinioAccessKeyID = viper.GetString("MinioAccessKeyID")
	MinioAccessSecretKey = viper.GetString("MinioAccessSecretKey")

	SecretKey = viper.GetString("SecretKey")
	IdentityKey = viper.GetString("IdentityKey")

	VideoFeedCount = viper.GetInt("VideoFeedCount")
	UserNameMaxLen = viper.GetInt("UserNameMaxLen")
	PassWordMaxLen = viper.GetInt("PassWordMaxLen")

	PasswordMinLen = viper.GetInt("PasswordMinLen")

	UserTableName = viper.GetString("UserTableName")
	RelationTableName = viper.GetString("RelationTableName")
	MessageTableName = viper.GetString("MessageTableName")
	PublishTableName = viper.GetString("PublishTableName")
	FavoriteTableName = viper.GetString("FavoriteTableName")
	CommentTableName = viper.GetString("CommentTableName")

	MinioVideoBucketName = viper.GetString("MinioVideoBucketName")
	MinioImageBucketName = viper.GetString("MinioImageBucketName")
	MinioAvatarBucketName = viper.GetString("MinioAvatarBucketName")
	MinioBackgroundBucketName = viper.GetString("MinioBackgroundBucketName")

	UserServiceName = viper.GetString("UserServiceName")
	RelationServiceName = viper.GetString("RelationServiceName")
	PublishServiceName = viper.GetString("PublishServiceName")
	MessageServiceName = viper.GetString("MessageServiceName")
	FavoriteServiceName = viper.GetString("FavoriteServiceName")
	CommentServiceName = viper.GetString("CommentServiceName")

	TestAva = viper.GetString("TestAva")
	TestBackground = viper.GetString("TestBackground")

	RabbitMqURI = os.Getenv("RABBIT_MQ_URI")
	if RabbitMqURI == "" {
		RabbitMqURI = viper.GetString("RabbitMqURI")
	}

	LimitsPerSecond = viper.GetInt("LimitsPerSecond")

	return nil
}

func GetOutBoundIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "", err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return strings.Split(localAddr.String(), ":")[0], nil
}

func init() {
	err := LoadConfig()
	if err != nil {
		panic(err)
	}
}

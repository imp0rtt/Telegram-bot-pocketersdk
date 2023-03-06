package repository

type Bucket string

const (
	AccessTokens  Bucket = "Access tokens"
	RequestTokens Bucket = "Request tokens"
)

type TokenRepository interface {
	Save(chatID int64, token string, bucket Bucket) error
	Get(chaID int64, bucket Bucket) (string, error)
}

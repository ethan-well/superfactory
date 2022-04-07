package uploader

import (
	"context"
	"github.com/ItsWewin/superfactory/aerror"
)

type Uploader interface {
	UploadBase64(ctx context.Context, base64str string, fileName string) (string, aerror.Error)
}

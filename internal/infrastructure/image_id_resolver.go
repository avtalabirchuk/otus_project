package infrastructure

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image-previewer/internal/domain"
	"image-previewer/internal/domain/valueObjects"
)

type ImageIDResolver struct {
}

func (r *ImageIDResolver) ResolveImageID(url string, dim valueObjects.ImageDimensions) domain.ImageID {
	hash := md5.Sum([]byte(url))
	ImageID := fmt.Sprintf("%s_%dx%d", hex.EncodeToString(hash[:]), dim.Width, dim.Height)

	return domain.ImageID(ImageID)
}

func NewImageIDResolver() *ImageIDResolver {
	return &ImageIDResolver{}
}

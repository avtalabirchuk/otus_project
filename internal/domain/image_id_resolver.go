package domain

import "image-previewer/internal/domain/valueObjects"

type ImageID string

type ImageIDResolver interface {
	ResolveImageID(url string, dim valueObjects.ImageDimensions) ImageID
}

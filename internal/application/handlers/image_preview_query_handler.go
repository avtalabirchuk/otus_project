package handlers

import (
	"errors"
	"image"
	"image-previewer/internal/application/queries"
	"image-previewer/internal/domain"
	"net/url"

	"go.uber.org/zap"
)

var ErrInvalidWidth = errors.New("width should be greater than 0")
var ErrInvalidHeight = errors.New("height should be greater than 0")
var ErrEmptyUrl = errors.New("url should not be empty")
var ErrInvalidUrl = errors.New("url should be valid")
var ErrNotFound = errors.New("img not found")

type imagePreviewQueryHandler struct {
	previewRepository domain.PreviewRepository
	downloader        domain.Downloader
	idResolver        domain.ImageIDResolver
}

func (h *imagePreviewQueryHandler) Handle(q queries.ImagePreviewQuery) (image.Image, error) {
	if err := h.checkQuery(q); err != nil {
		return nil, err
	}

	ImageID := h.idResolver.ResolveImageID(q.Url, q.Dimensions)

	zap.S().Debugf("started processing image %s", string(ImageID))

	img, err := h.previewRepository.FindOne(ImageID)

	if err != nil {
		if err == ErrNotFound {
			zap.S().Debug("not found in cache, downloading")

			img, err = h.downloader.Download(q.Url, q.Dimensions)

			if err != nil {
				return nil, err
			}

			zap.S().Debug("adding to repository")

			_, err = h.previewRepository.Add(ImageID, img)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		zap.S().Debug("using image from cache")
	}

	return img, nil
}

func (h *imagePreviewQueryHandler) checkQuery(q queries.ImagePreviewQuery) error {
	if q.Dimensions.Width < 1 {
		return ErrInvalidWidth
	}

	if q.Dimensions.Height < 1 {
		return ErrInvalidHeight
	}

	if q.Url == "" {
		return ErrEmptyUrl
	}

	if _, err := url.Parse(q.Url); err != nil {
		return ErrInvalidUrl
	}

	return nil
}

func NewImagePreviewQueryHandler(
	rep domain.PreviewRepository,
	downloader domain.Downloader,
	resolver domain.ImageIDResolver,
) *imagePreviewQueryHandler {
	return &imagePreviewQueryHandler{
		previewRepository: rep,
		downloader:        downloader,
		idResolver:        resolver,
	}
}

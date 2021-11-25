package infrastructure

import (
	"image-previewer/internal/domain"
	"image-previewer/internal/domain/valueObjects"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImageIDResolver_ResolveImageID(t *testing.T) {
	t.Run("resolve id should return valid string", func(t *testing.T) {
		actualId := NewImageIDResolver().ResolveImageID(
			"http://ya.ru/test.jpg",
			valueObjects.ImageDimensions{
				Width:  100,
				Height: 500,
			},
		)

		require.Equal(t, domain.ImageID("9508dfb97b74094e1b8134e15469fc0e_100x500"), actualId)
	})
}

package avatar_usecase

import (
	"bytes"
	"image"
	"image/png"
	"main/internal/pkg/avatar"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAvatar_Success(t *testing.T) {
	useCase := NewDefault()

	const mockUserId = uint64(1)

	width, height := 16, 9
	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	mockSrc := new(bytes.Buffer)
	err := png.Encode(mockSrc, img)
	if err != nil {
		t.Error(err)
	}

	mockSrcSize := int64(mockSrc.Len())
	_, err = useCase.GetAvatar(mockUserId, mockSrc, mockSrcSize)
	assert.NoError(t, err)
}
func TestGetAvatar_Failed(t *testing.T) {
	useCase := NewDefault()

	const mockUserId = uint64(1)

	mockSrc := new(bytes.Buffer)
	t.Run("too large file", func(t *testing.T) {
		_, err := useCase.GetAvatar(mockUserId, mockSrc, avatar.MaxAvatarSize+1)
		assert.EqualError(t, err, avatar.ErrAvatarIsTooLarge.Error())
	})

	t.Run("wrong avatar type", func(t *testing.T) {
		_, err := useCase.GetAvatar(mockUserId, mockSrc, int64(mockSrc.Len()))
		assert.EqualError(t, err, avatar.ErrWrongAvatarType.Error())
	})

}

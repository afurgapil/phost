package image

import (
	"errors"
	"reflect"
	"testing"

	"github.com/afurgapil/phost/backend/pkg/entities"
)

type MockRepository struct {
	images map[int]entities.Image
	nextID int
}

func (m *MockRepository) CreateImage(image entities.Image) (entities.Image, error) {
	m.nextID++
	image.ID = m.nextID
	m.images[image.ID] = image
	return image, nil
}

func (m *MockRepository) GetImageByID(id int) (entities.Image, error) {
	image, exists := m.images[id]
	if !exists {
		return entities.Image{}, errors.New("image not found")
	}
	return image, nil
}

func (m *MockRepository) DeleteImage(id int) error {
	_, exists := m.images[id]
	if !exists {
		return errors.New("image not found")
	}
	delete(m.images, id)
	return nil
}

func Test_service_CreateImage(t *testing.T) {
	mockRepo := &MockRepository{
		images: make(map[int]entities.Image),
	}

	type fields struct {
		repo Repository
	}
	type args struct {
		image entities.Image
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Image
		wantErr bool
	}{
		{
			name: "Valid Image Creation",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				image: entities.Image{Value: "test_image"},
			},
			want:    entities.Image{ID: 1, Value: "test_image"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.CreateImage(tt.args.image)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want.ID = got.ID
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetImageByID(t *testing.T) {
	mockRepo := &MockRepository{
		images: make(map[int]entities.Image),
	}
	createdImage, _ := mockRepo.CreateImage(entities.Image{Value: "test_image"})

	type fields struct {
		repo Repository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Image
		wantErr bool
	}{
		{
			name: "Valid Image Retrieval",
			fields: fields{
				repo: mockRepo,
			},
			args:    args{id: createdImage.ID},
			want:    createdImage,
			wantErr: false,
		},
		{
			name: "Invalid Image Retrieval - Non-existing ID",
			fields: fields{
				repo: mockRepo,
			},
			args:    args{id: 999},
			want:    entities.Image{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetImageByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetImageByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetImageByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteImage(t *testing.T) {
	mockRepo := &MockRepository{
		images: make(map[int]entities.Image),
	}
	createdImage, _ := mockRepo.CreateImage(entities.Image{Value: "test_image_to_delete"})

	type fields struct {
		repo Repository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Valid Image Deletion",
			fields: fields{
				repo: mockRepo,
			},
			args:    args{id: createdImage.ID},
			wantErr: false,
		},
		{
			name: "Invalid Image Deletion - Non-existing ID",
			fields: fields{
				repo: mockRepo,
			},
			args:    args{id: 999},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.DeleteImage(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

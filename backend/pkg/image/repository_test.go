package image

import (
	"net/http"
	"testing"

	"github.com/afurgapil/phost/backend/internal/config"
	"github.com/afurgapil/phost/backend/pkg/entities"
)

var baseURL string

func init() {
	var loadErr error
	baseURL, loadErr = config.LoadConfig("BASE_URL")
	if loadErr != nil || baseURL == "" {
		baseURL = "http://localhost:8080"
	}
}

func Test_repository_CreateImage(t *testing.T) {
	type fields struct {
		baseURL string
		client  *http.Client
	}
	type args struct {
		image entities.Image
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Valid Image Creation",
			fields: fields{
				baseURL: baseURL,
				client:  http.DefaultClient,
			},
			args: args{
				image: entities.Image{Value: "test_image"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				baseURL: tt.fields.baseURL,
				client:  tt.fields.client,
			}
			got, err := r.CreateImage(tt.args.image)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.ID <= 0 {
				t.Errorf("repository.CreateImage() = ID %v, expected positive integer", got.ID)
			}

			if err := r.DeleteImage(got.ID); err != nil {
				t.Errorf("repository.DeleteImage() error = %v", err)
			}
		})
	}
}

func Test_repository_GetImageByID(t *testing.T) {
	type fields struct {
		baseURL string
		client  *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid Image Retrieval",
			fields: fields{
				baseURL: baseURL,
				client:  http.DefaultClient,
			},
			wantErr: false,
		},
		{
			name: "Invalid Image Retrieval - Non-existing ID",
			fields: fields{
				baseURL: baseURL,
				client:  http.DefaultClient,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				baseURL: tt.fields.baseURL,
				client:  tt.fields.client,
			}

			if tt.name == "Valid Image Retrieval" {
				createdImage, err := r.CreateImage(entities.Image{Value: "test_image"})
				if err != nil {
					t.Fatalf("Failed to create image for test: %v", err)
				}

				got, err := r.GetImageByID(createdImage.ID)
				if (err != nil) != tt.wantErr {
					t.Errorf("repository.GetImageByID() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if got.Value != createdImage.Value {
					t.Errorf("repository.GetImageByID() = %v, want %v", got, createdImage)
				}

				if err := r.DeleteImage(createdImage.ID); err != nil {
					t.Errorf("Failed to delete image after retrieval: %v", err)
				}
			} else if tt.name == "Invalid Image Retrieval - Non-existing ID" {
				_, err := r.GetImageByID(999)
				if (err != nil) != tt.wantErr {
					t.Errorf("repository.GetImageByID() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
func Test_repository_DeleteImage(t *testing.T) {
	type fields struct {
		baseURL string
		client  *http.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid Image Deletion",
			fields: fields{
				baseURL: baseURL,
				client:  http.DefaultClient,
			},
			wantErr: false,
		},
		{
			name: "Invalid Image Deletion - Non-existing ID",
			fields: fields{
				baseURL: baseURL,
				client:  http.DefaultClient,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				baseURL: tt.fields.baseURL,
				client:  tt.fields.client,
			}

			if tt.name == "Valid Image Deletion" {
				createdImage, err := r.CreateImage(entities.Image{Value: "test_image_to_delete"})
				if err != nil {
					t.Fatalf("Failed to create image for deletion test: %v", err)
				}

				err = r.DeleteImage(createdImage.ID)
				if (err != nil) != tt.wantErr {
					t.Errorf("repository.DeleteImage() error = %v, wantErr %v", err, tt.wantErr)
				}

				_, err = r.GetImageByID(createdImage.ID)
				if err == nil {
					t.Errorf("repository.GetImageByID() should have failed for deleted ID %d", createdImage.ID)
				}
			} else if tt.name == "Invalid Image Deletion - Non-existing ID" {
				err := r.DeleteImage(9999999999)
				if (err != nil) != tt.wantErr {
					t.Errorf("repository.DeleteImage() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

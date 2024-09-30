package image

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/afurgapil/phost/backend/pkg/entities"
)

type Repository interface {
	CreateImage(image entities.Image) (entities.Image, error)
	GetImageByID(id int) (entities.Image, error)
	DeleteImage(id int) error
}

type repository struct {
	baseURL string
	client  *http.Client
}

func NewRepository(baseURL string) Repository {
	return &repository{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (r *repository) CreateImage(image entities.Image) (entities.Image, error) {
	command := fmt.Sprintf("INSERT INTO records VALUES ('%s')", image.Value)
	url := fmt.Sprintf("%s/execute?command=%s", r.baseURL, url.QueryEscape(command))
	resp, err := r.client.Get(url)
	if err != nil {
		return entities.Image{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entities.Image{}, fmt.Errorf("failed to execute command: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.Image{}, err
	}

	var createdImage entities.Image
	err = json.Unmarshal(body, &createdImage)
	if err != nil {
		return entities.Image{}, err
	}

	return createdImage, nil
}

func (r *repository) GetImageByID(id int) (entities.Image, error) {
	command := fmt.Sprintf("SELECT * FROM records WHERE id=%d", id)
	url := fmt.Sprintf("%s/execute?command=%s", r.baseURL, url.QueryEscape(command))
	resp, err := r.client.Get(url)
	if err != nil {
		return entities.Image{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entities.Image{}, fmt.Errorf("failed to execute command: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.Image{}, err
	}

	var images []entities.Image
	err = json.Unmarshal(body, &images)
	if err != nil {
		return entities.Image{}, err
	}

	if len(images) == 0 {
		return entities.Image{}, fmt.Errorf("no image found with ID %d", id)
	}

	return images[0], nil
}

func (r *repository) DeleteImage(id int) error {
	command := fmt.Sprintf("DELETE FROM records WHERE id=%d", id)
	url := fmt.Sprintf("%s/execute?command=%s", r.baseURL, url.QueryEscape(command))
	resp, err := r.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to execute command: %s", resp.Status)
	}

	return nil
}

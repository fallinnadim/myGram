package services

import (
	request "fp2/data/request/photo"
	response "fp2/data/response/photo"
	repository "fp2/repository/photo"
	"time"

	"github.com/go-playground/validator/v10"
)

type PhotoServiceImpl struct {
	PhotoRepository repository.PhotoRepository
	Validate        *validator.Validate
}

// Delete implements PhotoService.
func (*PhotoServiceImpl) Delete(id int) error {
	panic("unimplemented")
}

// GetAll implements PhotoService.
func (*PhotoServiceImpl) GetAll() []response.AllPhotoResponse {
	panic("unimplemented")
}

// Post implements PhotoService.
func (p *PhotoServiceImpl) Post(cp request.CreatePhotoRequest) (response.CreatedPhotoResponse, error) {
	// Validasi Struct
	errValidation := p.Validate.Struct(cp)
	if errValidation != nil {
		return response.CreatedPhotoResponse{}, errValidation
	}
	cp.Created_At = time.Now().Format("2006-01-02")
	cp.Updated_At = time.Now().Format("2006-01-02")
	// Panggil Repository
	result := p.PhotoRepository.Create(cp)
	// Return
	resp := response.CreatedPhotoResponse{
		Id:         result.Id,
		Title:      result.Title,
		Caption:    result.Caption,
		Photo_Url:  result.Photo_Url,
		User_Id:    result.User_Id,
		Created_At: result.Created_At,
	}
	return resp, nil
}

// Update implements PhotoService.
func (*PhotoServiceImpl) Update(p request.UpdatePhotoRequest) (response.UpdatedPhotoResponse, error) {
	panic("unimplemented")
}

func NewPhotoServiceImpl(p repository.PhotoRepository, v *validator.Validate) PhotoService {
	return &PhotoServiceImpl{
		PhotoRepository: p,
		Validate:        v,
	}
}

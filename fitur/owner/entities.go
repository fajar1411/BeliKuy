package owner

import "mime/multipart"

type OwnerEntities struct {
	ID        uint
	Nama_Toko string `validate:"required,min=5,required"`
	Alamat    string `validate:"required,min=5,required"`
	Ktp       string
	Status    string
	UserID    uint
}

type OwnerService interface {
	AddOwner(ktpfile *multipart.FileHeader, newOwner OwnerEntities) (OwnerEntities, error)
}

type OwnerData interface {
	AddOwner(newOwner OwnerEntities) (OwnerEntities, error)
}

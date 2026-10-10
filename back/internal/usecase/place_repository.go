package usecase

import placesusecase "naevis/internal/places/usecase"

// PlaceRepository is the compatibility export for the canonical place repository contract.
type PlaceRepository = placesusecase.PlaceRepository

// PlaceUseCase is the compatibility export for the canonical place use case contract.
type PlaceUseCase = placesusecase.PlaceUseCase

var NewPlaceUseCase = placesusecase.NewPlaceUseCase

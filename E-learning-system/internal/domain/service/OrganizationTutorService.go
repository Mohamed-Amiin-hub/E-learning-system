package service

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// OrganizationTutorService interface with CRUD methods
type OrganizationTutorService interface {
	CreateTutor(tutor *model.OrganizationTutor) (*model.OrganizationTutor, error)
	UpdateTutor(tutor *model.OrganizationTutor) error
	DeleteTutor(tutorID uuid.UUID) error
	GetTutorByID(tutorID uuid.UUID) (*model.OrganizationTutor, error)
	GetAllTutors() ([]*model.OrganizationTutor, error)
}

// organizationTutorServiceImpl struct implementing OrganizationTutorService
type organizationTutorServiceImpl struct {
	repo repository.OrganizationTutorRepository
}

// Constructor
func NewOrganizationTutorService(tutorRepo repository.OrganizationTutorRepository) OrganizationTutorService {
	return &organizationTutorServiceImpl{
		repo: tutorRepo,
	}
}

// CreateTutor creates a new tutor
func (s *organizationTutorServiceImpl) CreateTutor(tutor *model.OrganizationTutor) (*model.OrganizationTutor, error) {
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}

	tutor.ID = newID
	tutor.CreatedAt = time.Now()

	log.Printf("Creating organization tutor: %+v", tutor)

	if err := s.repo.Create(tutor); err != nil {
		return nil, fmt.Errorf("failed to create tutor: %v", err)
	}

	return tutor, nil
}

// UpdateTutor updates an existing tutor
func (s *organizationTutorServiceImpl) UpdateTutor(tutor *model.OrganizationTutor) error {
	// Check if tutor exists
	_, err := s.repo.GetByID(tutor.ID)
	if err != nil {
		return fmt.Errorf("tutor not found with ID %s: %v", tutor.ID, err)
	}

	if err := s.repo.Update(tutor); err != nil {
		return fmt.Errorf("failed to update tutor with ID %s: %v", tutor.ID, err)
	}

	log.Printf("Organization tutor updated: %+v", tutor)
	return nil
}

// DeleteTutor performs a soft delete
func (s *organizationTutorServiceImpl) DeleteTutor(tutorID uuid.UUID) error {
	// Check if tutor exists
	_, err := s.repo.GetByID(tutorID)
	if err != nil {
		return fmt.Errorf("tutor not found with ID %s: %v", tutorID, err)
	}

	if err := s.repo.Delete(tutorID); err != nil {
		return fmt.Errorf("failed to delete tutor with ID %s: %v", tutorID, err)
	}

	log.Printf("Organization tutor deleted: %v", tutorID)
	return nil
}

// GetTutorByID retrieves a single tutor by ID
func (s *organizationTutorServiceImpl) GetTutorByID(tutorID uuid.UUID) (*model.OrganizationTutor, error) {
	tutor, err := s.repo.GetByID(tutorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tutor by ID %s: %v", tutorID, err)
	}
	return tutor, nil
}

// GetAllTutors retrieves all organization tutors
func (s *organizationTutorServiceImpl) GetAllTutors() ([]*model.OrganizationTutor, error) {
	tutors, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all tutors: %v", err)
	}
	return tutors, nil
}

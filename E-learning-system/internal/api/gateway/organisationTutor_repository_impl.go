package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

type OrganizationTutorRepositoryImpl struct {
	db *sql.DB
}

// Create inserts a new tutor using the stored procedure
func (r *OrganizationTutorRepositoryImpl) Create(tutor *model.OrganizationTutor) error {
	_, err := r.db.Exec(`CALL create_organization_tutor($1, $2, $3)`,
		tutor.UserID, tutor.OrganizationID, tutor.Approved,
	)
	if err != nil {
		log.Printf("Error calling create_organization_tutor: %v", err)
		return err
	}

	log.Printf("OrganizationTutor created: %+v", tutor)
	return nil
}

// Update modifies an existing tutor using the stored procedure
func (r *OrganizationTutorRepositoryImpl) Update(tutor *model.OrganizationTutor) error {
	_, err := r.db.Exec(`CALL update_organization_tutor($1, $2)`,
		tutor.ID, tutor.Approved,
	)
	if err != nil {
		log.Printf("Error calling update_organization_tutor: %v", err)
		return err
	}

	log.Printf("OrganizationTutor updated: %+v", tutor)
	return nil
}

// Delete performs a soft delete of a tutor using the stored procedure
func (r *OrganizationTutorRepositoryImpl) Delete(tutorID uuid.UUID) error {
	_, err := r.db.Exec(`CALL delete_organization_tutor($1)`, tutorID)
	if err != nil {
		log.Printf("Error calling delete_organization_tutor for ID %v: %v", tutorID, err)
		return err
	}

	log.Printf("OrganizationTutor soft-deleted: %v", tutorID)
	return nil
}

// GetAll retrieves all tutors using the stored function
func (r *OrganizationTutorRepositoryImpl) GetAll() ([]*model.OrganizationTutor, error) {
	rows, err := r.db.Query(`SELECT * FROM get_all_organization_tutors()`)
	if err != nil {
		log.Printf("Error querying get_all_organization_tutors: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tutors []*model.OrganizationTutor
	for rows.Next() {
		var t model.OrganizationTutor
		err := rows.Scan(
			&t.ID,
			&t.UserID,
			&t.OrganizationID,
			&t.Approved,
			&t.CreatedAt,
		)
		if err != nil {
			log.Printf("Error scanning tutor row: %v", err)
			return nil, err
		}
		tutors = append(tutors, &t)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		return nil, err
	}

	log.Printf("OrganizationTutors retrieved: %d", len(tutors))
	return tutors, nil
}

// GetByID retrieves a single tutor by ID using the stored function
func (r *OrganizationTutorRepositoryImpl) GetByID(tutorID uuid.UUID) (*model.OrganizationTutor, error) {
	var t model.OrganizationTutor

	row := r.db.QueryRow(`SELECT * FROM get_organization_tutor_by_id($1)`, tutorID)
	err := row.Scan(
		&t.ID,
		&t.UserID,
		&t.OrganizationID,
		&t.Approved,
		&t.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("OrganizationTutor not found with ID: %v", tutorID)
			return nil, fmt.Errorf("organization tutor not found")
		}
		log.Printf("Error scanning tutor by ID: %v", err)
		return nil, err
	}

	log.Printf("OrganizationTutor retrieved by ID: %+v", t)
	return &t, nil
}

// Constructor
func NewOrganizationTutorRepository(db *sql.DB) repository.OrganizationTutorRepository {
	return &OrganizationTutorRepositoryImpl{db: db}
}

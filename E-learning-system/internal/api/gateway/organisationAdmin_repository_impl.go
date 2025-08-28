package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

type OrganizationAdminRepositoryImpl struct {
	db *sql.DB
}

// Create inserts a new admin using the stored procedure
func (r *OrganizationAdminRepositoryImpl) Create(admin *model.OrganizationAdmin) error {
	_, err := r.db.Exec(`CALL create_organization_admin($1, $2, $3)`,
		admin.UserID, admin.OrganizationID, admin.Role,
	)
	if err != nil {
		log.Printf("Error calling create_organization_admin: %v", err)
		return err
	}

	log.Printf("OrganizationAdmin created: %+v", admin)
	return nil
}

// Update modifies an existing admin using the stored procedure
func (r *OrganizationAdminRepositoryImpl) Update(admin *model.OrganizationAdmin) error {
	_, err := r.db.Exec(`CALL update_organization_admin($1, $2)`,
		admin.ID, admin.Role,
	)
	if err != nil {
		log.Printf("Error calling update_organization_admin: %v", err)
		return err
	}

	log.Printf("OrganizationAdmin updated: %+v", admin)
	return nil
}

// Delete performs a soft delete of an admin using the stored procedure
func (r *OrganizationAdminRepositoryImpl) Delete(adminID uuid.UUID) error {
	_, err := r.db.Exec(`CALL delete_organization_admin($1)`, adminID)
	if err != nil {
		log.Printf("Error calling delete_organization_admin for ID %v: %v", adminID, err)
		return err
	}

	log.Printf("OrganizationAdmin soft-deleted: %v", adminID)
	return nil
}

// GetAll retrieves all admins using the stored function
func (r *OrganizationAdminRepositoryImpl) GetAll() ([]*model.OrganizationAdmin, error) {
	rows, err := r.db.Query(`SELECT * FROM get_all_organization_admins()`)
	if err != nil {
		log.Printf("Error querying get_all_organization_admins: %v", err)
		return nil, err
	}
	defer rows.Close()

	var admins []*model.OrganizationAdmin
	for rows.Next() {
		var a model.OrganizationAdmin
		err := rows.Scan(
			&a.ID,
			&a.UserID,
			&a.OrganizationID,
			&a.Role,
			&a.CreatedAt,
		)
		if err != nil {
			log.Printf("Error scanning admin row: %v", err)
			return nil, err
		}
		admins = append(admins, &a)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		return nil, err
	}

	log.Printf("OrganizationAdmins retrieved: %d", len(admins))
	return admins, nil
}

// GetByID retrieves a single admin by ID using the stored function
func (r *OrganizationAdminRepositoryImpl) GetByID(adminID uuid.UUID) (*model.OrganizationAdmin, error) {
	var a model.OrganizationAdmin

	row := r.db.QueryRow(`SELECT * FROM get_organization_admin_by_id($1)`, adminID)
	err := row.Scan(
		&a.ID,
		&a.UserID,
		&a.OrganizationID,
		&a.Role,
		&a.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("OrganizationAdmin not found with ID: %v", adminID)
			return nil, fmt.Errorf("organization admin not found")
		}
		log.Printf("Error scanning admin by ID: %v", err)
		return nil, err
	}

	log.Printf("OrganizationAdmin retrieved by ID: %+v", a)
	return &a, nil
}

// Constructor
func NewOrganizationAdminRepository(db *sql.DB) repository.OrganizationAdminRepository {
	return &OrganizationAdminRepositoryImpl{db: db}
}

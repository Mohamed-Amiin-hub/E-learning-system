package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

type AdminRepositoryImpl struct {
	db *sql.DB
}

// Delete implements repository.AdminRepository.
func (r *AdminRepositoryImpl) Delete(AdminID uuid.UUID) error {
	//Delete performs a soft delete of an admin
	_, err := r.db.Exec(`Call Delete Admin($1)`, AdminID)
	if err != nil {
		log.Printf("error calling delete_admin for ID %v: %v", AdminID, err)
		return err
	}
	log.Printf("admin soft_deleted: %v", AdminID)
	return nil
}

// GetAll implements repository.AdminRepository.
func (r *AdminRepositoryImpl) GetAll() ([]*model.Admin, error) {
	rows, err := r.db.Query(`SELECT * FROM get_all_admins()`)
	if err != nil {
		log.Printf("Error querying get_all_admins: %v", err)
		return nil, err
	}
	defer rows.Close()

	var admins []*model.Admin
	for rows.Next() {
		var admin model.Admin
		err := rows.Scan(
			&admin.ID,
			&admin.UserID,
			&admin.OrganizationID,
			&admin.Role,
			&admin.Permission,
			&admin.Status,
			&admin.LastLogin,
			&admin.CreatedAt,
			&admin.UpdatedAt,
			&admin.DeletedAt,
		)
		if err != nil {
			log.Printf("Error scanning admin row: %v", err)
			return nil, err
		}
		admins = append(admins, &admin)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		return nil, err
	}

	log.Printf("Admins retrieved: %d", len(admins))
	return admins, nil
}

// GetByID implements repository.AdminRepository.
func (r *AdminRepositoryImpl) GetByID(AdminID uuid.UUID) (*model.Admin, error) {

	var admin model.Admin

	row := r.db.QueryRow(`SELECT * FROM get_admin_by_id($1)`, AdminID)
	err := row.Scan(
		&admin.ID,
		&admin.UserID,
		&admin.OrganizationID,
		&admin.Role,
		&admin.Permission,
		&admin.Status,
		&admin.LastLogin,
		&admin.CreatedAt,
		&admin.UpdatedAt,
		&admin.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Admin not found with ID: %v", AdminID)
			return nil, fmt.Errorf("admin not found")
		}
		log.Printf("Error scanning admin by ID: %v", err)
		return nil, err
	}

	log.Printf("Admin retrieved by ID: %+v", admin)
	return &admin, nil
	} 
	
// Create inserts a new admin using the stored procedure
func (r AdminRepositoryImpl) Create(admin *model.Admin) error {
	_, err := r.db.Exec(`CALL create_admin($1,$2, $3, $4, $5)`,
		admin.UserID, admin.OrganizationID, admin.Role, admin.Permission, admin.Status,
	)
	if err != nil {
		log.Printf("Error calling create_admin: %v", err)
		return err
	}

	log.Printf("Admin created: %+v", admin)
	return nil

}

// Update modifies an existing admin using the stored procedure
func (r AdminRepositoryImpl) Update(admin *model.Admin) error {
	_, err := r.db.Exec(`call update admin($1, $2, $3, $4, $5, $6)`,
		admin.ID, admin.OrganizationID, admin.Role, admin.Permission, admin.Status, admin.LastLogin)
	if err != nil {
		log.Printf("Error calling update_admin: %v", err)
		return err
	}

	log.Printf("Admin updated: %+v", admin)
	return nil
}

// Constructor
func NewAdminRepository(db *sql.DB) repository.AdminRepository {
	return &AdminRepositoryImpl{db: db}
}

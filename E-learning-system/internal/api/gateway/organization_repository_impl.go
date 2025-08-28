package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

type OrganizationRepositoryImpl struct {
	db *sql.DB
}

// Create inserts a new organization using the stored procedure
func (r *OrganizationRepositoryImpl) Create(org *model.Organization) error {
	_, err := r.db.Exec(`CALL create_organization($1,$2,$3,$4,$5,$6,$7,$8)`,
		org.Name, org.Description, org.LogoURL,
		org.PrimaryColor, org.SecondaryColor, org.Domain,
		org.Status, org.Plan,
	)
	if err != nil {
		log.Printf("Error calling create_organization: %v", err)
		return err
	}

	log.Printf("Organization created: %+v", org)
	return nil
}

// Update modifies an existing organization using the stored procedure
func (r *OrganizationRepositoryImpl) Update(org *model.Organization) error {
	_, err := r.db.Exec(`CALL update_organization($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		org.ID, org.Name, org.Description, org.LogoURL,
		org.PrimaryColor, org.SecondaryColor, org.Domain,
		org.Status, org.Plan,
	)
	if err != nil {
		log.Printf("Error calling update_organization: %v", err)
		return err
	}

	log.Printf("Organization updated: %+v", org)
	return nil
}

// Delete performs a soft delete of an organization using the stored procedure
func (r *OrganizationRepositoryImpl) Delete(orgID uuid.UUID) error {
	_, err := r.db.Exec(`CALL delete_organization($1)`, orgID)
	if err != nil {
		log.Printf("Error calling delete_organization for ID %v: %v", orgID, err)
		return err
	}

	log.Printf("Organization soft-deleted: %v", orgID)
	return nil
}

// GetAll retrieves all active organizations using the stored function
func (r *OrganizationRepositoryImpl) GetAll() ([]*model.Organization, error) {
	rows, err := r.db.Query(`SELECT * FROM get_all_organizations()`)
	if err != nil {
		log.Printf("Error querying get_all_organizations: %v", err)
		return nil, err
	}
	defer rows.Close()

	var orgs []*model.Organization
	for rows.Next() {
		var org model.Organization
		err := rows.Scan(
			&org.ID,
			&org.Name,
			&org.Description,
			&org.LogoURL,
			&org.PrimaryColor,
			&org.SecondaryColor,
			&org.Domain,
			&org.Status,
			&org.Plan,
			&org.CreatedAt,
			&org.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning organization row: %v", err)
			return nil, err
		}
		orgs = append(orgs, &org)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		return nil, err
	}

	log.Printf("Organizations retrieved: %d", len(orgs))
	return orgs, nil
}

// GetByID retrieves a single organization by ID using the stored function
func (r *OrganizationRepositoryImpl) GetByID(orgID uuid.UUID) (*model.Organization, error) {
	var org model.Organization

	row := r.db.QueryRow(`SELECT * FROM get_organization_by_id($1)`, orgID)

	err := row.Scan(
		&org.ID,
		&org.Name,
		&org.Description,
		&org.LogoURL,
		&org.PrimaryColor,
		&org.SecondaryColor,
		&org.Domain,
		&org.Status,
		&org.Plan,
		&org.CreatedAt,
		&org.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Organization not found with ID: %v", orgID)
			return nil, fmt.Errorf("organization not found")
		}
		log.Printf("Error scanning organization by ID: %v", err)
		return nil, err
	}

	log.Printf("Organization retrieved by ID: %+v", org)
	return &org, nil
}

// Constructor
func NewOrganizationRepository(db *sql.DB) repository.OrganizationRepository {
	return &OrganizationRepositoryImpl{db: db}
}

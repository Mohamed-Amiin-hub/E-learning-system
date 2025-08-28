package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"log"
	"fmt"

	"github.com/gofrs/uuid"
)

type OrganizationBrandingRepositoryImpl struct {
	db *sql.DB
}

// Create inserts a new branding record using the stored procedure
func (r *OrganizationBrandingRepositoryImpl) Create(branding *model.OrganizationBranding) error {
	_, err := r.db.Exec(`CALL create_organization_branding($1,$2,$3,$4,$5,$6)`,
		branding.OrganizationID,
		branding.LogoURL,
		branding.PrimaryColor,
		branding.SecondaryColor,
		branding.Theme,
		branding.EmailTemplate,
	)
	if err != nil {
		log.Printf("Error calling create_organization_branding: %v", err)
		return err
	}

	log.Printf("OrganizationBranding created: %+v", branding)
	return nil
}

// Update modifies an existing branding record using the stored procedure
func (r *OrganizationBrandingRepositoryImpl) Update(branding *model.OrganizationBranding) error {
	_, err := r.db.Exec(`CALL update_organization_branding($1,$2,$3,$4,$5,$6)`,
		branding.OrganizationID,
		branding.LogoURL,
		branding.PrimaryColor,
		branding.SecondaryColor,
		branding.Theme,
		branding.EmailTemplate,
	)
	if err != nil {
		log.Printf("Error calling update_organization_branding: %v", err)
		return err
	}

	log.Printf("OrganizationBranding updated: %+v", branding)
	return nil
}

// Delete performs a soft delete of a branding record using the stored procedure
func (r *OrganizationBrandingRepositoryImpl) Delete(orgID uuid.UUID) error {
	_, err := r.db.Exec(`CALL delete_organization_branding($1)`, orgID)
	if err != nil {
		log.Printf("Error calling delete_organization_branding for ID %v: %v", orgID, err)
		return err
	}

	log.Printf("OrganizationBranding soft-deleted for organization: %v", orgID)
	return nil
}

// GetAll retrieves all branding records using the stored function
func (r *OrganizationBrandingRepositoryImpl) GetAll() ([]*model.OrganizationBranding, error) {
	rows, err := r.db.Query(`SELECT * FROM get_all_organization_brandings()`)
	if err != nil {
		log.Printf("Error querying get_all_organization_brandings: %v", err)
		return nil, err
	}
	defer rows.Close()

	var brandings []*model.OrganizationBranding
	for rows.Next() {
		var b model.OrganizationBranding
		err := rows.Scan(
			&b.ID,
			&b.OrganizationID,
			&b.LogoURL,
			&b.PrimaryColor,
			&b.SecondaryColor,
			&b.Theme,
			&b.EmailTemplate,
			&b.CreatedAt,
			&b.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning branding row: %v", err)
			return nil, err
		}
		brandings = append(brandings, &b)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		return nil, err
	}

	log.Printf("OrganizationBrandings retrieved: %d", len(brandings))
	return brandings, nil
}

// GetByID retrieves a single branding record by organization ID using the stored function
func (r *OrganizationBrandingRepositoryImpl) GetByID(orgID uuid.UUID) (*model.OrganizationBranding, error) {
	var b model.OrganizationBranding

	row := r.db.QueryRow(`SELECT * FROM get_organization_branding_by_id($1)`, orgID)
	err := row.Scan(
		&b.ID,
		&b.OrganizationID,
		&b.LogoURL,
		&b.PrimaryColor,
		&b.SecondaryColor,
		&b.Theme,
		&b.EmailTemplate,
		&b.CreatedAt,
		&b.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("OrganizationBranding not found for organization ID: %v", orgID)
			return nil, fmt.Errorf("organization branding not found")
		}
		log.Printf("Error scanning branding by organization ID: %v", err)
		return nil, err
	}

	log.Printf("OrganizationBranding retrieved for organization ID: %+v", b)
	return &b, nil
}

// Constructor
func NewOrganizationBrandingRepository(db *sql.DB) repository.OrganizationBrandingRepository {
	return &OrganizationBrandingRepositoryImpl{db: db}
}

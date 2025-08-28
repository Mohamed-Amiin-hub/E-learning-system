package gateway

import (
	"database/sql"
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"log"
	"fmt"

	"github.com/gofrs/uuid"
)

type OrganizationBillingRepositoryImpl struct {
	db *sql.DB
}

// Create inserts a new billing record using the stored procedure
func (r *OrganizationBillingRepositoryImpl) Create(billing *model.OrganizationBilling) error {
	_, err := r.db.Exec(`CALL create_organization_billing($1,$2,$3,$4)`,
		billing.OrganizationID,
		billing.Plan,
		billing.PaymentMethod,
		billing.SubscriptionID,
	)
	if err != nil {
		log.Printf("Error calling create_organization_billing: %v", err)
		return err
	}

	log.Printf("OrganizationBilling created: %+v", billing)
	return nil
}

// Update modifies an existing billing record using the stored procedure
func (r *OrganizationBillingRepositoryImpl) Update(billing *model.OrganizationBilling) error {
	_, err := r.db.Exec(`CALL update_organization_billing($1,$2,$3,$4,$5)`,
		billing.ID,
		billing.Plan,
		billing.PaymentMethod,
		billing.SubscriptionID,
		billing.NextBillingAt,
	)
	if err != nil {
		log.Printf("Error calling update_organization_billing: %v", err)
		return err
	}

	log.Printf("OrganizationBilling updated: %+v", billing)
	return nil
}

// Delete performs a soft delete of a billing record using the stored procedure
func (r *OrganizationBillingRepositoryImpl) Delete(billingID uuid.UUID) error {
	_, err := r.db.Exec(`CALL delete_organization_billing($1)`, billingID)
	if err != nil {
		log.Printf("Error calling delete_organization_billing for ID %v: %v", billingID, err)
		return err
	}

	log.Printf("OrganizationBilling soft-deleted: %v", billingID)
	return nil
}

// GetAll retrieves all billing records using the stored function
func (r *OrganizationBillingRepositoryImpl) GetAll() ([]*model.OrganizationBilling, error) {
	rows, err := r.db.Query(`SELECT * FROM get_all_organization_billings()`)
	if err != nil {
		log.Printf("Error querying get_all_organization_billings: %v", err)
		return nil, err
	}
	defer rows.Close()

	var billings []*model.OrganizationBilling
	for rows.Next() {
		var b model.OrganizationBilling
		err := rows.Scan(
			&b.ID,
			&b.OrganizationID,
			&b.Plan,
			&b.PaymentMethod,
			&b.SubscriptionID,
			&b.NextBillingAt,
			&b.CreatedAt,
			&b.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning billing row: %v", err)
			return nil, err
		}
		billings = append(billings, &b)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
		return nil, err
	}

	log.Printf("OrganizationBillings retrieved: %d", len(billings))
	return billings, nil
}

// GetByID retrieves a single billing record by ID using the stored function
func (r *OrganizationBillingRepositoryImpl) GetByID(billingID uuid.UUID) (*model.OrganizationBilling, error) {
	var b model.OrganizationBilling

	row := r.db.QueryRow(`SELECT * FROM get_organization_billing_by_id($1)`, billingID)
	err := row.Scan(
		&b.ID,
		&b.OrganizationID,
		&b.Plan,
		&b.PaymentMethod,
		&b.SubscriptionID,
		&b.NextBillingAt,
		&b.CreatedAt,
		&b.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("OrganizationBilling not found with ID: %v", billingID)
			return nil, fmt.Errorf("organization billing not found")
		}
		log.Printf("Error scanning billing by ID: %v", err)
		return nil, err
	}

	log.Printf("OrganizationBilling retrieved by ID: %+v", b)
	return &b, nil
}

// Constructor
func NewOrganizationBillingRepository(db *sql.DB) repository.OrganizationBillingRepository {
	return &OrganizationBillingRepositoryImpl{db: db}
}

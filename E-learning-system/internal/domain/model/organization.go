package model

import (
	"time"

	"github.com/gofrs/uuid"
)

// Stores core info about each organization
type Organization struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name           string    `gorm:"size:255;not null"`
	Description    string    `gorm:"type:text"`
	LogoURL        string    `gorm:"size:512"`        // branding
	PrimaryColor   string    `gorm:"size:20"`         // branding
	SecondaryColor string    `gorm:"size:20"`         // branding
	Domain         string    `gorm:"size:255;unique"` // custom subdomain/domain
	Status         string    `gorm:"type:organization_status;default:'pending'"`
	Plan           string    `gorm:"size:50;default:'free'"` // free, pro, enterprise
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

// Links users with organization admin role.
type OrganizationAdmin struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         uuid.UUID `gorm:"type:uuid;not null;index"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	Role           string    `gorm:"size:50;default:'admin'"` // admin, manager
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

// Mapping of tutors to organizations.
type OrganizationTutor struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         uuid.UUID `gorm:"type:uuid;not null;index"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	Approved       bool      `gorm:"default:false"` // org admin must approve tutor
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

// Advanced branding config, separated for flexibility.
type OrganizationBranding struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	LogoURL        string    `gorm:"size:512"`
	PrimaryColor   string    `gorm:"size:20"`
	SecondaryColor string    `gorm:"size:20"`
	Theme          string    `gorm:"size:50;default:'light'"` // light, dark, custom
	EmailTemplate  string    `gorm:"type:text"`               // HTML email template
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

// Handles subscriptions & payments for organizations
type OrganizationBilling struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	Plan           string    `gorm:"size:50;default:'free'"` // free, pro, enterprise
	PaymentMethod  string    `gorm:"size:100"`               // Stripe, PayPal, Local
	SubscriptionID string    `gorm:"size:255"`               // external payment ID
	NextBillingAt  time.Time
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

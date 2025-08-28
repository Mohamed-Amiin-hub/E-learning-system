package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type AdminRole string

const (
	PlatFormAdmin      AdminRole = "admin_admin"
	OrganizationAdmins AdminRole = "organization_admins"
	CourseAdmin        AdminRole = "course_admin"
)

// AdminPermission defines optional granular permissions (can be extended)
type AdminPermission string

const (
	ManageUser          AdminPermission = "manage_User"
	ManageOrganizations AdminPermission = "manage_organizations"
	ManageCourses       AdminPermission = "manage_courses"
	ManagePayments      AdminPermission = "manage_payments"
	ViewAnalytics       AdminPermission = "view_analytics"
	ManageSettings      AdminPermission = "manage_settings"
)

// Admin represents platform or organization administrators
type Admin struct {
	ID             uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         uuid.UUID         `gorm:"type:uuid;not null;index"`
	OrganizationID *uuid.UUID        `gorm:"type:uuid;index"`
	Role           AdminRole         `gorm:"type:admin_role;not null;default:'platform_admin'"`
	Permission     []AdminPermission `gorm:"type:text[];default:'{}'"`
	Status         string            `gorm:"size:50;default:'active'"`
	LastLogin      *time.Time        `gorm:"index"`
	CreatedAt      time.Time         `gorm:"autoCreateTime"`
	UpdatedAt      time.Time         `gorm:"autoUpdateTime"`
	DeletedAt      *time.Time        `gorm:"index;default:null"`
}

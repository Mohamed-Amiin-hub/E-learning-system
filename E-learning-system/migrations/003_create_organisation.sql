-- =====================================================
-- ORGANIZATION ADMINS CRUD
-- =====================================================

-- Create Admin
CREATE OR REPLACE PROCEDURE create_organization_admin(
    IN p_user_id UUID,
    IN p_organization_id UUID,
    IN p_role VARCHAR
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO organization_admins (user_id, organization_id, role)
    VALUES (p_user_id, p_organization_id, p_role);
END;
$$;

-- Update Admin
CREATE OR REPLACE PROCEDURE update_organization_admin(
    IN p_id UUID,
    IN p_role VARCHAR
)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_admins
    SET role = p_role
    WHERE id = p_id AND deleted_at IS NULL;
END;
$$;

-- Soft Delete Admin
CREATE OR REPLACE PROCEDURE delete_organization_admin(IN p_id UUID)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_admins
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE id = p_id;
END;
$$;

-- Get All Admins for an Organization
CREATE OR REPLACE FUNCTION get_admins_by_organization(p_org_id UUID)
RETURNS TABLE (
    id UUID,
    user_id UUID,
    organization_id UUID,
    role VARCHAR,
    created_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, user_id, organization_id, role, created_at
    FROM organization_admins
    WHERE organization_id = p_org_id AND deleted_at IS NULL;
END;
$$;


-- =====================================================
-- ORGANIZATION TUTORS CRUD
-- =====================================================

-- Create Tutor
CREATE OR REPLACE PROCEDURE create_organization_tutor(
    IN p_user_id UUID,
    IN p_organization_id UUID,
    IN p_approved BOOLEAN
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO organization_tutors (user_id, organization_id, approved)
    VALUES (p_user_id, p_organization_id, p_approved);
END;
$$;

-- Update Tutor (approval status)
CREATE OR REPLACE PROCEDURE update_organization_tutor(
    IN p_id UUID,
    IN p_approved BOOLEAN
)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_tutors
    SET approved = p_approved
    WHERE id = p_id AND deleted_at IS NULL;
END;
$$;

-- Soft Delete Tutor
CREATE OR REPLACE PROCEDURE delete_organization_tutor(IN p_id UUID)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_tutors
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE id = p_id;
END;
$$;

-- Get All Tutors for an Organization
CREATE OR REPLACE FUNCTION get_tutors_by_organization(p_org_id UUID)
RETURNS TABLE (
    id UUID,
    user_id UUID,
    organization_id UUID,
    approved BOOLEAN,
    created_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, user_id, organization_id, approved, created_at
    FROM organization_tutors
    WHERE organization_id = p_org_id AND deleted_at IS NULL;
END;
$$;


-- =====================================================
-- ORGANIZATION BRANDING CRUD
-- =====================================================

-- Create Branding
CREATE OR REPLACE PROCEDURE create_organization_branding(
    IN p_organization_id UUID,
    IN p_logo_url VARCHAR,
    IN p_primary_color VARCHAR,
    IN p_secondary_color VARCHAR,
    IN p_theme VARCHAR,
    IN p_email_template TEXT
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO organization_brandings (organization_id, logo_url, primary_color, secondary_color, theme, email_template)
    VALUES (p_organization_id, p_logo_url, p_primary_color, p_secondary_color, p_theme, p_email_template);
END;
$$;

-- Update Branding
CREATE OR REPLACE PROCEDURE update_organization_branding(
    IN p_organization_id UUID,
    IN p_logo_url VARCHAR,
    IN p_primary_color VARCHAR,
    IN p_secondary_color VARCHAR,
    IN p_theme VARCHAR,
    IN p_email_template TEXT
)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_brandings
    SET logo_url = p_logo_url,
        primary_color = p_primary_color,
        secondary_color = p_secondary_color,
        theme = p_theme,
        email_template = p_email_template,
        updated_at = CURRENT_TIMESTAMP
    WHERE organization_id = p_organization_id AND deleted_at IS NULL;
END;
$$;

-- Soft Delete Branding
CREATE OR REPLACE PROCEDURE delete_organization_branding(IN p_organization_id UUID)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_brandings
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE organization_id = p_organization_id;
END;
$$;

-- Get Branding by Organization
CREATE OR REPLACE FUNCTION get_branding_by_organization(p_org_id UUID)
RETURNS TABLE (
    id UUID,
    organization_id UUID,
    logo_url VARCHAR,
    primary_color VARCHAR,
    secondary_color VARCHAR,
    theme VARCHAR,
    email_template TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, organization_id, logo_url, primary_color, secondary_color, theme, email_template, created_at, updated_at
    FROM organization_brandings
    WHERE organization_id = p_org_id AND deleted_at IS NULL;
END;
$$;


-- =====================================================
-- ORGANIZATION BILLING CRUD
-- =====================================================

-- Create Billing
CREATE OR REPLACE PROCEDURE create_organization_billing(
    IN p_organization_id UUID,
    IN p_plan plan_type,
    IN p_payment_method VARCHAR,
    IN p_subscription_id VARCHAR,
    IN p_next_billing_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO organization_billings (organization_id, plan, payment_method, subscription_id, next_billing_at)
    VALUES (p_organization_id, p_plan, p_payment_method, p_subscription_id, p_next_billing_at);
END;
$$;

-- Update Billing
CREATE OR REPLACE PROCEDURE update_organization_billing(
    IN p_organization_id UUID,
    IN p_plan plan_type,
    IN p_payment_method VARCHAR,
    IN p_subscription_id VARCHAR,
    IN p_next_billing_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_billings
    SET plan = p_plan,
        payment_method = p_payment_method,
        subscription_id = p_subscription_id,
        next_billing_at = p_next_billing_at,
        updated_at = CURRENT_TIMESTAMP
    WHERE organization_id = p_organization_id AND deleted_at IS NULL;
END;
$$;

-- Soft Delete Billing
CREATE OR REPLACE PROCEDURE delete_organization_billing(IN p_organization_id UUID)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_billings
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE organization_id = p_organization_id;
END;
$$;

-- Get Billing by Organization
CREATE OR REPLACE FUNCTION get_billing_by_organization(p_org_id UUID)
RETURNS TABLE (
    id UUID,
    organization_id UUID,
    plan VARCHAR,
    payment_method VARCHAR,
    subscription_id VARCHAR,
    next_billing_at TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, organization_id, plan::text, payment_method, subscription_id, next_billing_at, created_at, updated_at
    FROM organization_billings
    WHERE organization_id = p_org_id AND deleted_at IS NULL;
END;
$$;

-- =====================================================
-- ORGANIZATION ADMINS CRUD
-- =====================================================

-- Create Admin
CREATE OR REPLACE PROCEDURE create_organization_admin(
    IN p_user_id UUID,
    IN p_organization_id UUID,
    IN p_role VARCHAR
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO organization_admins (user_id, organization_id, role)
    VALUES (p_user_id, p_organization_id, p_role);
END;
$$;

-- Update Admin
CREATE OR REPLACE PROCEDURE update_organization_admin(
    IN p_id UUID,
    IN p_role VARCHAR
)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_admins
    SET role = p_role
    WHERE id = p_id AND deleted_at IS NULL;
END;
$$;

-- Soft Delete Admin
CREATE OR REPLACE PROCEDURE delete_organization_admin(IN p_id UUID)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_admins
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE id = p_id;
END;
$$;

-- Get All Admins for an Organization
CREATE OR REPLACE FUNCTION get_admins_by_organization(p_org_id UUID)
RETURNS TABLE (
    id UUID,
    user_id UUID,
    organization_id UUID,
    role VARCHAR,
    created_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, user_id, organization_id, role, created_at
    FROM organization_admins
    WHERE organization_id = p_org_id AND deleted_at IS NULL;
END;
$$;


-- =====================================================
-- ORGANIZATION TUTORS CRUD
-- =====================================================

-- Create Tutor
CREATE OR REPLACE PROCEDURE create_organization_tutor(
    IN p_user_id UUID,
    IN p_organization_id UUID,
    IN p_approved BOOLEAN
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO organization_tutors (user_id, organization_id, approved)
    VALUES (p_user_id, p_organization_id, p_approved);
END;
$$;

-- Update Tutor (approval status)
CREATE OR REPLACE PROCEDURE update_organization_tutor(
    IN p_id UUID,
    IN p_approved BOOLEAN
)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_tutors
    SET approved = p_approved
    WHERE id = p_id AND deleted_at IS NULL;
END;
$$;

-- Soft Delete Tutor
CREATE OR REPLACE PROCEDURE delete_organization_tutor(IN p_id UUID)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_tutors
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE id = p_id;
END;
$$;

-- Get All Tutors for an Organization
CREATE OR REPLACE FUNCTION get_tutors_by_organization(p_org_id UUID)
RETURNS TABLE (
    id UUID,
    user_id UUID,
    organization_id UUID,
    approved BOOLEAN,
    created_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, user_id, organization_id, approved, created_at
    FROM organization_tutors
    WHERE organization_id = p_org_id AND deleted_at IS NULL;
END;
$$;


-- =====================================================
-- ORGANIZATION BRANDING CRUD
-- =====================================================

-- Create Branding
CREATE OR REPLACE PROCEDURE create_organization_branding(
    IN p_organization_id UUID,
    IN p_logo_url VARCHAR,
    IN p_primary_color VARCHAR,
    IN p_secondary_color VARCHAR,
    IN p_theme VARCHAR,
    IN p_email_template TEXT
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO organization_brandings (organization_id, logo_url, primary_color, secondary_color, theme, email_template)
    VALUES (p_organization_id, p_logo_url, p_primary_color, p_secondary_color, p_theme, p_email_template);
END;
$$;

-- Update Branding
CREATE OR REPLACE PROCEDURE update_organization_branding(
    IN p_organization_id UUID,
    IN p_logo_url VARCHAR,
    IN p_primary_color VARCHAR,
    IN p_secondary_color VARCHAR,
    IN p_theme VARCHAR,
    IN p_email_template TEXT
)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_brandings
    SET logo_url = p_logo_url,
        primary_color = p_primary_color,
        secondary_color = p_secondary_color,
        theme = p_theme,
        email_template = p_email_template,
        updated_at = CURRENT_TIMESTAMP
    WHERE organization_id = p_organization_id AND deleted_at IS NULL;
END;
$$;

-- Soft Delete Branding
CREATE OR REPLACE PROCEDURE delete_organization_branding(IN p_organization_id UUID)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_brandings
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE organization_id = p_organization_id;
END;
$$;

-- Get Branding by Organization
CREATE OR REPLACE FUNCTION get_branding_by_organization(p_org_id UUID)
RETURNS TABLE (
    id UUID,
    organization_id UUID,
    logo_url VARCHAR,
    primary_color VARCHAR,
    secondary_color VARCHAR,
    theme VARCHAR,
    email_template TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, organization_id, logo_url, primary_color, secondary_color, theme, email_template, created_at, updated_at
    FROM organization_brandings
    WHERE organization_id = p_org_id AND deleted_at IS NULL;
END;
$$;


-- =====================================================
-- ORGANIZATION BILLING CRUD
-- =====================================================

-- Create Billing
CREATE OR REPLACE PROCEDURE create_organization_billing(
    IN p_organization_id UUID,
    IN p_plan plan_type,
    IN p_payment_method VARCHAR,
    IN p_subscription_id VARCHAR,
    IN p_next_billing_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    INSERT INTO organization_billings (organization_id, plan, payment_method, subscription_id, next_billing_at)
    VALUES (p_organization_id, p_plan, p_payment_method, p_subscription_id, p_next_billing_at);
END;
$$;

-- Update Billing
CREATE OR REPLACE PROCEDURE update_organization_billing(
    IN p_organization_id UUID,
    IN p_plan plan_type,
    IN p_payment_method VARCHAR,
    IN p_subscription_id VARCHAR,
    IN p_next_billing_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_billings
    SET plan = p_plan,
        payment_method = p_payment_method,
        subscription_id = p_subscription_id,
        next_billing_at = p_next_billing_at,
        updated_at = CURRENT_TIMESTAMP
    WHERE organization_id = p_organization_id AND deleted_at IS NULL;
END;
$$;

-- Soft Delete Billing
CREATE OR REPLACE PROCEDURE delete_organization_billing(IN p_organization_id UUID)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE organization_billings
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE organization_id = p_organization_id;
END;
$$;

-- Get Billing by Organization
CREATE OR REPLACE FUNCTION get_billing_by_organization(p_org_id UUID)
RETURNS TABLE (
    id UUID,
    organization_id UUID,
    plan VARCHAR,
    payment_method VARCHAR,
    subscription_id VARCHAR,
    next_billing_at TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, organization_id, plan::text, payment_method, subscription_id, next_billing_at, created_at, updated_at
    FROM organization_billings
    WHERE organization_id = p_org_id AND deleted_at IS NULL;
END;
$$;

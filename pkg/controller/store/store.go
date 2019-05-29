package store

import (
	"context"
	"errors"
	"time"

	"github.com/deviceplane/deviceplane/pkg/models"
)

type Users interface {
	CreateUser(ctx context.Context, email, passwordHash, firstName, lastName string) (*models.User, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	ValidateUser(ctx context.Context, email, passwordHash string) (*models.User, error)
	MarkRegistrationCompleted(ctx context.Context, id string) (*models.User, error)
}

var ErrUserNotFound = errors.New("user not found")

type RegistrationTokens interface {
	CreateRegistrationToken(ctx context.Context, userID, hash string) (*models.RegistrationToken, error)
	GetRegistrationToken(ctx context.Context, id string) (*models.RegistrationToken, error)
	ValidateRegistrationToken(ctx context.Context, hash string) (*models.RegistrationToken, error)
}

var ErrRegistrationTokenNotFound = errors.New("registration token not found")

type Sessions interface {
	CreateSession(ctx context.Context, userID string, hash string) (*models.Session, error)
	GetSession(ctx context.Context, id string) (*models.Session, error)
	ValidateSession(ctx context.Context, hash string) (*models.Session, error)
	DeleteSession(ctx context.Context, id string) error
}

var ErrSessionNotFound = errors.New("session not found")

type UserAccessKeys interface {
	CreateUserAccessKey(ctx context.Context, userID string, hash string) (*models.UserAccessKey, error)
	GetUserAccessKey(ctx context.Context, id string) (*models.UserAccessKey, error)
	ValidateUserAccessKey(ctx context.Context, hash string) (*models.UserAccessKey, error)
	DeleteUserAccessKey(ctx context.Context, id string) error
}

var ErrUserAccessKeyNotFound = errors.New("user access key not found")

type Projects interface {
	CreateProject(ctx context.Context, name string) (*models.Project, error)
	GetProject(ctx context.Context, id string) (*models.Project, error)
	LookupProject(ctx context.Context, name string) (*models.Project, error)
}

var ErrProjectNotFound = errors.New("project not found")

type ProjectDeviceCounts interface {
	GetProjectDeviceCounts(ctx context.Context, projectID string) (*models.ProjectDeviceCounts, error)
}

type ProjectApplicationCounts interface {
	GetProjectApplicationCounts(ctx context.Context, projectID string) (*models.ProjectApplicationCounts, error)
}

type Roles interface {
	CreateRole(ctx context.Context, projectID, name, description, config string) (*models.Role, error)
	GetRole(ctx context.Context, id, projectID string) (*models.Role, error)
	LookupRole(ctx context.Context, name, projectID string) (*models.Role, error)
	ListRoles(ctx context.Context, projectID string) ([]models.Role, error)
	UpdateRole(ctx context.Context, id, projectID, name, description, config string) (*models.Role, error)
	DeleteRole(ctx context.Context, id, projectID string) error
}

var ErrRoleNotFound = errors.New("role not found")

type Memberships interface {
	CreateMembership(ctx context.Context, userID, projectID string) (*models.Membership, error)
	GetMembership(ctx context.Context, userID, projectID string) (*models.Membership, error)
	ListMembershipsByUser(ctx context.Context, userID string) ([]models.Membership, error)
	ListMembershipsByProject(ctx context.Context, projectID string) ([]models.Membership, error)
	DeleteMembership(ctx context.Context, userID, projectID string) error
}

var ErrMembershipNotFound = errors.New("membership not found")

type MembershipRoleBindings interface {
	CreateMembershipRoleBinding(ctx context.Context, userID, projectID, roleID string) (*models.MembershipRoleBinding, error)
	GetMembershipRoleBinding(ctx context.Context, userID, projectID, roleID string) (*models.MembershipRoleBinding, error)
	ListMembershipRoleBindings(ctx context.Context, userID, projectID string) ([]models.MembershipRoleBinding, error)
	DeleteMembershipRoleBinding(ctx context.Context, userID, projectID, roleID string) error
}

var ErrMembershipRoleBindingNotFound = errors.New("membership role binding not found")

type ServiceAccounts interface {
	CreateServiceAccount(ctx context.Context, projectID, name, description string) (*models.ServiceAccount, error)
	GetServiceAccount(ctx context.Context, id, projectID string) (*models.ServiceAccount, error)
	LookupServiceAccount(ctx context.Context, name, projectID string) (*models.ServiceAccount, error)
	ListServiceAccounts(ctx context.Context, projectID string) ([]models.ServiceAccount, error)
	UpdateServiceAccount(ctx context.Context, id, projectID, name, description string) (*models.ServiceAccount, error)
	DeleteServiceAccount(ctx context.Context, id, projectID string) error
}

var ErrServiceAccountNotFound = errors.New("service account not found")

type ServiceAccountAccessKeys interface {
	CreateServiceAccountAccessKey(ctx context.Context, projectID, serviceAccountID string, hash string) (*models.ServiceAccountAccessKey, error)
	GetServiceAccountAccessKey(ctx context.Context, id, projectID string) (*models.ServiceAccountAccessKey, error)
	ValidateServiceAccountAccessKey(ctx context.Context, hash string) (*models.ServiceAccountAccessKey, error)
	DeleteServiceAccountAccessKey(ctx context.Context, id, projectID string) error
}

var ErrServiceAccountAccessKeyNotFound = errors.New("service account access key not found")

type ServiceAccountRoleBindings interface {
	CreateServiceAccountRoleBinding(ctx context.Context, serviceAccountID, roleID, projectID string) (*models.ServiceAccountRoleBinding, error)
	GetServiceAccountRoleBinding(ctx context.Context, serviceAccountID, roleID, projectID string) (*models.ServiceAccountRoleBinding, error)
	ListServiceAccountRoleBindings(ctx context.Context, serviceAccountID, projectID string) ([]models.ServiceAccountRoleBinding, error)
	DeleteServiceAccountRoleBinding(ctx context.Context, serviceAccountID, roleID, projectID string) error
}

var ErrServiceAccountRoleBindingNotFound = errors.New("service account role binding not found")

type Devices interface {
	CreateDevice(ctx context.Context, projectID string) (*models.Device, error)
	GetDevice(ctx context.Context, id, projectID string) (*models.Device, error)
	ListDevices(ctx context.Context, projectID string) ([]models.Device, error)
	SetDeviceInfo(ctx context.Context, id, projectID string, deviceInfo models.DeviceInfo) (*models.Device, error)
}

var ErrDeviceNotFound = errors.New("device not found")

type DeviceStatuses interface {
	ResetDeviceStatus(ctx context.Context, deviceID string, ttl time.Duration) error
	GetDeviceStatus(ctx context.Context, deviceID string) (models.DeviceStatus, error)
	GetDeviceStatuses(ctx context.Context, deviceIDs []string) ([]models.DeviceStatus, error)
}

type DeviceLabels interface {
	SetDeviceLabel(ctx context.Context, key, deviceID, projectID, value string) (*models.DeviceLabel, error)
	GetDeviceLabel(ctx context.Context, key, deviceID, projectID string) (*models.DeviceLabel, error)
	ListDeviceLabels(ctx context.Context, deviceID, projectID string) ([]models.DeviceLabel, error)
	DeleteDeviceLabel(ctx context.Context, key, deviceID, projectID string) error
}

var ErrDeviceLabelNotFound = errors.New("device label not found")

type DeviceRegistrationTokens interface {
	CreateDeviceRegistrationToken(ctx context.Context, projectID string) (*models.DeviceRegistrationToken, error)
	GetDeviceRegistrationToken(ctx context.Context, id, projectID string) (*models.DeviceRegistrationToken, error)
	BindDeviceRegistrationToken(ctx context.Context, id, projectID, deviceAccessKeyID string) (*models.DeviceRegistrationToken, error)
}

var ErrDeviceRegistrationTokenNotFound = errors.New("device registration token not found")

type DeviceAccessKeys interface {
	CreateDeviceAccessKey(ctx context.Context, projectID, deviceID, hash string) (*models.DeviceAccessKey, error)
	GetDeviceAccessKey(ctx context.Context, id, projectID string) (*models.DeviceAccessKey, error)
	ValidateDeviceAccessKey(ctx context.Context, projectID, hash string) (*models.DeviceAccessKey, error)
}

var ErrDeviceAccessKeyNotFound = errors.New("device access key not found")

type Applications interface {
	CreateApplication(ctx context.Context, projectID, name, description string, applicationSettings models.ApplicationSettings) (*models.Application, error)
	GetApplication(ctx context.Context, id, projectID string) (*models.Application, error)
	LookupApplication(ctx context.Context, name, projectID string) (*models.Application, error)
	ListApplications(ctx context.Context, projectID string) ([]models.Application, error)
	UpdateApplication(ctx context.Context, id, projectID, name, description string, applicationSettings models.ApplicationSettings) (*models.Application, error)
	DeleteApplication(ctx context.Context, id, projectID string) error
}

var ErrApplicationNotFound = errors.New("application not found")

type ApplicationDeviceCounts interface {
	GetApplicationDeviceCounts(ctx context.Context, projectID, applicationID string) (*models.ApplicationDeviceCounts, error)
}

type Releases interface {
	CreateRelease(ctx context.Context, projectID, applicationID, config string) (*models.Release, error)
	GetRelease(ctx context.Context, id, projectID, applicationID string) (*models.Release, error)
	GetLatestRelease(ctx context.Context, projectID, applicationID string) (*models.Release, error)
	ListReleases(ctx context.Context, projectID, applicationID string) ([]models.Release, error)
}

var ErrReleaseNotFound = errors.New("release not found")

type ReleaseDeviceCounts interface {
	GetReleaseDeviceCounts(ctx context.Context, projectID, applicationID, releaseID string) (*models.ReleaseDeviceCounts, error)
}

type DeviceApplicationStatuses interface {
	SetDeviceApplicationStatus(ctx context.Context, projectID, deviceID, applicationID, currentReleaseID string) error
	GetDeviceApplicationStatus(ctx context.Context, projectID, deviceID, applicationID string) (*models.DeviceApplicationStatus, error)
}

var ErrDeviceApplicationStatusNotFound = errors.New("device application status not found")

type DeviceServiceStatuses interface {
	SetDeviceServiceStatus(ctx context.Context, projectID, deviceID, applicationID, service, currentReleaseID string) error
	GetDeviceServiceStatus(ctx context.Context, projectID, deviceID, applicationID, service string) (*models.DeviceServiceStatus, error)
	GetDeviceServiceStatuses(ctx context.Context, projectID, deviceID, applicationID string) ([]models.DeviceServiceStatus, error)
}

var ErrDeviceServiceStatusNotFound = errors.New("device service status not found")

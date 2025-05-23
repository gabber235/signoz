package dao

import (
	"context"

	"github.com/SigNoz/signoz/pkg/query-service/model"
	"github.com/SigNoz/signoz/pkg/types"
	"github.com/SigNoz/signoz/pkg/types/authtypes"
)

type ModelDao interface {
	Queries
	Mutations
}

type Queries interface {
	GetInviteFromEmail(ctx context.Context, email string) (*types.Invite, *model.ApiError)
	GetInviteFromToken(ctx context.Context, token string) (*types.Invite, *model.ApiError)
	GetInvites(ctx context.Context, orgID string) ([]types.Invite, *model.ApiError)

	GetUser(ctx context.Context, id string) (*types.GettableUser, *model.ApiError)
	GetUserByEmail(ctx context.Context, email string) (*types.GettableUser, *model.ApiError)
	GetUsers(ctx context.Context) ([]types.GettableUser, *model.ApiError)
	GetUsersWithOpts(ctx context.Context, limit int) ([]types.GettableUser, *model.ApiError)

	GetResetPasswordEntry(ctx context.Context, token string) (*types.ResetPasswordRequest, *model.ApiError)
	GetUsersByOrg(ctx context.Context, orgId string) ([]types.GettableUser, *model.ApiError)
	GetUsersByRole(ctx context.Context, role authtypes.Role) ([]types.GettableUser, *model.ApiError)

	GetApdexSettings(ctx context.Context, orgID string, services []string) ([]types.ApdexSettings, *model.ApiError)

	PrecheckLogin(ctx context.Context, email, sourceUrl string) (*model.PrecheckResponse, model.BaseApiError)
}

type Mutations interface {
	CreateInviteEntry(ctx context.Context, req *types.Invite) *model.ApiError
	DeleteInvitation(ctx context.Context, orgID string, email string) *model.ApiError

	CreateUser(ctx context.Context, user *types.User, isFirstUser bool) (*types.User, *model.ApiError)
	EditUser(ctx context.Context, update *types.User) (*types.User, *model.ApiError)
	DeleteUser(ctx context.Context, id string) *model.ApiError

	CreateResetPasswordEntry(ctx context.Context, req *types.ResetPasswordRequest) *model.ApiError
	DeleteResetPasswordEntry(ctx context.Context, token string) *model.ApiError

	UpdateUserPassword(ctx context.Context, hash, userId string) *model.ApiError
	UpdateUserRole(ctx context.Context, userId string, role authtypes.Role) *model.ApiError

	SetApdexSettings(ctx context.Context, orgID string, set *types.ApdexSettings) *model.ApiError
}

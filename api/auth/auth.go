// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"xiao-yi-transit/api/auth/v1"
)

type IAuthV1 interface {
	AuthForgetPassword(ctx context.Context, req *v1.AuthForgetPasswordReq) (res *v1.AuthForgetPasswordRes, err error)
	AuthCurrent(ctx context.Context, req *v1.AuthCurrentReq) (res *v1.AuthCurrentRes, err error)
	AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error)
}

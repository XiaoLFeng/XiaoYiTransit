package cerror

import "github.com/XiaoLFeng/bamboo-utils/berror"

//
// 自定义错误码
//

var (
	ErrUserNotExist      = berror.NewErrorCode(40404, "用户不存在", nil)
	ErrPasswordIncorrect = berror.NewErrorCode(40106, "密码错误", nil)
	ErrUserTokenNotExist = berror.NewErrorCode(40107, "用户令牌不存在", nil)

	ErrPasswordMismatch = berror.NewErrorCode(40009, "新密码与确认密码不一致", nil)
)

package cerror

import "github.com/XiaoLFeng/bamboo-utils/berror"

//
// 自定义错误码
//

var (
	ErrUserNotExist      = berror.NewErrorCode(40404, "用户不存在", nil)
	ErrPasswordIncorrect = berror.NewErrorCode(40106, "密码错误", nil)
)

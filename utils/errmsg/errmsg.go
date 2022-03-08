package errmsg

const(
	SUCCESS = 200
	ERROR = 500
	//code = 1000... 用户模块的错误
	ERROR_USERNAME_USED = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USERNAME_NOT_EXIST = 1003
	ERROR_TOKEN_NOT_EXIST = 1004
	ERROR_TOKEN_RUNTIME = 1005
	ERROR_TOKEN_TYPE_WRONG = 1006
	ERROR_TOKEN_WRONG = 1007
	ERROR_FORMAT_WRONG = 1008
	ERROR_USER_NO_PERMISSION = 1009
	//code = 2000... 文章模块的错误
	ERROR_ARTICLE_NOT_EXIST= 2001
	ERROR_ARTICLE_OF_CATE_NOT_EXIST= 2002
	//code = 3000... 分类模块的错误
	ERROR_CATENAME_USED= 3001
	ERROR_CATENAME_NOT_EXIST= 3002
)

var CodeMsg = map[int]string{
	SUCCESS:"OK",
	ERROR:"FAIL",
	ERROR_USERNAME_USED:"用户名已存在",
	ERROR_PASSWORD_WRONG:"密码错误",
	ERROR_USERNAME_NOT_EXIST:"用户名不存在",
	ERROR_TOKEN_NOT_EXIST:"TOKEN不存在",
	ERROR_TOKEN_RUNTIME:"TOKEN已过期",
	ERROR_TOKEN_TYPE_WRONG:"TOKEN类型不正确",
	ERROR_TOKEN_WRONG:"TOKEN不正确",
	ERROR_FORMAT_WRONG:"TOKEN格式不正确",
	ERROR_USER_NO_PERMISSION:"用户没有权限",

	ERROR_ARTICLE_NOT_EXIST:"文章不存在",
	ERROR_ARTICLE_OF_CATE_NOT_EXIST:"该分类的文章不存在",

	ERROR_CATENAME_USED:"分类已存在",
	ERROR_CATENAME_NOT_EXIST:"分类不存在",


}

func GetErrMsg(code int)string{
	return CodeMsg[code]
}
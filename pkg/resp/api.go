package resp

type Response struct {
	Status   int         `json:"status,omitempty"`    // 状态码 200:成功   其他失败
	Code     int         `json:"code,omitempty"`      // 错误码 状态码不为200时返回
	Type     string      `json:"type,omitempty"`      // 错误类型 状态码不为200时返回
	Message  string      `json:"message,omitempty"`   // 错误信息 状态码不为200时返回
	MoreInfo interface{} `json:"more_info,omitempty"` // 错误扩展信息 状态码不为200时返回
	Data     interface{} `json:"data,omitempty"`      // 返回数据
}

// define Response status
const (
	StatusSuccess  = 200
	StatusUser     = 400
	StatusSystem   = 500
	StatusDatabase = 600
	StatusFS       = 700
	StatusCache    = 800
	StatusEtcd     = 900
	StatusProxy    = 1000
	StatusHttp     = 1100
)

func NewSuccess(data interface{}) (res *Response) {
	return &Response{
		Status: 200,
		Data:   data,
	}
}

var (
	RespSuccess = &Response{StatusSuccess, 90000, "操作成功", "操作成功", nil, nil}
)

var (
	Err404          = &Response{StatusUser, 404, "page not found", "page not found", nil, nil}
	ErrUserExist    = &Response{StatusUser, 10001, "用户操作错误", "用户账户已存在", nil, nil}
	ErrUserInput    = &Response{StatusUser, 10002, "用户操作错误", "用户输入参数错误", nil, nil}
	ErrUserModify   = &Response{StatusUser, 10003, "用户操作错误", "修改用户错误", nil, nil}
	ErrUserDelete   = &Response{StatusUser, 10004, "用户操作错误", "删除用户错误", nil, nil}
	ErrNoUser       = &Response{StatusUser, 10005, "用户操作错误", "用户账户不存在", nil, nil}
	ErrUserGet      = &Response{StatusUser, 10006, "用户操作错误", "获取所有用户错误", nil, nil}
	ErrUserGetLogs  = &Response{StatusUser, 10007, "用户操作错误", "获取用户登录日志错误", nil, nil}
	ErrAppid        = &Response{StatusUser, 10011, "用户认证错误", "无效的Appid或Secret", nil, nil}
	ErrTokenInValid = &Response{StatusUser, 10012, "Token 认证错误", "验证 token 无效或已过期", nil, nil}
	ErrPermission   = &Response{StatusUser, 10013, "没有权限", "没有操作权限", nil, nil}
	ErrUserLogin    = &Response{StatusUser, 10014, "用户未登录", "用户账户未登录", nil, nil}
	ErrInputData    = &Response{StatusUser, 20001, "数据输入错误", "客户端参数错误", nil, nil}
	ErrVersionCheck = &Response{StatusUser, 20002, "版本检查", "当前版本过低", nil, nil}
)

var (
	ErrOpenFile     = &Response{StatusSystem, 10001, "服务器错误", "打开文件出错", nil, nil}
	Err500          = &Response{StatusSystem, 10002, "服务器错误", "接口访问出错", nil, nil}
	ErrWriteFile    = &Response{StatusSystem, 10003, "服务器错误", "写文件出错", nil, nil}
	ErrSystem       = &Response{StatusSystem, 10004, "服务器错误", "操作系统错误", nil, nil}
	ErrTransferData = &Response{StatusSystem, 10005, "数据转换错误", "Json 字符串转 Map 错误", nil, nil}
	ErrNewTmpl      = &Response{StatusSystem, 10006, "数据转换错误", "Tmpl 模板文件解析错误", nil, nil}
	ErrTransferTmpl = &Response{StatusSystem, 10006, "数据转换错误", "Tmpl 模板文件转换错误", nil, nil}
)

var (
	ErrDb                  = &Response{StatusDatabase, 10001, "数据库错误", "数据库操作错误", nil, nil}
	ErrDupRecord           = &Response{StatusDatabase, 10002, "数据库错误", "数据库记录重复", nil, nil}
	ErrNoRecord            = &Response{StatusDatabase, 10003, "数据库错误", "数据库记录不存在", nil, nil}
	ErrUserPass            = &Response{StatusDatabase, 10004, "数据库错误", "用户名或密码不正确", nil, nil}
	ErrDbInsert            = &Response{StatusDatabase, 10005, "数据库错误", "数据添加错误", nil, nil}
	ErrDbRead              = &Response{StatusDatabase, 10006, "数据库错误", "数据读取错误", nil, nil}
	ErrDbUpdate            = &Response{StatusDatabase, 10007, "数据库错误", "数据更新错误", nil, nil}
	ErrDbDelete            = &Response{StatusDatabase, 10008, "数据库错误", "数据删除错误", nil, nil}
	ErrChangeAccountStatus = &Response{StatusDatabase, 10009, "数据库错误", "更新账户状态错误", nil, nil}
)

var (
	ErrFileSystem       = &Response{StatusFS, 10000, "文件系统错误", "文件系统连接错误！", nil, nil}
	ErrFileExist        = &Response{StatusFS, 10001, "文件系统错误", "文件已存在！", nil, nil}
	ErrFileNonExist     = &Response{StatusFS, 10002, "文件系统错误", "文件不存在！", nil, nil}
	ErrFileUpload       = &Response{StatusFS, 10003, "文件系统错误", "文件上传错误！", nil, nil}
	ErrFileDownload     = &Response{StatusFS, 10004, "文件系统错误", "文件下载错误！", nil, nil}
	ErrFileDelete       = &Response{StatusFS, 10005, "文件系统错误", "文件删除错误！", nil, nil}
	ErrFileList         = &Response{StatusFS, 10006, "文件系统错误", "获取文件列表错误！", nil, nil}
	ErrFileCopy         = &Response{StatusFS, 10007, "文件系统错误", "文件转存错误！", nil, nil}
	ErrFileSetLifeCycle = &Response{StatusFS, 10007, "文件系统错误", "设置存储桶生命周期错误！", nil, nil}
	ErrFileGetLifeCycle = &Response{StatusFS, 10007, "文件系统错误", "获取存储桶生命周期错误！", nil, nil}
)

var (
	ErrEtcdClient = &Response{StatusEtcd, 10001, "Etcd错误", "获取etcd连接错误！", nil, nil}
	ErrEtcdPut    = &Response{StatusEtcd, 10002, "Etcd错误", "etcd put 错误！", nil, nil}
	ErrEtcdGet    = &Response{StatusEtcd, 10003, "Etcd错误", "etcd get 错误！", nil, nil}
	ErrEtcdDelete = &Response{StatusEtcd, 10004, "Etcd错误", "etcd delete 错误！", nil, nil}
	ErrEtcdWatch  = &Response{StatusEtcd, 10005, "Etcd错误", "etcd watch 错误！", nil, nil}
)

var (
	ErrProxyInput = &Response{StatusProxy, 10001, "网关代理错误", "无效的代理地址！", nil, nil}
)

var (
	ErrHttpRequest      = &Response{StatusHttp, 10001, "Http代理错误", "发送Http请求错误！", nil, nil}
	ErrHttpRequestData  = &Response{StatusHttp, 10002, "Http代理错误", "设置Http请求数据错误！", nil, nil}
	ErrHttpResponse     = &Response{StatusHttp, 10003, "Http代理错误", "获取Http响应数据错误！", nil, nil}
	ErrHttpResponseData = &Response{StatusHttp, 10004, "Http代理错误", "解析Http响应数据错误！", nil, nil}
)

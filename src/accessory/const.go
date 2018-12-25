package accessory

const (
	//服务调用没有发生任何错误
	ErrorOKId           = "200000" //所有的服务都以这个编码为正确无误调用之后的返回值
	ErrorOKMsg          = "OK"     //所有的服务都以这个消息为200000错误码对应的消息
	ErrorOKInsertMsg    = "增加成功"
	ErrorOKDeleteMsg    = "删除成功"
	ErrorOKModifyMsg    = "修改成功"
	ErrorOKSelectMsg    = "查询成功"
	ErrorOKGetMsg       = "数据查询成功"
	ErrorOkOperationMsg = "操作服务成功"

	ComErrorId  = "200001"
	ComErrorMsg = "com参数错误"

	InsertErrorId  = "200002"
	InsertErrorMsg = "插入失败"

	CheckErrorId  = "200003"
	CheckErrorMsg = "数据未通过校验"

	//Json格式错误
	ErrorJsonErrId  = "200004"
	ErrorJsonErrMsg = "json格式错误"

	//get输入com判断
	ErrorComisNilId  = "200005"
	ErrorComisNilMsg = "查询命令标识为空"

	//数据异常
	ErrorId  = "200006"
	ErrorMsg = "数据异常"

	//数据已存在
	ErrorDataExistsErrId  = "200007"
	ErrorDataExistsErrMsg = "数据已存在"

	//系统错误
	ErrorSystemErrId  = "200008"
	ErrorSystemErrMsg = "系统错误"

	//数据不存在
	ErrorDataNotExistsErrId = "20009"
	ErrorDataNotExistsMsg   = "数据不存在"

	ErrorDiscoverCheckId  = "200010"
	ErrorDiscoverCheckMsg = "sns服务发现失败"

	ErrorAddInfoId  = "200011"
	ErrorAddInfoMsg = "添加错误"

	ErrorUpdataInfoFailedId  = "200012"
	ErrorUpdataInfoFailedMsg = "修改失败"

	ErrorParameterIsErrId  = "200014"
	ErrorParameterIsErrMsg = "参数错误"

	ErrorDelErrId  = "200015"
	ErrorDelErrMsg = "删除错误"

	ErrorJsonFmtErrId  = "2000018"
	ErrorJsonFmtErrMsg = "json格式错误!"

	ErrorQueryErrId  = "200021"
	ErrorQueryErrMsg = "查询错误"

	ErrorDataNullId  = "2000016"
	ErrorDataNullMsg = "请求参数为空!"

	ErrorGroupIdLenId  = "2000017"
	ErrorGroupIdLenMsg = "机构表示长度不为8位！"

	ErrorGroup_idNilId  = "400005"
	ErrorGroup_idNilMsg = "groupId不能为空"

	ErrorCodeNilId  = "400006"
	ErrorCodeNilMsg = "Code不能为空"

	ErrorGroup_idMust8lenId  = "400012"
	ErrorGroup_idMust8lenMsg = "group_id-长度必须是8位"

	ErrorDataIsNilId  = "3000008"
	ErrorDataIsNilMsg = "数据不存在"

	ErrorDeleteInfoId  = "400010"
	ErrorDeleteInfoMsg = "删除错误"

	ErrorSelectInfoId  = "400011"
	ErrorSelectInfoMsg = "查询错误"
)

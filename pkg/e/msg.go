package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_EXIST_TAG:                 "已经存在该标签名称",
	ERROR_NOT_EXIST_TAG:             "该标签不存在",
	ERROR_NOT_EXIST_ARICLE:          "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "上传保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "上传检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "上传图片格式错误",
}

func GetMsg(code int) string {
	message, ok := MsgFlags[code]
	if (!ok) {
		return MsgFlags[ERROR]
	}
	return message
}

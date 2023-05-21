package errorcode

import "simpleService/internal/response"

func Init() {
	response.Init()

	// 100-199
	response.AddListCodes(user)
}

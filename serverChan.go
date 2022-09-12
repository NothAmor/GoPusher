package provider

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	structs "github.com/NothAmor/GoPusher/structs"
)

func ServerChan(pusherParams structs.ServerChanRequestStruct) (structs.PusherResponse, error) {

	// 初始化返回值
	pusherReturn := structs.PusherResponse{
		Code:      500,
		Message:   "ERROR",
		Timestamp: time.Now().Unix(),
	}

	// 推送key和推送信息标题不可为空，如果为空，返回错误
	if pusherParams.Key == "" || pusherParams.Title == "" {
		return pusherReturn, errors.New("ServerChan推送的key和title不可为空！")
	}

	var serverChanUrl string = fmt.Sprintf("https://sctapi.ftqq.com/%s.send", pusherParams.Key)

	urlValues := url.Values{}

	// 遍历结构体，将非空参数加入post表单中
	structType := reflect.TypeOf(pusherParams)
	structValue := reflect.ValueOf(pusherParams)
	for i := 0; i < structType.NumField(); i++ {
		if structValue.Field(i).Interface() == "" {
			continue
		}

		urlValues.Add(strings.ToLower(structType.Field(i).Name), structValue.Field(i).String())
	}

	// 对ServerChan服务器进行POST请求，并解析response BODY
	response, err := http.PostForm(serverChanUrl, urlValues)
	if err != nil {
		return pusherReturn, errors.New("HTTP请求错误！")
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return pusherReturn, errors.New("解析HTT BODY错误！BODY：" + string(responseBody))
	}

	// 上面可能出现错误的地方都没问题，将返回值修改为成功
	pusherReturn.Code = 200
	pusherReturn.Message = "SUCCESS"

	return pusherReturn, nil
}

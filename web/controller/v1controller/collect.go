package v1controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"duoduo.com/micro/common/http/dto"
	tbServiceSrv "duoduo.com/micro/common/proto/tbServiceSrv"
	"duoduo.com/micro/tbservice/tb-service-web/controller"
)

func Push(c *gin.Context) {

}

var (
	Collect = CollectController{}
)

type CollectController struct {
	controller.ApiBaseController
}

func (t CollectController) Push(c *gin.Context) {
	signInfo :=  t.GetUserSinInfo(c)
	clientId := c.PostForm("clientId")
	did := c.Query("_did")
	data := make(map[string]string)
	data["did"] = did
	data["data"] = clientId
	
	webClient := tbServiceSrv.NewTbServiceSrvService("go.micro.srv.tbServiceSrv", client.DefaultClient)
	rsp, err := webClient.CollectPush(context.TODO(), &tbServiceSrv.PushClientRequest{
		UserId:   int64(signInfo.UserId),
		ClientId: clientId,
		Did:      did,
	})
	if err != nil {
		println(err.Error())
		return
	}
	var result = dto.ApiResponse{}
	if rsp.Result {
		result.Status.Code = 200
	}else{
		result.Status.Code = 500
	}
	result.Status.Msg = ""
	c.JSON(200, result)
}

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maybaby/gscheduler/pkg/app"
	"github.com/maybaby/gscheduler/pkg/e"
	"github.com/maybaby/gscheduler/services/process_service"
	"github.com/maybaby/gscheduler/services/task_service"
	"net/http"
	"time"
)

/**
 * 流程controller
 * 	pd := process_service.ProcessData{}
	td := &task_service.TaskNode{
		Id: "11111",
		Name: "firstNode",
		Params: task_service.TaskParams{
				ResourceList: []string{"1111","2222"},
				LocalParams: []string{"ccccccddd"},
				RawScript: "print(\"pythonNode1\")",
		},
		RunFlag: "1",
		Type: "PYTHON",

	}
	pd.Tasks = []*task_service.TaskNode{td}
	pd.Timeout = 123
	s, _:= pd.MarshalJSON()
	fmt.Println(string(s))
*/

// @Summary 测试任务定义接口 a rpc procedure
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/process/test [get]
func TestCreateProcessDefinition(c *gin.Context) {
	var (
		//appG = app.Gin{C: c}
		json_ AddProcessDefineJson
	)
	c.BindJSON(&json_)

	//appG.Response(httpCode, errCode, nil)
	processService := process_service.ProcessDefinition{}
	_ = processService.Save()
	pd := process_service.ProcessData{}
	td := &task_service.TaskNode{
		Id:   "11111",
		Name: "firstNode",
		Params: task_service.TaskParams{
			ResourceList: []string{"1111", "2222"},
			LocalParams:  []string{"ccccccddd"},
			RawScript:    "print(\"pythonNode1\")",
		},
		RunFlag: "1",
		Type:    "PYTHON",
	}
	pd.Tasks = []*task_service.TaskNode{td}
	pd.Timeout = 123
	s, _ := pd.MarshalJSON()
	fmt.Println(string(s))
	return

	//registryAddr := "http://localhost:9999/_gsrpc_/registry"
	////time.Sleep(time.Second*16)
	//call(registryAddr)

}

// @Summary 保存任务定义接口 a rpc procedure
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/process/save [post]
func CreateProcessDefinition(c *gin.Context) {
	var (
		appG  = app.Gin{C: c}
		json_ AddProcessDefineJson
	)
	c.BindJSON(&json_)

	processService := process_service.ProcessDefinition{
		Name:        json_.Name,
		Description: json_.Desc,
		GroupId:     json_.GroupId,
		ProcessData: json_.ProcessData,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	err := processService.Save()

	if err != nil {
		appG.Response(http.StatusForbidden, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.ERROR_AUTH_TOKEN, json_.Name)

	//registryAddr := "http://localhost:9999/_gsrpc_/registry"
	////time.Sleep(time.Second*16)
	//call(registryAddr)

}

type AddProcessDefineJson struct {
	Name        string                       `form:"name" json:"name"`
	GroupId     string                       `form:"groupId" json:"groupId"`
	ProcessData *process_service.ProcessData `form:"processData" json:"processData"`
	Desc        string                       `form:"desc" json:"desc"`
	//TODO 定义 dag的连接 暂时不支持dag
}

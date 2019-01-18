package controller

import (
	"../model"
	"../service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
)

// 根据emailId获取emialInfo数据
func GetEmailInfoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	emailId := params["emailId"]
	fmt.Print("GetEmailInfoById:" + emailId)
	emailInfo := model.GetEmailInfoById(emailId)
	json.NewEncoder(w).Encode(emailInfo)
}

func GetAllEmailInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Print("GetAllEmailInfo")
	emailInfos := model.GetAllEmailInfo()
	json.NewEncoder(w).Encode(emailInfos)
}

func SendEmail(w http.ResponseWriter, r *http.Request) {
	// parse JSON body
	var req map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)
	fmt.Print("post args:", req)
	// TODO 获取DUID
	emailInfo := &model.EmailInfo{
		EmailId:          "",
		EmailTitle:       req["emailTitle"].(string),
		EmailContent:     req["emailContent"].(string),
		SenderName:       "",
		SenderAddress:    "",
		RecipientName:    "",
		RecipientAddress: req["recipientAddress"].(string)}
	str, _ := json.Marshal(emailInfo)
	fmt.Print("SendEmail:", string(str))

	// 发送邮件
	err := service.SendEmail(emailInfo)
	if err != nil {
		fmt.Println("send fail, check it, ", err)
	} else {
		// 写入数据库
		res, _ := model.InsertEmail(emailInfo)
		json.NewEncoder(w).Encode(res)
	}
}

func UpdateEmailInfoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	emailInfo := &model.EmailInfo{
		0,
		params["EmailId"],
		params["EmailTitle"],
		params["EmailContent"],
		"",
		"",
		"",
		"",
		"",
		0}

	str, _ := json.Marshal(emailInfo)

	fmt.Print("UpdateEmailInfoById:", string(str))
	model.UpdateEmailInfoById(emailInfo)
	json.NewEncoder(w).Encode("")
}

func DeleteEmailInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	emailId := params["emailId"]
	fmt.Print("DeleteEmailInfo:" + emailId)
	emailInfo := model.DeleteEmailInfo(emailId)
	json.NewEncoder(w).Encode(emailInfo)
}

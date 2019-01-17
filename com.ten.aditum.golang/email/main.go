package main

import (
	"./model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

func send() {
	auth := smtp.PlainAuth("", "953637695@qq.com", "password", "smtp.qq.com")
	to := []string{"874560965@qq.com"}
	nickname := "test"
	user := "953637695@qq.com"
	subject := "test mail"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := "This is the email body."
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

//func mail(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("method:", r.Method)
//	// 解析url传递的参数，对于POST则解析响应包的主体（request body）
//	r.ParseForm()
//
//	if r.Method == "GET" {
//		t, _ := template.ParseFiles("mail.gtpl")
//		t.Execute(w, nil)
//		emails := model.GetAllEmailInfo()
//		for email := range emails {
//			w.Write(WriteJSON(email))
//		}
//	} else if r.Method == "POST" {
//		//请求的是登陆数据，那么执行登陆的逻辑判断
//		fmt.Println("username:", r.Form["username"])
//		fmt.Println("password:", r.Form["password"])
//	} else if r.Method == "PUT" {
//		t, _ := template.ParseFiles("mail.gtpl")
//		t.Execute(w, nil)
//	} else if r.Method == "DELETE" {
//		t, _ := template.ParseFiles("mail.gtpl")
//		t.Execute(w, nil)
//	} else {
//		t, _ := template.ParseFiles("mail.gtpl")
//		t.Execute(w, nil)
//	}
//}

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

func InsertEmail(w http.ResponseWriter, r *http.Request) {
	// parse JSON body
	var req map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)
	fmt.Print("post args:", req)
	emailInfo := &model.EmailInfo{
		0,
		req["emailId"].(string),
		req["emailTitle"].(string),
		req["emailContent"].(string),
		req["senderName"].(string),
		req["senderAddress"].(string),
		req["recipientName"].(string),
		req["recipientAddress"].(string),
		"",
		0}

	str, _ := json.Marshal(emailInfo)
	fmt.Print("InsertEmail:", string(str))
	emailId, _ := model.InsertEmail(emailInfo)
	json.NewEncoder(w).Encode(emailId)
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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/email/{emailId}", GetEmailInfoById).Methods("GET")
	router.HandleFunc("/email", GetAllEmailInfo).Methods("GET")
	router.HandleFunc("/email", InsertEmail).Methods("POST")
	router.HandleFunc("/email", UpdateEmailInfoById).Methods("PUT")
	router.HandleFunc("/email/{emailId}", DeleteEmailInfo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12001", router))
}

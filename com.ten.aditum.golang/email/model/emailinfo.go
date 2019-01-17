package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type EmailInfo struct {
	Id               int64  `json:"id"`
	EmailId          string `json:"emailId"`
	EmailTitle       string `json:"emailTitle"`
	EmailContent     string `json:"emailContent"`
	SenderName       string `json:"senderName"`
	SenderAddress    string `json:"senderAddress"`
	RecipientName    string `json:"recipientName"`
	RecipientAddress string `json:"recipientAddress"`
	EmailCreateTime  string `json:"emailCreateTime"`
	EmailIsDeleted   int    `json:"emailIsDeleted"`
}

// 通过邮件DUID(email_id)获取邮件数据
func GetEmailInfoById(emailId string) EmailInfo {
	var email EmailInfo
	err := DB.QueryRow("SELECT * FROM email_info WHERE email_id = ?", emailId).Scan(
		&email.Id,
		&email.EmailId,
		&email.EmailTitle,
		&email.EmailContent,
		&email.SenderName,
		&email.SenderAddress,
		&email.RecipientName,
		&email.RecipientAddress,
		&email.EmailCreateTime,
		&email.EmailIsDeleted)
	if err != nil {
		fmt.Println("查询出错了，emailId=" + emailId)
	}
	return email
}

// 通过分页参数查询所有邮件数据
func GetAllEmailInfo() []EmailInfo {
	// 执行查询语句
	rows, err := DB.Query("SELECT * from email_info")
	if err != nil {
		fmt.Println("查询出错了")
	}
	var emails []EmailInfo
	// 循环读取结果
	for rows.Next() {
		var email EmailInfo
		// 将每一行的结果都赋值到一个user对象中
		err := rows.Scan(
			&email.Id,
			&email.EmailId,
			&email.EmailTitle,
			&email.EmailContent,
			&email.SenderName,
			&email.SenderAddress,
			&email.RecipientName,
			&email.RecipientAddress,
			&email.EmailCreateTime,
			&email.EmailIsDeleted)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		emails = append(emails, email)
	}
	return emails
}

// 插入EmailInfo并返回主键ID
func InsertEmail(m *EmailInfo) (id int64, err error) {
	res, err := Transaction(
		func(tx *sql.Tx, res sql.Result) (sql.Result, error) {
			// 准备sql语句
			stmt, err := tx.Prepare(
				"INSERT INTO email_info(" +
					"email_id," +
					"email_title," +
					"email_content," +
					"sender_name," +
					"sender_address," +
					"recipient_name," +
					"recipient_address," +
					"email_is_deleted) " +
					"VALUES(?,?,?,?,?,?,?,?)")
			if err != nil {
				msg := "Prepare fail"
				fmt.Println(msg, err)
				return nil, err
			}
			// 将参数传递到sql语句中并且执行
			res, err = stmt.Exec(
				m.EmailId,
				m.EmailTitle,
				m.EmailContent,
				m.SenderName,
				m.SenderAddress,
				m.RecipientName,
				m.RecipientAddress,
				m.EmailIsDeleted)
			if err != nil {
				msg := "Exec fail"
				fmt.Println(msg, err)
				return nil, err
			}
			return res, nil
		})
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// 根据DUID更新标题和内容
func UpdateEmailInfoById(m *EmailInfo) (err error) {
	_, err = Transaction(
		func(tx *sql.Tx, res sql.Result) (sql.Result, error) {
			// 准备sql语句
			stmt, err := tx.Prepare("UPDATE email_info SET email_title = ?, email_content = ? WHERE email_id = ?")
			if err != nil {
				msg := "Prepare fail"
				fmt.Println(msg, err)
				return nil, err
			}
			// 将参数传递到sql语句中并且执行
			res, err = stmt.Exec(m.EmailTitle, m.EmailContent, m.EmailId)
			if err != nil {
				msg := "Exec fail"
				fmt.Println(msg, err)
				return nil, err
			}
			return res, nil
		})
	if err != nil {
		return err
	}
	return nil
}

// 根据ID删除相应记录
func DeleteEmailInfo(id string) (err error) {
	_, err = Transaction(
		func(tx *sql.Tx, res sql.Result) (sql.Result, error) {
			// 准备sql语句
			stmt, err := tx.Prepare("UPDATE email_info SET email_is_deleted = 1 WHERE email_id = ?")
			if err != nil {
				msg := "Prepare fail"
				fmt.Println(msg, err)
				return nil, err
			}
			// 将参数传递到sql语句中并且执行
			res, err = stmt.Exec(id)
			if err != nil {
				msg := "Exec fail"
				fmt.Println(msg, err)
				return nil, err
			}
			return res, nil
		})
	if err != nil {
		return err
	}
	return nil
}

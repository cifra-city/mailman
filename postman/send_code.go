package postman

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"gopkg.in/gomail.v2"
)

func (p *Postman) SendCode(to string, code string, templateList string) error {
	mes := gomail.NewMessage()
	mes.SetHeader("From", p.Address)
	mes.SetHeader("To", to)
	mes.SetHeader("Subject", "Email Confirmation")

	templatePath := filepath.Join(templateList)
	tmplContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New("emailTemplate").Parse(string(tmplContent))
	if err != nil {
		return err
	}

	var bodyContent bytes.Buffer
	err = tmpl.Execute(&bodyContent, map[string]string{"Code": code})
	if err != nil {
		return err
	}

	mes.SetBody("text/html", bodyContent.String())

	smtpPort, err := strconv.Atoi(p.Port)
	if err != nil {
		return err
	}

	d := gomail.NewDialer(p.Host, smtpPort, p.Address, p.Password)

	if err := d.DialAndSend(mes); err != nil {
		return err
	}

	return nil
}

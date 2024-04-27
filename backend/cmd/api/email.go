package api

import (
	"errors"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/liamt01/course-scope-cmu/backend/.gen/course_scope/public/model"
	"github.com/liamt01/course-scope-cmu/backend/internal/schemas"
	"github.com/matcornic/hermes/v2"
	"net/smtp"
)

type emailServer struct {
	host         string
	port         string
	hostUser     string
	hostPassword string
	from         string
	generator    *hermes.Hermes
	frontendLink string
}

func newEmailServer(host, port, hostUser, hostPassword, from, frontendLink string) *emailServer {
	// Remove trailing slash from frontendLink
	if frontendLink[len(frontendLink)-1] == '/' {
		frontendLink = frontendLink[:len(frontendLink)-1]
	}
	return &emailServer{
		host:         host,
		port:         port,
		hostUser:     hostUser,
		hostPassword: hostPassword,
		from:         from,
		generator: &hermes.Hermes{
			Product: hermes.Product{
				Name:      "CourseScope CMU",
				Link:      frontendLink,
				Copyright: "Copyright © 2024 CourseScope CMU. All rights reserved.",
			},
		},
		frontendLink: frontendLink,
	}
}

func (es *emailServer) send(to string, subject, html, text string) error {
	em := email.NewEmail()
	em.From = es.from
	em.To = []string{to}
	em.Subject = subject
	em.HTML = []byte(html)
	em.Text = []byte(text)
	return em.Send(
		fmt.Sprintf("%v:%v", es.host, es.port),
		smtp.PlainAuth("", es.hostUser, es.hostPassword, es.host),
	)
}

func (es *emailServer) sendTokenToUser(user *model.Users, tokenOut *schemas.TokenOut) error {
	if user == nil {
		user = &model.Users{}
	}
	if tokenOut == nil {
		tokenOut = &schemas.TokenOut{}
	}

	to := fmt.Sprintf("%v@andrew.cmu.edu", user.AndrewID)

	switch tokenOut.Scope {
	case model.ScopeType_Act:
		e := hermes.Email{
			Body: hermes.Body{
				Name: user.Username,
				Intros: []string{
					"Welcome to CourseScope CMU!",
				},
				Actions: []hermes.Action{
					{
						Instructions: "To activate your account, please click here:",
						Button: hermes.Button{
							Color: "#22BC66",
							Text:  "Activate your account",
							Link:  fmt.Sprintf("%v/account/activate?token=%v", es.frontendLink, tokenOut.Token),
						},
					},
				},
				Outros: []string{
					fmt.Sprintf("This link will expire in %v.", tokenOut.TTL.String()),
				},
			},
		}

		emailBody, err := es.generator.GenerateHTML(e)
		if err != nil {
			return err
		}

		emailText, err := es.generator.GeneratePlainText(e)
		if err != nil {
			return err
		}

		if err := es.send(to, "Activate Your Account", emailBody, emailText); err != nil {
			return err
		}

		return nil
	case model.ScopeType_Pwd:
		e := hermes.Email{
			Body: hermes.Body{
				Name: user.Username,
				Intros: []string{
					"You have requested to reset your password.",
				},
				Actions: []hermes.Action{
					{
						Instructions: "To reset your password, please click here:",
						Button: hermes.Button{
							Color: "#22BC66",
							Text:  "Reset your password",
							Link:  fmt.Sprintf("%v/account/password-reset?token=%v", es.frontendLink, tokenOut.Token),
						},
					},
				},
				Outros: []string{
					fmt.Sprintf("This link will expire in %v.", tokenOut.TTL.String()),
				},
			},
		}

		emailBody, err := es.generator.GenerateHTML(e)
		if err != nil {
			return err
		}

		emailText, err := es.generator.GeneratePlainText(e)
		if err != nil {
			return err
		}

		if err := es.send(to, "Reset Your Password", emailBody, emailText); err != nil {
			return err
		}

		return nil
	}

	return errors.New("unsupported token type")
}

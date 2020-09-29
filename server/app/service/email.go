package service

import "server/library/utils"

func EmailTest() error {
	subject := "test"
	body := "test"
	return utils.EmailTest(subject, body)
}

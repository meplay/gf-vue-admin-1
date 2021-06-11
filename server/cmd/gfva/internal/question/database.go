package question

import "github.com/AlecAivazis/survey/v2"

var DatabaseQuestion = []*survey.Question{
	{
		Name: "database",
		Prompt: &survey.Select{
			Message: "您的配置文件所配置的数据库不存在!",
			Options: []string{"Link Start! gfva 为您创建数据库", "闪开!我自己来", "退出程序"},
			Default: "Link Start! gfva 为您创建数据库",
		},
	},
}

type DatabaseAnswer struct {
	Database string
}

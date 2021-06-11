// Package gfva
/*
Copyright © 2020 SliverHorn

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package gfva

import (
	"flipped-aurora/gf-vue-admin/server/boot"
	"flipped-aurora/gf-vue-admin/server/cmd/gfva/internal"
	"flipped-aurora/gf-vue-admin/server/library/constant"
	"github.com/spf13/cobra"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "flipped-aurora/gf-vue-admin/server初始化数据",
	Long:  `flipped-aurora/gf-vue-admin/server初始化数据适配数据库情况: 1. mysql完美适配, 2. postgresql未适配, 3. sqlite未适配, 4. sqlserver未适配`,
	Run: func(cmd *cobra.Command, args []string) {
		frame, _ := cmd.Flags().GetString("frame")
		path, _ := cmd.Flags().GetString("path")
		if frame == "gf" {
			boot.Viper.Initialize(path)
			internal.DbResolver.Database()
			//internal.Mysql.Check()
			//boot.DbResolver.Initialize()
			//if global.Config.System.DbType == "mysql" {
			//	if err := data.Initialize(); err != nil {
			//		color.Info.Println("\n[Mysql] --> 初始化数据成功!\n")
			//	}
			//}
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)
	initdbCmd.Flags().StringP("path", "p", constant.ConfigPostgresFile, "自定配置文件路径(绝对路径)")
	initdbCmd.Flags().StringP("frame", "f", "gf", "可选参数为gin,gf")
	initdbCmd.Flags().StringP("type", "t", "mysql", "可选参数为mysql,postgresql,sqlite,sqlserver")
}

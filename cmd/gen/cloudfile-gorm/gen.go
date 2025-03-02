package main

import (
	"github.com/yanguiyuan/yuan/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	c := config.NewConfig()
	db, err := gorm.Open(mysql.Open(c.GetString("cloudfile.mysql.dsn")), &gorm.Config{})
	if err != nil {
		return
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      c.GetString("cloudfile.gorm-gen.dal"),
		ModelPkgPath: c.GetString("cloudfile.gorm-gen.model"),
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)

	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

package main

import (
	"context"

	eUser "github.com/kevin-luvian/goauth/server/entity/user"
	pdb "github.com/kevin-luvian/goauth/server/pkg/db"
	"github.com/kevin-luvian/goauth/server/pkg/logging"
	"github.com/kevin-luvian/goauth/server/pkg/setting"
	"github.com/kevin-luvian/goauth/server/pkg/util"
	"github.com/kevin-luvian/goauth/server/repositories/user"
)

func init() {
	util.Setup()
	setting.Setup()
	logging.Setup()
}

func main() {
	logging.Infoln("playground started")

	ctx := context.Background()

	db, err := pdb.New(pdb.Config{
		SourceURL: setting.Database.LocalURL,
		Retries:   setting.Database.Retries,
	})
	if err != nil {
		logging.Fatalln("connection refused")
	}

	err = db.Instance.Ping()
	if err != nil {
		logging.Fatalln("ping error")
	}

	userRepo := user.New(db)

	_, initialLen, _ := userRepo.Get(ctx, pdb.GetDBParam{})

	rstr := util.RandString(11)
	created, err := userRepo.Create(ctx, eUser.User{
		Tag:   rstr,
		Name:  "john",
		Email: rstr,
		HPass: "",
	})
	if err != nil {
		logging.Fatalln("user repo create error", err)
	}

	_, currLen, _ := userRepo.Get(ctx, pdb.GetDBParam{})
	if currLen != initialLen+1 {
		logging.Fatalln("different users length")
	}

	err = userRepo.DeleteByID(ctx, created.ID)
	if err != nil {
		logging.Fatalln("user repo delete error", err)
	}

	_, currLen, _ = userRepo.Get(ctx, pdb.GetDBParam{})
	if currLen != initialLen {
		logging.Fatalln("different users length")
	}
}

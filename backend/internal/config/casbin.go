package config

import (
	"path/filepath"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

func InitCasbin(db *gorm.DB) (*casbin.Enforcer, error) {
	modelPath := filepath.Join("internal", "config", "casbin", "model.conf")
	policyPath := filepath.Join("internal", "config", "casbin", "policy.csv")

	// adapter, err := gormadapter.NewAdapterByDB(db)
	// if err != nil {
	// 	return nil, err
	// }

	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		return nil, err
	}

	err = enforcer.LoadPolicy()

	return enforcer, err
}

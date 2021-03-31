package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

const up20 = `
		CREATE TYPE service AS ENUM (
			'directrequest',
			'fluxmonitor',
			'gasupdater',
			'job',
			'keeper',
			'log',
			'ocr',
			'backup',
			'pipeline',
			'postgres',
			'signatures',
			'synchronization',
			'telemetry',
			'vrf',
			'head_broadcaster',
			'head_tracker',
			'job_subscriber',
			'prom_reporter',
			'reaper',
			'run_executor',
			'run_manager',
			'run_queue',
			'scheduler',
			'subscription',
			'validators'
		);

		CREATE TYPE log_level AS ENUM (
			'debug',
			'info',
			'warn',
			'error',
			'panic'
		);

		CREATE TABLE log_configs (
			"id" BIGSERIAL PRIMARY KEY,
			"service_name" service NOT NULL,
			"log_level" log_level NOT NULL,
   		 	"created_at" timestamp with time zone,
    		"updated_at" timestamp with time zone
		);
	`

const down20 = `
	DROP TABLE IF EXISTS log_config;

	DROP TABLE IF EXISTS log_services;
`

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "0020_create_log_config_table",
		Migrate: func(db *gorm.DB) error {
			return db.Exec(up20).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Exec(down20).Error
		},
	})
}

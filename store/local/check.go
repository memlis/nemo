package boltdb

import (
	"encoding/json"
	"github.com/memlis/boat/types"
)

func (b *BoltStore) SaveCheck(task *types.Task, port uint32, appId string) error {
	tx, err := b.conn.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte("checks"))

	for _, healthCheck := range task.HealthChecks {

		check := types.Check{
			ID:       task.Name,
			Address:  *task.AgentHostname,
			Port:     int(port),
			TaskID:   task.Name,
			AppID:    appId,
			Protocol: healthCheck.Protocol,
			Interval: int(healthCheck.IntervalSeconds),
			Timeout:  int(healthCheck.TimeoutSeconds),
		}

		if healthCheck.Command != nil {
			check.Command = healthCheck.Command
		}

		if healthCheck.Path != nil {
			check.Path = *healthCheck.Path
		}

		if healthCheck.MaxConsecutiveFailures != nil {
			check.MaxFailures = *healthCheck.MaxConsecutiveFailures
		}

		data, err := json.Marshal(&check)
		if err != nil {
			return err
		}

		if err := bucket.Put([]byte(check.ID), data); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (b *BoltStore) ListChecks() ([]*types.Check, error) {
	tx, err := b.conn.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte("checks"))

	var checks []*types.Check
	if err := bucket.ForEach(func(k, v []byte) error {
		var check types.Check
		if err := json.Unmarshal(v, &check); err != nil {
			return err
		}
		checks = append(checks, &check)

		return nil
	}); err != nil {
		return nil, err
	}

	return checks, nil
}

func (b *BoltStore) DeleteCheck(checkId string) error {
	tx, err := b.conn.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte("checks"))

	if err := bucket.Delete([]byte(checkId)); err != nil {
		return err
	}

	return tx.Commit()
}

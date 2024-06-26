// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: Basic Test
---
*/
package endpoints

import (
	"context"

	"github.com/swaggest/usecase"

	"github.com/365admin/koksmat-test1/execution"
)

func TestBasicPost() usecase.Interactor {
	type Request struct {
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *string) error {

		_, err := execution.ExecutePowerShell("john", "*", "koksmat-test1", "50-test", "10-basic.ps1", "")
		if err != nil {
			return err
		}

		return err

	})
	u.SetTitle("Basic Test")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Test")
	return u
}

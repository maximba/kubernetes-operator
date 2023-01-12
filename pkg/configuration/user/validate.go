package user

import (
	"github.com/maximba/kubernetes-operator/api/v1alpha2"
	"github.com/maximba/kubernetes-operator/pkg/configuration/backuprestore"
	"github.com/maximba/kubernetes-operator/pkg/configuration/user/seedjobs"
)

// Validate validates Jenkins CR Spec section
func (r *reconcileUserConfiguration) Validate(jenkins *v1alpha2.Jenkins) ([]string, error) {
	backupAndRestore := backuprestore.New(r.Configuration, r.logger)
	if msg := backupAndRestore.Validate(); msg != nil {
		return msg, nil
	}

	seedJobs := seedjobs.New(r.jenkinsClient, r.Configuration)
	return seedJobs.ValidateSeedJobs(*jenkins)
}

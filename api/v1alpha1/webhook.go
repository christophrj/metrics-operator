package v1alpha1

import (
	"context"
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *Metric) Default(ctx context.Context, obj runtime.Object) error {
	m, ok := obj.(*Metric)
	if !ok {
		return errors.New("not a metric")
	}
	if m.Spec.DataSinkRef == nil || m.Spec.DataSinkRef.Name == "" {
		m.Spec.DataSinkRef = &DataSinkReference{
			Name: "default",
		}
	}
	return nil
}

func SetupMetricsWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(&Metric{}).
		WithDefaulter(&Metric{}).
		Complete()
}

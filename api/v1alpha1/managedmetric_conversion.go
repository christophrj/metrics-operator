package v1alpha1

import (
	"github.com/openmcp-project/metrics-operator/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *ManagedMetric) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.ManagedMetric)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.DataSinkRef = (*v1alpha2.DataSinkReference)(src.Spec.DataSinkRef)
	dst.Spec.Description = src.Spec.Description
	dst.Spec.FieldSelector = src.Spec.FieldSelector
	dst.Spec.Interval = src.Spec.Interval
	dst.Spec.LabelSelector = src.Spec.LabelSelector
	dst.Spec.Name = src.Spec.Name
	dst.Spec.RemoteClusterAccessRef = (*v1alpha2.RemoteClusterAccessRef)(src.Spec.RemoteClusterAccessRef)
	dst.Spec.Target = v1alpha2.GroupVersionKind{
		Kind:    src.Spec.Kind,
		Group:   src.Spec.Group,
		Version: src.Spec.Version,
	}

	// Status
	dst.Status.Conditions = src.Status.Conditions
	dst.Status.Observation = v1alpha2.ManagedObservation(src.Status.Observation)
	dst.Status.Ready = src.Status.Ready

	return nil
}

func (dst *ManagedMetric) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.ManagedMetric)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.DataSinkRef = (*DataSinkReference)(src.Spec.DataSinkRef)
	dst.Spec.Description = src.Spec.Description
	dst.Spec.FieldSelector = src.Spec.FieldSelector
	dst.Spec.Interval = src.Spec.Interval
	dst.Spec.LabelSelector = src.Spec.LabelSelector
	dst.Spec.Name = src.Spec.Name
	dst.Spec.RemoteClusterAccessRef = (*RemoteClusterAccessRef)(src.Spec.RemoteClusterAccessRef)
	dst.Spec.Group = src.Spec.Target.Group
	dst.Spec.Version = src.Spec.Target.Version
	dst.Spec.Kind = src.Spec.Target.Kind

	// Status
	dst.Status.Conditions = src.Status.Conditions
	dst.Status.Observation = ManagedObservation(src.Status.Observation)
	dst.Status.Ready = src.Status.Ready

	return nil
}

func (*ManagedMetric) CustomResourceDefinitionName() string {
	return "managedmetrics.metrics.openmcp.cloud"
}

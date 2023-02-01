/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta17 "github.com/upbound/provider-azure/apis/devices/v1beta1"
	v1beta16 "github.com/upbound/provider-azure/apis/eventhub/v1beta1"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	v1beta14 "github.com/upbound/provider-azure/apis/servicebus/v1beta1"
	v1beta13 "github.com/upbound/provider-azure/apis/sql/v1beta1"
	v1beta11 "github.com/upbound/provider-azure/apis/storage/v1beta1"
	v1beta15 "github.com/upbound/provider-azure/apis/synapse/v1beta1"
	v1beta12 "github.com/upbound/provider-azure/apis/web/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FunctionJavascriptUda.
func (mg *FunctionJavascriptUda) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamAnalyticsJobID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.StreamAnalyticsJobIDRef,
		Selector:     mg.Spec.ForProvider.StreamAnalyticsJobIDSelector,
		To: reference.To{
			List:    &JobList{},
			Managed: &Job{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamAnalyticsJobID")
	}
	mg.Spec.ForProvider.StreamAnalyticsJobID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamAnalyticsJobIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Job.
func (mg *Job) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ManagedPrivateEndpoint.
func (mg *ManagedPrivateEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamAnalyticsClusterName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StreamAnalyticsClusterNameRef,
		Selector:     mg.Spec.ForProvider.StreamAnalyticsClusterNameSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamAnalyticsClusterName")
	}
	mg.Spec.ForProvider.StreamAnalyticsClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamAnalyticsClusterNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetResourceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TargetResourceIDRef,
		Selector:     mg.Spec.ForProvider.TargetResourceIDSelector,
		To: reference.To{
			List:    &v1beta11.AccountList{},
			Managed: &v1beta11.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetResourceID")
	}
	mg.Spec.ForProvider.TargetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OutputBlob.
func (mg *OutputBlob) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageAccountNameRef,
		Selector:     mg.Spec.ForProvider.StorageAccountNameSelector,
		To: reference.To{
			List:    &v1beta11.AccountList{},
			Managed: &v1beta11.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountName")
	}
	mg.Spec.ForProvider.StorageAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageContainerName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageContainerNameRef,
		Selector:     mg.Spec.ForProvider.StorageContainerNameSelector,
		To: reference.To{
			List:    &v1beta11.ContainerList{},
			Managed: &v1beta11.Container{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageContainerName")
	}
	mg.Spec.ForProvider.StorageContainerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageContainerNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamAnalyticsJobName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StreamAnalyticsJobNameRef,
		Selector:     mg.Spec.ForProvider.StreamAnalyticsJobNameSelector,
		To: reference.To{
			List:    &JobList{},
			Managed: &Job{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamAnalyticsJobName")
	}
	mg.Spec.ForProvider.StreamAnalyticsJobName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamAnalyticsJobNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OutputFunction.
func (mg *OutputFunction) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FunctionApp),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.FunctionAppRef,
		Selector:     mg.Spec.ForProvider.FunctionAppSelector,
		To: reference.To{
			List:    &v1beta12.FunctionAppList{},
			Managed: &v1beta12.FunctionApp{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FunctionApp")
	}
	mg.Spec.ForProvider.FunctionApp = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FunctionAppRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamAnalyticsJobName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StreamAnalyticsJobNameRef,
		Selector:     mg.Spec.ForProvider.StreamAnalyticsJobNameSelector,
		To: reference.To{
			List:    &JobList{},
			Managed: &Job{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamAnalyticsJobName")
	}
	mg.Spec.ForProvider.StreamAnalyticsJobName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamAnalyticsJobNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OutputMSSQL.
func (mg *OutputMSSQL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Server),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServerRef,
		Selector:     mg.Spec.ForProvider.ServerSelector,
		To: reference.To{
			List:    &v1beta13.MSSQLServerList{},
			Managed: &v1beta13.MSSQLServer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Server")
	}
	mg.Spec.ForProvider.Server = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamAnalyticsJobName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StreamAnalyticsJobNameRef,
		Selector:     mg.Spec.ForProvider.StreamAnalyticsJobNameSelector,
		To: reference.To{
			List:    &JobList{},
			Managed: &Job{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamAnalyticsJobName")
	}
	mg.Spec.ForProvider.StreamAnalyticsJobName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamAnalyticsJobNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Table),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TableRef,
		Selector:     mg.Spec.ForProvider.TableSelector,
		To: reference.To{
			List:    &v1beta11.TableList{},
			Managed: &v1beta11.Table{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Table")
	}
	mg.Spec.ForProvider.Table = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TableRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OutputServiceBusQueue.
func (mg *OutputServiceBusQueue) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.QueueName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.QueueNameRef,
		Selector:     mg.Spec.ForProvider.QueueNameSelector,
		To: reference.To{
			List:    &v1beta14.QueueList{},
			Managed: &v1beta14.Queue{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.QueueName")
	}
	mg.Spec.ForProvider.QueueName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.QueueNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceBusNamespace),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceBusNamespaceRef,
		Selector:     mg.Spec.ForProvider.ServiceBusNamespaceSelector,
		To: reference.To{
			List:    &v1beta14.ServiceBusNamespaceList{},
			Managed: &v1beta14.ServiceBusNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceBusNamespace")
	}
	mg.Spec.ForProvider.ServiceBusNamespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceBusNamespaceRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamAnalyticsJobName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StreamAnalyticsJobNameRef,
		Selector:     mg.Spec.ForProvider.StreamAnalyticsJobNameSelector,
		To: reference.To{
			List:    &JobList{},
			Managed: &Job{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamAnalyticsJobName")
	}
	mg.Spec.ForProvider.StreamAnalyticsJobName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamAnalyticsJobNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OutputServiceBusTopic.
func (mg *OutputServiceBusTopic) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceBusNamespace),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceBusNamespaceRef,
		Selector:     mg.Spec.ForProvider.ServiceBusNamespaceSelector,
		To: reference.To{
			List:    &v1beta14.ServiceBusNamespaceList{},
			Managed: &v1beta14.ServiceBusNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceBusNamespace")
	}
	mg.Spec.ForProvider.ServiceBusNamespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceBusNamespaceRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamAnalyticsJobName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StreamAnalyticsJobNameRef,
		Selector:     mg.Spec.ForProvider.StreamAnalyticsJobNameSelector,
		To: reference.To{
			List:    &JobList{},
			Managed: &Job{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamAnalyticsJobName")
	}
	mg.Spec.ForProvider.StreamAnalyticsJobName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamAnalyticsJobNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TopicName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TopicNameRef,
		Selector:     mg.Spec.ForProvider.TopicNameSelector,
		To: reference.To{
			List:    &v1beta14.TopicList{},
			Managed: &v1beta14.Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TopicName")
	}
	mg.Spec.ForProvider.TopicName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OutputSynapse.
func (mg *OutputSynapse) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamAnalyticsJobName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StreamAnalyticsJobNameRef,
		Selector:     mg.Spec.ForProvider.StreamAnalyticsJobNameSelector,
		To: reference.To{
			List:    &JobList{},
			Managed: &Job{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamAnalyticsJobName")
	}
	mg.Spec.ForProvider.StreamAnalyticsJobName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamAnalyticsJobNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.User),
		Extract:      resource.ExtractParamPath("sql_administrator_login", false),
		Reference:    mg.Spec.ForProvider.UserRef,
		Selector:     mg.Spec.ForProvider.UserSelector,
		To: reference.To{
			List:    &v1beta15.WorkspaceList{},
			Managed: &v1beta15.Workspace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.User")
	}
	mg.Spec.ForProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ReferenceInputBlob.
func (mg *ReferenceInputBlob) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageAccountNameRef,
		Selector:     mg.Spec.ForProvider.StorageAccountNameSelector,
		To: reference.To{
			List:    &v1beta11.AccountList{},
			Managed: &v1beta11.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountName")
	}
	mg.Spec.ForProvider.StorageAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageContainerName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageContainerNameRef,
		Selector:     mg.Spec.ForProvider.StorageContainerNameSelector,
		To: reference.To{
			List:    &v1beta11.ContainerList{},
			Managed: &v1beta11.Container{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageContainerName")
	}
	mg.Spec.ForProvider.StorageContainerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageContainerNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamAnalyticsJobName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StreamAnalyticsJobNameRef,
		Selector:     mg.Spec.ForProvider.StreamAnalyticsJobNameSelector,
		To: reference.To{
			List:    &JobList{},
			Managed: &Job{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamAnalyticsJobName")
	}
	mg.Spec.ForProvider.StreamAnalyticsJobName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamAnalyticsJobNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StreamInputBlob.
func (mg *StreamInputBlob) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageAccountNameRef,
		Selector:     mg.Spec.ForProvider.StorageAccountNameSelector,
		To: reference.To{
			List:    &v1beta11.AccountList{},
			Managed: &v1beta11.Account{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountName")
	}
	mg.Spec.ForProvider.StorageAccountName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageContainerName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StorageContainerNameRef,
		Selector:     mg.Spec.ForProvider.StorageContainerNameSelector,
		To: reference.To{
			List:    &v1beta11.ContainerList{},
			Managed: &v1beta11.Container{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageContainerName")
	}
	mg.Spec.ForProvider.StorageContainerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageContainerNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamAnalyticsJobName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StreamAnalyticsJobNameRef,
		Selector:     mg.Spec.ForProvider.StreamAnalyticsJobNameSelector,
		To: reference.To{
			List:    &JobList{},
			Managed: &Job{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamAnalyticsJobName")
	}
	mg.Spec.ForProvider.StreamAnalyticsJobName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamAnalyticsJobNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StreamInputEventHub.
func (mg *StreamInputEventHub) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHubConsumerGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EventHubConsumerGroupNameRef,
		Selector:     mg.Spec.ForProvider.EventHubConsumerGroupNameSelector,
		To: reference.To{
			List:    &v1beta16.ConsumerGroupList{},
			Managed: &v1beta16.ConsumerGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventHubConsumerGroupName")
	}
	mg.Spec.ForProvider.EventHubConsumerGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventHubConsumerGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHubName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EventHubNameRef,
		Selector:     mg.Spec.ForProvider.EventHubNameSelector,
		To: reference.To{
			List:    &v1beta16.EventHubList{},
			Managed: &v1beta16.EventHub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventHubName")
	}
	mg.Spec.ForProvider.EventHubName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventHubNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceBusNamespace),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceBusNamespaceRef,
		Selector:     mg.Spec.ForProvider.ServiceBusNamespaceSelector,
		To: reference.To{
			List:    &v1beta16.EventHubNamespaceList{},
			Managed: &v1beta16.EventHubNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceBusNamespace")
	}
	mg.Spec.ForProvider.ServiceBusNamespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceBusNamespaceRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamAnalyticsJobName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StreamAnalyticsJobNameRef,
		Selector:     mg.Spec.ForProvider.StreamAnalyticsJobNameSelector,
		To: reference.To{
			List:    &JobList{},
			Managed: &Job{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamAnalyticsJobName")
	}
	mg.Spec.ForProvider.StreamAnalyticsJobName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamAnalyticsJobNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StreamInputIOTHub.
func (mg *StreamInputIOTHub) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHubConsumerGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EventHubConsumerGroupNameRef,
		Selector:     mg.Spec.ForProvider.EventHubConsumerGroupNameSelector,
		To: reference.To{
			List:    &v1beta16.ConsumerGroupList{},
			Managed: &v1beta16.ConsumerGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventHubConsumerGroupName")
	}
	mg.Spec.ForProvider.EventHubConsumerGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventHubConsumerGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IOTHubNamespace),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.IOTHubNamespaceRef,
		Selector:     mg.Spec.ForProvider.IOTHubNamespaceSelector,
		To: reference.To{
			List:    &v1beta17.IOTHubList{},
			Managed: &v1beta17.IOTHub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IOTHubNamespace")
	}
	mg.Spec.ForProvider.IOTHubNamespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IOTHubNamespaceRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StreamAnalyticsJobName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.StreamAnalyticsJobNameRef,
		Selector:     mg.Spec.ForProvider.StreamAnalyticsJobNameSelector,
		To: reference.To{
			List:    &JobList{},
			Managed: &Job{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StreamAnalyticsJobName")
	}
	mg.Spec.ForProvider.StreamAnalyticsJobName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StreamAnalyticsJobNameRef = rsp.ResolvedReference

	return nil
}

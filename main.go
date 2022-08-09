package ck8slibrary

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/coopnorge/cdk8slibrary/imports/k8s"
	"github.com/coopnorge/cdk8slibrary/imports/networkingistioio"
	"github.com/creasty/defaults"
)

type ChartProps struct {
	Name                         *string  `default:"-" mapstructure:"name"`                          // Name of the container
	Image                        *string  `default:"-" mapstructure:"image"`                         // Image of the container
	Replicas                     *float64 `default:"2" mapstructure:"replicas"`                      // Number of replicas
	MaxReplicas                  *float64 `default:"-" mapstructure:"maxReplicas"`                   // Set max number of replica's. This will enable autoscaling containers.
	AutoScaleTargetCPUPercentage *float64 `default:"50" mapstructure:"autoScaleTargetCPUPercentage"` // If autoscaling is enabled, this is the CPU percentage it will start to scale
	PdbMinAvailable              *string  `default:"1" mapstructure:"pdbMinAvailable"`               // Minimal availble pods when doing maintaince.

	Ports          []ChartPorts     `default:"[]" mapstructure:"ports"`
	ServiceEntries []ServiceEntries `default:"[]" mapstructure:"serviceEntries"`
}

type ChartPorts struct {
	Port          *float64 `default:"-" mapstructure:"port"`
	ContainerPort *float64 `default:"-" mapstructure:"containerPort"`
}

type ServiceEntries struct {
	Name  *string   `default:"-" mapstructure:"name"`
	Hosts []*string `default:"[]" mapstructure:"hosts"`
}

func Chart(scope constructs.Construct, id *string, props *ChartProps) constructs.Construct {

	if err := defaults.Set(props); err != nil {
		panic(err)
	}

	construct := constructs.NewConstruct(scope, id)

	label := map[string]*string{
		"app": constructs.Node_Of(construct).Id(),
	}

	metadata := k8s.ObjectMeta{
		Name:   props.Name,
		Labels: &label,
	}

	k8s.NewKubeDeployment(construct, jsii.String("deployment"), &k8s.KubeDeploymentProps{
		Metadata: &metadata,
		Spec: &k8s.DeploymentSpec{
			Replicas: props.replicaBuilder(),
			Selector: &k8s.LabelSelector{MatchLabels: &label},
			Template: &k8s.PodTemplateSpec{
				Metadata: &k8s.ObjectMeta{Labels: &label},
				Spec: &k8s.PodSpec{
					Containers: &[]*k8s.Container{{
						Name:  jsii.String(*id),
						Image: props.Image,
						Ports: props.containerPortBuilder(),
					}},
				},
			},
		},
	})

	k8s.NewKubeService(construct, jsii.String("service"), &k8s.KubeServiceProps{
		Metadata: &metadata,
		Spec: &k8s.ServiceSpec{
			Type:     jsii.String("ClusterIP"),
			Ports:    props.servicePortBuilder(),
			Selector: &label,
		},
	})

	if props.MaxReplicas != nil {
		k8s.NewKubeHorizontalPodAutoscaler(construct, jsii.String("hpa"), &k8s.KubeHorizontalPodAutoscalerProps{
			Metadata: &metadata,
			Spec: &k8s.HorizontalPodAutoscalerSpec{
				MinReplicas:                    props.Replicas,
				MaxReplicas:                    props.MaxReplicas,
				TargetCpuUtilizationPercentage: props.AutoScaleTargetCPUPercentage,
				ScaleTargetRef: &k8s.CrossVersionObjectReference{
					Kind: jsii.String("Deployment"),
					Name: props.Name,
				},
			},
		})
	}

	k8s.NewKubePodDisruptionBudget(construct, jsii.String("pdb"), &k8s.KubePodDisruptionBudgetProps{
		Metadata: &metadata,
		Spec: &k8s.PodDisruptionBudgetSpec{
			MinAvailable: k8s.IntOrString_FromString(props.PdbMinAvailable),
			Selector:     &k8s.LabelSelector{MatchLabels: &label},
		},
	})

	for _, serviceEntry := range props.ServiceEntries {
		networkingistioio.NewServiceEntry(construct, serviceEntry.Name, &networkingistioio.ServiceEntryProps{
			Metadata: &cdk8s.ApiObjectMetadata{
				Name:   serviceEntry.Name,
				Labels: &label,
			},
			Spec: &networkingistioio.ServiceEntrySpec{
				Hosts: &serviceEntry.Hosts,
			},
		})
	}
	return construct

}

func (props *ChartProps) containerPortBuilder() *[]*k8s.ContainerPort {
	var containerPorts []*k8s.ContainerPort
	for _, port := range props.Ports {

		if port.ContainerPort == nil {
			containerPorts = append(containerPorts, &k8s.ContainerPort{ContainerPort: port.Port})
		} else {
			containerPorts = append(containerPorts, &k8s.ContainerPort{ContainerPort: port.ContainerPort})
		}
	}
	return &containerPorts
}

func (props *ChartProps) servicePortBuilder() *[]*k8s.ServicePort {
	var servicePorts []*k8s.ServicePort
	for _, port := range props.Ports {

		if port.ContainerPort == nil {
			servicePorts = append(servicePorts, &k8s.ServicePort{Port: port.Port, TargetPort: k8s.IntOrString_FromNumber(port.Port)})
		} else {
			servicePorts = append(servicePorts, &k8s.ServicePort{Port: port.Port, TargetPort: k8s.IntOrString_FromNumber(port.ContainerPort)})
		}
	}
	return &servicePorts
}

func (props *ChartProps) replicaBuilder() *float64 {
	if props.MaxReplicas == nil {
		return props.Replicas
	}
	return nil
}

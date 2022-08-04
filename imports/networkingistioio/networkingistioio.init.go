package networkingistioio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"networkingistioio.DestinationRule",
		reflect.TypeOf((*DestinationRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DestinationRule{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleProps",
		reflect.TypeOf((*DestinationRuleProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpec",
		reflect.TypeOf((*DestinationRuleSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsets",
		reflect.TypeOf((*DestinationRuleSpecSubsets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicy",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyConnectionPool",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttp",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DEFAULT,
			"DO_NOT_UPGRADE": DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE,
			"UPGRADE": DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy_UPGRADE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolTcp",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolTcpTcpKeepalive",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyConnectionPoolTcpTcpKeepalive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancer",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerConsistentHash",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerConsistentHash)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerConsistentHashHttpCookie",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerConsistentHashHttpCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSetting",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingDistribute",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingDistribute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingFailover",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingFailover)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple)(nil)).Elem(),
		map[string]interface{}{
			"UNSPECIFIED": DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_UNSPECIFIED,
			"LEAST_CONN": DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_LEAST_CONN,
			"RANDOM": DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_RANDOM,
			"PASSTHROUGH": DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_PASSTHROUGH,
			"ROUND_ROBIN": DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_ROUND_ROBIN,
			"LEAST_REQUEST": DestinationRuleSpecSubsetsTrafficPolicyLoadBalancerSimple_LEAST_REQUEST,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyOutlierDetection",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettings",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPool",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttp",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DEFAULT,
			"DO_NOT_UPGRADE": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE,
			"UPGRADE": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_UPGRADE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcp",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancer",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple)(nil)).Elem(),
		map[string]interface{}{
			"UNSPECIFIED": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_UNSPECIFIED,
			"LEAST_CONN": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_LEAST_CONN,
			"RANDOM": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_RANDOM,
			"PASSTHROUGH": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_PASSTHROUGH,
			"ROUND_ROBIN": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_ROUND_ROBIN,
			"LEAST_REQUEST": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_LEAST_REQUEST,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsOutlierDetection",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsPort",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTls",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_DISABLE,
			"SIMPLE": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_SIMPLE,
			"MUTUAL": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_MUTUAL,
			"ISTIO_MUTUAL": DestinationRuleSpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyTls",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyTlsMode",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE": DestinationRuleSpecSubsetsTrafficPolicyTlsMode_DISABLE,
			"SIMPLE": DestinationRuleSpecSubsetsTrafficPolicyTlsMode_SIMPLE,
			"MUTUAL": DestinationRuleSpecSubsetsTrafficPolicyTlsMode_MUTUAL,
			"ISTIO_MUTUAL": DestinationRuleSpecSubsetsTrafficPolicyTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecSubsetsTrafficPolicyTunnel",
		reflect.TypeOf((*DestinationRuleSpecSubsetsTrafficPolicyTunnel)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicy",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyConnectionPool",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyConnectionPoolHttp",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DEFAULT,
			"DO_NOT_UPGRADE": DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE,
			"UPGRADE": DestinationRuleSpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy_UPGRADE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyConnectionPoolTcp",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyConnectionPoolTcpTcpKeepalive",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyConnectionPoolTcpTcpKeepalive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancer",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancerConsistentHash",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancerConsistentHash)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancerConsistentHashHttpCookie",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancerConsistentHashHttpCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSetting",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSettingDistribute",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSettingDistribute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSettingFailover",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancerLocalityLbSettingFailover)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecTrafficPolicyLoadBalancerSimple",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyLoadBalancerSimple)(nil)).Elem(),
		map[string]interface{}{
			"UNSPECIFIED": DestinationRuleSpecTrafficPolicyLoadBalancerSimple_UNSPECIFIED,
			"LEAST_CONN": DestinationRuleSpecTrafficPolicyLoadBalancerSimple_LEAST_CONN,
			"RANDOM": DestinationRuleSpecTrafficPolicyLoadBalancerSimple_RANDOM,
			"PASSTHROUGH": DestinationRuleSpecTrafficPolicyLoadBalancerSimple_PASSTHROUGH,
			"ROUND_ROBIN": DestinationRuleSpecTrafficPolicyLoadBalancerSimple_ROUND_ROBIN,
			"LEAST_REQUEST": DestinationRuleSpecTrafficPolicyLoadBalancerSimple_LEAST_REQUEST,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyOutlierDetection",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettings",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPool",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttp",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DEFAULT,
			"DO_NOT_UPGRADE": DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE,
			"UPGRADE": DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_UPGRADE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolTcp",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancer",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple)(nil)).Elem(),
		map[string]interface{}{
			"UNSPECIFIED": DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_UNSPECIFIED,
			"LEAST_CONN": DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_LEAST_CONN,
			"RANDOM": DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_RANDOM,
			"PASSTHROUGH": DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_PASSTHROUGH,
			"ROUND_ROBIN": DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_ROUND_ROBIN,
			"LEAST_REQUEST": DestinationRuleSpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_LEAST_REQUEST,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsOutlierDetection",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsPort",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsTls",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE": DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode_DISABLE,
			"SIMPLE": DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode_SIMPLE,
			"MUTUAL": DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode_MUTUAL,
			"ISTIO_MUTUAL": DestinationRuleSpecTrafficPolicyPortLevelSettingsTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyTls",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleSpecTrafficPolicyTlsMode",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE": DestinationRuleSpecTrafficPolicyTlsMode_DISABLE,
			"SIMPLE": DestinationRuleSpecTrafficPolicyTlsMode_SIMPLE,
			"MUTUAL": DestinationRuleSpecTrafficPolicyTlsMode_MUTUAL,
			"ISTIO_MUTUAL": DestinationRuleSpecTrafficPolicyTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecTrafficPolicyTunnel",
		reflect.TypeOf((*DestinationRuleSpecTrafficPolicyTunnel)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleSpecWorkloadSelector",
		reflect.TypeOf((*DestinationRuleSpecWorkloadSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.DestinationRuleV1Beta1",
		reflect.TypeOf((*DestinationRuleV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DestinationRuleV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1Props",
		reflect.TypeOf((*DestinationRuleV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1Spec",
		reflect.TypeOf((*DestinationRuleV1Beta1Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsets",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicy",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPool",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPoolHttp",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DEFAULT,
			"DO_NOT_UPGRADE": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE,
			"UPGRADE": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPoolHttpH2UpgradePolicy_UPGRADE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPoolTcp",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPoolTcpTcpKeepalive",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyConnectionPoolTcpTcpKeepalive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancer",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerConsistentHash",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerConsistentHash)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerConsistentHashHttpCookie",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerConsistentHashHttpCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerLocalityLbSetting",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerLocalityLbSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingDistribute",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingDistribute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingFailover",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerLocalityLbSettingFailover)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerSimple",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerSimple)(nil)).Elem(),
		map[string]interface{}{
			"UNSPECIFIED": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerSimple_UNSPECIFIED,
			"LEAST_CONN": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerSimple_LEAST_CONN,
			"RANDOM": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerSimple_RANDOM,
			"PASSTHROUGH": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerSimple_PASSTHROUGH,
			"ROUND_ROBIN": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerSimple_ROUND_ROBIN,
			"LEAST_REQUEST": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyLoadBalancerSimple_LEAST_REQUEST,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyOutlierDetection",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettings",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPool",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttp",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DEFAULT,
			"DO_NOT_UPGRADE": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE,
			"UPGRADE": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_UPGRADE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcp",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancer",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple)(nil)).Elem(),
		map[string]interface{}{
			"UNSPECIFIED": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_UNSPECIFIED,
			"LEAST_CONN": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_LEAST_CONN,
			"RANDOM": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_RANDOM,
			"PASSTHROUGH": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_PASSTHROUGH,
			"ROUND_ROBIN": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_ROUND_ROBIN,
			"LEAST_REQUEST": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsLoadBalancerSimple_LEAST_REQUEST,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsOutlierDetection",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsPort",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsTls",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsTlsMode",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_DISABLE,
			"SIMPLE": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_SIMPLE,
			"MUTUAL": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_MUTUAL,
			"ISTIO_MUTUAL": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyPortLevelSettingsTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyTls",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyTlsMode",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyTlsMode_DISABLE,
			"SIMPLE": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyTlsMode_SIMPLE,
			"MUTUAL": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyTlsMode_MUTUAL,
			"ISTIO_MUTUAL": DestinationRuleV1Beta1SpecSubsetsTrafficPolicyTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecSubsetsTrafficPolicyTunnel",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecSubsetsTrafficPolicyTunnel)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicy",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyConnectionPool",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyConnectionPoolHttp",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": DestinationRuleV1Beta1SpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DEFAULT,
			"DO_NOT_UPGRADE": DestinationRuleV1Beta1SpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE,
			"UPGRADE": DestinationRuleV1Beta1SpecTrafficPolicyConnectionPoolHttpH2UpgradePolicy_UPGRADE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyConnectionPoolTcp",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyConnectionPoolTcpTcpKeepalive",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyConnectionPoolTcpTcpKeepalive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancer",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerConsistentHash",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerConsistentHash)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerConsistentHashHttpCookie",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerConsistentHashHttpCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerLocalityLbSetting",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerLocalityLbSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerLocalityLbSettingDistribute",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerLocalityLbSettingDistribute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerLocalityLbSettingFailover",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerLocalityLbSettingFailover)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerSimple",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerSimple)(nil)).Elem(),
		map[string]interface{}{
			"UNSPECIFIED": DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerSimple_UNSPECIFIED,
			"LEAST_CONN": DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerSimple_LEAST_CONN,
			"RANDOM": DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerSimple_RANDOM,
			"PASSTHROUGH": DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerSimple_PASSTHROUGH,
			"ROUND_ROBIN": DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerSimple_ROUND_ROBIN,
			"LEAST_REQUEST": DestinationRuleV1Beta1SpecTrafficPolicyLoadBalancerSimple_LEAST_REQUEST,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyOutlierDetection",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettings",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPool",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPoolHttp",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DEFAULT,
			"DO_NOT_UPGRADE": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_DO_NOT_UPGRADE,
			"UPGRADE": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPoolHttpH2UpgradePolicy_UPGRADE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPoolTcp",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsConnectionPoolTcpTcpKeepalive)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancer",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHash)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerConsistentHashHttpCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSetting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingDistribute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerLocalityLbSettingFailover)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerSimple",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerSimple)(nil)).Elem(),
		map[string]interface{}{
			"UNSPECIFIED": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_UNSPECIFIED,
			"LEAST_CONN": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_LEAST_CONN,
			"RANDOM": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_RANDOM,
			"PASSTHROUGH": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_PASSTHROUGH,
			"ROUND_ROBIN": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_ROUND_ROBIN,
			"LEAST_REQUEST": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsLoadBalancerSimple_LEAST_REQUEST,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsOutlierDetection",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsPort",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsTls",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsTlsMode",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsTlsMode_DISABLE,
			"SIMPLE": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsTlsMode_SIMPLE,
			"MUTUAL": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsTlsMode_MUTUAL,
			"ISTIO_MUTUAL": DestinationRuleV1Beta1SpecTrafficPolicyPortLevelSettingsTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyTls",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyTlsMode",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"DISABLE": DestinationRuleV1Beta1SpecTrafficPolicyTlsMode_DISABLE,
			"SIMPLE": DestinationRuleV1Beta1SpecTrafficPolicyTlsMode_SIMPLE,
			"MUTUAL": DestinationRuleV1Beta1SpecTrafficPolicyTlsMode_MUTUAL,
			"ISTIO_MUTUAL": DestinationRuleV1Beta1SpecTrafficPolicyTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecTrafficPolicyTunnel",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecTrafficPolicyTunnel)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.DestinationRuleV1Beta1SpecWorkloadSelector",
		reflect.TypeOf((*DestinationRuleV1Beta1SpecWorkloadSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.EnvoyFilter",
		reflect.TypeOf((*EnvoyFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EnvoyFilter{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterProps",
		reflect.TypeOf((*EnvoyFilterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpec",
		reflect.TypeOf((*EnvoyFilterSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatches",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatches)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.EnvoyFilterSpecConfigPatchesApplyTo",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesApplyTo)(nil)).Elem(),
		map[string]interface{}{
			"INVALID": EnvoyFilterSpecConfigPatchesApplyTo_INVALID,
			"LISTENER": EnvoyFilterSpecConfigPatchesApplyTo_LISTENER,
			"FILTER_CHAIN": EnvoyFilterSpecConfigPatchesApplyTo_FILTER_CHAIN,
			"NETWORK_FILTER": EnvoyFilterSpecConfigPatchesApplyTo_NETWORK_FILTER,
			"HTTP_FILTER": EnvoyFilterSpecConfigPatchesApplyTo_HTTP_FILTER,
			"ROUTE_CONFIGURATION": EnvoyFilterSpecConfigPatchesApplyTo_ROUTE_CONFIGURATION,
			"VIRTUAL_HOST": EnvoyFilterSpecConfigPatchesApplyTo_VIRTUAL_HOST,
			"HTTP_ROUTE": EnvoyFilterSpecConfigPatchesApplyTo_HTTP_ROUTE,
			"CLUSTER": EnvoyFilterSpecConfigPatchesApplyTo_CLUSTER,
			"EXTENSION_CONFIG": EnvoyFilterSpecConfigPatchesApplyTo_EXTENSION_CONFIG,
			"BOOTSTRAP": EnvoyFilterSpecConfigPatchesApplyTo_BOOTSTRAP,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatch",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchCluster",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchCluster)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchContext",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchContext)(nil)).Elem(),
		map[string]interface{}{
			"ANY": EnvoyFilterSpecConfigPatchesMatchContext_ANY,
			"SIDECAR_INBOUND": EnvoyFilterSpecConfigPatchesMatchContext_SIDECAR_INBOUND,
			"SIDECAR_OUTBOUND": EnvoyFilterSpecConfigPatchesMatchContext_SIDECAR_OUTBOUND,
			"GATEWAY": EnvoyFilterSpecConfigPatchesMatchContext_GATEWAY,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchListener",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchListener)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchListenerFilterChain",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchListenerFilterChain)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilter",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilterSubFilter",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilterSubFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchProxy",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchProxy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchRouteConfiguration",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchRouteConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhost",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhost)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRoute",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRoute)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction)(nil)).Elem(),
		map[string]interface{}{
			"ANY": EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction_ANY,
			"ROUTE": EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction_ROUTE,
			"REDIRECT": EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction_REDIRECT,
			"DIRECT_RESPONSE": EnvoyFilterSpecConfigPatchesMatchRouteConfigurationVhostRouteAction_DIRECT_RESPONSE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecConfigPatchesPatch",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesPatch)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.EnvoyFilterSpecConfigPatchesPatchFilterClass",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesPatchFilterClass)(nil)).Elem(),
		map[string]interface{}{
			"UNSPECIFIED": EnvoyFilterSpecConfigPatchesPatchFilterClass_UNSPECIFIED,
			"AUTHN": EnvoyFilterSpecConfigPatchesPatchFilterClass_AUTHN,
			"AUTHZ": EnvoyFilterSpecConfigPatchesPatchFilterClass_AUTHZ,
			"STATS": EnvoyFilterSpecConfigPatchesPatchFilterClass_STATS,
		},
	)
	_jsii_.RegisterEnum(
		"networkingistioio.EnvoyFilterSpecConfigPatchesPatchOperation",
		reflect.TypeOf((*EnvoyFilterSpecConfigPatchesPatchOperation)(nil)).Elem(),
		map[string]interface{}{
			"INVALID": EnvoyFilterSpecConfigPatchesPatchOperation_INVALID,
			"MERGE": EnvoyFilterSpecConfigPatchesPatchOperation_MERGE,
			"ADD": EnvoyFilterSpecConfigPatchesPatchOperation_ADD,
			"REMOVE": EnvoyFilterSpecConfigPatchesPatchOperation_REMOVE,
			"INSERT_BEFORE": EnvoyFilterSpecConfigPatchesPatchOperation_INSERT_BEFORE,
			"INSERT_AFTER": EnvoyFilterSpecConfigPatchesPatchOperation_INSERT_AFTER,
			"INSERT_FIRST": EnvoyFilterSpecConfigPatchesPatchOperation_INSERT_FIRST,
			"REPLACE": EnvoyFilterSpecConfigPatchesPatchOperation_REPLACE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.EnvoyFilterSpecWorkloadSelector",
		reflect.TypeOf((*EnvoyFilterSpecWorkloadSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.Gateway",
		reflect.TypeOf((*Gateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Gateway{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewayProps",
		reflect.TypeOf((*GatewayProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewaySpec",
		reflect.TypeOf((*GatewaySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewaySpecServers",
		reflect.TypeOf((*GatewaySpecServers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewaySpecServersPort",
		reflect.TypeOf((*GatewaySpecServersPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewaySpecServersTls",
		reflect.TypeOf((*GatewaySpecServersTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.GatewaySpecServersTlsMaxProtocolVersion",
		reflect.TypeOf((*GatewaySpecServersTlsMaxProtocolVersion)(nil)).Elem(),
		map[string]interface{}{
			"TLS_AUTO": GatewaySpecServersTlsMaxProtocolVersion_TLS_AUTO,
			"TLSV1_0": GatewaySpecServersTlsMaxProtocolVersion_TLSV1_0,
			"TLSV1_1": GatewaySpecServersTlsMaxProtocolVersion_TLSV1_1,
			"TLSV1_2": GatewaySpecServersTlsMaxProtocolVersion_TLSV1_2,
			"TLSV1_3": GatewaySpecServersTlsMaxProtocolVersion_TLSV1_3,
		},
	)
	_jsii_.RegisterEnum(
		"networkingistioio.GatewaySpecServersTlsMinProtocolVersion",
		reflect.TypeOf((*GatewaySpecServersTlsMinProtocolVersion)(nil)).Elem(),
		map[string]interface{}{
			"TLS_AUTO": GatewaySpecServersTlsMinProtocolVersion_TLS_AUTO,
			"TLSV1_0": GatewaySpecServersTlsMinProtocolVersion_TLSV1_0,
			"TLSV1_1": GatewaySpecServersTlsMinProtocolVersion_TLSV1_1,
			"TLSV1_2": GatewaySpecServersTlsMinProtocolVersion_TLSV1_2,
			"TLSV1_3": GatewaySpecServersTlsMinProtocolVersion_TLSV1_3,
		},
	)
	_jsii_.RegisterEnum(
		"networkingistioio.GatewaySpecServersTlsMode",
		reflect.TypeOf((*GatewaySpecServersTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"PASSTHROUGH": GatewaySpecServersTlsMode_PASSTHROUGH,
			"SIMPLE": GatewaySpecServersTlsMode_SIMPLE,
			"MUTUAL": GatewaySpecServersTlsMode_MUTUAL,
			"AUTO_PASSTHROUGH": GatewaySpecServersTlsMode_AUTO_PASSTHROUGH,
			"ISTIO_MUTUAL": GatewaySpecServersTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterClass(
		"networkingistioio.GatewayV1Beta1",
		reflect.TypeOf((*GatewayV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GatewayV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewayV1Beta1Props",
		reflect.TypeOf((*GatewayV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewayV1Beta1Spec",
		reflect.TypeOf((*GatewayV1Beta1Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewayV1Beta1SpecServers",
		reflect.TypeOf((*GatewayV1Beta1SpecServers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewayV1Beta1SpecServersPort",
		reflect.TypeOf((*GatewayV1Beta1SpecServersPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.GatewayV1Beta1SpecServersTls",
		reflect.TypeOf((*GatewayV1Beta1SpecServersTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.GatewayV1Beta1SpecServersTlsMaxProtocolVersion",
		reflect.TypeOf((*GatewayV1Beta1SpecServersTlsMaxProtocolVersion)(nil)).Elem(),
		map[string]interface{}{
			"TLS_AUTO": GatewayV1Beta1SpecServersTlsMaxProtocolVersion_TLS_AUTO,
			"TLSV1_0": GatewayV1Beta1SpecServersTlsMaxProtocolVersion_TLSV1_0,
			"TLSV1_1": GatewayV1Beta1SpecServersTlsMaxProtocolVersion_TLSV1_1,
			"TLSV1_2": GatewayV1Beta1SpecServersTlsMaxProtocolVersion_TLSV1_2,
			"TLSV1_3": GatewayV1Beta1SpecServersTlsMaxProtocolVersion_TLSV1_3,
		},
	)
	_jsii_.RegisterEnum(
		"networkingistioio.GatewayV1Beta1SpecServersTlsMinProtocolVersion",
		reflect.TypeOf((*GatewayV1Beta1SpecServersTlsMinProtocolVersion)(nil)).Elem(),
		map[string]interface{}{
			"TLS_AUTO": GatewayV1Beta1SpecServersTlsMinProtocolVersion_TLS_AUTO,
			"TLSV1_0": GatewayV1Beta1SpecServersTlsMinProtocolVersion_TLSV1_0,
			"TLSV1_1": GatewayV1Beta1SpecServersTlsMinProtocolVersion_TLSV1_1,
			"TLSV1_2": GatewayV1Beta1SpecServersTlsMinProtocolVersion_TLSV1_2,
			"TLSV1_3": GatewayV1Beta1SpecServersTlsMinProtocolVersion_TLSV1_3,
		},
	)
	_jsii_.RegisterEnum(
		"networkingistioio.GatewayV1Beta1SpecServersTlsMode",
		reflect.TypeOf((*GatewayV1Beta1SpecServersTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"PASSTHROUGH": GatewayV1Beta1SpecServersTlsMode_PASSTHROUGH,
			"SIMPLE": GatewayV1Beta1SpecServersTlsMode_SIMPLE,
			"MUTUAL": GatewayV1Beta1SpecServersTlsMode_MUTUAL,
			"AUTO_PASSTHROUGH": GatewayV1Beta1SpecServersTlsMode_AUTO_PASSTHROUGH,
			"ISTIO_MUTUAL": GatewayV1Beta1SpecServersTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterClass(
		"networkingistioio.ProxyConfig",
		reflect.TypeOf((*ProxyConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ProxyConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ProxyConfigProps",
		reflect.TypeOf((*ProxyConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ProxyConfigSpec",
		reflect.TypeOf((*ProxyConfigSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ProxyConfigSpecImage",
		reflect.TypeOf((*ProxyConfigSpecImage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ProxyConfigSpecSelector",
		reflect.TypeOf((*ProxyConfigSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.ServiceEntry",
		reflect.TypeOf((*ServiceEntry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceEntry{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntryProps",
		reflect.TypeOf((*ServiceEntryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntrySpec",
		reflect.TypeOf((*ServiceEntrySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntrySpecEndpoints",
		reflect.TypeOf((*ServiceEntrySpecEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.ServiceEntrySpecLocation",
		reflect.TypeOf((*ServiceEntrySpecLocation)(nil)).Elem(),
		map[string]interface{}{
			"MESH_EXTERNAL": ServiceEntrySpecLocation_MESH_EXTERNAL,
			"MESH_INTERNAL": ServiceEntrySpecLocation_MESH_INTERNAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntrySpecPorts",
		reflect.TypeOf((*ServiceEntrySpecPorts)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.ServiceEntrySpecResolution",
		reflect.TypeOf((*ServiceEntrySpecResolution)(nil)).Elem(),
		map[string]interface{}{
			"NONE": ServiceEntrySpecResolution_NONE,
			"STATIC": ServiceEntrySpecResolution_STATIC,
			"DNS": ServiceEntrySpecResolution_DNS,
			"DNS_ROUND_ROBIN": ServiceEntrySpecResolution_DNS_ROUND_ROBIN,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntrySpecWorkloadSelector",
		reflect.TypeOf((*ServiceEntrySpecWorkloadSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.ServiceEntryV1Beta1",
		reflect.TypeOf((*ServiceEntryV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceEntryV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntryV1Beta1Props",
		reflect.TypeOf((*ServiceEntryV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntryV1Beta1Spec",
		reflect.TypeOf((*ServiceEntryV1Beta1Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntryV1Beta1SpecEndpoints",
		reflect.TypeOf((*ServiceEntryV1Beta1SpecEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.ServiceEntryV1Beta1SpecLocation",
		reflect.TypeOf((*ServiceEntryV1Beta1SpecLocation)(nil)).Elem(),
		map[string]interface{}{
			"MESH_EXTERNAL": ServiceEntryV1Beta1SpecLocation_MESH_EXTERNAL,
			"MESH_INTERNAL": ServiceEntryV1Beta1SpecLocation_MESH_INTERNAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntryV1Beta1SpecPorts",
		reflect.TypeOf((*ServiceEntryV1Beta1SpecPorts)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.ServiceEntryV1Beta1SpecResolution",
		reflect.TypeOf((*ServiceEntryV1Beta1SpecResolution)(nil)).Elem(),
		map[string]interface{}{
			"NONE": ServiceEntryV1Beta1SpecResolution_NONE,
			"STATIC": ServiceEntryV1Beta1SpecResolution_STATIC,
			"DNS": ServiceEntryV1Beta1SpecResolution_DNS,
			"DNS_ROUND_ROBIN": ServiceEntryV1Beta1SpecResolution_DNS_ROUND_ROBIN,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.ServiceEntryV1Beta1SpecWorkloadSelector",
		reflect.TypeOf((*ServiceEntryV1Beta1SpecWorkloadSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.Sidecar",
		reflect.TypeOf((*Sidecar)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Sidecar{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarProps",
		reflect.TypeOf((*SidecarProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpec",
		reflect.TypeOf((*SidecarSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecEgress",
		reflect.TypeOf((*SidecarSpecEgress)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarSpecEgressCaptureMode",
		reflect.TypeOf((*SidecarSpecEgressCaptureMode)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": SidecarSpecEgressCaptureMode_DEFAULT,
			"IPTABLES": SidecarSpecEgressCaptureMode_IPTABLES,
			"NONE": SidecarSpecEgressCaptureMode_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecEgressPort",
		reflect.TypeOf((*SidecarSpecEgressPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecIngress",
		reflect.TypeOf((*SidecarSpecIngress)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarSpecIngressCaptureMode",
		reflect.TypeOf((*SidecarSpecIngressCaptureMode)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": SidecarSpecIngressCaptureMode_DEFAULT,
			"IPTABLES": SidecarSpecIngressCaptureMode_IPTABLES,
			"NONE": SidecarSpecIngressCaptureMode_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecIngressPort",
		reflect.TypeOf((*SidecarSpecIngressPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecIngressTls",
		reflect.TypeOf((*SidecarSpecIngressTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarSpecIngressTlsMaxProtocolVersion",
		reflect.TypeOf((*SidecarSpecIngressTlsMaxProtocolVersion)(nil)).Elem(),
		map[string]interface{}{
			"TLS_AUTO": SidecarSpecIngressTlsMaxProtocolVersion_TLS_AUTO,
			"TLSV1_0": SidecarSpecIngressTlsMaxProtocolVersion_TLSV1_0,
			"TLSV1_1": SidecarSpecIngressTlsMaxProtocolVersion_TLSV1_1,
			"TLSV1_2": SidecarSpecIngressTlsMaxProtocolVersion_TLSV1_2,
			"TLSV1_3": SidecarSpecIngressTlsMaxProtocolVersion_TLSV1_3,
		},
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarSpecIngressTlsMinProtocolVersion",
		reflect.TypeOf((*SidecarSpecIngressTlsMinProtocolVersion)(nil)).Elem(),
		map[string]interface{}{
			"TLS_AUTO": SidecarSpecIngressTlsMinProtocolVersion_TLS_AUTO,
			"TLSV1_0": SidecarSpecIngressTlsMinProtocolVersion_TLSV1_0,
			"TLSV1_1": SidecarSpecIngressTlsMinProtocolVersion_TLSV1_1,
			"TLSV1_2": SidecarSpecIngressTlsMinProtocolVersion_TLSV1_2,
			"TLSV1_3": SidecarSpecIngressTlsMinProtocolVersion_TLSV1_3,
		},
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarSpecIngressTlsMode",
		reflect.TypeOf((*SidecarSpecIngressTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"PASSTHROUGH": SidecarSpecIngressTlsMode_PASSTHROUGH,
			"SIMPLE": SidecarSpecIngressTlsMode_SIMPLE,
			"MUTUAL": SidecarSpecIngressTlsMode_MUTUAL,
			"AUTO_PASSTHROUGH": SidecarSpecIngressTlsMode_AUTO_PASSTHROUGH,
			"ISTIO_MUTUAL": SidecarSpecIngressTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecOutboundTrafficPolicy",
		reflect.TypeOf((*SidecarSpecOutboundTrafficPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecOutboundTrafficPolicyEgressProxy",
		reflect.TypeOf((*SidecarSpecOutboundTrafficPolicyEgressProxy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecOutboundTrafficPolicyEgressProxyPort",
		reflect.TypeOf((*SidecarSpecOutboundTrafficPolicyEgressProxyPort)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarSpecOutboundTrafficPolicyMode",
		reflect.TypeOf((*SidecarSpecOutboundTrafficPolicyMode)(nil)).Elem(),
		map[string]interface{}{
			"REGISTRY_ONLY": SidecarSpecOutboundTrafficPolicyMode_REGISTRY_ONLY,
			"ALLOW_ANY": SidecarSpecOutboundTrafficPolicyMode_ALLOW_ANY,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarSpecWorkloadSelector",
		reflect.TypeOf((*SidecarSpecWorkloadSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.SidecarV1Beta1",
		reflect.TypeOf((*SidecarV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SidecarV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarV1Beta1Props",
		reflect.TypeOf((*SidecarV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarV1Beta1Spec",
		reflect.TypeOf((*SidecarV1Beta1Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarV1Beta1SpecEgress",
		reflect.TypeOf((*SidecarV1Beta1SpecEgress)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarV1Beta1SpecEgressCaptureMode",
		reflect.TypeOf((*SidecarV1Beta1SpecEgressCaptureMode)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": SidecarV1Beta1SpecEgressCaptureMode_DEFAULT,
			"IPTABLES": SidecarV1Beta1SpecEgressCaptureMode_IPTABLES,
			"NONE": SidecarV1Beta1SpecEgressCaptureMode_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarV1Beta1SpecEgressPort",
		reflect.TypeOf((*SidecarV1Beta1SpecEgressPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarV1Beta1SpecIngress",
		reflect.TypeOf((*SidecarV1Beta1SpecIngress)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarV1Beta1SpecIngressCaptureMode",
		reflect.TypeOf((*SidecarV1Beta1SpecIngressCaptureMode)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": SidecarV1Beta1SpecIngressCaptureMode_DEFAULT,
			"IPTABLES": SidecarV1Beta1SpecIngressCaptureMode_IPTABLES,
			"NONE": SidecarV1Beta1SpecIngressCaptureMode_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarV1Beta1SpecIngressPort",
		reflect.TypeOf((*SidecarV1Beta1SpecIngressPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarV1Beta1SpecIngressTls",
		reflect.TypeOf((*SidecarV1Beta1SpecIngressTls)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarV1Beta1SpecIngressTlsMaxProtocolVersion",
		reflect.TypeOf((*SidecarV1Beta1SpecIngressTlsMaxProtocolVersion)(nil)).Elem(),
		map[string]interface{}{
			"TLS_AUTO": SidecarV1Beta1SpecIngressTlsMaxProtocolVersion_TLS_AUTO,
			"TLSV1_0": SidecarV1Beta1SpecIngressTlsMaxProtocolVersion_TLSV1_0,
			"TLSV1_1": SidecarV1Beta1SpecIngressTlsMaxProtocolVersion_TLSV1_1,
			"TLSV1_2": SidecarV1Beta1SpecIngressTlsMaxProtocolVersion_TLSV1_2,
			"TLSV1_3": SidecarV1Beta1SpecIngressTlsMaxProtocolVersion_TLSV1_3,
		},
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarV1Beta1SpecIngressTlsMinProtocolVersion",
		reflect.TypeOf((*SidecarV1Beta1SpecIngressTlsMinProtocolVersion)(nil)).Elem(),
		map[string]interface{}{
			"TLS_AUTO": SidecarV1Beta1SpecIngressTlsMinProtocolVersion_TLS_AUTO,
			"TLSV1_0": SidecarV1Beta1SpecIngressTlsMinProtocolVersion_TLSV1_0,
			"TLSV1_1": SidecarV1Beta1SpecIngressTlsMinProtocolVersion_TLSV1_1,
			"TLSV1_2": SidecarV1Beta1SpecIngressTlsMinProtocolVersion_TLSV1_2,
			"TLSV1_3": SidecarV1Beta1SpecIngressTlsMinProtocolVersion_TLSV1_3,
		},
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarV1Beta1SpecIngressTlsMode",
		reflect.TypeOf((*SidecarV1Beta1SpecIngressTlsMode)(nil)).Elem(),
		map[string]interface{}{
			"PASSTHROUGH": SidecarV1Beta1SpecIngressTlsMode_PASSTHROUGH,
			"SIMPLE": SidecarV1Beta1SpecIngressTlsMode_SIMPLE,
			"MUTUAL": SidecarV1Beta1SpecIngressTlsMode_MUTUAL,
			"AUTO_PASSTHROUGH": SidecarV1Beta1SpecIngressTlsMode_AUTO_PASSTHROUGH,
			"ISTIO_MUTUAL": SidecarV1Beta1SpecIngressTlsMode_ISTIO_MUTUAL,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarV1Beta1SpecOutboundTrafficPolicy",
		reflect.TypeOf((*SidecarV1Beta1SpecOutboundTrafficPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarV1Beta1SpecOutboundTrafficPolicyEgressProxy",
		reflect.TypeOf((*SidecarV1Beta1SpecOutboundTrafficPolicyEgressProxy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarV1Beta1SpecOutboundTrafficPolicyEgressProxyPort",
		reflect.TypeOf((*SidecarV1Beta1SpecOutboundTrafficPolicyEgressProxyPort)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.SidecarV1Beta1SpecOutboundTrafficPolicyMode",
		reflect.TypeOf((*SidecarV1Beta1SpecOutboundTrafficPolicyMode)(nil)).Elem(),
		map[string]interface{}{
			"REGISTRY_ONLY": SidecarV1Beta1SpecOutboundTrafficPolicyMode_REGISTRY_ONLY,
			"ALLOW_ANY": SidecarV1Beta1SpecOutboundTrafficPolicyMode_ALLOW_ANY,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.SidecarV1Beta1SpecWorkloadSelector",
		reflect.TypeOf((*SidecarV1Beta1SpecWorkloadSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.VirtualService",
		reflect.TypeOf((*VirtualService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VirtualService{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceProps",
		reflect.TypeOf((*VirtualServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpec",
		reflect.TypeOf((*VirtualServiceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttp",
		reflect.TypeOf((*VirtualServiceSpecHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpCorsPolicy",
		reflect.TypeOf((*VirtualServiceSpecHttpCorsPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpCorsPolicyAllowOrigins",
		reflect.TypeOf((*VirtualServiceSpecHttpCorsPolicyAllowOrigins)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpDelegate",
		reflect.TypeOf((*VirtualServiceSpecHttpDelegate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpDirectResponse",
		reflect.TypeOf((*VirtualServiceSpecHttpDirectResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpDirectResponseBody",
		reflect.TypeOf((*VirtualServiceSpecHttpDirectResponseBody)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpFault",
		reflect.TypeOf((*VirtualServiceSpecHttpFault)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpFaultAbort",
		reflect.TypeOf((*VirtualServiceSpecHttpFaultAbort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpFaultAbortPercentage",
		reflect.TypeOf((*VirtualServiceSpecHttpFaultAbortPercentage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpFaultDelay",
		reflect.TypeOf((*VirtualServiceSpecHttpFaultDelay)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpFaultDelayPercentage",
		reflect.TypeOf((*VirtualServiceSpecHttpFaultDelayPercentage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpHeaders",
		reflect.TypeOf((*VirtualServiceSpecHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpHeadersRequest",
		reflect.TypeOf((*VirtualServiceSpecHttpHeadersRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpHeadersResponse",
		reflect.TypeOf((*VirtualServiceSpecHttpHeadersResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatch",
		reflect.TypeOf((*VirtualServiceSpecHttpMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchAuthority",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchAuthority)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchHeaders",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchMethod",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchMethod)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchQueryParams",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchQueryParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchScheme",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchScheme)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchUri",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchUri)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMatchWithoutHeaders",
		reflect.TypeOf((*VirtualServiceSpecHttpMatchWithoutHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMirror",
		reflect.TypeOf((*VirtualServiceSpecHttpMirror)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMirrorPercentage",
		reflect.TypeOf((*VirtualServiceSpecHttpMirrorPercentage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpMirrorPort",
		reflect.TypeOf((*VirtualServiceSpecHttpMirrorPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRedirect",
		reflect.TypeOf((*VirtualServiceSpecHttpRedirect)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"networkingistioio.VirtualServiceSpecHttpRedirectDerivePort",
		reflect.TypeOf((*VirtualServiceSpecHttpRedirectDerivePort)(nil)).Elem(),
		map[string]interface{}{
			"FROM_PROTOCOL_DEFAULT": VirtualServiceSpecHttpRedirectDerivePort_FROM_PROTOCOL_DEFAULT,
			"FROM_REQUEST_PORT": VirtualServiceSpecHttpRedirectDerivePort_FROM_REQUEST_PORT,
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRetries",
		reflect.TypeOf((*VirtualServiceSpecHttpRetries)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRewrite",
		reflect.TypeOf((*VirtualServiceSpecHttpRewrite)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRoute",
		reflect.TypeOf((*VirtualServiceSpecHttpRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRouteDestination",
		reflect.TypeOf((*VirtualServiceSpecHttpRouteDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRouteDestinationPort",
		reflect.TypeOf((*VirtualServiceSpecHttpRouteDestinationPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRouteHeaders",
		reflect.TypeOf((*VirtualServiceSpecHttpRouteHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRouteHeadersRequest",
		reflect.TypeOf((*VirtualServiceSpecHttpRouteHeadersRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecHttpRouteHeadersResponse",
		reflect.TypeOf((*VirtualServiceSpecHttpRouteHeadersResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTcp",
		reflect.TypeOf((*VirtualServiceSpecTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTcpMatch",
		reflect.TypeOf((*VirtualServiceSpecTcpMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTcpRoute",
		reflect.TypeOf((*VirtualServiceSpecTcpRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTcpRouteDestination",
		reflect.TypeOf((*VirtualServiceSpecTcpRouteDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTcpRouteDestinationPort",
		reflect.TypeOf((*VirtualServiceSpecTcpRouteDestinationPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTls",
		reflect.TypeOf((*VirtualServiceSpecTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTlsMatch",
		reflect.TypeOf((*VirtualServiceSpecTlsMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTlsRoute",
		reflect.TypeOf((*VirtualServiceSpecTlsRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTlsRouteDestination",
		reflect.TypeOf((*VirtualServiceSpecTlsRouteDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.VirtualServiceSpecTlsRouteDestinationPort",
		reflect.TypeOf((*VirtualServiceSpecTlsRouteDestinationPort)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.WorkloadEntry",
		reflect.TypeOf((*WorkloadEntry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_WorkloadEntry{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadEntryProps",
		reflect.TypeOf((*WorkloadEntryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadEntrySpec",
		reflect.TypeOf((*WorkloadEntrySpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.WorkloadEntryV1Beta1",
		reflect.TypeOf((*WorkloadEntryV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_WorkloadEntryV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadEntryV1Beta1Props",
		reflect.TypeOf((*WorkloadEntryV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadEntryV1Beta1Spec",
		reflect.TypeOf((*WorkloadEntryV1Beta1Spec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.WorkloadGroup",
		reflect.TypeOf((*WorkloadGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_WorkloadGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupProps",
		reflect.TypeOf((*WorkloadGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpec",
		reflect.TypeOf((*WorkloadGroupSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecMetadata",
		reflect.TypeOf((*WorkloadGroupSpecMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecProbe",
		reflect.TypeOf((*WorkloadGroupSpecProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecProbeExec",
		reflect.TypeOf((*WorkloadGroupSpecProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecProbeHttpGet",
		reflect.TypeOf((*WorkloadGroupSpecProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkloadGroupSpecProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecProbeTcpSocket",
		reflect.TypeOf((*WorkloadGroupSpecProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupSpecTemplate",
		reflect.TypeOf((*WorkloadGroupSpecTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"networkingistioio.WorkloadGroupV1Beta1",
		reflect.TypeOf((*WorkloadGroupV1Beta1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_WorkloadGroupV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupV1Beta1Props",
		reflect.TypeOf((*WorkloadGroupV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupV1Beta1Spec",
		reflect.TypeOf((*WorkloadGroupV1Beta1Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupV1Beta1SpecMetadata",
		reflect.TypeOf((*WorkloadGroupV1Beta1SpecMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupV1Beta1SpecProbe",
		reflect.TypeOf((*WorkloadGroupV1Beta1SpecProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupV1Beta1SpecProbeExec",
		reflect.TypeOf((*WorkloadGroupV1Beta1SpecProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupV1Beta1SpecProbeHttpGet",
		reflect.TypeOf((*WorkloadGroupV1Beta1SpecProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupV1Beta1SpecProbeHttpGetHttpHeaders",
		reflect.TypeOf((*WorkloadGroupV1Beta1SpecProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupV1Beta1SpecProbeTcpSocket",
		reflect.TypeOf((*WorkloadGroupV1Beta1SpecProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingistioio.WorkloadGroupV1Beta1SpecTemplate",
		reflect.TypeOf((*WorkloadGroupV1Beta1SpecTemplate)(nil)).Elem(),
	)
}

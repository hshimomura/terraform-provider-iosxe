package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestAPPortfastDeleteOwnsOnlyTrunkAttribute(t *testing.T) {
	model := InterfaceEthernet{
		Type: types.StringValue("GigabitEthernet"), Name: types.StringValue("1/0/23"),
		SpanningTreePortfastTrunk: types.BoolValue(true),
	}
	created := model.toBodyXML(context.Background(), model)
	if !strings.Contains(created, "<portfast-v2><trunk>") || strings.Contains(created, "<portfast>") {
		t.Fatalf("AP trunk must use the effective current YANG path: %s", created)
	}
	body := model.addDeletePathsXML(context.Background(), "")
	if strings.Count(body, `operation="remove"`) != 1 ||
		!strings.Contains(body, `trunk operation="remove"`) ||
		!strings.Contains(body, "<portfast-v2>") ||
		strings.Contains(body, `GigabitEthernet operation="remove"`) ||
		strings.Contains(body, "shutdown") || strings.Contains(body, "bpduguard") {
		t.Fatalf("AP destroy must preserve the interface and all common settings: %s", body)
	}
}

func TestAPStaticTrunkDisablesOnlyAutoconfInheritance(t *testing.T) {
	model := InterfaceEthernet{
		Type: types.StringValue("GigabitEthernet"), Name: types.StringValue("1/0/23"),
		AccessSessionInheritDisableAutoconf: types.BoolValue(true),
	}
	body := model.toBodyXML(context.Background(), model)
	if !strings.Contains(body, "<access-session><inherit><disable><autoconf>") {
		t.Fatalf("static AP trunk must exclude AutoConf inheritance: %s", body)
	}
	removed := model.addDeletePathsXML(context.Background(), "")
	if strings.Count(removed, `operation="remove"`) != 1 || !strings.Contains(removed, `autoconf operation="remove"`) || strings.Contains(removed, "service-policy") {
		t.Fatalf("destroy must only remove this port's AutoConf override: %s", removed)
	}
}

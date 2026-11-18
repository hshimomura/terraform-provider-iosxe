package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func TestPrefixListDeleteModeSchema(t *testing.T) {
	response := resource.SchemaResponse{}
	NewPrefixListResource().Schema(context.Background(), resource.SchemaRequest{}, &response)
	attribute, ok := response.Schema.Attributes["delete_mode"]
	if !ok || !attribute.IsOptional() {
		t.Fatal("shared prefix-list container requires an optional delete_mode")
	}
}

func TestPrefixListDeletePathsAreKeyed(t *testing.T) {
	model := PrefixList{Prefixes: []PrefixListPrefixes{
		{Name: types.StringValue("EVPN-Campus-TO-FUSION"), Seq: types.Int64Value(10)},
		{Name: types.StringValue("EVPN-IoT-TO-FUSION"), Seq: types.Int64Value(10)},
	}}
	body := model.addDeletePathsXML(context.Background(), "")
	if strings.Count(body, `operation="remove"`) != 2 ||
		!strings.Contains(body, "EVPN-Campus-TO-FUSION") ||
		!strings.Contains(body, "EVPN-IoT-TO-FUSION") ||
		strings.Contains(body, `<prefix-lists nc:operation="remove"`) {
		t.Fatalf("delete must target only keyed managed entries: %s", body)
	}
}

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestProviderConfiguresActions(t *testing.T) {
	ctx := context.Background()
	p := New("test")()
	schemaResponse := provider.SchemaResponse{}
	p.Schema(ctx, provider.SchemaRequest{}, &schemaResponse)
	objectType := schemaResponse.Schema.Type().TerraformType(ctx).(tftypes.Object)
	values := make(map[string]tftypes.Value, len(objectType.AttributeTypes))
	for name, typ := range objectType.AttributeTypes {
		values[name] = tftypes.NewValue(typ, nil)
	}
	for name, value := range map[string]string{"host": "127.0.0.1", "username": "test", "password": "test"} {
		values[name] = tftypes.NewValue(tftypes.String, value)
	}
	response := provider.ConfigureResponse{}
	p.Configure(ctx, provider.ConfigureRequest{Config: tfsdk.Config{Raw: tftypes.NewValue(objectType, values), Schema: schemaResponse.Schema}}, &response)
	if response.Diagnostics.HasError() {
		t.Fatalf("configure failed: %v", response.Diagnostics)
	}
	if response.ActionData == nil || response.ActionData != response.ResourceData {
		t.Fatal("actions must receive the same configured device clients as resources")
	}
}

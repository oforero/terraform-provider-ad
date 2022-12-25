package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	kcapi "kopicloud/api"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		// First some scafolding
		DataSourcesMap: dataSources(),
		// ResourcesMap: map[string]*schema.Resource{
		// 	"kopicloud_computer": resourceComputer(),
		// },
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Kopiclcoud Server URL",
				DefaultFunc: schema.EnvDefaultFunc("API_HOST", ""),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Bearer (JWT) or Basic Authentication Token",
				DefaultFunc: schema.EnvDefaultFunc("API_AUTHENTICATION_TOKEN", ""),
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func schemaOfStringList(list_name string, field_name string, description string) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		list_name: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					field_name: {
						Type:        schema.TypeString,
						Computed:    true,
						Description: description,
					},
				},
			},
		},
	}
}

func stringToTerraform(field_name string, obj string) map[string]interface{} {
	result := make(map[string]interface{})
	result["name"] = obj
	return result
}

func flattenStringList(list *[]string) []interface{} {
	if list != nil {
		results := make([]interface{}, len(*list))
		for i, value := range *list {
			results[i] = stringToTerraform("name", value)
		}
		return results
	}
	return make([]interface{}, 0)
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return NewApiClient(d)

}

type ApiClient struct {
	data   *schema.ResourceData
	client *kcapi.ClientWithResponses
}

func NewApiClient(d *schema.ResourceData) (*ApiClient, diag.Diagnostics) {
	c := &ApiClient{data: d}
	client, err := c.NewKopiCloudClient()
	if err != nil {
		return c, diag.FromErr(err)
	}
	c.client = client
	return c, nil

}

func (a *ApiClient) NewKopiCloudClient() (*kcapi.ClientWithResponses, error) {
	host := a.data.Get("host").(string)
	c, err := kcapi.NewClientWithResponses(host)
	if err != nil {
		return c, err
	}
	return c, nil
}
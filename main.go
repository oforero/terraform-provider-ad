package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	kopicloud "gitlab.com/KopiCloud/kopicloud-ad-tf-provider/kopicloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return kopicloud.Provider()
		},
	})
}

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceComputeFirewallPolicyAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeFirewallPolicyAssociationCreate,
		Read:   resourceComputeFirewallPolicyAssociationRead,
		Delete: resourceComputeFirewallPolicyAssociationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeFirewallPolicyAssociationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"attachment_target": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The target that the firewall policy is attached to.`,
			},
			"firewall_policy": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `The firewall policy of the resource.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name for an association.`,
			},
			"short_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The short name of the firewall policy of the association.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeFirewallPolicyAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeFirewallPolicyAssociationName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	attachmentTargetProp, err := expandComputeFirewallPolicyAssociationAttachmentTarget(d.Get("attachment_target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("attachment_target"); !tpgresource.IsEmptyValue(reflect.ValueOf(attachmentTargetProp)) && (ok || !reflect.DeepEqual(v, attachmentTargetProp)) {
		obj["attachmentTarget"] = attachmentTargetProp
	}
	firewallPolicyProp, err := expandComputeFirewallPolicyAssociationFirewallPolicy(d.Get("firewall_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("firewall_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(firewallPolicyProp)) && (ok || !reflect.DeepEqual(v, firewallPolicyProp)) {
		obj["firewallPolicy"] = firewallPolicyProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}locations/global/firewallPolicies/{{firewall_policy}}/addAssociation")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FirewallPolicyAssociation: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating FirewallPolicyAssociation: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	parent := d.Get("firewall_policy").(string)
	var opRes map[string]interface{}
	err = ComputeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, parent, "Creating FirewallPolicyAssociation", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create FirewallPolicyAssociation: %s", err)
	}

	log.Printf("[DEBUG] Finished creating FirewallPolicyAssociation %q: %#v", d.Id(), res)

	return resourceComputeFirewallPolicyAssociationRead(d, meta)
}

func resourceComputeFirewallPolicyAssociationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}locations/global/firewallPolicies/{{firewall_policy}}/getAssociation?name={{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	expandComputeFirewallPolicyAssociationFirewallPolicy(d.Get("firewall_policy"), d, config)
	url, err = tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}locations/global/firewallPolicies/{{firewall_policy}}/getAssociation?name={{name}}")
	if err != nil {
		return err
	}
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeFirewallPolicyAssociation %q", d.Id()))
	}

	if err := d.Set("name", flattenComputeFirewallPolicyAssociationName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyAssociation: %s", err)
	}
	if err := d.Set("attachment_target", flattenComputeFirewallPolicyAssociationAttachmentTarget(res["attachmentTarget"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyAssociation: %s", err)
	}
	if err := d.Set("short_name", flattenComputeFirewallPolicyAssociationShortName(res["shortName"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallPolicyAssociation: %s", err)
	}

	return nil
}

func resourceComputeFirewallPolicyAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}locations/global/firewallPolicies/{{firewall_policy}}/removeAssociation?name={{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting FirewallPolicyAssociation %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "FirewallPolicyAssociation")
	}

	parent := d.Get("firewall_policy").(string)
	var opRes map[string]interface{}
	err = ComputeOrgOperationWaitTimeWithResponse(
		config, res, &opRes, parent, "Deleting FirewallPolicyAssociation", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to delete FirewallPolicyAssociation: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting FirewallPolicyAssociation %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeFirewallPolicyAssociationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^locations/global/firewallPolicies/(?P<firewall_policy>[^/]+)/associations/(?P<name>[^/]+)$",
		"^(?P<firewall_policy>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeFirewallPolicyAssociationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyAssociationAttachmentTarget(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeFirewallPolicyAssociationShortName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandComputeFirewallPolicyAssociationName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyAssociationAttachmentTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyAssociationFirewallPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	firewallPolicyId := tpgresource.GetResourceNameFromSelfLink(v.(string))
	if err := d.Set("firewall_policy", firewallPolicyId); err != nil {
		return nil, fmt.Errorf("Error setting firewall_policy: %s", err)
	}
	return firewallPolicyId, nil
}

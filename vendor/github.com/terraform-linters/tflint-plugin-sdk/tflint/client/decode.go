package client

import (
	"fmt"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform"
)

// Attribute is an intermediate representation of hcl.Attribute.
type Attribute struct {
	Name      string
	Expr      []byte
	ExprRange hcl.Range
	Range     hcl.Range
	NameRange hcl.Range
}

func decodeAttribute(attribute *Attribute) (*hcl.Attribute, hcl.Diagnostics) {
	expr, diags := parseExpression(attribute.Expr, attribute.ExprRange.Filename, attribute.ExprRange.Start)
	if diags.HasErrors() {
		return nil, diags
	}

	return &hcl.Attribute{
		Name:      attribute.Name,
		Expr:      expr,
		Range:     attribute.Range,
		NameRange: attribute.NameRange,
	}, nil
}

// Block is an intermediate representation of hcl.Block.
type Block struct {
	Type      string
	Labels    []string
	Body      []byte
	BodyRange hcl.Range

	DefRange    hcl.Range
	TypeRange   hcl.Range
	LabelRanges []hcl.Range
}

func decodeBlock(block *Block) (*hcl.Block, hcl.Diagnostics) {
	file, diags := parseConfig(block.Body, block.BodyRange.Filename, block.BodyRange.Start)
	if diags.HasErrors() {
		return nil, diags
	}

	return &hcl.Block{
		Type:        block.Type,
		Labels:      block.Labels,
		Body:        file.Body,
		DefRange:    block.DefRange,
		TypeRange:   block.TypeRange,
		LabelRanges: block.LabelRanges,
	}, nil
}

// Resource is an intermediate representation of terraform.Resource.
type Resource struct {
	Mode         terraform.ResourceMode
	Name         string
	Type         string
	Config       []byte
	ConfigRange  hcl.Range
	Count        []byte
	CountRange   hcl.Range
	ForEach      []byte
	ForEachRange hcl.Range

	ProviderConfigRef *terraform.ProviderConfigRef

	Managed *ManagedResource

	DeclRange hcl.Range
	TypeRange hcl.Range
}

func decodeResource(resource *Resource) (*terraform.Resource, hcl.Diagnostics) {
	file, diags := parseConfig(resource.Config, resource.ConfigRange.Filename, resource.ConfigRange.Start)
	if diags.HasErrors() {
		return nil, diags
	}

	var count hcl.Expression
	if resource.Count != nil {
		count, diags = parseExpression(resource.Count, resource.CountRange.Filename, resource.CountRange.Start)
		if diags.HasErrors() {
			return nil, diags
		}
	}

	var forEach hcl.Expression
	if resource.ForEach != nil {
		forEach, diags = parseExpression(resource.ForEach, resource.ForEachRange.Filename, resource.ForEachRange.Start)
		if diags.HasErrors() {
			return nil, diags
		}
	}

	managed, diags := decodeManagedResource(resource.Managed)
	if diags.HasErrors() {
		return nil, diags
	}

	return &terraform.Resource{
		Mode:    resource.Mode,
		Name:    resource.Name,
		Type:    resource.Type,
		Config:  file.Body,
		Count:   count,
		ForEach: forEach,

		ProviderConfigRef: resource.ProviderConfigRef,

		Managed: managed,

		DeclRange: resource.DeclRange,
		TypeRange: resource.TypeRange,
	}, nil
}

// ManagedResource is an intermediate representation of terraform.ManagedResource.
type ManagedResource struct {
	Connection   *Connection
	Provisioners []*Provisioner

	CreateBeforeDestroy bool
	PreventDestroy      bool
	IgnoreAllChanges    bool

	CreateBeforeDestroySet bool
	PreventDestroySet      bool
}

func decodeManagedResource(resource *ManagedResource) (*terraform.ManagedResource, hcl.Diagnostics) {
	connection, diags := decodeConnection(resource.Connection)
	if diags.HasErrors() {
		return nil, diags
	}

	provisioners := make([]*terraform.Provisioner, len(resource.Provisioners))
	for i, p := range resource.Provisioners {
		provisioner, diags := decodeProvisioner(p)
		if diags.HasErrors() {
			return nil, diags
		}
		provisioners[i] = provisioner
	}

	return &terraform.ManagedResource{
		Connection:   connection,
		Provisioners: provisioners,

		CreateBeforeDestroy: resource.CreateBeforeDestroy,
		PreventDestroy:      resource.PreventDestroy,
		IgnoreAllChanges:    resource.IgnoreAllChanges,

		CreateBeforeDestroySet: resource.CreateBeforeDestroySet,
		PreventDestroySet:      resource.PreventDestroySet,
	}, nil
}

// Connection is an intermediate representation of terraform.Connection.
type Connection struct {
	Config      []byte
	ConfigRange hcl.Range

	DeclRange hcl.Range
}

func decodeConnection(connection *Connection) (*terraform.Connection, hcl.Diagnostics) {
	if connection == nil {
		return nil, hcl.Diagnostics{}
	}

	file, diags := parseConfig(connection.Config, connection.ConfigRange.Filename, connection.ConfigRange.Start)
	if diags.HasErrors() {
		return nil, diags
	}

	return &terraform.Connection{
		Config:    file.Body,
		DeclRange: connection.DeclRange,
	}, nil
}

// Provisioner is an intermediate representation of terraform.Provisioner.
type Provisioner struct {
	Type        string
	Config      []byte
	ConfigRange hcl.Range
	Connection  *Connection
	When        terraform.ProvisionerWhen
	OnFailure   terraform.ProvisionerOnFailure

	DeclRange hcl.Range
	TypeRange hcl.Range
}

func decodeProvisioner(provisioner *Provisioner) (*terraform.Provisioner, hcl.Diagnostics) {
	file, diags := parseConfig(provisioner.Config, provisioner.ConfigRange.Filename, provisioner.ConfigRange.Start)
	if diags.HasErrors() {
		return nil, diags
	}

	connection, diags := decodeConnection(provisioner.Connection)
	if diags.HasErrors() {
		return nil, diags
	}

	return &terraform.Provisioner{
		Type:       provisioner.Type,
		Config:     file.Body,
		Connection: connection,
		When:       provisioner.When,
		OnFailure:  provisioner.OnFailure,

		DeclRange: provisioner.DeclRange,
		TypeRange: provisioner.TypeRange,
	}, nil
}

func parseExpression(src []byte, filename string, start hcl.Pos) (hcl.Expression, hcl.Diagnostics) {
	if strings.HasSuffix(filename, ".tf") {
		return hclsyntax.ParseExpression(src, filename, start)
	}

	if strings.HasSuffix(filename, ".tf.json") {
		return nil, hcl.Diagnostics{
			&hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "JSON configuration syntax is not supported",
				Subject:  &hcl.Range{Filename: filename, Start: start, End: start},
			},
		}
	}

	panic(fmt.Sprintf("Unexpected file: %s", filename))
}

func parseConfig(src []byte, filename string, start hcl.Pos) (*hcl.File, hcl.Diagnostics) {
	if strings.HasSuffix(filename, ".tf") {
		return hclsyntax.ParseConfig(src, filename, start)
	}

	if strings.HasSuffix(filename, ".tf.json") {
		return nil, hcl.Diagnostics{
			&hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "JSON configuration syntax is not supported",
				Subject:  &hcl.Range{Filename: filename, Start: start, End: start},
			},
		}
	}

	panic(fmt.Sprintf("Unexpected file: %s", filename))
}

// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.

package ucloudimport

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName       *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType     *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion     *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug           *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce           *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError         *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars        map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars   []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	PublicKey             *string           `mapstructure:"public_key" required:"true" cty:"public_key" hcl:"public_key"`
	PrivateKey            *string           `mapstructure:"private_key" required:"true" cty:"private_key" hcl:"private_key"`
	Region                *string           `mapstructure:"region" required:"true" cty:"region" hcl:"region"`
	ProjectId             *string           `mapstructure:"project_id" required:"true" cty:"project_id" hcl:"project_id"`
	BaseUrl               *string           `mapstructure:"base_url" required:"false" cty:"base_url" hcl:"base_url"`
	Profile               *string           `mapstructure:"profile" required:"false" cty:"profile" hcl:"profile"`
	SharedCredentialsFile *string           `mapstructure:"shared_credentials_file" required:"false" cty:"shared_credentials_file" hcl:"shared_credentials_file"`
	UFileBucket           *string           `mapstructure:"ufile_bucket_name" required:"true" cty:"ufile_bucket_name" hcl:"ufile_bucket_name"`
	UFileKey              *string           `mapstructure:"ufile_key_name" required:"false" cty:"ufile_key_name" hcl:"ufile_key_name"`
	SkipClean             *bool             `mapstructure:"skip_clean" required:"false" cty:"skip_clean" hcl:"skip_clean"`
	ImageName             *string           `mapstructure:"image_name" required:"true" cty:"image_name" hcl:"image_name"`
	ImageDescription      *string           `mapstructure:"image_description" required:"false" cty:"image_description" hcl:"image_description"`
	OSType                *string           `mapstructure:"image_os_type" required:"true" cty:"image_os_type" hcl:"image_os_type"`
	OSName                *string           `mapstructure:"image_os_name" required:"true" cty:"image_os_name" hcl:"image_os_name"`
	Format                *string           `mapstructure:"format" required:"true" cty:"format" hcl:"format"`
	WaitImageReadyTimeout *int              `mapstructure:"wait_image_ready_timeout" required:"false" cty:"wait_image_ready_timeout" hcl:"wait_image_ready_timeout"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"public_key":                 &hcldec.AttrSpec{Name: "public_key", Type: cty.String, Required: false},
		"private_key":                &hcldec.AttrSpec{Name: "private_key", Type: cty.String, Required: false},
		"region":                     &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"project_id":                 &hcldec.AttrSpec{Name: "project_id", Type: cty.String, Required: false},
		"base_url":                   &hcldec.AttrSpec{Name: "base_url", Type: cty.String, Required: false},
		"profile":                    &hcldec.AttrSpec{Name: "profile", Type: cty.String, Required: false},
		"shared_credentials_file":    &hcldec.AttrSpec{Name: "shared_credentials_file", Type: cty.String, Required: false},
		"ufile_bucket_name":          &hcldec.AttrSpec{Name: "ufile_bucket_name", Type: cty.String, Required: false},
		"ufile_key_name":             &hcldec.AttrSpec{Name: "ufile_key_name", Type: cty.String, Required: false},
		"skip_clean":                 &hcldec.AttrSpec{Name: "skip_clean", Type: cty.Bool, Required: false},
		"image_name":                 &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"image_description":          &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"image_os_type":              &hcldec.AttrSpec{Name: "image_os_type", Type: cty.String, Required: false},
		"image_os_name":              &hcldec.AttrSpec{Name: "image_os_name", Type: cty.String, Required: false},
		"format":                     &hcldec.AttrSpec{Name: "format", Type: cty.String, Required: false},
		"wait_image_ready_timeout":   &hcldec.AttrSpec{Name: "wait_image_ready_timeout", Type: cty.Number, Required: false},
	}
	return s
}

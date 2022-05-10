// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cellsapi-idm.proto

package rest

import (
	fmt "fmt"
	math "math"
	proto "google.golang.org/protobuf/proto"
	_ "github.com/pydio/cells/v4/common/proto/service"
	_ "github.com/pydio/cells/v4/common/proto/idm"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ResourcePolicyQuery) Validate() error {
	return nil
}
func (this *SearchRoleRequest) Validate() error {
	for _, item := range this.Queries {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Queries", err)
			}
		}
	}
	if this.ResourcePolicyQuery != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResourcePolicyQuery); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResourcePolicyQuery", err)
		}
	}
	return nil
}
func (this *RolesCollection) Validate() error {
	for _, item := range this.Roles {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	return nil
}
func (this *SearchUserRequest) Validate() error {
	for _, item := range this.Queries {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Queries", err)
			}
		}
	}
	if this.ResourcePolicyQuery != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResourcePolicyQuery); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResourcePolicyQuery", err)
		}
	}
	return nil
}
func (this *UsersCollection) Validate() error {
	for _, item := range this.Groups {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Groups", err)
			}
		}
	}
	for _, item := range this.Users {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
			}
		}
	}
	return nil
}
func (this *BindResponse) Validate() error {
	return nil
}
func (this *SearchACLRequest) Validate() error {
	for _, item := range this.Queries {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Queries", err)
			}
		}
	}
	return nil
}
func (this *ACLCollection) Validate() error {
	for _, item := range this.ACLs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ACLs", err)
			}
		}
	}
	return nil
}
func (this *SearchWorkspaceRequest) Validate() error {
	for _, item := range this.Queries {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Queries", err)
			}
		}
	}
	if this.ResourcePolicyQuery != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResourcePolicyQuery); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResourcePolicyQuery", err)
		}
	}
	return nil
}
func (this *WorkspaceCollection) Validate() error {
	for _, item := range this.Workspaces {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Workspaces", err)
			}
		}
	}
	return nil
}
func (this *UserMetaCollection) Validate() error {
	for _, item := range this.Metadatas {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Metadatas", err)
			}
		}
	}
	return nil
}
func (this *UserMetaNamespaceCollection) Validate() error {
	for _, item := range this.Namespaces {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Namespaces", err)
			}
		}
	}
	return nil
}
func (this *ListUserMetaTagsRequest) Validate() error {
	return nil
}
func (this *ListUserMetaTagsResponse) Validate() error {
	return nil
}
func (this *PutUserMetaTagRequest) Validate() error {
	return nil
}
func (this *PutUserMetaTagResponse) Validate() error {
	return nil
}
func (this *DeleteUserMetaTagsRequest) Validate() error {
	return nil
}
func (this *DeleteUserMetaTagsResponse) Validate() error {
	return nil
}
func (this *UserBookmarksRequest) Validate() error {
	return nil
}
func (this *RevokeRequest) Validate() error {
	return nil
}
func (this *RevokeResponse) Validate() error {
	return nil
}
func (this *ResetPasswordTokenRequest) Validate() error {
	return nil
}
func (this *ResetPasswordTokenResponse) Validate() error {
	return nil
}
func (this *ResetPasswordRequest) Validate() error {
	return nil
}
func (this *ResetPasswordResponse) Validate() error {
	return nil
}
func (this *DocumentAccessTokenRequest) Validate() error {
	return nil
}
func (this *DocumentAccessTokenResponse) Validate() error {
	return nil
}

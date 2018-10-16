// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/identifiers.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import regexp "regexp"
import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/gogo/protobuf/proto"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_ApplicationIdentifiers_ApplicationID = regexp.MustCompile(`^[a-z0-9](?:[-]?[a-z0-9]){1,35}$|^$`)

func (this *ApplicationIdentifiers) Validate() error {
	if !_regex_ApplicationIdentifiers_ApplicationID.MatchString(this.ApplicationID) {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationID", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9](?:[-]?[a-z0-9]){1,35}$|^$"`, this.ApplicationID))
	}
	return nil
}

var _regex_ClientIdentifiers_ClientID = regexp.MustCompile(`^[a-z0-9](?:[-]?[a-z0-9]){1,35}$|^$`)

func (this *ClientIdentifiers) Validate() error {
	if !_regex_ClientIdentifiers_ClientID.MatchString(this.ClientID) {
		return github_com_mwitkow_go_proto_validators.FieldError("ClientID", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9](?:[-]?[a-z0-9]){1,35}$|^$"`, this.ClientID))
	}
	return nil
}

var _regex_EndDeviceIdentifiers_DeviceID = regexp.MustCompile(`^[a-z0-9](?:[-]?[a-z0-9]){1,35}$|^$`)

func (this *EndDeviceIdentifiers) Validate() error {
	if !_regex_EndDeviceIdentifiers_DeviceID.MatchString(this.DeviceID) {
		return github_com_mwitkow_go_proto_validators.FieldError("DeviceID", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9](?:[-]?[a-z0-9]){1,35}$|^$"`, this.DeviceID))
	}
	if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(&(this.ApplicationIdentifiers)); err != nil {
		return github_com_mwitkow_go_proto_validators.FieldError("ApplicationIdentifiers", err)
	}
	if !(len(this.DevEUI) == 8) {
		return github_com_mwitkow_go_proto_validators.FieldError("DevEUI", fmt.Errorf(`value '%v' must length be not equal '8'`, this.DevEUI))
	}
	if !(len(this.JoinEUI) == 8) {
		return github_com_mwitkow_go_proto_validators.FieldError("JoinEUI", fmt.Errorf(`value '%v' must length be not equal '8'`, this.JoinEUI))
	}
	if !(len(this.DevAddr) == 4) {
		return github_com_mwitkow_go_proto_validators.FieldError("DevAddr", fmt.Errorf(`value '%v' must length be not equal '4'`, this.DevAddr))
	}
	return nil
}

var _regex_GatewayIdentifiers_GatewayID = regexp.MustCompile(`^[a-z0-9](?:[-]?[a-z0-9]){1,35}$|^$`)

func (this *GatewayIdentifiers) Validate() error {
	if !_regex_GatewayIdentifiers_GatewayID.MatchString(this.GatewayID) {
		return github_com_mwitkow_go_proto_validators.FieldError("GatewayID", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9](?:[-]?[a-z0-9]){1,35}$|^$"`, this.GatewayID))
	}
	if !(len(this.EUI) == 8) {
		return github_com_mwitkow_go_proto_validators.FieldError("EUI", fmt.Errorf(`value '%v' must length be not equal '8'`, this.EUI))
	}
	return nil
}

var _regex_OrganizationIdentifiers_OrganizationID = regexp.MustCompile(`^[a-z0-9](?:[-]?[a-z0-9]){1,35}$|^$`)

func (this *OrganizationIdentifiers) Validate() error {
	if !_regex_OrganizationIdentifiers_OrganizationID.MatchString(this.OrganizationID) {
		return github_com_mwitkow_go_proto_validators.FieldError("OrganizationID", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9](?:[-]?[a-z0-9]){1,35}$|^$"`, this.OrganizationID))
	}
	return nil
}

var _regex_UserIdentifiers_UserID = regexp.MustCompile(`^[a-z0-9](?:[-]?[a-z0-9]){1,35}$|^$`)

func (this *UserIdentifiers) Validate() error {
	if !_regex_UserIdentifiers_UserID.MatchString(this.UserID) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserID", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9](?:[-]?[a-z0-9]){1,35}$|^$"`, this.UserID))
	}
	return nil
}
func (this *OrganizationOrUserIdentifiers) Validate() error {
	if oneOfNester, ok := this.GetIds().(*OrganizationOrUserIdentifiers_OrganizationIDs); ok {
		if oneOfNester.OrganizationIDs != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.OrganizationIDs); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OrganizationIDs", err)
			}
		}
	}
	if oneOfNester, ok := this.GetIds().(*OrganizationOrUserIdentifiers_UserIDs); ok {
		if oneOfNester.UserIDs != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.UserIDs); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UserIDs", err)
			}
		}
	}
	return nil
}
func (this *EntityIdentifiers) Validate() error {
	if oneOfNester, ok := this.GetIds().(*EntityIdentifiers_ApplicationIDs); ok {
		if oneOfNester.ApplicationIDs != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ApplicationIDs); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ApplicationIDs", err)
			}
		}
	}
	if oneOfNester, ok := this.GetIds().(*EntityIdentifiers_ClientIDs); ok {
		if oneOfNester.ClientIDs != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ClientIDs); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ClientIDs", err)
			}
		}
	}
	if oneOfNester, ok := this.GetIds().(*EntityIdentifiers_DeviceIDs); ok {
		if oneOfNester.DeviceIDs != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.DeviceIDs); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DeviceIDs", err)
			}
		}
	}
	if oneOfNester, ok := this.GetIds().(*EntityIdentifiers_GatewayIDs); ok {
		if oneOfNester.GatewayIDs != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.GatewayIDs); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("GatewayIDs", err)
			}
		}
	}
	if oneOfNester, ok := this.GetIds().(*EntityIdentifiers_OrganizationIDs); ok {
		if oneOfNester.OrganizationIDs != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.OrganizationIDs); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("OrganizationIDs", err)
			}
		}
	}
	if oneOfNester, ok := this.GetIds().(*EntityIdentifiers_UserIDs); ok {
		if oneOfNester.UserIDs != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.UserIDs); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UserIDs", err)
			}
		}
	}
	return nil
}
func (this *CombinedIdentifiers) Validate() error {
	for _, item := range this.EntityIdentifiers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("EntityIdentifiers", err)
			}
		}
	}
	return nil
}
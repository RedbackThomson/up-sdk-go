// Copyright 2022 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package organizations

import (
	"time"

	"github.com/google/uuid"
)

// Organization is an organization on Upbound.
type Organization struct {
	ID                   uint                        `json:"id"`
	Name                 string                      `json:"name"`
	DisplayName          string                      `json:"displayName"`
	CreatorID            uint                        `json:"creatorId"`
	Role                 OrganizationPermissionGroup `json:"role"`
	ReservedEnvironments int                         `json:"reservedEnvironments"`
}

// Robot is a robot account on Upbound.
type Robot struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	TeamIDs     []uuid.UUID `json:"teamIDs"`
	TokenIDs    []uuid.UUID `json:"tokenIDs"`
	CreatedAt   time.Time   `json:"createdAt"`
}

// Team is a team on Upbound.
type Team struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// OrganizationPermissionGroup is the type of permission a user has in the
// organization.
type OrganizationPermissionGroup string

const (
	// OrganizationMember denotes basic permission on an organization.
	OrganizationMember OrganizationPermissionGroup = "member"
	// OrganizationOwner denotes full access permission on an organization.
	OrganizationOwner OrganizationPermissionGroup = "owner"
)

// OrganizationCreateParameters are the parameters for creating an organization.
type OrganizationCreateParameters struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

// User is a user on Upbound.
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// Member is a member of an organization.
type Member struct {
	Permission OrganizationPermissionGroup `json:"permission"`
	User       User                        `json:"user"`
}

// Invite is a pending organization member.
type Invite struct {
	ID         uint                        `json:"id"`
	Email      string                      `json:"email"`
	Permission OrganizationPermissionGroup `json:"permission"`
	CreatedAt  string                      `json:"createdAt"`
}

// OrganizationInviteCreateParameters are the parameters for creating a user invite to an organization.
type OrganizationInviteCreateParameters struct {
	Email      string                      `json:"email"`
	Permission OrganizationPermissionGroup `json:"organizationPermission"`
}

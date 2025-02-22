package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/sourcehub/x/acp/policy"
	"github.com/sourcenetwork/sourcehub/x/acp/testutil"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

func TestMsgCreatePolicy_ValidPolicyIsCreated(t *testing.T) {
	policyStr := `
name: policy
description: ok
resources:
  file:
    relations: 
      owner:
        doc: owner owns
        types:
          - actor-resource
      reader:
      admin:
        manages:
          - reader
    permissions: 
      own:
        expr: owner
        doc: own doc
      read: 
        expr: owner + reader
actor:
  name: actor-resource
  doc: my actor
          `

	ctx, msgServer, accKeeper := setupMsgServer(t)

	key := secp256k1.GenPrivKeyFromSecret(nil).PubKey()
	accKeeper.NewAccount(key)
	creator := "cosmos1346fyal5a9xxwlygkqmkkqf7g3q3zwzpdmkam8"

	_ = accKeeper.GenAccount().GetAddress().String()

	msg := types.MsgCreatePolicy{
		Creator:      creator,
		Policy:       policyStr,
		MarshalType:  types.PolicyMarshalingType_SHORT_YAML,
		CreationTime: timestamp,
	}
	resp, err := msgServer.CreatePolicy(ctx, &msg)

	require.Nil(t, err)

	require.Equal(t, resp.Policy, &types.Policy{
		Id:           "4419a8abb886c641bc794b9b3289bc2118ab177542129627b6b05d540de03e46",
		Name:         "policy",
		Description:  "ok",
		CreationTime: timestamp,
		Creator:      creator,
		Resources: []*types.Resource{
			&types.Resource{
				Name: "file",
				Relations: []*types.Relation{
					&types.Relation{
						Name: "admin",
						Manages: []string{
							"reader",
						},
						VrTypes: []*types.Restriction{},
					},
					&types.Relation{
						Name: "owner",
						Doc:  "owner owns",
						VrTypes: []*types.Restriction{
							&types.Restriction{
								ResourceName: "actor-resource",
								RelationName: "",
							},
						},
					},
					&types.Relation{
						Name: "reader",
					},
				},
				Permissions: []*types.Permission{
					&types.Permission{
						Name:       "own",
						Expression: "owner",
						Doc:        "own doc",
					},
					&types.Permission{
						Name:       "read",
						Expression: "owner + reader",
					},
					&types.Permission{
						Name:       "_can_manage_admin",
						Expression: "owner",
						Doc:        "permission controls actors which are allowed to create relationships for the admin relation (permission was auto-generated by SourceHub).",
					},
					&types.Permission{
						Name:       "_can_manage_owner",
						Expression: "owner",
						Doc:        "permission controls actors which are allowed to create relationships for the owner relation (permission was auto-generated by SourceHub).",
					},
					&types.Permission{
						Name:       "_can_manage_reader",
						Expression: "(admin + owner)",
						Doc:        "permission controls actors which are allowed to create relationships for the reader relation (permission was auto-generated by SourceHub).",
					},
				},
			},
		},
		ActorResource: &types.ActorResource{
			Name: "actor-resource",
			Doc:  "my actor",
		},
	})

	event := &types.EventPolicyCreated{
		Creator:    creator,
		PolicyId:   "4419a8abb886c641bc794b9b3289bc2118ab177542129627b6b05d540de03e46",
		PolicyName: "policy",
	}
	testutil.AssertEventEmmited(t, ctx, event)
}

func TestMsgCreatePolicy_PolicyResourcesRequiresOwnerRelation(t *testing.T) {
	pol := `
name: policy
description: ok
resources:
  file:
    relations: 
      reader:
    permissions: 
  foo:
    relations:
      owner:
    permissions:
`

	ctx, msgServer, accKeeper := setupMsgServer(t)
	creator := accKeeper.GenAccount().GetAddress().String()
	msg := types.NewMsgCreatePolicyNow(creator, pol, types.PolicyMarshalingType_SHORT_YAML)
	resp, err := msgServer.CreatePolicy(ctx, msg)

	require.Nil(t, resp)
	require.ErrorIs(t, err, policy.ErrResourceMissingOwnerRelation)
}

func TestMsgCreatePolicy_ManagementReferencingUndefinedRelationReturnsError(t *testing.T) {
	pol := `
name: policy
description: ok
resources:
  file:
    relations: 
      owner:
      admin:
        manages:
          - deleter
    permissions: 
`

	ctx, msgServer, accKeeper := setupMsgServer(t)
	creator := accKeeper.GenAccount().GetAddress().String()

	msg := types.NewMsgCreatePolicyNow(creator, pol, types.PolicyMarshalingType_SHORT_YAML)
	resp, err := msgServer.CreatePolicy(ctx, msg)

	require.Nil(t, resp)
	require.ErrorIs(t, err, policy.ErrInvalidManagementRule)
}

func TestMsgCreatePolicy_InvalidCreatorAddressCausesError(t *testing.T) {
	pol := `
name: policy
resources:
`

	ctx, msgServer, _ := setupMsgServer(t)
	msg := types.NewMsgCreatePolicyNow("creator", pol, types.PolicyMarshalingType_SHORT_YAML)
	resp, err := msgServer.CreatePolicy(ctx, msg)

	require.Nil(t, resp)
	require.ErrorIs(t, err, policy.ErrInvalidCreator)
}

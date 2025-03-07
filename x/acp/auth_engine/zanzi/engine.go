package zanzi

import (
	"context"
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	rcdb "github.com/sourcenetwork/raccoondb"
	"github.com/sourcenetwork/zanzi"
	"github.com/sourcenetwork/zanzi/pkg/api"
	"github.com/sourcenetwork/zanzi/pkg/domain"

	"github.com/sourcenetwork/sourcehub/utils"
	"github.com/sourcenetwork/sourcehub/x/acp/auth_engine"
	"github.com/sourcenetwork/sourcehub/x/acp/types"
)

var _ auth_engine.AuthEngine = (*Zanzi)(nil)

type RecordFound = auth_engine.RecordFound

// NewZanzi builds an AuthEngine with zanzi as backend
func NewZanzi(kv storetypes.KVStore, logger log.Logger) (*Zanzi, error) {
	store := rcdb.KvFromCosmosKv(kv)
	wrappedLogger := &loggerWrapper{logger}

	z, err := zanzi.New(
		zanzi.WithKVStore(store),
		zanzi.WithLogger(wrappedLogger),
	)
	if err != nil {
		return nil, err
	}

	return &Zanzi{
		zanzi:        z,
		policyMapper: policyMapper{},
	}, nil
}

// Zanzi implements AuthEngine from zanzi's PolicyService
type Zanzi struct {
	zanzi        zanzi.Zanzi
	policyMapper policyMapper
}

func (z *Zanzi) GetRelationship(ctx context.Context, policy *types.Policy, rel *types.Relationship) (*types.RelationshipRecord, error) {
	serv := z.zanzi.GetPolicyService()
	mapper := newRelationshipMapper(policy.ActorResource.Name)

	req := &api.GetRelationshipRequest{
		PolicyId:     policy.Id,
		Relationship: mapper.ToZanziRelationship(rel),
	}

	result, err := serv.GetRelationship(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetRelationship: %w", err)
	}

	fetchedRel, err := mapper.FromZanziRelationship(result.Record)
	if err != nil {
		return nil, fmt.Errorf("GetRelationship: %w", err)
	}

	return fetchedRel, nil
}

func (z *Zanzi) SetRelationship(ctx context.Context, policy *types.Policy, rec *types.RelationshipRecord) (RecordFound, error) {
	serv := z.zanzi.GetPolicyService()
	mapper := newRelationshipMapper(policy.ActorResource.Name)

	rec.PolicyId = policy.Id
	zanziRecord, err := mapper.ToZanziRelationshipRecord(rec)
	if err != nil {
		return false, fmt.Errorf("SetRelationship: %w", err)
	}

	req := &api.SetRelationshipRequest{
		PolicyId:     policy.Id,
		Relationship: zanziRecord.Relationship,
		AppData:      zanziRecord.AppData,
	}

	response, err := serv.SetRelationship(ctx, req)
	if err != nil {
		return false, fmt.Errorf("SetRelationship: %w: %w", err, types.ErrAcpInput)
	}

	return RecordFound(response.RecordOverwritten), nil
}

func (z *Zanzi) GetPolicy(ctx context.Context, policyId string) (*types.PolicyRecord, error) {
	serv := z.zanzi.GetPolicyService()

	req := api.GetPolicyRequest{
		Id: policyId,
	}
	res, err := serv.GetPolicy(ctx, &req)
	if err != nil {
		return nil, err
	}
	if res.Record == nil {
		return nil, nil
	}

	mapped, err := z.policyMapper.FromZanzi(res.Record)
	if err != nil {
		return nil, err
	}

	return mapped, nil
}

func (z *Zanzi) SetPolicy(ctx context.Context, record *types.PolicyRecord) error {
	serv := z.zanzi.GetPolicyService()

	zanziRecord, err := z.policyMapper.ToZanziRecord(record)
	if err != nil {
		return err
	}

	req := api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_Policy{
				Policy: zanziRecord.Policy,
			},
		},
		AppData: zanziRecord.AppData,
	}
	_, err = serv.CreatePolicy(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}

func (z *Zanzi) FilterRelationships(ctx context.Context, policy *types.Policy, selector *types.RelationshipSelector) ([]*types.RelationshipRecord, error) {
	serv := z.zanzi.GetPolicyService()
	relationshipMapper := newRelationshipMapper(policy.ActorResource.Name)
	selectorMapper := newSelectorMapper(relationshipMapper)

	zanziSelector, err := selectorMapper.ToZanziSelector(selector)
	if err != nil {
		return nil, fmt.Errorf("FilterRelationships: %v", err)
	}

	req := api.FindRelationshipRecordsRequest{
		PolicyId: policy.Id,
		Selector: zanziSelector,
	}

	resp, err := serv.FindRelationshipRecords(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("FilterRelationships: %w: %v", types.ErrAcpInput, err)
	}

	records, err := utils.MapFailableSlice(resp.Result.Records, relationshipMapper.FromZanziRelationship)
	if err != nil {
		return nil, fmt.Errorf("FilterRelationships: %v", err)
	}

	return records, nil
}

func (z *Zanzi) Check(ctx context.Context, policy *types.Policy, operation *types.Operation, actor *types.Actor) (bool, error) {
	service := z.zanzi.GetRelationGraphService()
	mapper := newRelationshipMapper(policy.ActorResource.Name)

	req := &api.CheckRequest{
		PolicyId: policy.Id,
		AccessRequest: &domain.AccessRequest{
			Object:   mapper.MapObject(operation.Object),
			Relation: operation.Permission,
			Subject: &domain.Entity{
				Resource: policy.ActorResource.Name,
				Id:       actor.Id,
			},
		},
	}
	response, err := service.Check(ctx, req)
	if err != nil {
		return false, fmt.Errorf("Check: %w", err)
	}

	return response.Result.Authorized, nil
}

func (z *Zanzi) DeleteRelationship(ctx context.Context, policy *types.Policy, relationship *types.Relationship) (RecordFound, error) {
	service := z.zanzi.GetPolicyService()
	mapper := newRelationshipMapper(policy.ActorResource.Name)

	req := api.DeleteRelationshipRequest{
		PolicyId:     policy.Id,
		Relationship: mapper.ToZanziRelationship(relationship),
	}
	response, err := service.DeleteRelationship(ctx, &req)
	if err != nil {
		return false, err
	}

	return RecordFound(response.Found), nil
}

func (z *Zanzi) DeleteRelationships(ctx context.Context, policy *types.Policy, selector *types.RelationshipSelector) (uint, error) {
	service := z.zanzi.GetPolicyService()
	relationshipMapper := newRelationshipMapper(policy.ActorResource.Name)
	selectorMapper := newSelectorMapper(relationshipMapper)

	zanziSelector, err := selectorMapper.ToZanziSelector(selector)
	if err != nil {
		return 0, fmt.Errorf("DeleteRelationships: %w", err)
	}

	request := api.DeleteRelationshipsRequest{
		PolicyId: policy.Id,
		Selector: zanziSelector,
	}
	response, err := service.DeleteRelationships(ctx, &request)
	if err != nil {
		return 0, fmt.Errorf("DeleteRelationships: %w", err)
	}

	return uint(response.RecordsAffected), nil
}

func (z *Zanzi) ListPolicyIds(ctx context.Context) ([]string, error) {
	service := z.zanzi.GetPolicyService()

	req := api.ListPolicyIdsRequest{}
	resp, err := service.ListPolicyIds(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("ListPolicyIds: %w", err)
	}

	return utils.MapSlice(resp.Records, func(rec *api.ListPolicyIdsResponse_Record) string {
		return rec.Id
	}), nil
}

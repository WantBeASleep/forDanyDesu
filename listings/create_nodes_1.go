package node_segment

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"uzi/internal/domain"
	nodeEntity "uzi/internal/repository/node/entity"
	segmentEntity "uzi/internal/repository/segment/entity"
)

const (
	avgSegmentPerNode = 3
)

func (s *service) createNodesWithSegments(
	ctx context.Context,
	uziID uuid.UUID,
	ai bool,
	arg []CreateNodesWithSegmentsArg,
	opts ...CreateNodesWithSegmentsOption,
) ([]CreateNodesWithSegmentsID, error) {
	opt := &createNodesWithSegmentsOption{}
	for _, o := range opts {
		o(opt)
	}

	nodes, segments, ids := s.createDomainNodeSegmentsFromArgs(uziID, ai, arg)

	if opt.setNodesValidation != nil {
		for i := range nodes {
			nodes[i].Validation = opt.setNodesValidation
		}
	}

	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = s.dao.RollbackTx(ctx) }()

	nodeQuery := s.dao.NewNodeQuery(ctx)
	segmentQuery := s.dao.NewSegmentQuery(ctx)

	if err := nodeQuery.InsertNodes(nodeEntity.Node{}.SliceFromDomain(nodes)...); err != nil {
		return nil, fmt.Errorf("insert nodes: %w", err)
	}

	if err := segmentQuery.InsertSegments(segmentEntity.Segment{}.SliceFromDomain(segments)...); err != nil {
		return nil, fmt.Errorf("insert segments: %w", err)
	}
	if opt.newUziStatus != nil {
		if err := s.dao.NewUziQuery(ctx).UpdateUziStatus(uziID, opt.newUziStatus.String()); err != nil {
			return nil, fmt.Errorf("update uzi status: %w", err)
		}
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return nil, fmt.Errorf("commit transaction: %w", err)
	}

	return ids, nil
}

func (s *service) SaveProcessedNodesWithSegments(
	ctx context.Context,
	uziID uuid.UUID,
	arg []CreateNodesWithSegmentsArg,
) error {
	_, err := s.createNodesWithSegments(
		ctx,
		uziID,
		true,
		arg,
		WithNewUziStatus(domain.UziStatusCompleted),
		WithSetNodesValidation(domain.NodeValidationNull),
	)

	return err
}
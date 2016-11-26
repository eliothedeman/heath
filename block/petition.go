package block

func (p *Petition) Valid(transactions []*Transaction) bool {
	var scratchA []byte
	var scratchB []byte
	for _, t := range transactions {
		scratchA = hash(t.GetPayload())
		scratchB = append(scratchB, scratchA...)

	}

	return true
}

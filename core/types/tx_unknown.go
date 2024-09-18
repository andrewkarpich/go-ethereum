package types

type UnknownTx struct {
	LegacyTx
}

func (tx *UnknownTx) txType() byte {
	return UnknownTxType
}

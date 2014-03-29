package index

type DocFieldConsumer interface {
	finishDocument() error
}

/*
This is a DocFieldCosumer that inverts each field, separately, from a
Document, and accepts an InvertedTermsConsumer to process those terms.
*/
type DocInverter struct {
	consumer    InvertedDocConsumer
	endConsumer InvertedDocEndConsumer
	docState    *docState
}

func newDocInverter(docState *docState, consumer InvertedDocConsumer,
	endConsumer InvertedDocEndConsumer) *DocInverter {

	return &DocInverter{consumer, endConsumer, docState}
}

func (di *DocInverter) finishDocument() error {
	panic("not implemented yet")
}
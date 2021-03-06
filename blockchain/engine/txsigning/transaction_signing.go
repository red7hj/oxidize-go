package txsigning

import (
	log "github.com/sirupsen/logrus"

	"github.com/tclchiam/oxidize-go/crypto"
	"github.com/tclchiam/oxidize-go/blockchain/entity"
	"github.com/tclchiam/oxidize-go/identity"
)

func GenerateSignature(input *entity.UnsignedInput, outputs []*entity.Output, spender *identity.Identity, encoder entity.TransactionEncoder) *crypto.Signature {
	signatureData := serializeSignatureData(input, outputs, encoder)
	signature, err := spender.Sign(signatureData)
	if err != nil {
		log.Panic(err)
	}

	return signature
}

func VerifySignature(input *entity.SignedInput, outputs []*entity.Output, encoder entity.TransactionEncoder) (verified bool) {
	unsignedInput := &entity.UnsignedInput{
		PublicKey:       input.PublicKey,
		OutputReference: input.OutputReference,
	}
	signatureData := serializeSignatureData(unsignedInput, outputs, encoder)

	return input.PublicKey.Verify(signatureData, input.Signature)
}

func serializeSignatureData(input *entity.UnsignedInput, outputs []*entity.Output, encoder entity.TransactionEncoder) []byte {
	var data []byte

	bytes, err := encoder.EncodeUnsignedInput(input)
	if err != nil {
		log.Panic(err)
	}

	data = append(data, bytes...)
	for _, output := range outputs {
		bytes, err := encoder.EncodeOutput(output)
		if err != nil {
			log.Panic(err)
		}
		data = append(data, bytes...)
	}

	return data
}

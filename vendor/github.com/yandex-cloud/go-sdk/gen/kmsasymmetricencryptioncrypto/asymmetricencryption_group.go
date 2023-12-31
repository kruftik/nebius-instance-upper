// Code generated by sdkgen. DO NOT EDIT.

package asymmetricencryption

import (
	"context"

	"google.golang.org/grpc"
)

// KMSAsymmetricEncryptionCrypto provides access to "asymmetricencryption" component of Yandex.Cloud
type KMSAsymmetricEncryptionCrypto struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewKMSAsymmetricEncryptionCrypto creates instance of KMSAsymmetricEncryptionCrypto
func NewKMSAsymmetricEncryptionCrypto(g func(ctx context.Context) (*grpc.ClientConn, error)) *KMSAsymmetricEncryptionCrypto {
	return &KMSAsymmetricEncryptionCrypto{g}
}

// AsymmetricEncryptionCrypto gets AsymmetricEncryptionCryptoService client
func (k *KMSAsymmetricEncryptionCrypto) AsymmetricEncryptionCrypto() *AsymmetricEncryptionCryptoServiceClient {
	return &AsymmetricEncryptionCryptoServiceClient{getConn: k.getConn}
}

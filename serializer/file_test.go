package serializer_test

import (
	"testing"

	"github.com/YuichiNAGAO/grpc_sample_app/sample"
	"github.com/YuichiNAGAO/grpc_sample_app/serializer"
	"github.com/stretchr/testify/require"
)

func TestFileSerialiser(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	require.NoError(t, err)
}

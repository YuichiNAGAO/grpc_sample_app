package serializer_test

import (
	"testing"

	"github.com/YuichiNAGAO/grpc_sample_app/pb"
	"github.com/YuichiNAGAO/grpc_sample_app/sample"
	"github.com/YuichiNAGAO/grpc_sample_app/serializer"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestFileSerialiser(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)

	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}

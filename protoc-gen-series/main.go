package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	pbfeatures := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.
		Init(pgs.DebugEnv("DEBUG"), pgs.SupportedFeatures(&pbfeatures)).
		RegisterModule(newSeries()).Render()
}

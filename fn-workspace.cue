module: "namespacelabs.dev/foundation"
requirements: {
	api: 37
}
prebuilts: {
	digest: {
		"namespacelabs.dev/foundation/cmd/ns":                                                         "sha256:c18378b6ea6c0bdec0525638c38fcde7c2f0f5db94b95a5c4448125f27130ed5"
		"namespacelabs.dev/foundation/cmd/nspipelines":                                                "sha256:0810371a59e47ad07fb5c86145d46d52151b7c64f4f6ee3961de498b051e5b5b"
		"namespacelabs.dev/foundation/devworkflow/web":                                                "sha256:3b29d1e923e58799326818549d70cb2a5b76ca63aca8d05cec1d1cf87ad6a29a"
		"namespacelabs.dev/foundation/internal/sdk/buf/image/prebuilt":                                "sha256:271b1c4185353dfcaf93cfae52f5446bc8eef74effde38de94de84021da209ec"
		"namespacelabs.dev/foundation/std/dev/controller":                                             "sha256:795643a10f044e9987a89fdf577b02864e65a45fd1785b215687bd36610a17e5"
		"namespacelabs.dev/foundation/std/development/filesync/controller":                            "sha256:84f793519683bc3fbab216362442727e1d2078918b7db320a40229b8ab19728f"
		"namespacelabs.dev/foundation/std/grpc/httptranscoding/configure":                             "sha256:fbfb7ea15ec348ec2fff10c66acf2626dbbf7ff2ba6143cd24452f5be5949169"
		"namespacelabs.dev/foundation/std/monitoring/grafana/tool":                                    "sha256:f10c7525e286c936ed4b0ee4e18c9f0095555a268fab9b9f987658cff938b8ab"
		"namespacelabs.dev/foundation/std/monitoring/prometheus/tool":                                 "sha256:eca44f5294efd69e719e9c0baa049f8c1d11475a6297924629f01b7be853597c"
		"namespacelabs.dev/foundation/std/networking/gateway/controller":                              "sha256:9012332c31fa50ca69587452eca8572121b8cb643035f990c851d948740588c7"
		"namespacelabs.dev/foundation/std/networking/gateway/server/configure":                        "sha256:a20beb03bb10e0be198dd82954965b8abe9036ff053987d3ad1b80716c9dfaf7"
		"namespacelabs.dev/foundation/std/runtime/kubernetes/controller/img":                          "sha256:5ab27cab12d01b1e2b7a03e022658ccbf1d0d07ef8756c69a1d0cf9169412fbf"
		"namespacelabs.dev/foundation/std/runtime/kubernetes/controller/tool":                         "sha256:bf7f919692141ab63a9002fe6ba1ae9641a60f770ff97f41edb8fbfa09178537"
		"namespacelabs.dev/foundation/std/runtime/kubernetes/kube-state-metrics/configure":            "sha256:4bfcba46c96c4d046f70f8b7ad2a7b2263606c9fa0cc17fdc939b95dfb6c3850"
		"namespacelabs.dev/foundation/std/secrets/kubernetes":                                         "sha256:5a16ec45ad9818b4b46cc495a7e41c1999d8d34e50ecb487e9881c8835e8eba3"
		"namespacelabs.dev/foundation/std/startup/testdriver":                                         "sha256:59b589f655843a9eae4f3c63df03de8e216ac1811fe6710b491c8375a8489a09"
		"namespacelabs.dev/foundation/std/testdata/datastore/keygen":                                  "sha256:390f9dcd59d3c59184a6eb86935d90b2ce5578aecbca3607db80972be2a79bf7"
		"namespacelabs.dev/foundation/std/web/http/configure":                                         "sha256:e3a311fb02a4c73c23413dc8c6998fac8d9dd1a6e66c1322c6ea4c4768a47ee9"
		"namespacelabs.dev/foundation/universe/aws/ecr/configure":                                     "sha256:b6cd04e4791478d881e576f10e864e330fd32f504573e3ccb1e60614450da0fd"
		"namespacelabs.dev/foundation/universe/aws/irsa/prepare":                                      "sha256:2c128fb8236730b8aaf399dd62dbdce6d5bced2fa545f1c5a753cac22f4fbb5d"
		"namespacelabs.dev/foundation/universe/aws/s3/internal/configure":                             "sha256:2920d9650f070f736d2bb643c9067b3f7c51b1c10468dbfcb2c8b015a2e1234a"
		"namespacelabs.dev/foundation/universe/aws/s3/internal/managebuckets/init":                    "sha256:af5dec5af48a76c516edd01fa058c8c8cad5e2a77d84307db50898594f9e96bc"
		"namespacelabs.dev/foundation/universe/db/maria/incluster/tool":                               "sha256:cd2b8934b79d3f85cf713f8ecbaa021c0148cb2042333a3c347d6a7521775d08"
		"namespacelabs.dev/foundation/universe/db/maria/init":                                         "sha256:4b301fcb0b5619d4aafaf608c0257ee5c00507a8016b308586934b17ad16e348"
		"namespacelabs.dev/foundation/universe/db/maria/server/creds/tool":                            "sha256:7def752d6316f42969ddd08fef5a87dec2a45cd6dc6231b91b55da216960b34f"
		"namespacelabs.dev/foundation/universe/db/postgres/incluster/tool":                            "sha256:fd907ebcf944359c0256d65d9f373e6e0ca369b1d39c7cce84424a7be4d01d12"
		"namespacelabs.dev/foundation/universe/db/postgres/init":                                      "sha256:08f62e761d82cdc8b85f1a038e9390fd48619966af40d789d1cf6ec67ef61e8b"
		"namespacelabs.dev/foundation/universe/db/postgres/opaque/tool":                               "sha256:1c25e88a71c091d98d5386e58620101b01e1250b63b2ed53ba93d1f7858ca81f"
		"namespacelabs.dev/foundation/universe/db/postgres/server/creds/tool":                         "sha256:c97d887b68a20562d6f60b8362cf69e95248f90131fc30f0f2fc02ad5207faec"
		"namespacelabs.dev/foundation/universe/db/postgres/server/img":                                "sha256:5299fbc48f11fb9de0a86380260187259ff26e4b79840b30dfec137c1b5725a3"
		"namespacelabs.dev/foundation/universe/development/localstack/s3/internal/configure":          "sha256:68bca065e40b79693d3c1360af9ebc83dc2ec9678ae8faacf407dcd0aabddc95"
		"namespacelabs.dev/foundation/universe/development/localstack/s3/internal/managebuckets/init": "sha256:177e457b6d82dcacb9c08136c33b8d8bbc6cdeb6599d2699522a8376f57a6220"
		"namespacelabs.dev/foundation/universe/networking/tailscale/image":                            "sha256:195801cf0aa421c6957caec3e136a5476cb9a43549dafad17ad9794fd3145a4b"
		"namespacelabs.dev/foundation/universe/storage/s3/internal/managebuckets":                     "sha256:8cf113f222359186f2cc8af167ea969ff40577a1f09f34818168b2973a797390"
		"namespacelabs.dev/foundation/universe/storage/s3/internal/prepare":                           "sha256:886a0312019e9ae9bdba381756a754aa5d93904e32f88136008191d4b275b797"
	}
	baseRepository: "us-docker.pkg.dev/foundation-344819/prebuilts/"
}

package storage

const (
	StorageTypeAzure         = "azure"
	StorageTypeOSS           = "oss"
	StorageTypeS3            = "s3"
	StorageTypeSwift         = "swift"
	StorageTypeCOS           = "cos"
	StorageTypeOBS           = "obs"
	StorageTypeNull          = "null"
	StorageTypeServer        = "server"
	StorageTypeES            = "es"
	StorageTypeElasticSearch = "elasticsearch"
)

type StorageType interface {
	TypeName() string
}

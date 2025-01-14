package dms

const (
	endpointStatusDeleting = "deleting"
)

const (
	engineNameAurora                     = "aurora"
	engineNameAuroraPostgresql           = "aurora-postgresql"
	engineNameAuroraPostgresqlServerless = "aurora-postgresql-serverless"
	engineNameAuroraServerless           = "aurora-serverless"
	engineNameAzuredb                    = "azuredb"
	engineNameDB2                        = "db2"
	engineNameTransfer                   = "dms-transfer"
	engineNameDocDB                      = "docdb"
	engineNameDynamoDB                   = "dynamodb"
	engineNameElasticSearch              = "elasticsearch"
	engineNameKafka                      = "kafka"
	engineNameKinesis                    = "kinesis"
	engineNameMariadb                    = "mariadb"
	engineNameMongodb                    = "mongodb"
	engineNameMySQL                      = "mysql"
	engineNameNeptune                    = "neptune"
	engineNameOpensearch                 = "opensearch"
	engineNameOracle                     = "oracle"
	engineNamePostgres                   = "postgres"
	engineNameRedis                      = "redis"
	engineNameRedshift                   = "redshift"
	engineNameS3                         = "s3"
	engineNameSQLServer                  = "sqlserver"
	engineNameSybase                     = "sybase"
)

func engineName_Values() []string {
	return []string{
		engineNameAurora,
		engineNameAuroraPostgresql,
		engineNameAuroraPostgresqlServerless,
		engineNameAuroraServerless,
		engineNameAzuredb,
		engineNameDB2,
		engineNameTransfer,
		engineNameDocDB,
		engineNameDynamoDB,
		engineNameElasticSearch,
		engineNameKafka,
		engineNameKinesis,
		engineNameMariadb,
		engineNameMongodb,
		engineNameMySQL,
		engineNameNeptune,
		engineNameOpensearch,
		engineNameOracle,
		engineNamePostgres,
		engineNameRedis,
		engineNameRedshift,
		engineNameS3,
		engineNameSQLServer,
		engineNameSybase,
	}
}

const (
	kafkaDefaultTopic = "kafka-default-topic"
)

// https://github.com/aws/aws-sdk-go/issues/2522.
const (
	mongoDBAuthMechanismValueDefault   = "default"
	mongoDBAuthMechanismValueMongodbCr = "mongodb-cr"
	mongoDBAuthMechanismValueScramSha1 = "scram-sha-1"
)

func mongoDBAuthMechanismValue_Values() []string {
	return []string{
		mongoDBAuthMechanismValueDefault,
		mongoDBAuthMechanismValueMongodbCr,
		mongoDBAuthMechanismValueScramSha1,
	}
}

const (
	mongoDBAuthSourceAdmin = "admin"
)

const (
	s3SettingsCompressionTypeGzip = "GZIP"
	s3SettingsCompressionTypeNone = "NONE"
)

func s3SettingsCompressionType_Values() []string {
	return []string{
		s3SettingsCompressionTypeGzip,
		s3SettingsCompressionTypeNone,
	}
}

const (
	s3SettingsEncryptionModeSseKMS = "SSE_KMS"
	s3SettingsEncryptionModeSseS3  = "SSE_S3"
)

func s3SettingsEncryptionMode_Values() []string {
	return []string{
		s3SettingsEncryptionModeSseKMS,
		s3SettingsEncryptionModeSseS3,
	}
}

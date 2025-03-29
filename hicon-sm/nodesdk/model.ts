// Type definitions
type Option<T> = (config: T) => void;

interface TLS {
	certPEM: string;
	privateKeyPEM: string;
	rootCAPEM: string;
}

interface Column {
	name: string;
	type: string;
	nullable: boolean;
	isPrimaryKey: boolean;
	softDelete: boolean;
}

interface RelationColumn {
	name: string;
	refTable: string;
	type: string;
	join: string;
}

interface DBConfig {
	type?: string;
	host?: string;
	port?: number;
	username?: string;
	password?: string;
	database?: string;
	maxCons?: number;
	tls?: TLS;
}

interface Redis {
	host?: string;
	port?: number;
	username?: string;
	password?: string;
	db?: number;
	poolSize?: number;
}

interface TableConfig {
	name: string;
	columns: Column[];
	relationColumns: RelationColumn[];
}

interface UpsertConfig {
	dbConfig?: DBConfig;
	redis?: Redis;
	tableConfigs?: TableConfig[];
	debug?: boolean;
	disableCache?: boolean;
}

// Configuration Builder Class
// Configuration Builder
class ConfigBuilder {
	private config: UpsertConfig;

	constructor() {}

	// Debug Options
	withDebug(debug: boolean): this {
		this.config.debug = debug;
		return this;
	}

	withDisableCache(disable: boolean): this {
		this.config.disableCache = disable;
		return this;
	}

	// Database Configuration
	withDBConfig(dbConfig: DBConfig): this {
		this.config.dbConfig = dbConfig;
		return this;
	}

	// Shorthand database configuration methods
	withDatabase(dbConfig: DBConfig): this {
		this.config.dbConfig = dbConfig;
		return this;
	}

	// Redis Configuration
	withRedis(redis: Redis): this {
		this.config.redis = redis;
		return this;
	}

	// Table Configuration
	withTable(name: string, columns: Column[], relations: RelationColumn[]): this {
		const tableConfig = new TableConfig({
			name,
			columns,
			relationColumns: relations
		});
		this.config.tableConfigs.push(tableConfig);
		return this;
	}

	// Build and Validate
	build(): UpsertConfig {
		if (this.config.dbConfig) {
			this.config.dbConfig.validate();
		}
		return this.config;
	}
}
// Example Usage
function exampleConfiguration() {
	// Example column configurations
	const userColumns: Column[] = [
		{
			name: 'id',
			type: 'uuid',
			isPrimaryKey: true,
			nullable: false,
			softDelete: false
		},
		{
			name: 'username',
			type: 'string',
			nullable: false,
			isPrimaryKey: false,
			softDelete: false
		},
		{
			name: 'email',
			type: 'string',
			nullable: true,
			isPrimaryKey: false,
			softDelete: false
		}
	];

	const userRelations: RelationColumn[] = [
		{
			name: 'role_id',
			refTable: 'roles',
			type: 'uuid',
			join: 'LEFT'
		}
	];

	// Creating configuration using fluent interface
	const configBuilder = new ConfigurationBuilder();
	const config = configBuilder
		.withDebug(true)
		.withDatabase('postgres', 'localhost', 5432)
		.withDatabaseCredentials('user', 'pass')
		.withDatabaseName('myapp')
		.withMaxConnections(10)
		.withRedis('localhost', 6379)
		.withRedisCredentials('', 'redispass')
		.withRedisDB(0)
		.withRedisPoolSize(50)
		.withTable('users', userColumns, userRelations)
		.build();

	return config;
}

// Export the builder for use
export {
	ConfigurationBuilder,
	UpsertConfig,
	DBConfig,
	Redis,
	TableConfig,
	Column,
	RelationColumn
};

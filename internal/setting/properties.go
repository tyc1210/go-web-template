package setting

type ServerProperties struct {
	RunMode  string
	HttpPort string
}

type DataBaseProperties struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTProperties struct {
	Secret string
	Expire int
}

type AppProperties struct {
	LoggerFileName   string
	LoggerLevel      string
	LoggerMaxSize    int
	LoggerMaxBackups int
	LoggerMaxAge     int
	UploadSavePath   string
	UploadServerUrl  string
	UploadMaxSize    int64
	UploadAllowExt   []interface{}
}

local Apps = [
    "fs",
];

local Environment = {
	"MARIADB_HOST":"mariadb",
	"MARIADB_USER": "estore",
	"MARIADB_PASS": "password",
	"PREFIX_DBNAME": "e",
	"DBNAME": "store",

	"MIGRATE_FOLDER": "file://migrate",
	"RESOURCE_ROOT": "/src",

	"REDIS_ADDR": "redis",
	"REDIS_DB": "0",
	"RESOURCE_HOST": "http://localhost:8080",

	"GRPC_PORT": "8090",
	"HTTP_PORT": "8080",
};

local Services = [
    {
        "name": "mariadb",
        "image": "mariadb",
        "environment": {
            "MYSQL_ALLOW_EMPTY_PASSWORD": "yes",
        },
    },
];

local MysqlHealthCheck = {
    "name": "mysql-health-check",
    "image": "mysql",
    "commands": [
        "while ! mysqladmin ping -h mariadb -u root --silent; do sleep 1; done",
        "mysql -h mariadb -u root -e \"CREATE USER 'estore'@'%' IDENTIFIED BY 'password';\"",
        "mysql -h mariadb -u root -e \"GRANT ALTER, CREATE, CREATE VIEW, DELETE, DROP, INDEX, INSERT, SELECT, SHOW VIEW, UPDATE ON *.* TO 'estore'@'%' IDENTIFIED BY 'password';\"",
    ],
};

local FromSecret(key,) = {
    from_secret: key ,
};

local Testing(branch,) = {
    name: "testing-" + branch,
    image: "golang",
    environment: Environment,
    commands: [
        "printenv | grep MARIADB",
        "go run cmd/auto-migrate/main.go",
        "go test -v ./...",
    ],
};

local Builder(branch,) = {
    name: "builder-" + branch,
    image: "golang",
    commands: [
        
    ] + [
        "CGO_ENABLED=0 GOOS=linux GORACH=amd64 go build -o "+ x +" ./cmd/"+x
        for x in Apps
    ],
};

local Publish(trigger,) = {
    name: "publish-" + trigger,
    image: "plugins/docker",
    settings: {
      dockerfile: "deployment/docker/all.Dockerfile",
      repo: "jigadhirasu/estore-"+trigger,
      tags: [ "latest" ],
      username: FromSecret("docker_username"),
      password: FromSecret("docker_password"),
    }
};

local Helm(trigger,) = {
    name: "helm-" + trigger,
    image: "pelotech/drone-helm3",
    settings: {
      values_files: "./deployment/charts/"+trigger+".yaml",
      namespace: trigger,
      mode: "upgrade",
      chart: "./deployment/charts",
      release: "estore",
      api_server: if trigger == "main" then FromSecret("host") else FromSecret("local-host"),
      kube_token: if trigger == "main" then FromSecret("token") else FromSecret("local-token"),
      kube_certificate: if trigger == "main" then FromSecret("ca") else FromSecret("local-ca"),
      kube_service_account: FromSecret("account"),
    },
};

// local Telegram() = {
//     image: "appleboy/drone-telegram",
//     name: "telegram-after",
//     settings: {
//         token: FromSecret("tg_token"),
//         to: FromSecret("tg_to"),
//         format: "markdown",
//         message: "[estore New Commit]\n✅ Build #{{build.number}} of `{{repo.name}}`.\n📝 Commit by {{commit.author}} on `{{commit.branch}}`:\n```\n{{commit.message}}\n```\n🌐 {{ build.link }} ",
//     }
// };

local Pipeline(branch,trigger,) = {
    kind: "pipeline",
    name: trigger,
    type: "kubernetes",
    platform: {
        os: "linux",
        arch: "amd64",
    },
    services: Services,
    steps: [
        MysqlHealthCheck,
        Testing(trigger),
        Builder(trigger),
        Publish(trigger),
        Helm(trigger),
    ],
    trigger: {
        branch:  [ branch ],
    },
};

local Secret(path,name,) = {
    kind: "secret",
    name: name,
    get: {
        path: path,
        name: name,
    },
};

[
    // Pipeline("develop","develop"),
    // Pipeline("release/*","release"),
    Pipeline("main","main"),
]
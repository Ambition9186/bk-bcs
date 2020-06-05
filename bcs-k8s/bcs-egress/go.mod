module bk-bcs/bcs-k8s/bcs-egress

go 1.13

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/go-openapi/spec v0.19.4 // indirect
	github.com/gobuffalo/packr v1.30.1 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/helm/helm-2to3 v0.5.1 // indirect
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/markbates/inflect v1.0.4 // indirect
	github.com/martinlindhe/base36 v1.0.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mikefarah/yq/v2 v2.4.1 // indirect
	github.com/mitchellh/hashstructure v1.0.0 // indirect
	github.com/openshift/api v0.0.0-20200205133042-34f0ec8dab87 // indirect
	github.com/openshift/client-go v0.0.0-20190923180330-3b6373338c9b // indirect
	github.com/operator-framework/api v0.1.1 // indirect
	github.com/operator-framework/operator-sdk v0.17.1
	github.com/otiai10/copy v1.0.2 // indirect
	github.com/prometheus/client_golang v1.5.1 // indirect
	github.com/rogpeppe/go-internal v1.5.0 // indirect
	github.com/rubenv/sql-migrate v0.0.0-20191025130928-9355dd04f4b3 // indirect
	github.com/sirupsen/logrus v1.5.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.4.0 // indirect
	github.com/thanos-io/thanos v0.11.0 // indirect
	github.com/ziutek/mymysql v1.5.4 // indirect
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/tools v0.0.0-20200327195553-82bb89366a1e // indirect
	gomodules.xyz/jsonpatch/v3 v3.0.1 // indirect
	gopkg.in/gorp.v1 v1.7.2 // indirect
	helm.sh/helm/v3 v3.1.2 // indirect
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	k8s.io/apiserver v0.17.4 // indirect
	k8s.io/cli-runtime v0.17.4 // indirect
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/gengo v0.0.0-20191010091904-7fa3014cb28f // indirect
	k8s.io/kube-state-metrics v1.7.2 // indirect
	rsc.io/letsencrypt v0.0.3 // indirect
	sigs.k8s.io/controller-runtime v0.5.2
	sigs.k8s.io/controller-tools v0.2.8 // indirect
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	k8s.io/client-go => k8s.io/client-go v0.17.4 // Required by prometheus-operator
)

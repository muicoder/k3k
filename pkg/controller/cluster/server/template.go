package server

var singleServerTemplate string = `
if [ -d "{{.ETCD_DIR}}" ]; then
	# if directory exists then it means its not an initial run
	/bin/k3s server --cluster-reset --config  {{.INIT_CONFIG}} {{.EXTRA_ARGS}}
fi
rm -f /var/lib/rancher/k3s/server/db/reset-flag
/bin/k3s server --config {{.INIT_CONFIG}} {{.EXTRA_ARGS}}`

var HAServerTemplate string = ` 
if [ ${POD_NAME: -1} == 0 ] && [ ! -d "{{.ETCD_DIR}}" ]; then
	/bin/k3s server --config {{.INIT_CONFIG}} {{.EXTRA_ARGS}}
else 
	SERVER=$(grep server: {{.SERVER_CONFIG}} | awk '{print $NF}')
	until kubectl --username k3k --insecure-skip-tls-verify --server $SERVER version 2>&1 | grep -q credentials; do sleep 1; done
	/bin/k3s server --config {{.SERVER_CONFIG}} {{.EXTRA_ARGS}}
fi`

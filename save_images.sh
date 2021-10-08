rm -rf devops-images.tar.gz

docker save -o devops-images.tar.gz devops:5.1.2.0 devops-ui:5.1.2.0 prom/prometheus:v2.29.2

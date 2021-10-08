chmod +x devops
rm -rf devops.tar.gz
tar -czvf devops.tar.gz devops config.yaml

docker build -t devops:5.1.2.0 .

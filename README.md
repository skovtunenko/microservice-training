# microservice-training
29.sep.2018 
Sigma software, Kharkiv

AWS Free tier: https://aws.amazon.com/free/

----

### launch ec2 instance, but before that download *.pem 
> chmod 400 sigma-aws.pem

----

### To Build application for AWS Linux:
> ./awsBuild.sh

----

### To upload application to EC2:
> scp -i sigma-aws.pem ./uploadDocuments ec2-user@ec2-34-201-167-90.compute-1.amazonaws.com:~

----

### Bastion Host how to:
> ssh -i "sigma-aws.pem" ec2-user@ec2-34-201-167-90.compute-1.amazonaws.com -L 8080:localhost:8080

Where `8080:localhost:8080` will map remote EC2 host to localhost:8080.

----

### Elastic search installation:
> docker run -d -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" --name es  docker.elastic.co/elasticsearch/elasticsearch-oss:6.3.2

! Keep in mind `-d` , `--name es`, `elasticsearch-oss`

### Kibana installation:
> docker run -d --name kibana -p 5601:5601 -e "ELASTICSEARCH_URL=http://es:9200" --link es docker.elastic.co/kibana/kibana-oss:6.3.2

! Keep in mind `-e "ELASTICSEARCH_URL=http://es:9200"` and `--link es` (to point to another docker).

----

## Bastion Host to map localhost to EC2 and use ElasticSearch cluster:
> ssh -i "sigma-aws.pem" ec2-user@ec2-34-201-167-90.compute-1.amazonaws.com -L 8080:localhost:8080 -L 9080:vpc-sigma-aws-gnurpfe3gkw4y7h7pzgcy6ldly.us-east-1.es.amazonaws.com:80

Where `-L 9080:vpc-sigma-aws-gnurpfe3gkw4y7h7pzgcy6ldly.us-east-1.es.amazonaws.com:80` will map remote AWS Elasticsearch host to localhost:9080.

----

## S3:
Bucket name: `sigma-aws-skovtunenko`

----

## Intall LocalStack
> docker run -d -e "SERVICES=s3,sqs" -p 4572:4572 -p 4576:4576 --name localstack localstack/localstack

...and check logs:
> docker logs localstack

Then you can interact with localstack this way:
> aws s3 --endpoint-url=http://localhost:4572 mb s3://sigma.files

where `http://localhost:4572` will point to localstack's S3 service! 

----


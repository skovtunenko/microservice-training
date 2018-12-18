# AWS deployment workshop
19.dec.2018 

AWS Free tier: https://aws.amazon.com/free/

----

### Launch ec2 instance, but before that download *.pem 
> chmod 400 exadel-demo.pem

----

### To Build application for local environment:
> ./localBuild.sh

----

### To Build application for AWS Linux:
> ./awsBuild.sh

----

### To upload application to EC2:
> scp -i exadel-demo.pem ./usersManager-aws ec2-user@ec2-34-201-167-90.compute-1.amazonaws.com:~

----

### Bastion Host how to:
> ssh -i "exadel-demo.pem" ec2-user@ec2-34-201-167-90.compute-1.amazonaws.com -L 8080:localhost:8080

Where `8080:localhost:8080` will map remote EC2 host to localhost:8080.

# aws-demo-webserver
A simple web server in Go to help learn AWS.

To use this, create a new EC2 instance and paste the script below in the 
user data. It will download and start the webserver on the instance.

From there you can visit the webserver or create a load balancer and 
visualize the X-Forwarded-For, and sticky cookies for each request.

```bash
#!/bin/bash
# Change this to the desired release 
RELEASE_TAG=v0.1.1
# Don't edit below
sudo mkdir /usr/local/aws-demo-webserver
cd /usr/local/aws-demo-webserver
sudo wget https://github.com/bmcculley/aws-demo-webserver/releases/download/${RELEASE_TAG}/server-${RELEASE_TAG}-linux-amd64.tar.gz
sudo chown -R ec2-user:ec2-user /usr/local/aws-demo-webserver
tar -xvf server-${RELEASE_TAG}-linux-amd64.tar.gz
sudo rm server-${RELEASE_TAG}-linux-amd64.tar.gz
sudo mv demo-server.service /lib/systemd/system/
sudo chmod 644 /lib/systemd/system/demo-server.service
sudo systemctl daemon-reload
sudo systemctl enable demo-server.service
sudo systemctl start demo-server.service
```

Feel free to open an issue or pull request with any suggestions, updates, etc.
set -ex

# Talk to the metadata server to get the project id and location of application binary.
PROJECTID=$(curl -s "http://metadata.google.internal/computeMetadata/v1/project/project-id" -H "Metadata-Flavor: Google")

# Install logging monitor. The monitor will automatically pickup logs send to
# syslog.
curl -s "https://storage.googleapis.com/signals-agents/logging/google-fluentd-install.sh" | bash
service google-fluentd restart &

# Install dependencies from apt
apt-get update
apt-get install -yq ca-certificates supervisor golang git

# Install the app.
export GOPATH=/root
go get github.com/ijt/tour/gotour

# Configure supervisor to run the Go app.
cat >/etc/supervisor/conf.d/tour.conf << EOF
[program:gotour]
directory=/root/src/github.com/ijt/tour
command=/root/bin/gotour -http=:80
autostart=true
autorestart=true
user=root
environment=HOME="/root",USER="root",GOROOT='/root',GOPATH='/root'
stdout_logfile=syslog
stderr_logfile=syslog
EOF

supervisorctl reread
supervisorctl update

# Application should now be running under supervisor

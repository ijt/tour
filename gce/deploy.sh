# This script is meant to be run from the directory containing it.

instance=$1
if [[ -z "$instance" ]]; then
  echo "usage: $0 instance-name"
  exit 1
fi
gcloud compute instances create ${instance?} \
    --image-family=debian-9 \
    --image-project=debian-cloud \
    --machine-type=g1-small \
    --scopes userinfo-email,cloud-platform \
    --metadata-from-file startup-script=./startup-script.sh \
    --zone us-central1-f \
    --tags http-server

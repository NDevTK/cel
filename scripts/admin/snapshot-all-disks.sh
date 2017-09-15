#!/bin/bash

PROJECT=google.com:chrome-auth-lab
ZONE=us-east1-b

disks=( `gcloud compute disks list --project=$PROJECT --zones=$ZONE --format="value(name)"` )
if [ $? -ne 0 ]; then
  exit 1
fi

snapshot_suffix=-on-$(date +%F)
snapshot_names=$(printf "%s," ${disks[@]/%/$snapshot_suffix})

gcloud compute disks snapshot ${disks[@]} --snapshot-names=$snapshot_names \
  --project=$PROJECT --zone=$ZONE --quiet --verbosity=info

#!/bin/sh
function istioCrds()
{
  while [[ $(kubectl -n istio-system get jobs --no-headers| grep istio-init-crd |grep -v 1/1) ]]; do
  echo "Waiting istio crd jobs to complete"
  sleep 5
  done
}

function istioPods()
{
  while [[ $(kubectl -n istio-system get deployments --no-headers | grep istio | grep -v 1/1) ]]; do
  echo "Waiting for istio pods running"
  sleep 5
  done
}

function knativePods()
{
  while [[ $(kubectl -n knative-serving get deployments --no-headers | grep -v 1/1) ]] || [[ $(kubectl -n knative-eventing get deployments --no-headers | grep eventing | grep -v 1/1) ]]; do
  echo "Waiting for knative pods running"
  sleep 5
  done
}


function defaultBrokerPods()
{
  while [[ $(kubectl -n default get deployments --no-headers -l eventing.knative.dev/broker=default | grep -v 1/1) ]]; do
  echo "Waiting for knative broker in default namespace running"
  sleep 5
  done
}

function postgresqlcrd()
{
  while [[ -z $(kubectl get crd --ignore-not-found=true postgresqls.acid.zalan.do) ]] ; do
  echo "Waiting for postgresql crd"
  sleep 5
  done

}
function vaultcrd()
{
  while [[ -z $(kubectl get crd --ignore-not-found=true vaults.vault.banzaicloud.com) ]] ; do
  echo "Waiting for vault crd"
  sleep 5
  done

}
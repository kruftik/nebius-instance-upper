# nebius-instance-upper :: rolling restart controller for a set of preemptible instances in nebius-cloud

## What is the app intended for

`nebius-instance-upper` continuously checks your preemptible instances in nebius cloud and immediately start them up in case of shutdown by preemptible-instances-controller of the cloud.

Also, in order to prevent unexpected service disruptions, the app performs a rolling instance restart precedure on a daily basis:
* gracefully drains node in a node group through a configurable command (either if it detect that is running on the node to be restarted or alwaysNodeDrain option is set)
* run a configurable pre-stop script and wait for completion
* performs stop-wait-start-wait power cycle that allows the cloud scheduler to migrate the instance to a different server and thus prevents unexpected shutdown by preemtible-instance-controller
* run a configurable post-start script which can do whatever you want, revert the draining for example.
* goes to the next node in the group (if any)

## Configuration 

See [details](./internal/config/config.go) in the source.

## Helm chart

Is also [available](./charts/nebius-instance-upper).
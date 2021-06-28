# Buildkit

Used for build cross-platform container images

### Prereqisites

* Install [ECR Credential Helper][1] and configure `~/.docker/config.json`
* Install [Buildkit][2]

[1]: https://github.com/awslabs/amazon-ecr-credential-helper
[2]: https://github.com/moby/buildkit

```bash
docker-compose up -d
docker-compose logs buildkitd
export BUILDKIT_HOST="tcp://127.0.0.1:1234"
buildctl debug workers
```

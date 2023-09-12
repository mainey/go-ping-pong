## Golang Ping Pong

runs on port 3000, use /ping to get response.

## CI/CD
Consists of 3 stage:
    - Test
    - Build
    - Rollout

* Seperated some common variables as environment variables to ease changes.
* Use dockerx for faster build.
* Rollout deployment after build stage is done instead of setting the image manually (k8s deployment must image pull policy must be always).

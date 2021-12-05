> NOTE: under contruction

# SecureX
Spawning up Decoy Server in case of any fraudulent activity and directing the intruder towards the decoy. Auto Killing the decoy if it is idle for too long.


> Here Decoy is Dockerised Golang API and the attack that is considered is Password Spray Attack

## Links
- Docker Hub Image - https://hub.docker.com/repository/docker/ankithans/securex_app

<!-- <img src="./mockups/workflow.png" /> -->
## Process Diagram
<img src="./mockups/process.png" />

## Tech Stack
Golang, Docker, SMTP, PostgreSQL

## Tasks
1. ~~Create a golang API for login~~
2. ~~System to detect the port scan attack or password spray attack~~
3. ~~Docker container with the same Golang API~~
4. ~~Spawn the Docker container programmatically on a specific port~~
5. ~~Redirect the intruder to server created by docker container~~
6. ~~Shut down container if it is idle~~
7. ~~Audit Logging service in the container API~~
8. Notification service in the container API
9. Improve dataset for logins
---
> Made with love by [Ankit Hans](https://www.github.com/ankithans)
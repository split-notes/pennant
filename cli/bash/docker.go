package bash

import "fmt"

var (
	DockerListContainersByName = "docker ps --format '{{.Names}}'"
)

func DockerListShellsOfContainer(container string) string { return fmt.Sprintf("docker exec -it %s -- cat /etc/shells", container) }
func DockerSSH(container string, shell string) string { return fmt.Sprintf("docker exec -it %s %s", container, shell)}
func DockerStart(projectLocation string) string { return fmt.Sprintf("docker-compose -f %s/docker-compose.yml up", projectLocation) }
func DockerStop(projectLocation string) string { return fmt.Sprintf("docker-compose -f %s/docker-compose.yml down", projectLocation) }
func DockerWreck() string { return "docker rm $(docker ps -a -q)" }

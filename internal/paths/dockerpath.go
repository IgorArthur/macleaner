package paths

var DockerPaths = map[string][]string{
	"darwin":  MacOSDockerPaths,
	"linux":   LinuxDockerPaths,
	"windows": WindowsDockerPaths,
}

var MacOSDockerPaths = []string{
	"~/Library/Containers/com.docker.docker",
	"~/Library/Group Containers/group.com.docker",
	"~/Library/Application Support/Docker Desktop",
	"~/Library/Logs/Docker Desktop",
	"~/Library/Caches/com.docker.docker",
}

var LinuxDockerPaths = []string{
	"/var/lib/docker",
	"/etc/docker",
	"~/.docker",
	"/var/run/docker.sock",
	"~/.local/share/docker",
}

var WindowsDockerPaths = []string{
	`%USERPROFILE%\.docker`,
	`%LOCALAPPDATA%\Docker`,
	`%PROGRAMDATA%\Docker`,
	`%PROGRAMDATA%\DockerDesktop`,
}

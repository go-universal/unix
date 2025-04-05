package nginx

import (
	"fmt"
	"os"
	"os/exec"
)

const reverseTemplate = `server {
	listen 80;
	listen [::]:80;
	server_name {domains};

	location / {
		client_max_body_size 1M;
		proxy_pass http://localhost:{port};
		proxy_http_version 1.1;
		proxy_set_header Upgrade $http_upgrade;
		proxy_set_header Connection 'upgrade';
		proxy_set_header Host $host;
		proxy_set_header Referer $http_referer;
		proxy_set_header X-Forwarded-Proto $scheme;
		proxy_set_header X-Forwarded-For $remote_addr;
		proxy_set_header X-Forwarded-Referer $http_referer;
		proxy_cache_bypass $http_upgrade;
	}
}`

// cmdError handles execution errors, including extracting the exit code and stderr output.
func cmdError(err error) error {
	if err == nil {
		return nil
	}

	if exitErr, ok := err.(*exec.ExitError); ok {
		return fmt.Errorf("exit %d, %s", exitErr.ExitCode(), string(exitErr.Stderr))
	}

	return err
}

// fileExists check if file exists.
func fileExists(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// restart restarts the nginx service.
func restart() error {
	return cmdError(exec.Command("sudo", "systemctl", "restart", "nginx").Run())
}

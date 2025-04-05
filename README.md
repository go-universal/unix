# Unix

`unix` is a Go library for managing cron jobs, Nginx server blocks, systemd services, and system information. It provides a simple and consistent API for creating, installing, and managing these services.

## Installation

To install the library, use the following command:

```sh
go get github.com/go-universal/unix
```

## Features

- **Cron Jobs**: Manage cron schedules with timezone support.
- **Nginx Server Blocks**: Create and manage Nginx server configurations.
- **Systemd Services**: Manage systemd services for Linux systems.
- **System Information**: Retrieve system information such as CPU, memory, disk, and network stats.
- **Template Engine**: Use a lightweight template engine for string replacements.

### Cron Jobs

The `Cron` interface provides methods for scheduling and managing cron jobs.

- `Raw() string`: Returns the raw cron expression.
- `Exists() (bool, error)`: Checks whether the cron job is installed.
- `Install() (bool, error)`: Installs the cron job. Returns `false` if it already exists.
- `Uninstall() error`: Removes the cron job.

#### Options

- `WithTimezone(tz *CronTZ) Option` sets the timezone for the cron schedule.
- `RunAtReboot() Option`: Schedules the cron to run at system reboot.
- `RunAtReboot() Option`: Schedules the cron to run at system reboot.
- `RunYearly() Option`: Schedules the cron to run once a year (January 1st at midnight).
- `RunMonthly() Option`: Schedules the cron to run once a month (1st day at midnight).
- `RunWeekly(wd Weekday) Option`: Schedules the cron to run once a week on the specified weekday.
- `RunDaily() Option`: Schedules the cron to run once a day at midnight.
- `EveryXMinutes(minutes int) Option`: Runs the cron every X minutes.
- `EveryXHours(hours int) Option`: Runs the cron every X hours.
- `Minute(minute int) Option`: Sets the specific minute for the cron schedule.
- `Hour(hour int) Option`: Sets the specific hour for the cron schedule.
- `DayOfMonth(day int) Option`: Sets the specific day of the month for the cron schedule.
- `Month(month int) Option`: Sets the specific month for the cron schedule.
- `DayOfWeek(wd Weekday) Option`: Sets the specific day of the week for the cron schedule.

```go
package main

import (
    "fmt"
    "github.com/go-universal/unix/cron"
)

func main() {
    job := cron.New("echo 'Hello, World!'", cron.RunDaily(), cron.Hour(2), cron.Minute(30))

    if installed, err := job.Install(); err != nil {
        fmt.Println("Error installing cron job:", err)
    } else if installed {
        fmt.Println("Cron job installed successfully")
    } else {
        fmt.Println("Cron job already exists")
    }
}
```

### Nginx Server Blocks

The `ServerBlock` and `ReverseProxy` interfaces provide methods for managing Nginx server configurations.

- `Exists() (bool, error)`: Checks if the configuration file exists.
- `Enabled() (bool, error)`: Checks if the configuration is enabled.
- `Disable() error`: Disables the configuration.
- `Enable() error`: Enables the configuration.
- `Install(override bool) (bool, error)`: Installs the configuration. Returns `false` if it already exists and `override` is `false`.
- `Uninstall() error`: Removes the configuration.

```go
package main

import (
    "fmt"
    "github.com/go-universal/unix/nginx"
)

func main() {
    proxy := nginx.NewReverseProxy("example", "8080", []string{"example.com", "www.example.com"})

    if installed, err := proxy.Install(true); err != nil {
        fmt.Println("Error installing reverse proxy:", err)
    } else if installed {
        fmt.Println("Reverse proxy installed successfully")
    } else {
        fmt.Println("Reverse proxy already exists")
    }
}
```

### Systemd Services

The `SystemdService` interface provides methods for managing systemd services.

- `Exists() bool`: Checks if the service exists.
- `Enabled() bool`: Checks if the service is enabled.
- `Disable() error`: Disables the service.
- `Install(override bool) (bool, error)`: Installs the service. Returns `false` if it already exists and `override` is `false`.
- `Uninstall() error`: Removes the service.

```go
package main

import (
    "fmt"
    "github.com/go-universal/unix/systemd"
)

func main() {
    service := systemd.NewService("example-service", "/path/to/root", "example-command")

    if installed, err := service.Install(true); err != nil {
        fmt.Println("Error installing systemd service:", err)
    } else if installed {
        fmt.Println("Systemd service installed successfully")
    } else {
        fmt.Println("Systemd service already exists")
    }
}
```

### System Information

The `sysinfo` package provides methods for retrieving system information.

- `CPUInfo() (cores int, used, free float64, err error)`: Retrieves CPU information.
- `MemoryInfo() (total, used, free uint64, err error)`: Retrieves memory statistics.
- `DiskInfo() (total, used, free uint64, err error)`: Retrieves disk usage statistics.
- `NetworkInfo() (sent, recv uint64, err error)`: Retrieves network I/O statistics.
- `Uptime() (time.Duration, error)`: Retrieves system uptime.
- `UptimeParts() (days, hours, minutes uint64, err error)`: Retrieves uptime as days, hours, and minutes.
- `UptimeI18n(dayTitle, hourTitle, minuteTitle, separator string) (string, error)`: Retrieves localized uptime.

```go
package main

import (
    "fmt"
    "github.com/go-universal/unix/sysinfo"
)

func main() {
    total, used, free, err := sysinfo.MemoryInfo()
    if err != nil {
        fmt.Println("Error retrieving memory info:", err)
        return
    }

    fmt.Printf("Total: %d, Used: %d, Free: %d\n", total, used, free)
}
```

### Utility Functions

#### `IsSudo`

Checks if the program is running with sudo privileges.

```go
package main

import (
    "fmt"
    "github.com/go-universal/unix"
)

func main() {
    if unix.IsSudo() {
        fmt.Println("Running with sudo privileges")
    } else {
        fmt.Println("Not running with sudo privileges")
    }
}
```

---

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.

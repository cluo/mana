PACKAGE

package info
    import "mana/info"

    简单的获取cpu利用率，系统内存使用情况，磁盘读写情况. 查看cpu温度命令sensors输出内容，判断是否高于（70）等. 获取硬盘温度.
    使用package net检查tcp/udp等服务是否在线. 使用shell脚本检查特定进程。

FUNCTIONS

func GetAdapters() ([]Adapter, error)

func GetHddtemps() (temps []Hddtemp, err error)

func GetPcpus() ([]Pcpu, error)

func NewLogger(name string) *log.Logger


TYPES

type Adapter struct {
    Name     string
    Receive  int64
    Transmit int64
    Time     time.Time
}

func (a Adapter) String() string

type Agent struct{}

func (a *Agent) Process(name string) (*Process, error)

func (a *Agent) Shell(name, path string) (*Shell, error)

func (a *Agent) System() (*Server, error)

func (a *Agent) Tcp(name, addr, port string) (*Tcp, error)

func (a *Agent) Top(n string, sort string) (*TopProcess, error)

func (a *Agent) Udp(name, addr, port string) (*Udp, error)
    Udp only check adr='127.0.0.1'

type ByteSize float64

const (
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
)

func (b ByteSize) String() string

type Free struct {
    Mem  Mem
    Swap Swap
}
    free -o

func GetFree() (*Free, error)

func (m *Free) Format() *Free
    格式化为易读的数据

func (m *Free) Real() float32
    未使用的内存加上缓存

func (f *Free) String() string

type Hddtemp struct {
    Dev  string
    Desc string
    Temp string
}

func (t Hddtemp) String() string

type Hostname struct {
    Name   string
    Boot   time.Time
    Uptime string
}

func GetHostname() (*Hostname, error)
    服务器主机名、启动时间以及运行时间:/proc/uptime

func (h *Hostname) String() string

type Iostat string

func GetIostat() (Iostat, error)

type Load struct {
    Cpu  []Pcpu
    Free *Free
    Load *Loadavg
    IO   Iostat
}

func GetLoad() (*Load, error)

func (l *Load) String() string

type Loadavg struct {
    La1, La5, La15 string
    Processes      string
}

func GetLoadavg() (*Loadavg, error)
    系统负载情况

func (la *Loadavg) Overload() bool

func (l *Loadavg) String() string

type Mem struct {
    Total   string
    Used    string
    Free    string
    Buffers string
    Cached  string
}
    物理内存

type Pcpu struct {
    Us float64
    Sy float64
    Wa float64
    Id float64
}
    CPU使用率

func (p Pcpu) String() string

type Process struct {
    Name string
    Pid  string
}
    自定义的进程检查,

func (p *Process) String() string

type Sensor string

func GetSensor() (Sensor, error)

type Server struct {
    Hostname *Hostname
    Load     *Load
    Temp     *Temp
}

func (s *Server) String() string

type Shell struct {
    Name, Path string
    Result     string
}

func (sh *Shell) String() string

type Swap struct {
    Total string
    Used  string
    Free  string
}
    交换分区

type Tcp struct {
    Name    string
    Address string
    Status  bool
}
    tcp/udp 系统服务

func (t *Tcp) String() string

type Temp struct {
    Disks   []Hddtemp
    Sensors Sensor
}

func GetTemp() (*Temp, error)

func (temp *Temp) String() string

type TopProcess struct {
    Sort   string
    Num    string
    Result string
}
    CPU使用率最高的进程

func (top *TopProcess) String() string

type Udp struct {
    Name    string
    Address string
    Status  bool
}

func (u *Udp) String() string


SUBDIRECTORIES

	bin


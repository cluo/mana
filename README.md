PACKAGE

package info
    import "mana/info"

    sensors命令:"lm-sensors" mpstat, iostat命令:"sysstat"

FUNCTIONS

func GetAdapters() ([]Adapter, error)
    读取/proc/net/dev

func GetHddtemps() (temps []Hddtemp, err error)
    需要hddtemp以守护进程运行，如："sudo hddtemp -d /dev/sd[a-d]"

func GetPcpus() ([]Pcpu, error)
    cpu使用率，多cpu的情况，第一个是平均值

func NewLogger(name string) *log.Logger


TYPES

type Adapter struct {
    Name     string
    Receive  float64
    Transmit float64
    Time     time.Time
}
    上传和下载的流量，从系统启动累加

func (a Adapter) String() string

type Agent struct{}

func (a *Agent) Process(name string) (*Process, error)
    名称=name的脚本放置在bin目录下

func (a *Agent) Shell(name, path string) (*Shell, error)
    name如果为""，使用filepath.Base(path)

func (a *Agent) System() (*Server, error)

func (a *Agent) Tcp(name, addr, port string) (*Tcp, error)
    name服务名称，addr、port ip地址和端口

func (a *Agent) Top(n string, sort string) (*TopProcess, error)
    n 是一个整数，脚本"bin/mem_top","bin/cpu_top"中取值为5-10 sort只能是"cpu"或者"mem"

func (a *Agent) Udp(name, addr, port string) (*Udp, error)
    Udp 只用于检查本机地址,如"127.0.0.1"、"192.168.1.10"等，port端口

type ByteSize float64
    ByteSize格式化内存或者流量数据为易读的格式

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
    使用"free -o -b"命令

func (m *Free) Format() *Free
    内存信息转换

func (m *Free) Real() float32
    未使用的内存加上缓存

func (f *Free) String() string

type Hddtemp struct {
    Dev  string
    Desc string
    Temp string
}
    硬盘温度

func (t Hddtemp) String() string

type Hostname struct {
    Name   string
    Boot   time.Time
    Uptime string
}
    服务器主机名、启动时间以及运行时间

func GetHostname() (*Hostname, error)
    读取/proc/uptime，第一数值即系统运行时间，单位为秒(s)

func (h *Hostname) String() string

type Iostat string
    获取系统的IO状态

func GetIostat() (Iostat, error)
    需要命令"iostat"

type Load struct {
    Cpu  []Pcpu
    Free *Free
    Load *Loadavg
    IO   Iostat
}
    Load struct 包含了cpu、内存、系统负载、以及io状态

func GetLoad() (*Load, error)

func (l *Load) String() string

type Loadavg struct {
    La1, La5, La15 string
    Processes      string
}
    系统负载情况

func GetLoadavg() (*Loadavg, error)
    读取/proc/loadavg

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
    自定义的进程检查脚本

func (p *Process) String() string

type Sensor string
    cpu温度

func GetSensor() (Sensor, error)
    需要命令sensors

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
    方便以后使用shell脚本获取特定信息

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
    tcp 服务

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
    CPU或者MEM使用率最高的进程

func (top *TopProcess) String() string

type Udp struct {
    Name    string
    Address string
    Status  bool
}
    udp服务

func (u *Udp) String() string


SUBDIRECTORIES


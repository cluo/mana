<!--
	Copyright 2009 The Go Authors. All rights reserved.
	Use of this source code is governed by a BSD-style
	license that can be found in the LICENSE file.
-->

	
		<div id="short-nav">
			<dl>
			<dd><code>import "mana/jk"</code></dd>
			</dl>
			<dl>
			<dd><a href="#overview" class="overviewLink">Overview</a></dd>
			<dd><a href="#index">Index</a></dd>
			
			
				<dd><a href="#subdirectories">Subdirectories</a></dd>
			
			</dl>
		</div>
		<!-- The package's Name is printed as title by the top-level template -->
		<div id="overview" class="toggleVisible">
			<div class="collapsed">
				<h2 class="toggleButton" title="Click to show Overview section">Overview ▹</h2>
			</div>
			<div class="expanded">
				<h2 class="toggleButton" title="Click to hide Overview section">Overview ▾</h2>
				
			</div>
		</div>
		
	
		<h2 id="index">Index</h2>
		<!-- Table of contents for API; must be named manual-nav to turn off auto nav. -->
		<div id="manual-nav">
			<dl>
			
			
			
				
				<dd><a href="#GetHddtemps">func GetHddtemps() (temps []Hddtemp, err error)</a></dd>
			
				
				<dd><a href="#SetDuration">func SetDuration(t Timer, out chan&lt;- Timer)</a></dd>
			
				
				<dd><a href="#Topcpu">func Topcpu() (string, error)</a></dd>
			
				
				<dd><a href="#Topmem">func Topmem() (string, error)</a></dd>
			
			
				
				<dd><a href="#ByteSize">type ByteSize</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#ByteSize.String">func (b ByteSize) String() string</a></dd>
				
			
				
				<dd><a href="#Hddtemp">type Hddtemp</a></dd>
				
				
			
				
				<dd><a href="#Iostat">type Iostat</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Getiostat">func Getiostat() (Iostat, error)</a></dd>
				
				
			
				
				<dd><a href="#Loadavg">type Loadavg</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Getloadavg">func Getloadavg() (*Loadavg, error)</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#Loadavg.Check">func (la Loadavg) Check() bool</a></dd>
				
			
				
				<dd><a href="#Log">type Log</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#Log.Init">func (l *Log) Init() (e, i *log.Logger)</a></dd>
				
			
				
				<dd><a href="#Memory">type Memory</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Getmemory">func Getmemory() (memory *Memory, err error)</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#Memory.Format">func (m *Memory) Format() *Memory</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Memory.Realfree">func (m *Memory) Realfree() float32</a></dd>
				
			
				
				<dd><a href="#Pcpu">type Pcpu</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Getpcpu">func Getpcpu() (*Pcpu, error)</a></dd>
				
				
			
				
				<dd><a href="#Process">type Process</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Getpid">func Getpid(name string) (p *Process, err error)</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#Process.String">func (p Process) String() string</a></dd>
				
			
				
				<dd><a href="#Realm">type Realm</a></dd>
				
				
			
				
				<dd><a href="#Sensors">type Sensors</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Getsensors">func Getsensors() (Sensors, error)</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#Sensors.Check">func (s Sensors) Check() (bool, error)</a></dd>
				
			
				
				<dd><a href="#Server">type Server</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#Server.Check">func (t *Server) Check()</a></dd>
				
			
				
				<dd><a href="#Swapd">type Swapd</a></dd>
				
				
			
				
				<dd><a href="#Timer">type Timer</a></dd>
				
				
			
				
				<dd><a href="#Uptime">type Uptime</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#GetUptime">func GetUptime() (u *Uptime, err error)</a></dd>
				
				
			
			
		</dl>

		

		
			<h4>Package files</h4>
			<p>
			<span style="font-size:90%">
			
				<a href="/target/host.go">host.go</a>
			
				<a href="/target/load.go">load.go</a>
			
				<a href="/target/log.go">log.go</a>
			
				<a href="/target/process.go">process.go</a>
			
				<a href="/target/service.go">service.go</a>
			
				<a href="/target/temp.go">temp.go</a>
			
				<a href="/target/time.go">time.go</a>
			
			</span>
			</p>
		
	
		
		
		
			
			
			<h2 id="GetHddtemps">func <a href="/target/temp.go?s=220:267#L11">GetHddtemps</a></h2>
			<pre>func GetHddtemps() (temps []Hddtemp, err error)</pre>
			
			
		
			
			
			<h2 id="SetDuration">func <a href="/target/time.go?s=133:176#L4">SetDuration</a></h2>
			<pre>func SetDuration(t Timer, out chan&lt;- Timer)</pre>
			
			
		
			
			
			<h2 id="Topcpu">func <a href="/target/process.go?s=386:415#L19">Topcpu</a></h2>
			<pre>func Topcpu() (string, error)</pre>
			
			
		
			
			
			<h2 id="Topmem">func <a href="/target/process.go?s=531:560#L25">Topmem</a></h2>
			<pre>func Topmem() (string, error)</pre>
			
			
		
		
			
			
			<h2 id="ByteSize">type <a href="/target/load.go?s=1512:1533#L76">ByteSize</a></h2>
			<pre>type ByteSize float64</pre>
			

			
				<pre>const (
    KB ByteSize = 1 &lt;&lt; (10 * iota)
    MB
    GB
    TB
)</pre>
				
			

			

			

			

			
				
				<h3 id="ByteSize.String">func (ByteSize) <a href="/target/load.go?s=1610:1643#L86">String</a></h3>
				<pre>func (b ByteSize) String() string</pre>
				
				
				
			
		
			
			
			<h2 id="Hddtemp">type <a href="/target/temp.go?s=72:134#L1">Hddtemp</a></h2>
			<pre>type Hddtemp struct {
    Dev  string
    Desc string
    Temp string
}</pre>
			

			

			

			

			

			
		
			
			
			<h2 id="Iostat">type <a href="/target/load.go?s=712:730#L33">Iostat</a></h2>
			<pre>type Iostat string</pre>
			

			

			

			

			
				
				<h3 id="Getiostat">func <a href="/target/load.go?s=732:764#L35">Getiostat</a></h3>
				<pre>func Getiostat() (Iostat, error)</pre>
				
				
			

			
		
			
			
			<h2 id="Loadavg">type <a href="/target/load.go?s=915:984#L45">Loadavg</a></h2>
			<pre>type Loadavg struct {
    La1, La5, La15 string
    Processes      string
}</pre>
			

			

			

			

			
				
				<h3 id="Getloadavg">func <a href="/target/load.go?s=986:1021#L50">Getloadavg</a></h3>
				<pre>func Getloadavg() (*Loadavg, error)</pre>
				
				
			

			
				
				<h3 id="Loadavg.Check">func (Loadavg) <a href="/target/load.go?s=1230:1260#L61">Check</a></h3>
				<pre>func (la Loadavg) Check() bool</pre>
				
				
				
			
		
			
			
			<h2 id="Log">type <a href="/target/log.go?s=43:95#L1">Log</a></h2>
			<pre>type Log struct {
    Error string
    Info  string
}</pre>
			

			

			

			

			

			
				
				<h3 id="Log.Init">func (*Log) <a href="/target/log.go?s=97:136#L3">Init</a></h3>
				<pre>func (l *Log) Init() (e, i *log.Logger)</pre>
				
				
				
			
		
			
			
			<h2 id="Memory">type <a href="/target/load.go?s=2096:2142#L117">Memory</a></h2>
			<pre>type Memory struct {
    Mem  Realm
    Swap Swapd
}</pre>
			<p>
free -o
</p>


			

			

			

			
				
				<h3 id="Getmemory">func <a href="/target/load.go?s=2144:2188#L122">Getmemory</a></h3>
				<pre>func Getmemory() (memory *Memory, err error)</pre>
				
				
			

			
				
				<h3 id="Memory.Format">func (*Memory) <a href="/target/load.go?s=2855:2888#L152">Format</a></h3>
				<pre>func (m *Memory) Format() *Memory</pre>
				
				
				
			
				
				<h3 id="Memory.Realfree">func (*Memory) <a href="/target/load.go?s=2515:2550#L136">Realfree</a></h3>
				<pre>func (m *Memory) Realfree() float32</pre>
				
				
				
			
		
			
			
			<h2 id="Pcpu">type <a href="/target/load.go?s=88:144#L2">Pcpu</a></h2>
			<pre>type Pcpu struct {
    Us float64
    Sy float64
    Id float64
}</pre>
			

			

			

			

			
				
				<h3 id="Getpcpu">func <a href="/target/load.go?s=146:175#L8">Getpcpu</a></h3>
				<pre>func Getpcpu() (*Pcpu, error)</pre>
				
				
			

			
		
			
			
			<h2 id="Process">type <a href="/target/process.go?s=35:84#L1">Process</a></h2>
			<pre>type Process struct {
    Name string
    Pid  string
}</pre>
			

			

			

			

			
				
				<h3 id="Getpid">func <a href="/target/process.go?s=162:210#L7">Getpid</a></h3>
				<pre>func Getpid(name string) (p *Process, err error)</pre>
				
				
			

			
				
				<h3 id="Process.String">func (Process) <a href="/target/process.go?s=86:118#L2">String</a></h3>
				<pre>func (p Process) String() string</pre>
				
				
				
			
		
			
			
			<h2 id="Realm">type <a href="/target/load.go?s=1909:2010#L101">Realm</a></h2>
			<pre>type Realm struct {
    Total   string
    Used    string
    Free    string
    Buffers string
    Cached  string
}</pre>
			<p>
mem
</p>


			

			

			

			

			
		
			
			
			<h2 id="Sensors">type <a href="/target/temp.go?s=706:725#L32">Sensors</a></h2>
			<pre>type Sensors string</pre>
			

			

			

			

			
				
				<h3 id="Getsensors">func <a href="/target/temp.go?s=727:761#L34">Getsensors</a></h3>
				<pre>func Getsensors() (Sensors, error)</pre>
				
				
			

			
				
				<h3 id="Sensors.Check">func (Sensors) <a href="/target/temp.go?s=839:877#L39">Check</a></h3>
				<pre>func (s Sensors) Check() (bool, error)</pre>
				
				
				
			
		
			
			
			<h2 id="Server">type <a href="/target/service.go?s=50:181#L1">Server</a></h2>
			<pre>type Server struct {
    Name string
    <span class="comment">// tcp/udp/unix</span>
    Net string
    <span class="comment">// 127.0.0.1:80 /var/run/example.sock</span>
    Addr   string
    Status bool
}</pre>
			<p>
tcp/udp
</p>


			

			

			

			

			
				
				<h3 id="Server.Check">func (*Server) <a href="/target/service.go?s=215:239#L10">Check</a></h3>
				<pre>func (t *Server) Check()</pre>
				
				
				
			
		
			
			
			<h2 id="Swapd">type <a href="/target/load.go?s=2020:2083#L110">Swapd</a></h2>
			<pre>type Swapd struct {
    Total string
    Used  string
    Free  string
}</pre>
			<p>
swap
</p>


			

			

			

			

			
		
			
			
			<h2 id="Timer">type <a href="/target/time.go?s=32:131#L1">Timer</a></h2>
			<pre>type Timer interface {
    Stat() []byte
    String() string
    Interval() time.Duration
    Last(time.Time)
}</pre>
			

			

			

			

			

			
		
			
			
			<h2 id="Uptime">type <a href="/target/host.go?s=62:138#L1">Uptime</a></h2>
			<pre>type Uptime struct {
    Hostname string
    Boot     time.Time
    Duration string
}</pre>
			

			

			

			

			
				
				<h3 id="GetUptime">func <a href="/target/host.go?s=219:258#L10">GetUptime</a></h3>
				<pre>func GetUptime() (u *Uptime, err error)</pre>
				
				
			

			
		
		</div>
	

	







	
	
		<h2 id="subdirectories">Subdirectories</h2>
	
	<table class="dir">
	<tr>
	<th>Name</th>
	<th>&nbsp;&nbsp;&nbsp;&nbsp;</th>
	<th style="text-align: left; width: auto">Synopsis</th>
	</tr>
	
		<tr>
		<td><a href="..">..</a></td>
		</tr>
	
	
		
			<tr>
			<td class="name"><a href="tmp">tmp</a></td>
			<td>&nbsp;&nbsp;&nbsp;&nbsp;</td>
			<td style="width: auto"></td>
			</tr>
		
	
	</table>
	


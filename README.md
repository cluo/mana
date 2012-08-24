# -------- mana_group ----------------------------
# -h
# -p
# -mail

# -------- mana_agent (use -i 参数) data ---------
# /status
    主机名: debian
    启动时间: 2012-08-13 17:12:15 +0800 CST
    已运行: 23m47.15s
    ---------------------------------------------
    系统负载: 0.02 0.03 0.00	2/202
# /system
    {
      "Hostname": {
        "Name": "debian",
        "Boot": "2012-08-13T17:12:15.935594+08:00",
        "Uptime": "25m0.35s"
      },
      "Load": {
        "Cpu": [
          {
            "ID": "all",
            "Us": 0.7,
            "Sy": 0.42,
            "Wa": 1.81,
            "Idle": 96.96
          },
          {
            "ID": "0",
            "Us": 0.44,
            "Sy": 0.55,
            "Wa": 1.9,
            "Idle": 96.88
          },
          {
            "ID": "1",
            "Us": 0.96,
            "Sy": 0.29,
            "Wa": 1.71,
            "Idle": 97.04
          }
        ],
        "Free": {
          "Mem": {
            "Total": "526991360",
            "Used": "462036992",
            "Free": "64954368",
            "Buffers": "10502144",
            "Cached": "190382080"
          },
          "Swap": {
            "Total": "526376960",
            "Used": "0",
            "Free": "526376960"
          }
        },
        "Loadavg": {
          "La1": "0.01",
          "La5": "0.02",
          "La15": "0.00",
          "Processes": "1/202"
        },
        "IO": "Linux 2.6.32-5-686 (debian) \u00092012年08月13日 \u0009_i686_\u0009(2 CPU)\nDevice:         rrqm/s   wrqm/s     r/s     w/s    rkB/s    wkB/s avgrq-sz avgqu-sz   await  svctm  %util\nsda               0.58     4.13    4.85    0.78   128.41    19.66    52.54     0.18   32.11   5.51   3.10\n"
      },
      "Traffic": [
        {
          "Name": "eth0",
          "Receive": 24064,
          "Transmit": 186,
          "Time": "2012-08-13T17:37:16.294606+08:00"
        },
        {
          "Name": "eth1",
          "Receive": 0,
          "Transmit": 6,
          "Time": "2012-08-13T17:37:16.294606+08:00"
        }
      ],
      "Temp": {
        "Disks": [
          {
            "Dev": "/dev/sda",
            "Desc": "VBOX HARDDISK",
            "Temp": "UNK"
          }
        ],
        "Sensors": "No sensors found"
      }
    }
# /service?q=tcp
    [
      {
        "Name": "hddtemp",
        "Net": "tcp",
        "Address": "127.0.0.1:7634",
        "Status": true
      },
      {
        "Name": "https",
        "Net": "tcp",
        "Address": "192.168.56.101:443",
        "Status": true
      },
      {
        "Name": "http",
        "Net": "tcp",
        "Address": "127.0.0.1:80",
        "Status": true
      },
      {
        "Name": "ssh",
        "Net": "tcp",
        "Address": "192.168.56.101:22",
        "Status": true
      },
      {
        "Name": "godoc",
        "Net": "tcp",
        "Address": "127.0.0.1:6060",
        "Status": false
      },
      {
        "Name": "amavisd",
        "Net": "tcp",
        "Address": "127.0.0.1:10024",
        "Status": true
      }
    ]
# /service?q=tcp&name=godoc
    {
      "Name": "godoc",
      "Net": "tcp",
      "Address": "127.0.0.1:6060",
      "Status": false
    }
# /service?q=udp
    [
      {
        "Name": "portmap",
        "Net": "udp",
        "Address": "127.0.0.1:111",
        "Status": true
      }
    ]
# /process
    [
      {
        "Name": "nginx_master",
        "Pid": "1476"
      },
      {
        "Name": "clamd_unix",
        "Pid": "1590"
      }
    ]
# /process?q=clamd_unix
    {
          "Name": "clamd_unix",
            "Pid": "1590"
    }
# /custom
    [
      {
        "Name": "netstat",
        "Result": "tcp 127.0.0.1:10024\ntcp 0.0.0.0:111\ntcp 0.0.0.0:80\ntcp 127.0.0.1:7634\ntcp 192.168.56.101:22\ntcp 0.0.0.0:48154\ntcp 192.168.56.101:443\ntcp6 :::12345\nudp 127.0.0.1:161\nudp 0.0.0.0:68\nudp 0.0.0.0:57041\nudp 0.0.0.0:39763\nudp 0.0.0.0:111\nudp 0.0.0.0:626\n"
      },
      {
        "Name": "myip",
        "Result": "123.233.148.82"
      }
    ]
# /custom?q=myip
    {
          "Name": "myip",
            "Result": "123.233.148.82"
    }


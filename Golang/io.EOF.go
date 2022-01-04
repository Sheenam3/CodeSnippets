# To check end of file error and break out of the loop

for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				logger.Log.Error(err)
			}
		}
		parsedLine := strings.Fields(string(line))
		if parsedLine[0] != "TIME(s)" {
			ppid, err := strconv.ParseInt(parsedLine[3], 10, 64)
			if err != nil {
				logger.Log.Error("TCPConnect PID Error")
			}

			timest := 0.00
			n := Log{Fulllog: string(line), Pid: ppid, Time: timest, Probe: tool}
			logtcpconnect <- n
		}
	}

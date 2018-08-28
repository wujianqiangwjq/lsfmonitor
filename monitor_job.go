package main

import (
	"errors"
	"fmt"
	"log"
	"lsfmonitor/common"
	"lsfmonitor/job"
	"regexp"
	"strconv"
	"strings"
	"time"

	"runtime"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	grpc_port    = "127.0.0.1:5354"
	monitor_time = 10
)

func Monitor_jobs(jobclient job.JobClient) {
	for {
		out := common.ExcuteWithoutpw("bjobs", "-a", "-u", "all")
		alljobsin := strings.Split(out, "\n")

		for _, value := range alljobsin[1:] {
			itemone := strings.Fields(value)
			e, _ := GetData(itemone[0])
			fmt.Println(e)
			go Backrun(jobclient, e)
		}
		time.Sleep(monitor_time * time.Second)
	}
}
func Backrun(jobclient job.JobClient, data *job.Requestjob) {
	res, _ := jobclient.UpdateAndCreateJob(context.Background(), data)
	if res.Status == 1 {
		jobclient.UpdateAndCreateJob(context.Background(), data)
	}
}
func GetData(jobid string) (*job.Requestjob, error) {
	out := common.ExcuteWithoutpw("bjobs", "-UF", jobid)
	if len(out) <= 0 {
		log.Printf("data is empty for jobid:%d", jobid)
		return &job.Requestjob{}, errors.New("job data donot exist")
	} else {
		return RegxData(out, jobid)
	}
}

func TimeConvert(date string) int64 {
	formate := "Mon Jan _2 15:04:05 2006"
	year := strconv.Itoa(time.Now().Year())
	t, err := time.ParseInLocation(formate, date+" "+year, time.Local)
	if err != nil {
		return 0
	} else {
		return t.Unix()
	}
}

func Split(data string) []string {
	str := strings.Replace(data, ">;", ">,", -1)
	return strings.Split(str, ">,")
}

func GetHosts(data string) (hosts []string, cpus []int64) {
	hdata := strings.Split(data, "> <")
	nodes := strings.Split(hdata[0], "<")[1]
	item := strings.Split(nodes, "*")
	slen := len(item)
	if slen == 2 {
		hosts = append(hosts, item[1])
		cpu, _ := strconv.ParseInt(item[0], 0, 64)
		cpus = append(cpus, cpu)
	} else if slen == 1 {
		hosts = append(hosts, item[0])
		cpus = append(cpus, 1)
	}
	for _, val := range hdata[1:] {
		item = strings.Split(val, "*")
		slen = len(item)
		if slen == 2 {
			hosts = append(hosts, item[1])
			cpu, _ := strconv.ParseInt(item[0], 0, 64)
			cpus = append(cpus, cpu)
		} else if slen == 1 {
			hosts = append(hosts, item[0])
			cpus = append(cpus, 1)
		}
	}
	return
}
func RegxData(out string, jobid string) (*job.Requestjob, error) {
	var onejob job.Requestjob
	jobidint, _ := strconv.ParseInt(jobid, 0, 64)
	onejob.Jobid = jobidint
	jobstart := regexp.MustCompile(`^Job <`)
	start := regexp.MustCompile(`Started (\d+) Task\(s\)`)
	gpu := regexp.MustCompile(`\[ngpus=(\d+)`)
	splitsdata := strings.Split(out, "\n")
	gpudata := int64(0)
	for index, value := range splitsdata {
		if len(jobstart.FindAllStringSubmatch(value, -1)) > 0 {
			datas := strings.Split(value, ">, ")
			for _, item := range datas {
				slitem := strings.Split(item, " <")
				if len(slitem) != 2 {
					continue
				}

				key := strings.Trim(slitem[0], " ")
				value := strings.Trim(slitem[1], " ")
				switch key {
				case "Job Name":
					onejob.Jobname = value
				case "User":
					onejob.Submiter = value
				case "Status":
					onejob.Status = value

				}
			}
			for _, value2 := range splitsdata[index:] {

				if strings.Contains(value2, "Submitted from host") {
					qtimedata := Split(value2)
					qtimes := strings.Split(qtimedata[0], ": ")[0]
					onejob.Qtime = TimeConvert(qtimes)

				}
				if start.MatchString(value2) {
					startdata := Split(value2)
					starttime := strings.Split(startdata[0], ": ")[0]
					onejob.Starttime = TimeConvert(starttime)
					onejob.Nodes, onejob.Cpus = GetHosts(startdata[0])

				}
				if strings.Contains(value2, ": Completed") || strings.Contains(value2, ": Done successfully") {
					if strings.Contains(value2, ": Completed") {
						cdata := strings.Split(value2, "; ")
						if len(cdata) >= 2 {
							ctime := strings.Split(cdata[1], ": ")
							onejob.Endtime = TimeConvert(ctime[0])
						} else {
							ctime := strings.Split(cdata[0], ": ")
							onejob.Endtime = TimeConvert(ctime[0])
						}
					} else {
						ctime := strings.Split(value2, ": ")
						onejob.Endtime = TimeConvert(ctime[0])
					}
				}
				if gpu.MatchString(value2) {
					gval := gpu.FindAllStringSubmatch(value2, -1)
					sgpud, er := strconv.ParseInt(gval[0][1], 0, 64)
					if er != nil {
						gpudata = 0
					} else {
						gpudata = sgpud
					}
				}
			}
			break
		}

	}
	if gpudata != 0 {
		nlen := len(onejob.Nodes)
		for i := 0; i < nlen; i++ {
			onejob.Gpus = append(onejob.Gpus, gpudata)
		}

	}
	return &onejob, nil
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	conn, err := grpc.Dial(grpc_port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := job.NewJobClient(conn)

	Monitor_jobs(client)
}

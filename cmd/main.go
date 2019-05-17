package main

import (
	"log"
	"os"
	"strconv"

	"github.com/forthelols/go-slurm/slurm"
)

func main() {
	nodeInfo := slurm.LoadNode(0, 0)
	nodeInfo.Deref()
	log.Printf("Node info: %+v", nodeInfo)
	slurm.FreeNodeInfoMsg(nodeInfo)

	jobId, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	jobInfo := slurm.LoadJob(uint32(jobId), 0)
	jobInfo.Deref()
	log.Printf("Info for job %d", jobId)
	log.Printf("%+v", jobInfo)
	log.Printf("%+v", jobInfo.Ref())
	slurm.FreeJobInfoMsg(jobInfo)
}

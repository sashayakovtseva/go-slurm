package main

import (
	"log"
	"os"
	"strconv"

	"github.com/forthelols/go-slurm/slurm"
)

func main() {
	nodeInfo := slurm.LoadNode(0, slurm.SHOW_ALL)
	nodeInfo.Deref()
	log.Printf("Node info: %+v", nodeInfo)
	log.Printf("Node array: %+v", nodeInfo.NodeArray)
	slurm.FreeNodeInfoMsg(nodeInfo)

	if len(os.Args) > 1 {
		jobId, err := strconv.ParseUint(os.Args[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		jobInfo := slurm.LoadJob(uint32(jobId), slurm.SHOW_ALL)
		jobInfo.Deref()
		log.Printf("Info for job %d", jobId)
		log.Printf("%+v", jobInfo)
		log.Printf("%+v", jobInfo.Ref())
		slurm.FreeJobInfoMsg(jobInfo)
	}
}

package main

import (
	"fmt"
	"log"
	"lsfmonitor/job"
	"net"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port       = ":5354"
	mg_dburl   = "mongodb://127.0.0.1:27017"
	mgo_dbname = "scheduler"
	mgo_dbcn   = "job"
)

type Job struct{}

var collection *mongo.Collection

func (p *Job) UpdateAndCreateJob(ctx context.Context, request *job.Requestjob) (*job.Responsejob, error) {
	var dbjob job.Requestjob
	jobid := request.Jobid
	err := collection.FindOne(context.Background(), map[string]interface{}{"jobid": jobid}).Decode(&dbjob)
	fmt.Println(err)
	if err != nil {
		d, _ := bson.NewDocumentEncoder().EncodeDocument(request)
		d.Delete("xxx_nounkeyedliteral")
		d.Delete("xxx_unrecognized")
		d.Delete("xxx_sizecache")
		_, inerr := collection.InsertOne(context.Background(), d)
		if inerr == nil {

			return &job.Responsejob{Status: 0}, nil
		} else {
			return &job.Responsejob{Status: 1}, inerr
		}
	} else {
		update := Compare(&dbjob, request)
		if update.Len() > 0 {
			fmt.Println(update)
			_, uperr := collection.UpdateOne(context.Background(), bson.NewDocument(bson.EC.Int64("jobid", dbjob.Jobid)), update)
			if uperr != nil {
				return &job.Responsejob{Status: 1}, uperr
			}
		}

		return &job.Responsejob{Status: 0}, nil
	}
}
func CheckStringarray(a []string, b []string) bool {

	if len(a) != len(b) {
		return false
	}

	tem := false
	for _, val := range a {
		tem = false
		for _, val2 := range b {
			if val2 == val {
				tem = true
				break
			}
		}
		if !tem {
			break
			return false
		}

	}
	return true
}
func CheckIntarray(a []int64, b []int64) bool {

	if len(a) != len(b) {
		return false
	}

	tem := false
	for _, val := range a {
		tem = false
		for _, val2 := range b {
			if val2 == val {
				tem = true
				break
			}
		}
		if !tem {
			break
			return false
		}

	}
	return true
}
func Compare(old *job.Requestjob, news *job.Requestjob) *bson.Document {
	update := bson.NewDocument()
	if old.Status != news.Status {
		update.Append(bson.EC.SubDocumentFromElements("$set", bson.EC.String("status", news.Status)))
	} else {
		return update
	}
	if old.Qtime != news.Qtime {
		update.Append(bson.EC.SubDocumentFromElements("$set", bson.EC.Int64("qtime", news.Qtime)))
	}
	if old.Starttime != news.Starttime {
		update.Append(bson.EC.SubDocumentFromElements("$set", bson.EC.Int64("starttime", news.Starttime)))
	}
	if old.Endtime != news.Endtime {
		update.Append(bson.EC.SubDocumentFromElements("$set", bson.EC.Int64("endtime", news.Endtime)))
	}
	if !CheckStringarray(old.Nodes, news.Nodes) {
		var nodesValue []*bson.Value
		nodesValue = make([]*bson.Value, 0, len(news.Nodes))
		for _, val := range news.Nodes {
			nodesValue = append(nodesValue, bson.VC.String(val))
		}
		update.Append(bson.EC.SubDocumentFromElements("$set", bson.EC.ArrayFromElements("nodes", nodesValue...)))
	}
	if !CheckIntarray(old.Cpus, news.Cpus) {
		var nodesValue []*bson.Value
		nodesValue = make([]*bson.Value, 0, len(news.Cpus))
		for _, val := range news.Cpus {
			nodesValue = append(nodesValue, bson.VC.Int64(val))
		}
		update.Append(bson.EC.SubDocumentFromElements("$set", bson.EC.ArrayFromElements("cpus", nodesValue...)))
	}
	if !CheckIntarray(old.Gpus, news.Gpus) {
		var nodesValue []*bson.Value
		nodesValue = make([]*bson.Value, 0, len(news.Gpus))
		for _, val := range news.Gpus {
			nodesValue = append(nodesValue, bson.VC.Int64(val))
		}
		update.Append(bson.EC.SubDocumentFromElements("$set", bson.EC.ArrayFromElements("gpus", nodesValue...)))
	}
	return update
}
func main() {
	client, err := mongo.NewClient(mg_dburl)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(mgo_dbname).Collection(mgo_dbcn)
	//	var dbjob *construct.CJob
	//	err = collection.FindOne(context.Background(), map[string]interface{}{"jobid": 18}).Decode(dbjob)
	//	fmt.Println(err)
	//	fmt.Println(dbjob)
	listen, lerr := net.Listen("tcp", port)
	if lerr != nil {
		log.Fatal(lerr)
	}
	gserver := grpc.NewServer()
	job.RegisterJobServer(gserver, &Job{})
	gserver.Serve(listen)

}

package main

import (
	"log"
	"lsfproject/jobsearch"
	"net"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Job struct{}

var collection *mongo.Collection

func (j *Job) FindJob(ctx context.Context, request *jobsearch.Requestjob) (*jobsearch.Responsejob, error) {
	var jobitem *jobsearch.Jobitem
	var jobres jobsearch.Responsejob
	var err error
	searchdoc := GetSearchDoument(request)
	log.Println(searchdoc)
	cur, ok := collection.Find(context.Background(), searchdoc)
	if ok != nil {
		log.Fatal(ok)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		jobitem = new(jobsearch.Jobitem)
		err = cur.Decode(jobitem)
		if err != nil {
			return &jobres, err
		}
		jobres.Jobs = append(jobres.Jobs, jobitem)
	}
	return &jobres, nil
}
func GetSearchDoument(job *jobsearch.Requestjob) *bson.Document {
	search := bson.NewDocument()

	if len(job.User) > 0 {
		var users []*bson.Value
		for _, item := range job.User {
			users = append(users, bson.VC.String(item))
		}
		search.Append(bson.EC.SubDocumentFromElements("submiter", bson.EC.ArrayFromElements("$in", users...)))
	}
	if len(job.Status) > 0 {
		var status []*bson.Value
		for _, item := range job.Status {
			status = append(status, bson.VC.String(item))
		}
		search.Append(bson.EC.SubDocumentFromElements("status", bson.EC.ArrayFromElements("$in", status...)))
	}
	return search
}

const (
	port       = ":5355"
	mg_dburl   = "mongodb://127.0.0.1:27017"
	mgo_dbname = "scheduler"
	mgo_dbcn   = "job"
)

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
	listen, lerr := net.Listen("tcp", port)
	if lerr != nil {
		log.Fatal(lerr)
	}
	gserver := grpc.NewServer()
	jobsearch.RegisterJobServer(gserver, &Job{})
	gserver.Serve(listen)

}

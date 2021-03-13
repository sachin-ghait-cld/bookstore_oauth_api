package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// Connect to cassandra
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 3
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}

}

// GetSession Get db Session cassandra
func GetSession() *gocql.Session {
	return session
}

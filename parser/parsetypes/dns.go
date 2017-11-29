package parsetypes

import (
	"github.com/ocmdev/rita/config"
	"gopkg.in/mgo.v2/bson"
)

// DNS provides a data structure for entries in the bro DNS log
type DNS struct {
	// ID contians the id set by mongodb
	ID bson.ObjectId `bson:"_id,omitempty"`
	// TimeStamp of this connection
	TimeStamp int64 `bson:"ts" bro:"ts" brotype:"time"`
	// UID is the Unique Id for this connection (generated by Bro)
	UID string `bson:"uid" bro:"uid" brotype:"string"`
	// Source is the source address for this connection
	Source string `bson:"id_orig_h" bro:"id.orig_h" brotype:"addr"`
	// SourcePort is the source port of this connection
	SourcePort int `bson:"id_orig_p" bro:"id.orig_p" brotype:"port"`
	// Destination is the destination of the connection
	Destination string `bson:"id_resp_h" bro:"id.resp_h" brotype:"addr"`
	// DestinationPort is the port at the destination host
	DestinationPort int `bson:"id_resp_p" bro:"id.resp_p" brotype:"port"`
	// Proto is the string protocol identifier for this connection
	Proto string `bson:"proto" bro:"proto" brotype:"enum"`
	// TransID contains a 16 bit identifier assigned by the program that generated
	// the query
	TransID int64 `bson:"trans_id" bro:"trans_id" brotype:"count"`
	// RTT contains the round trip time of this request / response
	RTT float64 `bson:"RTT" bro:"rtt" brotype:"interval"`
	// Query contians the query string
	Query string `bson:"query" bro:"query" brotype:"string"`
	// QClass contains a the qclass of the query
	QClass int64 `bson:"qclass" bro:"qclass" brotype:"count"`
	// QClassName contains a descriptive name for the query
	QClassName string `bson:"qclass_name" bro:"qclass_name" brotype:"string"`
	// QType contains the value of the query type
	QType int64 `bson:"qtype" bro:"qtype" brotype:"count"`
	// QTypeName provides a descriptive name for the query
	QTypeName string `bson:"qtype_name" bro:"qtype_name" brotype:"string"`
	// RCode contains the response code value from the DNS messages
	RCode int64 `bson:"rcode" bro:"rcode" brotype:"count"`
	// RCodeName provides a descriptive name for RCode
	RCodeName string `bson:"rcode_name" bro:"rcode_name" brotype:"string"`
	// AA represents the state of the authoritive answer bit of the resp messages
	AA bool `bson:"AA" bro:"AA" brotype:"bool"`
	// TC represents the truncation bit of the message
	TC bool `bson:"TC" bro:"TC" brotype:"bool"`
	// RD represens the recursion desired bit of the message
	RD bool `bson:"RD" bro:"RD" brotype:"bool"`
	// RA represents the recursion available bit of the message
	RA bool `bson:"RA" bro:"RA" brotype:"bool"`
	// Z represents the state of a reseverd field that should be zero in qll queries
	Z int64 `bson:"Z" bro:"Z" brotype:"count"`
	// Answers contains the set of resource descriptions in the query answer
	Answers []string `bson:"answers" bro:"answers" brotype:"vector[string]"`
	// TTLs contians a vector of interval type time to live values
	TTLs []float64 `bson:"TTLs" bro:"TTLs" brotype:"vector[interval]"`
	// Rejected indicates if this query was rejected or not
	Rejected bool `bson:"rejected" bro:"rejected" brotype:"bool"`
}

//TargetCollection returns the mongo collection this entry should be inserted
//into
func (in *DNS) TargetCollection(config *config.StructureTableCfg) string {
	return config.DNSTable
}

//Indices gives MongoDB indices that should be used with the collection
func (in *DNS) Indices() []string {
	return []string{"$hashed:id_orig_h", "$hashed:id_resp_h", "$hashed:query"}
}

//Normalize pre processes this type of entry before it is imported by rita
func (in *DNS) Normalize() {}

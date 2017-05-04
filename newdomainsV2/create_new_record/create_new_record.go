package create_new_record

import (
	"github.com/remotejob/odomains/domains"
	"github.com/remotejob/odomains/httpgetter"
)

func Create(token string, domaincsv domains.Domaincsv) {

	newdomain := domains.NewDomain{

		Name: domaincsv.Name,
		Ip:   domaincsv.Ip,
	}

	record := domains.NewRecord{

		Type: "CNAME",
		Name: "*",
		Data: domaincsv.Name + ".",
	}

	//	fmt.Println(record)
	httpgetter.PostNewRecord(token, newdomain, record)

}

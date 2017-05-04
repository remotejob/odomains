package addnewdomain

import (
	"github.com/remotejob/odomains/domains"
	"github.com/remotejob/odomains/httpgetter"
)

func Add(token string, domaincsv domains.Domaincsv) {

	newdomain := domains.NewDomain{

		Name: domaincsv.Name,
		Ip:   domaincsv.Ip,
	}

	httpgetter.PostNewDomain(token, newdomain)

}

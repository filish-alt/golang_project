package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
)

type Metadata struct {
	Useragent string
	ClientIp  string
}
const(
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent" //for grpc user agent
	xForwardedForHeader        = "x-forwarded-for"
)

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt :=&Metadata{}

      if mtd, ok := metadata.FromIncomingContext(ctx); ok{
   log.Printf("m :=%v/n",mtd)

   if useragents := mtd.Get(grpcGatewayUserAgentHeader); len(useragents)>0{
	    mtdt.Useragent = useragents[0]

   }
   
   if useragents := mtd.Get(userAgentHeader); len(useragents)>0{
	mtdt.Useragent = useragents[0]

}

   if clientips := mtd.Get(xForwardedForHeader); len(clientips)>0{
	mtdt.Useragent = clientips[0]

}
	  }
	return mtdt

}
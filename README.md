### Customized logging framework

[Select Logging Framework](https://github.com/go-kratos/kratos/tree/main/contrib/log)


#### using logrus

````go
import (
"flag"
logdef "github.com/go-kratos/kratos/contrib/log/logrus/v2"
"github.com/go-kratos/kratos/v2/middleware/tracing"
"os"

"github.com/go-kratos/kratos-layout/internal/conf"
"github.com/go-kratos/kratos/v2"
"github.com/go-kratos/kratos/v2/config"
"github.com/go-kratos/kratos/v2/config/file"
"github.com/go-kratos/kratos/v2/log"
"github.com/go-kratos/kratos/v2/transport/grpc"
"github.com/go-kratos/kratos/v2/transport/http"
"github.com/sirupsen/logrus"
)


func main(){
    logFmt := logrus.New()
    logFmt.Formatter = &logrus.JSONFormatter{}
    logFmt.SetLevel(logrus.InfoLevel)
    logger := logdef.NewLogger(logFmt)
    logger = log.With(
        logger,
        "ts", log.DefaultTimestamp,
        "caller", log.DefaultCaller,
        "service.id", id,
        "service.name", Name,
        "service.version", Version,
        "trace.id", tracing.TraceID(),
        "span.id", tracing.SpanID(),
    )
    
    app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
}

````
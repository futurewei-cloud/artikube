package artikube

import (
	"github.com/gin-gonic/gin"
)

var rootPageHTML = []byte(`<!DOCTYPE html>
<html>
<head>
<title>Welcome to Artikube, Golang implemenation of Artifact Repository Manager</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to Artikube, Golang implemenation of Artifact Repository Manager</h1>
<p>Aritpie is running.</p>

<p>For more information, please refer to the
<a href="https://github.com/futurewei-cloud/artikube">GitHub project</a>.<br/>

<p><em>Next thing you can try uploading an artifact to the server.</em></p>
</body>
</html>
`)

func (svr *Server) getRootPageHandler(c *gin.Context) {
	c.Data(200, "text/html", rootPageHTML)
}

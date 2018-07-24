# CereVoice Go
CereVoice Go is a [Go](https://golang.org) package for accessing the [CereProc](https://www.cereproc.com/) 
[CereVoice Cloud API](https://www.cereproc.com/en/products/cloud). Currently this package 
implements all available funcitons using the REST API. SOAP has not yet been implemented.

## Notice
Please note that is package is not complete and is considered pre-release, meaning it
is subject to change. Safety not guaranteed.

## Usage

Assuming you have Go already setup and working, grab the latest version of the package 
from master.

```sh
go get github.com/bganderson/cerevoicego
```

Import the package into your project.

```go
import "github.com/bganderson/cerevoicego"
```

Create a cerevoice client. Please note that `AccountID` and `Password` are not the same 
credentials you use to login to the website.

For the REST API URL you can use the exported const provided by the package.

```go
cerevoice := cerevoicego.Client{
    CereVoiceAPIURL: cerevoicego.DefaultRESTAPIURL,
    AccountID:       "<YOUR_ACCOUNTID>",
    Password:        "<YOUR_PASSWORD",
}
```

Make an API request and do something with the response.

```go
res := cerevoice.SpeakSimple(&cerevoicego.SpeakSimpleInput{
    Voice: "Jess",
    Text:  "Hello world!",
})
if res.Error != nil {
    log.Fatalln(res.Error)
}

fmt.Printf("The sound file is available at: %s\n", res.FileURL)
```





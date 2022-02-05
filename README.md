# Trafikverket Client
Client library written in Go for consuming Trafikverket's api.

## Scope

This library provides the constructs used for communicating with Trafikverket.

## Usage

The request object is built using a builder-like approach:

```go
func GetExampleRequest(authenticationKey string) *trv.Request {
	request := trv.
		NewRequest(authenticationKey)
	or := request.
		NewQuery("TrainAnnouncement", "1.6").
		WithOrderBy("AdvertisedTimeAtLocation").
		AddInclude("AdvertisedTrainIdent").
		AddInclude("AdvertisedTimeAtLocation").
		AddInclude("TrackAtLocation").
		AddInclude("ToLocation").
		NewFilter().
		AddEQ("ActivityType", "Avgang").
		AddEQ("LocationSignature", "Cst").
		NewOr()
	or.
		NewAnd().
		AddGT("AdvertisedTimeAtLocation", "$dateadd(-00:15:00)").
		AddLT("AdvertisedTimeAtLocation", "$dateadd(14:00:00)")
	or.
		NewAnd().
		AddLT("AdvertisedTimeAtLocation", "$dateadd(00:30:00)").
		AddGT("EstimatedTimeAtLocation", "$dateadd(-00:15:00)")
	return request
}
```
Note that methods prefixed with `New` returns new objects while the other methods return the object for which the method was called. Thus, the return value of `NewRequest` and `NewOr` are stored in variables so that they can be used later on. 

Marshaling the Request returned by `GetExampleRequest` generates the following xml string:
```xml
<REQUEST>
  <LOGIN authenticationkey="-"></LOGIN>
  <QUERY objecttype="TrainAnnouncement" orderby="AdvertisedTimeAtLocation" schemaversion="1.6">
    <FILTER>
      <EQ name="ActivityType" value="Avgang"></EQ>
      <EQ name="LocationSignature" value="Cst"></EQ>
      <OR>
        <AND>
          <GT name="AdvertisedTimeAtLocation" value="$dateadd(-00:15:00)"></GT>
          <LT name="AdvertisedTimeAtLocation" value="$dateadd(14:00:00)"></LT>
        </AND>
        <AND>
          <LT name="AdvertisedTimeAtLocation" value="$dateadd(00:30:00)"></LT>
          <GT name="EstimatedTimeAtLocation" value="$dateadd(-00:15:00)"></GT>
        </AND>
      </OR>
    </FILTER>
    <INCLUDE>AdvertisedTrainIdent</INCLUDE>
    <INCLUDE>AdvertisedTimeAtLocation</INCLUDE>
    <INCLUDE>TrackAtLocation</INCLUDE>
    <INCLUDE>ToLocation</INCLUDE>
  </QUERY>
</REQUEST>
```
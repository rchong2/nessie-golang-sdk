package merchant

import(
    "fmt"
    "math"
    "net/http"
    "io/ioutil"
    "strconv"
    "bytes"
    "../shared"
)

const baseUrl = "http://api.reimaginebanking.com/merchants"
var apiKey = shared.ApiKey
const blankNumber = math.SmallestNonzeroFloat64

//GET: Returns the merchants that have been assigned to you
func GetAllMerchants(lat float64, lng float64, rad int){
	
    latString := strconv.FormatFloat(lat,'f',4,64)
    lngString := strconv.FormatFloat(lng,'f',4,64)
    radString := strconv.Itoa(rad)

	url := baseUrl + "?lat=" + latString + "&lng=" + lngString + "&rad=" + radString + "&key=" + apiKey

	req, err := http.NewRequest("GET", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

//GET: Returns the merchant with the specific id
func GetMerchantInfo(merchantId string){

    var url = baseUrl + "/" + merchantId + "?key=" + apiKey

    req, err := http.NewRequest("GET", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

//POST: Creates a merchant
//For optional Params, use empty string "" and blankNumber for empty lat/lng
func CreateMerchant(merchantName string, category string, street_number string, street_name string, city string, state string, zip string, lat float64, lng float64){

    url := baseUrl + "?key=" + apiKey

    fmt.Println("URL:>", url)

    var latString = strconv.FormatFloat(lat,'f',4,64)
    var lngString = strconv.FormatFloat(lng,'f',4,64)

    var geocode = `{"lat": ` + latString + `, "lng": ` + lngString + `}`
    var address = `{"street_number": "` + street_number + `", "street_name":"` + street_name + `", "city": "` + city + `", "state":"` + state + `", "zip":"` + zip + `"}`

    if lat == blankNumber || lng == blankNumber {
    	geocode = ""
    }

    var payloadStr = `{"name":"` + merchantName + `"`

    if len(category) > 0{
    	payloadStr = payloadStr + `, "category":"` + category + `"`
    }

    if len(street_number) > 0{
    	payloadStr = payloadStr + `,"address":` + address
    }
    
    if len(geocode) > 0{
    	payloadStr = payloadStr + `, "geocode": ` + geocode
    } 
    
    payloadStr = payloadStr + `}`

    fmt.Println("geocode payload:", string(geocode))
    fmt.Println("address payload:", string(address))
    fmt.Println("payload:", string(payloadStr))

    var jsonStr = []byte(payloadStr)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response Status:", resp.Status)
    fmt.Println("Response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response Body:", string(body))
}

//PUT: Updates a specific merchant
//For optional Params, use empty string "" and blankNumber for empty lat/lng
func UpdateMerchant(merchantId string, merchantName string, category string, street_number string, street_name string, city string, state string, zip string, lat float64, lng float64){

    url := baseUrl + "/" + merchantId + "?key=" + apiKey

    fmt.Println("URL:>", url)

    var latString = strconv.FormatFloat(lat,'f',4,64)
    var lngString = strconv.FormatFloat(lng,'f',4,64)

    var geocode = `{"lat": ` + latString + `, "lng": ` + lngString + `}`
    var address = `{"street_number": "` + street_number + `", "street_name":"` + street_name + `", "city": "` + city + `", "state":"` + state + `", "zip":"` + zip + `"}`

    if lat == blankNumber || lng == blankNumber {
    	geocode = ""
    }

    var payloadStr = `{"name":"` + merchantName + `"`

    if len(category) > 0{
    	payloadStr = payloadStr + `, "category":"` + category + `"`
    }

    if len(street_number) > 0{
    	payloadStr = payloadStr + `,"address":` + address
    }
    
    if len(geocode) > 0{
    	payloadStr = payloadStr + `, "geocode": ` + geocode
    } 
    
    payloadStr = payloadStr + `}`
    
    fmt.Println("geocode payload:", string(geocode))
    fmt.Println("address payload:", string(address))
    fmt.Println("payload:", string(payloadStr))

    var jsonStr = []byte(payloadStr)
    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response Status:", resp.Status)
    fmt.Println("Response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response Body:", string(body))
}
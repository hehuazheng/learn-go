package temporary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

/*
{
	"status": 1,
	"code": 200,
	"message": "请求成功",
	"data": {
		"driverId": 1,
		"driverName": "xxx",
		"locations": [{
			"latlng": {
				"latitude": "39.82373519712668",
				"longitude": "116.56605988044419"
			},
			"#timestamp": 1567887172871
		}]
	}
}
*/
func TestParse(t *testing.T) {
	if contents, err := ioutil.ReadFile("/Users/xxx/json.txt"); err == nil {
		jsonData := string(contents)
		var v DriverTrackers
		json.Unmarshal([]byte(jsonData), &v)
		for _, r := range v.Data.Locations {
			fmt.Println(r)
		}
	}
}

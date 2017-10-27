package schedule

import (
	"peaberry/domain/entity"
	"peaberry/adaptor/ghttp"
	"net/http"
	"fmt"
	"encoding/json"
)

var HOST = "https://imt.services.isca.jp"
var COOKIE = "schSchDialogReferWidth=725; schSchDialogReferHeight=600; jp.co.intra_mart.session.cookie=11qfyzzx|1509066128212; __zlcmid=iognBu9o4TNeii; _ga=GA1.2.2011196724.1507019323; _gid=GA1.2.129791155.1508720447; AWSELB=0DFB4119120A7C8E000FD729A065620BDD13FCDC87B5EA90D77EBEB2DDCF22CB4F694EA114BE6FBAC628DF42FA281DFAC56D9567F961ED0AB0174320D95C201674279E2DFE; opentoken=T1RLAQII9p2oL52ijptzAxSkRKQx1RFhFBDNyeQmTBAHWsrFjONwyIelAACAqy_eCsRORCdH0U73rflFgR0BJEFDECtSUQ5gJuMIuXtFxKGmLi1cYoTco64oZESRFSHEtSSLNXULfuVyivDbgJWc8gwe69RbCD_nOCsLFNrwZL23YapMf-DzYaEYS0o-YqX2uOtFoH8AC-wCzMOojSxdXVKKuzACGQsnv_v6iNo*; JSESSIONID=aaah1uK2D6Q4XfokXtj9v"
var TODAY_SCHEDULE_PATH = "/imart/collaboration/schedule/user/calendar/find_personal_day"
var URL = HOST + TODAY_SCHEDULE_PATH

func FetchTodaySchedule() []entity.Schedule {

	println("Fetching today schedules...")

	req, err := http.NewRequest("POST", URL, nil)
	req.Header.Set("Cookie", COOKIE)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	body := ghttp.Do(req)

	var dto Response
	if err := json.Unmarshal(body, &dto); err != nil {
		fmt.Println(err)
		panic(err)
	}

	var schedules = []entity.Schedule{}
	for _, s := range dto.Data.Schedules {
		schedules = append(schedules, toEntity(s))
	}

	return schedules
}

func toEntity(s *Schedule) entity.Schedule {
	return entity.Schedule{
		Id:    s.ScheduleKey.Code,
		Title: s.Title,
		Start: s.Start.Time,
		End:   s.End.Time,
	}
}

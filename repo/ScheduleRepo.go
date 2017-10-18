package repo

import (
	"peaberry/repo/ghttp"
	"peaberry/repo/dto"
	"peaberry/domain/entity"
)

var TODAY_SCHEDULE_PATH = "/imart/collaboration/schedule/user/calendar/find_personal_day"

func FetchTodaySchedule() []entity.Schedule {
	println("Fetching today schedules...")
	res := ghttp.Post(TODAY_SCHEDULE_PATH)

	var schedules = []entity.Schedule{}
	for _, s := range res.Data.Schedules {
		schedules = append(schedules, toEntity(s))
	}

	return schedules
}

func toEntity(s *dto.Schedule) entity.Schedule {
	return entity.Schedule{
		Title: s.Title,
		Start: s.Start.Time,
		End:   s.End.Time,
	}
}

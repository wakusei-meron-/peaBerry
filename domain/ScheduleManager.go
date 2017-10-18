package domain

import (
	//"peaberry/repo"
	"peaberry/domain/entity"
	"peaberry/util"
	"time"
)

type NotificationSchedules []entity.Schedule

func StartApplication() {
	notificationSchedules := NotificationSchedules{}
	//latestSchedules := repo.FetchTodaySchedule()
	//util.PrettyPrint(latestSchedules)
	// 最新のスケジュールの取得
	latestSchedules := []entity.Schedule{}
	start, _ := time.Parse("2006-01-02T15:04:05", "2017-10-18T20:00:00")
	s := entity.Schedule{
		Title: "title1",
		Start: start,
		End: time.Now(),
	}
	latestSchedules = append(latestSchedules, s)

	// 通知予定のスケジュールと比較
	newSchedule := notificationSchedules.diff(latestSchedules)
	util.PrettyPrint(newSchedule)

	remindSchedule(latestSchedules)
}

/**
 * 予定をリマインドします
 */
func remindSchedule(schedules []entity.Schedule) {
	//d, _ :=time.ParseDuration("9h")
	//now := time.Now().Add(d)
	//util.PrettyPrint(now)
	//
	//for _, s := range schedules{
	//	util.PrettyPrint(s.Start)
	//	duration := now.Sub(s.Start).Minutes()
	//
	//	util.PrettyPrint(duration)
	//}

	// Slack通知
}

/**
 * 通知予定と最新予定の比較
 */
 func (n *NotificationSchedules) diff(schedules []entity.Schedule) ([]entity.Schedule){
 	newSchedules := []entity.Schedule{}
 	for _, s := range schedules {
 		if !n.contains(s) {
			newSchedules = append(newSchedules, s)
		}
	}

	//deletedSchedules := []entity.Schedule{}
	//for _, ns := range *n{
	//	if s.conhhh
	//}

	return newSchedules
 }

 func (n *NotificationSchedules) contains(schedule entity.Schedule) bool {
 	for _, ns := range *n{
 		if ns == schedule {
 			return true
		}
	}

	return false
 }

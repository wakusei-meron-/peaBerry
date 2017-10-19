package domain

import (
	"peaberry/domain/entity"
	"peaberry/adaptor/mac"
	"time"
	"fmt"
	"peaberry/config"
)

type NotificationSchedules []entity.Schedule

func StartApplication(config config.Config) {
	notificationSchedules := NotificationSchedules{}
	//latestSchedules := repo.FetchTodaySchedule()
	//util.PrettyPrint(latestSchedules)
	// 最新のスケジュールの取得
	latestSchedules := []entity.Schedule{}
	start, _ := time.Parse("2006-01-02T15:04:05", "2017-10-18T20:00:00")
	s := entity.Schedule{
		Title: "title1",
		Start: start,
		End:   time.Now(),
	}
	latestSchedules = append(latestSchedules, s)

	// 通知予定のスケジュールと比較
	newSchedules, deletedSchedules := notificationSchedules.diff(latestSchedules)

	notificationMsg := formatUpdateMessage(newSchedules, deletedSchedules, config.Notification)
	mac.Notify(config.Notification.UpdatedTitle, notificationMsg, config.Notification.SoundFlag)

	//remindSchedule(latestSchedules)
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

}

/**
 * 通知予定と最新予定の比較
 */
func (n *NotificationSchedules) diff(schedules []entity.Schedule) ([]entity.Schedule, []entity.Schedule) {
	newSchedules := []entity.Schedule{}
	for _, s := range schedules {
		if !contains(*n, s) {
			newSchedules = append(newSchedules, s)
		}
	}

	deletedSchedules := []entity.Schedule{}
	for _, ns := range *n {
		if !contains(schedules, ns) {
			deletedSchedules = append(deletedSchedules, ns)
		}
	}

	return newSchedules, deletedSchedules
}

func contains(schedules []entity.Schedule, schedule entity.Schedule) bool {
	for _, s := range schedules {
		if s == schedule {
			return true
		}
	}
	return false
}

/**
 * 変更予定の通知用の文字列を生成します
 */
func formatUpdateMessage(newSchedules []entity.Schedule, deletedSchedules []entity.Schedule, config config.NotificationConfig) string {
	dateFormat := "15:04"
	var msg string
	for _, s := range newSchedules {

		msg += fmt.Sprintln(config.PrefixNewSchedule,
			s.Start.Format(dateFormat),
			" ~ ",
			s.End.Format(dateFormat),
			s.Title)
	}
	for _, s := range deletedSchedules {
		msg += fmt.Sprintln(config.PrefixDeletedSchedule,
			s.Start.Format(dateFormat),
			" ~ ",
			s.End.Format(dateFormat),
			s.Title)
	}
	return msg
}

package domain

import (
	"github.com/robfig/cron"
	"peaberry/domain/entity"
	"peaberry/config"
	"time"
	"fmt"
	"peaberry/util"
	"math"
)

type ScheduleManager struct {
	NotiSchedules []entity.Schedule
}

var schedMgr = newInstance()

func newInstance() *ScheduleManager {
	s := []entity.Schedule{}
	sm := ScheduleManager{s}
	return &sm
}

func GetInstance() *ScheduleManager {
	return schedMgr
}

var notiConf = config.GetInstance().Notification

func StartApplication() {
	//latestSchedules := repo.FetchTodaySchedule()
	//util.PrettyPrint(latestSchedules)
	// 最新のスケジュールの取得
	latestSchedules := []entity.Schedule{}
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	duration, _ := time.ParseDuration("-9h")
	start, _ := time.Parse("2006-01-02T15:04:05", "2017-10-24T00:55:00")
	s := entity.Schedule{
		Title: "title1",
		Start: start.Add(duration).In(jst),
		End:   time.Now(),
	}
	latestSchedules = append(latestSchedules, s)

	// 通知予定のスケジュールと比較
	newSchedules, deletedSchedules := schedMgr.diff(latestSchedules)

	// 予定変更の通知
	notiMsg := formatUpdateMessage(newSchedules, deletedSchedules)
	noti := entity.Notification{notiConf.UpdatedTitle, notiMsg}
	noti.Fire()

	// 予定変更の反映
	schedMgr.add(newSchedules)
	schedMgr.delete(deletedSchedules)

	// 通知予定のものがあるかチェック
	inTimeSchedule := schedMgr.check()
	if len(inTimeSchedule) > 0 {
		remind(inTimeSchedule)
		schedMgr.done(inTimeSchedule)
	}
}

/**
 * 通知予定時間内の予定を通知
 */
func remind(schedules []entity.Schedule) {
	dateFormat := "15:04"
	var msg string
	for _, s := range schedules {
		msg += fmt.Sprintln(notiConf.PrefixNewSchedule,
			s.Start.Format(dateFormat),
			" ~ ",
			s.End.Format(dateFormat),
			s.Title)
	}

	n := entity.Notification{
		"リマインド",
		msg,
	}
	n.Fire()
}

/**
 * 指定時間以内に通知予定のものがある予定を取得
 * 一度通知しているものは除外
 */
func (sm *ScheduleManager) check() []entity.Schedule {
	now := time.Now()
	inTimeSchedule := []entity.Schedule{}

	for _, s := range sm.NotiSchedules {
		if s.Done {
			continue
		}

		duration := s.Start.Sub(now)

		m := int(duration.Minutes())

		if (m >= 0) && (m < notiConf.MinuteBefore) {
			inTimeSchedule = append(inTimeSchedule, s)
		}
	}
	return inTimeSchedule
}

/**
 * 通知予定の登録
 */
func register(schdules []entity.Schedule) {

	c := cron.New()
	for _, s := range schdules {
		c.AddFunc("@every 5s", func() {
			util.PrettyPrint(s)
		})
	}

	c.Start()
	util.PrettyPrint(c.Entries())

	for {
		time.Sleep(math.MaxInt64)
		fmt.Println("sleep")
	}
}

/**
 * 通知予定と最新予定の比較
 */
func (sm *ScheduleManager) diff(schedules []entity.Schedule) ([]entity.Schedule, []entity.Schedule) {
	newSchedules := []entity.Schedule{}
	for _, s := range schedules {
		if !contains(sm.NotiSchedules, s) {
			newSchedules = append(newSchedules, s)
		}
	}

	deletedSchedules := []entity.Schedule{}
	for _, ns := range sm.NotiSchedules {
		if !contains(schedules, ns) {
			deletedSchedules = append(deletedSchedules, ns)
		}
	}

	return newSchedules, deletedSchedules
}

func (sm *ScheduleManager) add(schedules []entity.Schedule) {
	sm.NotiSchedules = append(sm.NotiSchedules, schedules...)
}

func (sm *ScheduleManager) delete(schedules []entity.Schedule) {
	notiSchedules := []entity.Schedule{}
	for _, s := range schedules {
		for i, sms := range sm.NotiSchedules {
			if s == sms {
				continue
			}
			sm.NotiSchedules = append(notiSchedules, sm.NotiSchedules[i])
		}
	}
	sm.NotiSchedules = notiSchedules
}

/**
 * 通知が完了した予定のステータス変更
 */
func (sm *ScheduleManager) done(schedules []entity.Schedule) {
	for i, s := range sm.NotiSchedules {
		if contains(schedules, s) {
			sm.NotiSchedules[i].Done = true
		}
	}
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
func formatUpdateMessage(newSchedules []entity.Schedule, deletedSchedules []entity.Schedule) string {
	dateFormat := "15:04"

	var msg string
	for _, s := range newSchedules {

		msg += fmt.Sprintln(notiConf.PrefixNewSchedule,
			s.Start.Format(dateFormat),
			" ~ ",
			s.End.Format(dateFormat),
			s.Title)
	}
	for _, s := range deletedSchedules {
		msg += fmt.Sprintln(notiConf.PrefixDeletedSchedule,
			s.Start.Format(dateFormat),
			" ~ ",
			s.End.Format(dateFormat),
			s.Title)
	}
	return msg
}

package ui

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/skanehira/ght/domain"
	"github.com/skanehira/ght/utils"
)

var MilestoneUI *SelectUI

func NewMilestoneUI() {
	//getList := func(cursor *string) ([]List, github.PageInfo) {
	//	v := map[string]interface{}{
	//		"owner":  githubv4.String(config.GitHub.Owner),
	//		"name":   githubv4.String(config.GitHub.Repo),
	//		"first":  githubv4.Int(100),
	//		"cursor": (*githubv4.String)(cursor),
	//	}
	//	resp, err := github.GetRepoMillestones(v)
	//	if err != nil {
	//		return nil, github.PageInfo{}
	//	}

	//	milestones := make([]List, len(resp.Nodes))
	//	for i, m := range resp.Nodes {
	//		milestones[i] = &Milestone{
	//			Title: string(m.Title),
	//		}
	//	}

	//	return milestones, resp.PageInfo
	//}

	setOpt := func(ui *SelectUI) {

		ui.capture = func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyCtrlO:
				var urls []string
				if len(MilestoneUI.selected) == 0 {
					data := MilestoneUI.GetSelect()
					if data != nil {
						urls = append(urls, data.(*domain.Milestone).URL)
					}
				} else {
					for _, s := range MilestoneUI.selected {
						urls = append(urls, s.(*domain.Milestone).URL)
					}
				}

				for _, url := range urls {
					if err := utils.Open(url); err != nil {
						log.Println(err)
					}
				}
			}
			return event
		}
	}

	ui := NewSelectListUI(UIKindMilestones, tcell.ColorGreen, setOpt)
	MilestoneUI = ui
}

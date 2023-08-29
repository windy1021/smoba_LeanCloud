## {{.RoleName}}{{.Sign}}
* {{.Sign}}{{if eq .SignB true}}

## 账户信息
* 游戏昵称: {{.UserName}}
* 营地ID: {{.UserId}}
* 段位: {{.RoleJob}}

## 王者日历 
* 农历: {{.Lunar}} 
* 适宜：{{.Good}}
* 禁忌：{{.Bad}}
{{end}}

## 完成任务 {{range $i, $v := .DoTask}}
* {{$v}}{{end}}

## 领取奖励 {{range $i, $v := .DoGift}}
* {{$v}}{{end}}

## 时间
* {{.Time}}

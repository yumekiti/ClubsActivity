/*
## Activity Data（活動データ）

| カラム名         | 説明                    | 型        | Unique | Nullable |
| ---------------- | ----------------------- | --------- | ------ | -------- |
| activity_id      | 活動 ID                 | Integer   | Yes    | No       |
| activity_place   | 活動場所                | String    | No     | No       |
| activity_detail  | 活動詳細                | Text      | No     | Yes      |
| activity_user_id | 活動に参加したユーザー  | Integer[] | No     | Yes      |
| activity_club_id | 活動が行われた部活動 ID | Integer   | No     | No       |
| created_at       | 作成日時                | DateTime  | No     | No       |
| updated_at       | 更新日時                | DateTime  | No     | No       |
*/

package domain

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Place  string
	Detail string
	Users  []User `gorm:"many2many:activity_users;"`
	ClubID uint
	Club   Club `gorm:"foreignKey:ClubID"`
}

type ActivityUser struct {
	gorm.Model
	ActivityID uint
	UserID     uint
}

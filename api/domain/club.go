/*
## Club Data（部活動データ）

| カラム名         | 説明             | 型       | Unique | Nullable |
| ---------------- | ---------------- | -------- | ------ | -------- |
| club_id          | 部活動 ID        | Integer  | Yes    | No       |
| club_name        | 部活動名         | String   | Yes    | No       |
| club_description | 部活動の説明     | String   | No     | Yes      |
| club_category    | 部活動のカテゴリ | String   | No     | Yes      |
| created_at       | 作成日時         | DateTime | No     | No       |
| updated_at       | 更新日時         | DateTime | No     | No       |
*/

package domain

import (
	"gorm.io/gorm"
)

type Club struct {
	gorm.Model
	Name     string
	Readme   string
	Category string
}

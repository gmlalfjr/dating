package profile

import (
	"dating/domains/entities"
	"time"
)

func (p *ProfileRepository) FindProfileExcludeUserId(userId int, limit, offset int) ([]*entities.User, int64, error) {
	result := []*entities.User{}
	var countData int64

	subquery := p.db.Table("swipes").
		Select("DISTINCT swiped_user_id").
		Where("created_at >= ?", time.Now().AddDate(0, 0, -1).Format("2006-01-02"))

	if err := p.db.Where("id NOT IN (?) AND id != ?", subquery, 2).Find(&result).Error; err != nil {
		return nil, 0, err
	}

	if err := p.db.Model(&entities.User{}).
		Where("id NOT IN (?) AND id != ?", subquery, 2).
		Count(&countData).Error; err != nil {
		return nil, 0, err
	}

	return result, countData, nil
}

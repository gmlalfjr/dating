package profile

import "dating/domains/entities"

func (p *ProfileRepository) FindProfileExcludeUserId(userId int, limit, offset int) ([]*entities.User, int64, error) {
	result := []*entities.User{}
	var countData int64
	if err := p.db.Model(&entities.User{}).Count(&countData).Error; err != nil {
		return nil, 0, err
	}
	if err := p.db.Where("id != ?", userId).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, 0, err
	}
	return result, countData, nil
}

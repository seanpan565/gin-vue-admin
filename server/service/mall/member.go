package mall

import (
	"errors"
	"regexp"
	"sync"
	"time"

	"mall-admin/server/global"
	"mall-admin/server/model/mall"
	mallReq "mall-admin/server/model/mall/request"
	"mall-admin/server/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var mobilePattern = regexp.MustCompile(`^1[3-9]\d{9}$`)

var (
	cachedDefaultLevelID uint
	defaultLevelOnce     sync.Once
	defaultLevelErr      error
)

// MemberService 商城会员业务逻辑。
type MemberService struct{}

// ResetDefaultLevelCache 等级数据变更后可调用以刷新缓存。
func (s *MemberService) ResetDefaultLevelCache() {
	cachedDefaultLevelID = 0
	defaultLevelOnce = sync.Once{}
	defaultLevelErr = nil
}

// EnsureDefaultLevels 初始化默认会员等级。
func (s *MemberService) EnsureDefaultLevels() error {
	var count int64
	if err := global.GVA_DB.Model(&mall.MallMemberLevel{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	levels := []mall.MallMemberLevel{
		{Name: "普通会员", Level: 1, MinGrowth: 0, DiscountRate: 100, Status: 1},
		{Name: "银卡会员", Level: 2, MinGrowth: 1000, DiscountRate: 98, Status: 1},
		{Name: "金卡会员", Level: 3, MinGrowth: 5000, DiscountRate: 95, Status: 1},
	}
	return global.GVA_DB.Create(&levels).Error
}

func (s *MemberService) defaultLevelID() (uint, error) {
	defaultLevelOnce.Do(func() {
		var level mall.MallMemberLevel
		defaultLevelErr = global.GVA_DB.Where("level = ? AND status = ?", 1, 1).First(&level).Error
		if defaultLevelErr == nil {
			cachedDefaultLevelID = level.ID
		}
	})
	if defaultLevelErr != nil {
		return 0, defaultLevelErr
	}
	return cachedDefaultLevelID, nil
}

// Register 会员注册：创建账号 + 扩展资料。
func (s *MemberService) Register(req mallReq.MemberRegister, clientIP string) (mall.MallMember, error) {
	if !mobilePattern.MatchString(req.Mobile) {
		return mall.MallMember{}, errors.New("手机号格式不正确")
	}

	var exists mall.MallMember
	if err := global.GVA_DB.Where("mobile = ?", req.Mobile).First(&exists).Error; err == nil {
		return mall.MallMember{}, errors.New("该手机号已注册")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return mall.MallMember{}, err
	}

	levelID, err := s.defaultLevelID()
	if err != nil {
		return mall.MallMember{}, errors.New("会员等级未初始化")
	}

	source := req.RegisterSource
	if source == "" {
		source = "h5"
	}
	nickname := req.Nickname
	if nickname == "" {
		nickname = "用户" + req.Mobile[len(req.Mobile)-4:]
	}

	member := mall.MallMember{
		UUID:           uuid.New(),
		Mobile:         req.Mobile,
		Password:       utils.BcryptHash(req.Password),
		Nickname:       nickname,
		Status:         1,
		RegisterSource: source,
		RegisterIP:     clientIP,
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&member).Error; err != nil {
			return err
		}
		profile := mall.MallMemberProfile{
			MemberID: member.ID,
			LevelID:  levelID,
		}
		return tx.Create(&profile).Error
	})
	if err != nil {
		return mall.MallMember{}, err
	}

	member.Profile = mall.MallMemberProfile{MemberID: member.ID, LevelID: levelID}
	var level mall.MallMemberLevel
	if e := global.GVA_DB.First(&level, levelID).Error; e == nil {
		member.Profile.Level = level
	}
	return member, nil
}

// Login 会员密码登录，返回含扩展资料与等级的完整信息。
func (s *MemberService) Login(mobile, password string) (*mall.MallMember, error) {
	var member mall.MallMember
	if err := global.GVA_DB.Preload("Profile.Level").Where("mobile = ?", mobile).First(&member).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("手机号或密码错误")
		}
		return nil, err
	}
	if member.Status != 1 {
		return nil, errors.New("账号已被冻结")
	}
	if !utils.BcryptCheck(password, member.Password) {
		return nil, errors.New("手机号或密码错误")
	}
	return &member, nil
}

// UpdateLastLogin 更新最后登录信息。
func (s *MemberService) UpdateLastLogin(memberID uint, clientIP string) error {
	now := time.Now()
	return global.GVA_DB.Model(&mall.MallMember{}).Where("id = ?", memberID).Updates(map[string]interface{}{
		"last_login_at": now,
		"last_login_ip": clientIP,
	}).Error
}

// GetMemberByID 查询会员（含扩展资料与等级）。
func (s *MemberService) GetMemberByID(id uint) (mall.MallMember, error) {
	var member mall.MallMember
	err := global.GVA_DB.Preload("Profile.Level").Where("id = ?", id).First(&member).Error
	return member, err
}

// GetMemberList 后台会员分页列表。
func (s *MemberService) GetMemberList(info mallReq.MemberSearch) (list []mall.MallMember, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&mall.MallMember{})

	if info.Mobile != "" {
		db = db.Where("mobile LIKE ?", "%"+info.Mobile+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", *info.Status)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Preload("Profile.Level").Limit(limit).Offset(offset).Order("id desc").Find(&list).Error
	return
}

// CreateLoginLog 记录会员登录日志。
func (s *MemberService) CreateLoginLog(log mall.MallMemberLoginLog) error {
	return global.GVA_DB.Create(&log).Error
}

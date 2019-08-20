package model

import "time"

type Service struct {
	Id  	int64 		`json:"id"`
	Gid     int64 		`json:"gid"`
	Name 	string 		`json:"name"` // 服务名称

	Path  	string 		`json:"path"` // 服务所在路径
	Conf    string 		`json:"conf"` // 配置文件路径
	Start 	string 		`json:"start"` // 启动命令
	Restart string 		`json:"restart"` // 重启命令
	Stop    string 		`json:"stop"`
	Ctime   time.Time 	`json:"ctime"`
}

func (s *Service)Save() error {
	
	_,err := Orm.InsertOne(s)
	if err != nil{
		return err
	}
	return nil
}

func (s *Service)GetForId(id int64) (*Service,error) {
	bean := new(Service)
	has,err := Orm.Id(id).Get(bean)
	if err != nil{
		return nil,err
	}
	if !has{
		return nil,nil
	}
	return bean,nil
}

func(s *Service)GetListForGid(gid int64) ([]*Service,error)  {
	beans := make([]*Service,0)
	err := Orm.Where("gid = ?",gid).Find(&beans)
	if err != nil{
		return nil,err
	}
	return beans,nil
}
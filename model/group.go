package model

import "fmt"

type ServiceGroup struct {
	Id  int64 `json:"id"`
	Name string `json:"name" xorm:"unique index"`
}


func (s *ServiceGroup)Save() error {
	if len(s.Name) == 0{
		return fmt.Errorf("组名不能为空")
	}
	_,err := Orm.InsertOne(s)
	if err != nil{
		return err
	}
	return nil
}

func (s *ServiceGroup)Update() error {
	if s.Id == 0{
		return fmt.Errorf("ID不能为空")
	}
	if len(s.Name) == 0{
		return fmt.Errorf("组名不能为空")
	}
	_,err := Orm.Id(s.Id).Update(s)
	if err != nil{
		return err
	}
	return nil
}

func (s *ServiceGroup)Del(id int64) error {
	if s.Id == 0{
		return fmt.Errorf("ID不能为空")
	}

	data,_ := new(Service).GetListForGid(id)
	if len(data) > 0{
		return fmt.Errorf("组下还有服务，无法删除")
	}
	_,err := Orm.Id(id).Delete(ServiceGroup{})
	if err != nil{
		return err
	}
	return nil
}

func (s *ServiceGroup)GetForId(id int64) (*ServiceGroup,error) {
	bean := new(ServiceGroup)
	has,err := Orm.Id(id).Get(bean)
	if err != nil{
		return nil,err
	}
	if !has{
		return nil,nil
	}
	return bean,nil
}

func(s *ServiceGroup)GetList() ([]*ServiceGroup,error)  {
	beans := make([]*ServiceGroup,0)
	err := Orm.Find(&beans)
	if err != nil{
		return nil,err
	}
	return beans,nil
}
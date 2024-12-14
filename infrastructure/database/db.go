package database

import (
	"api-pedidos/adapter/database"
	"api-pedidos/core/domain"
	"api-pedidos/core/erros"
	"api-pedidos/infrastructure/database/entity"
	"api-pedidos/infrastructure/logger"
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLConnection struct {
	db *gorm.DB
}

var _ database.DbPedidos = (*SQLConnection)(nil)

func NewSQLConnection(dns string) (*SQLConnection, error) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return &SQLConnection{}, err
	}
	db.AutoMigrate(&entity.Pedido{})
	return &SQLConnection{db: db}, nil

}

func (s *SQLConnection) Save(ctx *context.Context, pedido *domain.Pedido) (domain.Pedido, error) {
	p, err := entity.NewPedidoEntity(pedido)
	if err != nil {
		return domain.Pedido{}, err
	}
	result := s.db.Create(p)

	if result.Error != nil {
		return domain.Pedido{}, result.Error
	}
	pedido.Id = p.Id
	return *pedido, nil
}

func (s *SQLConnection) FindByUserId(ctx *context.Context, uId string) ([]domain.Pedido, error) {
	var pedidos []entity.Pedido
	result := s.db.Where("user_id=?", uId).Find(&pedidos)
	if result.Error != nil {
		return []domain.Pedido{}, result.Error
	}

	if len(pedidos) == 0 {
		return []domain.Pedido{}, erros.NewNotFoundErr("Usuario", uId)
	}

	dList, err := entityToPedidoList(&pedidos)
	if err != nil {
		return []domain.Pedido{}, err
	}

	return dList, nil
}

func (s *SQLConnection) FindById(ctx *context.Context, id string) (domain.Pedido, error) {
	var pedido entity.Pedido
	result := s.db.Where("id=?", id).Find(&pedido)
	logger.Info("Erro", result.Error)

	if result.Error != nil {
		return domain.Pedido{}, result.Error
	}
	if pedido.Id == "" {
		return domain.Pedido{}, erros.NewNotFoundErr("Pedido", id)
	}
	logger.Info("Pedido", pedido)

	d, err := pedido.ToDomain()
	if err != nil {
		return domain.Pedido{}, err
	}

	return *d, nil
}

func entityToPedidoList(list *[]entity.Pedido) ([]domain.Pedido, error) {
	var domainList []domain.Pedido
	for _, p := range *list {
		p, err := p.ToDomain()
		if err != nil {
			return []domain.Pedido{}, err
		}
		domainList = append(domainList, *p)
	}
	return domainList, nil
}

package service

import (
	"context"
	"fmt"
	"github.com/PGabrielDev/grpc-go/interal/database"
	"github.com/PGabrielDev/grpc-go/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB *database.Category
}

func NewCategoryService(categoryDB *database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, input *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(input.Nome, input.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}
func (c *CategoryService) ListCategories(context.Context, *pb.Blank) (*pb.ListCategory, error) {

	categories, err := c.CategoryDB.FindAll()
	fmt.Println(categories)
	if err != nil {
		return nil, err
	}
	var listCategories pb.ListCategory
	fmt.Println("Testee")
	for _, c := range categories {
		listCategories.Categories = append(listCategories.Categories, &pb.Category{
			Id:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		})
	}
	fmt.Println("Testee")
	return &listCategories, err
}

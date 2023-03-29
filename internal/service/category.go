package service

import (
	"context"
	"github.com/PGabrielDev/grpc-go/interal/database"
	"github.com/PGabrielDev/grpc-go/internal/pb"
	"io"
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
	if err != nil {
		return nil, err
	}
	var listCategories pb.ListCategory
	for _, c := range categories {
		listCategories.Categories = append(listCategories.Categories, &pb.Category{
			Id:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		})
	}
	return &listCategories, err
}

func (c *CategoryService) CreateCAtegoryStram(stream pb.CategoryService_CreateCAtegoryStramServer) error {
	categories := &pb.ListCategory{}
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}
		cate, err := c.CategoryDB.Create(category.Nome, category.Description)
		if err != nil {
			return err
		}
		categories.Categories = append(categories.Categories, &pb.Category{Id: cate.ID, Name: cate.Name, Description: cate.Description})

	}
}

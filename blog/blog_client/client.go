package main

import (
	"context"
	"fmt"
	"log"

	"github.com/S-ign/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Blog Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	blog, _ := readBlog("60299453e504d424437c91c3", c)
	fmt.Println((*blog).GetContent())
	//fmt.Println(updateBlog("hello", "60299453e504d424437c91c3", c))

}

func createBlog(author, title, content string, c blogpb.BlogServiceClient) (string, error) {
	fmt.Println("creating the blog...")
	blog := &blogpb.Blog{
		AuthorId: author,
		Title:    title,
		Content:  content,
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("unexpected error: %v\n", err)
		return "", err
	}
	fmt.Printf("blog posted: %v\n", createBlogRes)

	return createBlogRes.GetBlog().GetId(), nil
}

func readBlog(blogID string, c blogpb.BlogServiceClient) (*blogpb.Blog, error) {
	fmt.Printf("Reading blog ID: %v\n", blogID)
	blog, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: blogID,
	})
	if err != nil {
		fmt.Printf("error happened while reading: %v\n", err)
		return &blogpb.Blog{}, err
	}
	return blog.GetBlog(), nil
}

func updateBlog(content, blogID string, c blogpb.BlogServiceClient) (*blogpb.Blog, error) {
	// update Blog
	blog, err := readBlog(blogID, c)
	if err != nil {
		return &blogpb.Blog{}, err
	}
	newBlog := &blogpb.Blog{
		Id:       blog.GetId(),
		AuthorId: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  content,
	}
	res, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: newBlog,
	})
	if err != nil {
		return &blogpb.Blog{}, status.Errorf(
			codes.Internal,
			fmt.Sprintf("error updating blog: %v\n", err),
		)
	}
	fmt.Printf("response sent: %v\n", res)
	return newBlog, nil
}

func deleteBlog(blogID string, c blogpb.BlogServiceClient) error {
	res, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{
		BlogId: blogID,
	})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("error deleting blog: %v\n", err),
		)
	}
	fmt.Printf("response sent: %v\n", res)
	return nil
}

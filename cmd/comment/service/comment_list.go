package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/comment/dal/db"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"sync"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{
		ctx: ctx,
	}
}

func (s *CommentListService) CommentList(req *comment.CommentListRequest) ([]*comment.Comment, error) {
	var comments []*comment.Comment

	exist, _ := rpc.PublishExist(s.ctx, req.VideoId)
	if !exist {
		return comments, errno.VideoIsNotExistErr
	}
	dbComments, err := db.GetCommentListByVideoID(s.ctx, req.VideoId)
	if err != nil {
		return comments, errno.CommentIsNotExistErr
	}

	num := len(dbComments)
	commentChan := make(chan comment.Comment, num)
	errChan := make(chan error, num)
	doneChan := make(chan struct{})

	go func() {
		for {
			select {
			case c := <-commentChan:
				comments = append(comments, &c)
			case <-doneChan:
				return
			}
		}
	}()

	var wg sync.WaitGroup
	for _, c := range dbComments {
		wg.Add(1)
		go func(cmt *db.Comment) {
			defer wg.Done()
			resp, e := rpc.UserInfo(s.ctx, &user.UserInfoRequest{
				CurrentUserId: req.UserId,
				UserId:        cmt.UserId,
			})
			if e != nil {
				errChan <- e
			} else {
				commentChan <- comment.Comment{
					Id:         cmt.ID,
					User:       resp.User,
					Content:    &cmt.CommentText,
					CreateDate: cmt.CreatedAt.Format("01-02"),
				}
			}
		}(c)
	}

	wg.Wait()
	doneChan <- struct{}{}

	select {
	case err = <-errChan:
		return comments, err
	default:
	}

	return comments, nil
}

package data

import (
	"context"
	"errors"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dburl = os.Getenv("PG_URI")

var ctx = context.Background()

var db, _ = pgxpool.New(ctx, dburl)

type HTTPError struct {
	Status  string
	Message string
}

func QueryAndScan(ctx context.Context, query string, dest interface{}, args ...interface{}) error {
	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	sliceVal := reflect.ValueOf(dest)
	if sliceVal.Kind() != reflect.Ptr || sliceVal.Elem().Kind() != reflect.Slice {
		return errors.New("dest must be a pointer to a slice")
	}

	sliceElemType := sliceVal.Elem().Type().Elem()
	isPointer := sliceElemType.Kind() == reflect.Ptr
	if isPointer {
		sliceElemType = sliceElemType.Elem() // Get the type that the pointer points to
	}
	if sliceElemType.Kind() != reflect.Struct {
		return errors.New("slice elements must be structs or pointers to structs")
	}

	for rows.Next() {
		var elem reflect.Value
		if isPointer {
			elem = reflect.New(sliceElemType)
		} else {
			elem = reflect.New(sliceElemType).Elem()
		}

		fields := make([]interface{}, elem.Elem().NumField())
		for i := range fields {
			fields[i] = elem.Elem().Field(i).Addr().Interface()
		}

		err := rows.Scan(fields...)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}

		if isPointer {
			sliceVal.Elem().Set(reflect.Append(sliceVal.Elem(), elem))
		} else {
			sliceVal.Elem().Set(reflect.Append(sliceVal.Elem(), elem.Elem()))
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

type Post struct {
	ID              int32     `json:"id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	Type            string    `json:"type"`
	Description     string    `json:"description"`
	ImageUrl        string    `json:"image_url"`
	Slug            string    `json:"slug"`
	DateCreated     time.Time `json:"date_created"`
	TableOfContents bool      `json:"table_of_contents"`
}
type Task struct {
	Task     string `json:"task"`
	Complete bool   `json:"complete"`
}
type CurrentProject struct {
	ID              int32     `json:"id"`
	Name            string    `json:"name"`
	Progress        int32     `json:"progress"`
	Tasks           []Task    `json:"tasks"`
	ImageUrl        string    `json:"image_url"`
	ExpectedRelease time.Time `json:"expected_release"`
}

func GetSinglePost(slug string) *Post {
	var posts []*Post
	QueryAndScan(ctx, "SELECT * FROM posts WHERE slug = $1", &posts, slug)
	if len(posts) > 0 {
		return posts[0]
	}
	var post *Post
	return post
}

func wildcard(input string) string {
	return "%" + input + "%"
}

func SearchPosts(title string, args ...string) ([]*Post, error) {
	var posts []*Post
	//Checks length of args to prevent panics
	if len(args) > 0 {
		//Checks if args[0] (limit) is an empty string, because if it is then we don't have a limit
		if args[0] != "" {
			limit, err := strconv.Atoi(args[0])
			if err != nil {
				return posts, err
			}
			var page int = 1
			//Checks if args[1] (page) is an empty string, because if it is then we don't have a page
			if args[1] != "" {
				page, err = strconv.Atoi(args[1])
				if err != nil {
					return posts, err
				}
			}
			page-- //Decrement page to make sure the offset is 1 less than the page
			err = QueryAndScan(ctx, "SELECT * FROM posts WHERE title ILIKE $1 ORDER BY id LIMIT $2 OFFSET $3 ", &posts, wildcard(title), limit, page*limit)
			if err != nil {
				return posts, err
			}
		} else {
			err := QueryAndScan(ctx, "SELECT * FROM posts WHERE title ILIKE $1 ORDER BY id", &posts, wildcard(title))
			if err != nil {
				return posts, err
			}
		}
	} else {
		err := QueryAndScan(ctx, "SELECT * FROM posts WHERE title ILIKE $1 ORDER BY id", &posts, wildcard(title))
		if err != nil {
			return posts, err
		}
	}

	return posts, nil
}

func GetAllPosts(args ...string) ([]*Post, error) {
	var posts []*Post
	//Checks length of args to prevent panics
	if len(args) > 0 {
		//Checks if args[0] (limit) is an empty string, because if it is then we don't have a limit
		if args[0] != "" {
			limit, err := strconv.Atoi(args[0])
			if err != nil {
				return posts, err
			}
			var page int = 1
			//Checks if args[1] (page) is an empty string, because if it is then we don't have a page
			if args[1] != "" {
				page, err = strconv.Atoi(args[1])
				if err != nil {
					return posts, err
				}
			}
			page-- //Decrement page to make sure the offset is 1 less than the page
			err = QueryAndScan(ctx, "SELECT * FROM posts ORDER BY id LIMIT $1 OFFSET $2 ", &posts, limit, page*limit)
			if err != nil {
				return posts, err
			}
		} else {
			err := QueryAndScan(ctx, "SELECT * FROM posts ORDER BY id  ", &posts)
			if err != nil {
				return posts, err
			}
		}
	} else {
		err := QueryAndScan(ctx, "SELECT * FROM posts ORDER BY id", &posts)
		if err != nil {
			return posts, err
		}
	}

	return posts, nil
}

func GetCurrentProjects(args ...string) ([]*CurrentProject, error) {
	var current_projects []*CurrentProject
	//Checks length of args to prevent panics
	if len(args) > 0 {
		//Checks if args[0] (limit) is an empty string, because if it is then we don't have a limit
		if args[0] != "" {
			limit, err := strconv.Atoi(args[0])
			if err != nil {
				return current_projects, err
			}
			var page int = 1
			//Checks if args[1] (page) is an empty string, because if it is then we don't have a page
			if args[1] != "" {
				page, err = strconv.Atoi(args[1])
				if err != nil {
					return current_projects, err
				}
			}
			page-- //Decrement page to make sure the offset is 1 less than the page
			err = QueryAndScan(ctx, "SELECT * FROM current_projects ORDER BY id LIMIT $1 OFFSET $2 ", &current_projects, limit, page*limit)
			if err != nil {
				return current_projects, err
			}
		} else {
			err := QueryAndScan(ctx, "SELECT * FROM current_projects ORDER BY id  ", &current_projects)
			if err != nil {
				return current_projects, err
			}
		}
	} else {
		err := QueryAndScan(ctx, "SELECT * FROM current_projects ORDER BY id", &current_projects)
		if err != nil {
			return current_projects, err
		}
	}

	return current_projects, nil
}

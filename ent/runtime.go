// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/a6e5h1/ent-sample/ent/comment"
	"github.com/a6e5h1/ent-sample/ent/post"
	"github.com/a6e5h1/ent-sample/ent/schema"
	"github.com/a6e5h1/ent-sample/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescUUID is the schema descriptor for uuid field.
	commentDescUUID := commentFields[1].Descriptor()
	// comment.DefaultUUID holds the default value on creation for the uuid field.
	comment.DefaultUUID = commentDescUUID.Default.(func() uuid.UUID)
	// commentDescContent is the schema descriptor for content field.
	commentDescContent := commentFields[4].Descriptor()
	// comment.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	comment.ContentValidator = func() func(string) error {
		validators := commentDescContent.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(content string) error {
			for _, fn := range fns {
				if err := fn(content); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// commentDescCreatedAt is the schema descriptor for created_at field.
	commentDescCreatedAt := commentFields[5].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the created_at field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(func() time.Time)
	// commentDescModifiedAt is the schema descriptor for modified_at field.
	commentDescModifiedAt := commentFields[6].Descriptor()
	// comment.UpdateDefaultModifiedAt holds the default value on update for the modified_at field.
	comment.UpdateDefaultModifiedAt = commentDescModifiedAt.UpdateDefault.(func() time.Time)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescUUID is the schema descriptor for uuid field.
	postDescUUID := postFields[1].Descriptor()
	// post.DefaultUUID holds the default value on creation for the uuid field.
	post.DefaultUUID = postDescUUID.Default.(func() uuid.UUID)
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[3].Descriptor()
	// post.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	post.TitleValidator = func() func(string) error {
		validators := postDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// postDescBody is the schema descriptor for body field.
	postDescBody := postFields[4].Descriptor()
	// post.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	post.BodyValidator = func() func(string) error {
		validators := postDescBody.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(body string) error {
			for _, fn := range fns {
				if err := fn(body); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postFields[5].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescModifiedAt is the schema descriptor for modified_at field.
	postDescModifiedAt := postFields[6].Descriptor()
	// post.UpdateDefaultModifiedAt holds the default value on update for the modified_at field.
	post.UpdateDefaultModifiedAt = postDescModifiedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUUID is the schema descriptor for uuid field.
	userDescUUID := userFields[1].Descriptor()
	// user.DefaultUUID holds the default value on creation for the uuid field.
	user.DefaultUUID = userDescUUID.Default.(func() uuid.UUID)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
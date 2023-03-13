package model

import (
	"fmt"
	"io"

	"github.com/gofrs/uuid"
)

type UUID struct {
	uuid.UUID
}

func (u UUID) MarshalGQL(w io.Writer) {
	io.WriteString(w, fmt.Sprintf(`"%s"`, u.UUID.String()))
}

func (u *UUID) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		if result, err := uuid.FromString(v); err != nil {
			return err
		} else {
			u = &UUID{result}
		}
		return nil

	case []byte:
		// if err := uuid.FromBytes(v); err != nil {
		// 	return err
		// }
		result, err := uuid.FromBytes(v)
		if err != nil {
			return err
		}
		u = &UUID{result}
		return nil

	default:
		return fmt.Errorf("%T is not a uuid.UUID", v)
	}
}

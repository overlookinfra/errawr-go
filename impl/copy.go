package impl

import (
	"fmt"
	"strings"

	errawr "github.com/puppetlabs/errawr-go"
)

func Copy(e errawr.Error) *Error {
	if ce, ok := e.(*Error); ok {
		return ce
	}

	code := strings.TrimPrefix(e.Code(), fmt.Sprintf("%s_%s_", e.Domain().Key(), e.Section().Key()))

	arguments := e.Arguments()
	eas := make(ErrorArguments, len(arguments))
	for name, value := range arguments {
		eas[name] = &ErrorArgument{
			Value: value,
		}
	}

	metadata := &ErrorMetadata{}
	if hm, ok := e.Metadata().HTTP(); ok {
		metadata.HTTPErrorMetadata = &HTTPErrorMetadata{
			ErrorStatus:  hm.Status(),
			ErrorHeaders: HTTPErrorMetadataHeaders(hm.Headers()),
		}
	}

	return &Error{
		Version: errawr.Version,
		ErrorDomain: &ErrorDomain{
			Key:   e.Domain().Key(),
			Title: e.Domain().Title(),
		},
		ErrorSection: &ErrorSection{
			Key:   e.Section().Key(),
			Title: e.Section().Title(),
		},
		ErrorCode:  code,
		ErrorTitle: e.Title(),
		ErrorDescription: &ErrorDescription{
			Friendly:  e.Description().Friendly(),
			Technical: e.Description().Technical(),
		},
		ErrorArguments: eas,
		ErrorMetadata:  metadata,

		causes: e.Causes(),
		buggy:  e.IsBug(),
	}
}

package matcher

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

// Cmp is a Matcher that matches using go-cmp.
//
// Example usage:
//
//	dbMock.EXPECT().
//		Insert(t, matcher.CmpProto(&someComplex).
func CmpProto(t *testing.T, want interface{}, opts ...cmp.Option) gomock.Matcher {
	opts = append(opts, protocmp.Transform())
	return &cmpMatcher{t, want, opts}
}

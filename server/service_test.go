package server

import (
	"context"
	pb "corpuser/user"
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetUser(t *testing.T) {
	s := NewService()

	testCases := []struct {
		name        string
		req         *pb.UserRequest
		result      bool
		expectedErr bool
	}{
		{
			name:        "Req ok",
			req:         &pb.UserRequest{Id: 1},
			result:      true,
			expectedErr: false,
		},
		{
			name:        "Req with empty Id",
			req:         &pb.UserRequest{},
			expectedErr: true,
		},
		{
			name:        "Nil request",
			req:         nil,
			expectedErr: true,
		},
		{
			name:        "Id not exist request",
			req:         &pb.UserRequest{Id: 100},
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ctx := context.Background()

			// call
			response, err := s.GetUser(ctx, testCase.req)

			t.Log("Got : ", response)

			// assert results expectations
			if testCase.expectedErr {
				g.Expect(err).ToNot(BeNil(), "Result should be nil")
			} else {
				g.Expect(response).ToNot(BeNil(), "Result should not be nil")
			}
		})
	}
}

func TestGetUsersById(t *testing.T) {
	s := NewService()

	testCases := []struct {
		name        string
		req         *pb.UsersRequest
		result      bool
		expectedErr bool
	}{
		{
			name:        "Req ok",
			req:         &pb.UsersRequest{Id: []int64{1, 2}},
			result:      true,
			expectedErr: false,
		},
		{
			name:        "Req with empty Ids",
			req:         &pb.UsersRequest{},
			expectedErr: false,
		},
		{
			name:        "Nil request",
			req:         nil,
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ctx := context.Background()

			// call
			response, err := s.GetUsersById(ctx, testCase.req)

			t.Log("Got : ", response)

			// assert results expectations
			if testCase.expectedErr {
				g.Expect(err).ToNot(BeNil(), "Result should be nil")
			} else {
				g.Expect(response).ToNot(BeNil(), "Result should not be nil")
			}
		})
	}
}

func TestSeed(t *testing.T) {
	g := NewGomegaWithT(t)
	r := Seed()
	g.Expect(r).ToNot(BeNil())

}

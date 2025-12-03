package service_test

import (
	service "ITMO-students/lecture-16/8-testify"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"ITMO-students/lecture-16/8-testify/mocks"
)

var errTest = errors.New("test error")

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context.
type ServiceSuite struct {
	suite.Suite

	ctrl *gomock.Controller
	uRep *mocks.MockUserRepository
	aRep *mocks.MockAutoRepository

	// logger  internal.Logger
	// metrics internal.Metrics
}

// Executes before each test case
// Make sure that same
// service & service mock before each test.
func (suite *ServiceSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())

	suite.uRep = mocks.NewMockUserRepository(suite.ctrl)
	suite.aRep = mocks.NewMockAutoRepository(suite.ctrl)
}

// Executes after each test case.
func (suite *ServiceSuite) TearDownTest() {
	// Verify that all methods that were expected to be called were called.
	suite.ctrl.Finish()
	suite.ctrl = nil

	suite.uRep = nil
	suite.aRep = nil
}

// TestSuiteSecrets runs all suite tests.
func TestSuiteHandle(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (suite *ServiceSuite) TestHandler_Check() {
	tests := map[string]struct {
		expErr           error
		in               struct{}
		prepareMockCalls func()
		want             []string
	}{
		"Normalize errorMessage": {
			prepareMockCalls: func() {
				suite.uRep.EXPECT().
					GetAllBy().
					Return(nil, errTest)
			},
			expErr: errTest,
		},

		"Success": {
			expErr: nil,
			prepareMockCalls: func() {
				suite.uRep.EXPECT().
					GetAllBy().
					Return([]string{"amogus"}, nil)

				suite.aRep.EXPECT().
					Search(gomock.Any()).
					Return([]string{"amogus-pt2"}, nil)
			},
			want: []string{"amogus-pt2"},
		},
	}
	for tn, tc := range tests {
		suite.Run(tn, func() {
			tc.prepareMockCalls()

			svc := service.New(
				suite.uRep,
				suite.aRep,
			)

			got, err := svc.Check()
			suite.Require().ErrorIs(err, tc.expErr)
			suite.Require().Equal(tc.want, got)
		})
	}
}

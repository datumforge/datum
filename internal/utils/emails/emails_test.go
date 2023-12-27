package emails_test

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	sgmail "github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/trisacrypto/directory/pkg/utils/logger"

	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/emails/mock"
	"github.com/datumforge/datum/internal/utils/sendgrid"
)

// If the eyeball flag is set, then the tests will write MIME emails to the testdata directory
var eyeball = flag.Bool("eyeball", false, "Generate MIME emails for eyeball testing")

// This suite mocks the SendGrid email client to verify that email metadata is
// populated corectly and emails can be marshaled into bytes for transmission.
func TestEmailSuite(t *testing.T) {
	suite.Run(t, &EmailTestSuite{})
}

type EmailTestSuite struct {
	suite.Suite
	conf emails.Config
}

func (suite *EmailTestSuite) SetupSuite() {
	// TODO: figure out how to get zap to work for this purpose, or replace directly with zerolog - this package was used because the "MIME" write function was pulled directly from that upstream repo
	logger.Discard()

	suite.conf = emails.Config{
		Testing:    true,
		FromEmail:  "mitb@vandelayindustries.com",
		AdminEmail: "funkytown@sfunk.com",
		Archive:    filepath.Join("fixtures", "emails"),
	}
}

func (suite *EmailTestSuite) BeforeTest(suiteName, testName string) {
	setupMIMEDir(suite.T())
}

func (suite *EmailTestSuite) AfterTest(suiteName, testName string) {
	mock.ResetEmailMock()
}

func (suite *EmailTestSuite) TearDownSuite() {
	logger.ResetLogger()
}

// If eyeball testing is enabled, this removes and recreates the eyeball directory for
// this test
func setupMIMEDir(t *testing.T) {
	if *eyeball {
		path := filepath.Join("testdata", fmt.Sprintf("eyeball%s", t.Name()))
		err := os.RemoveAll(path)
		require.NoError(t, err)
		err = os.MkdirAll(path, 0755)
		require.NoError(t, err)
	}
}

// If eyeball testing is enabled, this writes an SGMailV3 email to a MIME file for manual inspection - you can inspect these files locally with a text editor
// but you can also use tools online to convert MIME -> PDF or similar if you're worried about specifics around rendering; there are also local email utilities / viewer but none are setup as a part of this codebase
func generateMIME(t *testing.T, msg *sgmail.SGMailV3, name string) {
	if *eyeball {
		err := sendgrid.WriteMIME(msg, filepath.Join("testdata", fmt.Sprintf("eyeball%s", t.Name()), name))
		require.NoError(t, err)
	}
}
